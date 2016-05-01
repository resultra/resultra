package datastoreWrapper

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
)

// This file contains the public functions for the datastoreWrapper package.
// Basically a standardized insert, get, update (and delete) wrapper function is
// included for both child and root entity types. These functions encapsulate access
// to the GAE datastore by using opaque/encoded IDs instead of raw datastore keys.
// The only place where a raw datastore key is needed is when one entity refers to
// another and needs to create a pointer to it using a *datastore.Key.

func InsertNewRootEntity(appEngContext appengine.Context, entityKind string,
	src interface{}) (string, error) {

	return insertNewEntity(appEngContext, entityKind, nil, src)
}

// Get an entity key for an existing entity,but first verify the entity exits
func GetExistingRootEntityKey(appEngContext appengine.Context,
	entityKind string, encodedID string) (*datastore.Key, error) {

	rootKey, keyErr := newRootEntityKey(appEngContext, entityKind, encodedID)
	if keyErr != nil {
		return nil, keyErr
	}

	if err := verifyEntityExists(appEngContext, entityKind, rootKey); err != nil {
		return nil, err
	}

	return rootKey, nil

}

func GetRootEntity(appEngContext appengine.Context, entityKind string, encodedID string, dest interface{}) error {

	decodedID, err := decodeUniqueEntityIDStrToInt(encodedID)
	if err != nil {
		return err
	}

	// nil argument for parentKey (no parent in this case)
	getKey := datastore.NewKey(appEngContext, entityKind, "", decodedID, nil) // nil for parent key
	log.Printf("GET root entity: kind=%v, id (base36)=%v key=%+v", entityKind, encodedID, getKey)

	getErr := datastore.Get(appEngContext, getKey, dest)
	if getErr != nil {
		return getErr
	}

	return nil
}

// GetRootEntityFromKey retrieve a single root entity given a key to that entity. This wrapper function is generally
// intended for use with entities which store keys to other other entities as a way of referencing those entities.
// Although the given rootKey exposes the Google datastore implementation, storing references to other entities using
// *datastore.Key is the most straightforward way to store permanent references to other objects.
//
// The alternative would be to store an opaque string key instead of the *datastore.Key, but this would mean we'd no longer be able
// to change the ID format; however, this design alternative would technically hide more of the implementation of the
// Google datastore. In fact, this is how records' values are stored into different fields; the property name is overridden
// as the fields' opaque string ID.
//
// TODO(IMPORTANT) - Given the above,evaluate whether to shift entirely off *datastoreKey. It is currently only used for field references
// and it wouldn't be too expensive to switch over to unique string IDs at this point. This would have the benefit of
// completely abstracting the application from Google's datastore. With that in place, any back-end datastore could be used
// with the application for performance or cost reasons. Even a simple JSON storage format could be used for entities.
//
// Some considerations:
//    - This would make evaluation of other database backends much easier (or potentially shifting to another)
//    - When saving a version of the database as a template or for backup, this would make "swizzling" of the
//      unique IDs much easier.
func GetRootEntityFromKey(appEngContext appengine.Context, entityKind string, rootKey *datastore.Key, dest interface{}) (string, error) {

	rootID, encodeErr := encodeUniqueEntityIDToStr(rootKey)
	if encodeErr != nil {
		return "", fmt.Errorf("GetRootEntityFromKey: Failed to encode unique ID: key=%+v, encode err=%v", rootKey, encodeErr)
	}

	if getErr := GetRootEntity(appEngContext, entityKind, rootID, dest); getErr != nil {
		return "", getErr
	}

	return rootID, nil

}

// GetAllRootEntities wraps a call to a datastore GetAll() query, given datastore IDs and their entity kinds.
// It also converts the keys to opaque IDs before returning the results.
func GetAllRootEntities(appEngContext appengine.Context, entityKind string, destSlice interface{}) ([]string, error) {

	query := datastore.NewQuery(entityKind)
	keys, err := query.GetAll(appEngContext, destSlice)
	if err != nil {
		return nil, fmt.Errorf("GetAllRootEntities: Unable to retrieve layouts from datastore: datastore error =%v", err)
	}

	rootIDs := make([]string, len(keys))
	for i, currKey := range keys {
		rootID, encodeErr := encodeUniqueEntityIDToStr(currKey)
		if encodeErr != nil {
			return nil, fmt.Errorf("GetAllRootEntities: Failed to encode unique ID for layout: key=%+v, encode err=%v", currKey, encodeErr)
		}
		rootIDs[i] = rootID
	}
	return rootIDs, nil

}

func UpdateExistingRootEntity(appEngContext appengine.Context, entityKind string,
	encodedID string, src interface{}) error {

	rootKey, keyErr := newRootEntityKey(appEngContext, entityKind, encodedID)
	if keyErr != nil {
		return fmt.Errorf("updateExistingRootEntity failed: err = %v", keyErr)
	}

	_, putErr := datastore.Put(appEngContext, rootKey, src)
	if putErr != nil {
		return fmt.Errorf("updateExistingRootEntity Put() failed: entity kind=%v,root key=%+v, datastore error=%v",
			entityKind, rootKey, putErr)
	}

	return nil

}

func InsertNewChildEntity(appEngContext appengine.Context,
	parentID string, entityRel ChildParentEntityRel, newEntity interface{}) (string, error) {

	log.Printf("InsertNewChildEntity: Updating child entity: parent=(id=%v,kind=%v) child=(kind=%v): new entity = %+v)",
		parentID, entityRel.ParentEntityKind, entityRel.ChildEntityKind, newEntity)

	parentKey, getErr := GetExistingRootEntityKey(appEngContext, entityRel.ParentEntityKind, parentID)
	if getErr != nil {
		return "", fmt.Errorf("InsertNewChildEntity: Unable to retrieve parent entity: %v", getErr)
	}

	childID, insertErr := insertNewEntity(appEngContext, entityRel.ChildEntityKind, parentKey, newEntity)
	if insertErr != nil {
		return "", insertErr
	}

	return encodeChildEntityIDToStr(parentID, childID), nil

}

// GetChildEntity retrieves a child entity for the given unique ID and associated entity kind for both the child
// and parent entitiy.
func GetChildEntity(appEngContext appengine.Context,
	encodedChildID string, entityRel ChildParentEntityRel, getDest interface{}) error {

	uniqueID, decodeErr := decodeUniqueChildID(encodedChildID)
	if decodeErr != nil {
		return fmt.Errorf("getChildEntity: unable to decode child id: %v", decodeErr)
	}

	log.Printf("GetChildEntity: Getting child entity: parent=(id=%v,kind=%v) child=(kind=%v,id=%v)",
		uniqueID.parentID, entityRel.ParentEntityKind, entityRel.ChildEntityKind, uniqueID.childID)

	parentKey, parentKeyErr := newRootEntityKey(appEngContext, entityRel.ParentEntityKind, uniqueID.parentID)
	if parentKeyErr != nil {
		return fmt.Errorf("GetChildEntity: unable to retrieve parent key for entity: parent id = %v, parent kind = %v",
			uniqueID.parentID, entityRel.ParentEntityKind)
	}

	if getErr := getChildEntityByID(uniqueID.childID, appEngContext, entityRel.ChildEntityKind, parentKey, getDest); getErr != nil {
		return fmt.Errorf("getChildEntity: Unable to get child entity from datastore: child id = %v, error = %v",
			encodedChildID, getErr)
	}

	return nil
}

// GetAllChildEntities wraps a call to a datastore GetAll() query, given datastore IDs and their entity kinds.
// It also converts the keys to opaque IDs before returning the results.
func GetAllChildEntities(appEngContext appengine.Context, parentID string,
	entityRel ChildParentEntityRel, destSlice interface{}) ([]string, error) {

	parentKey, parentErr := GetExistingRootEntityKey(appEngContext, entityRel.ParentEntityKind, parentID)
	if parentErr != nil {
		return nil, fmt.Errorf("GetAllChildEntities: Unable to retrieve parent: parent id = %v, parent entity kind = %v, error=%v",
			parentID, entityRel.ParentEntityKind, parentErr)
	}

	getAllQuery := datastore.NewQuery(entityRel.ChildEntityKind).Ancestor(parentKey)
	keys, getErr := getAllQuery.GetAll(appEngContext, destSlice)

	if getErr != nil {
		return nil, fmt.Errorf("GetAllChildEntities: Unable to get all child entities:  parent id=%+v, entity kinds=%+v, error=%v",
			parentID, entityRel, getErr)
	}

	childIDs := make([]string, len(keys))
	for keyIter, currKey := range keys {
		childID, encodeErr := encodeUniqueEntityIDToStr(currKey)
		if encodeErr != nil {
			return nil, fmt.Errorf("GetAllChildEntities: Failed to encode unique ID: key=%+v, encode err=%v", currKey, encodeErr)
		}
		childIDs[keyIter] = encodeChildEntityIDToStr(parentID, childID)
	}

	return childIDs, nil
}

func UpdateExistingChildEntity(appEngContext appengine.Context, childID string,
	entityRel ChildParentEntityRel, entityToUpdate interface{}) error {

	uniqueID, decodeErr := decodeUniqueChildID(childID)
	if decodeErr != nil {
		return fmt.Errorf("UpdateExistingChildEntity: Invalid child ID: %v", decodeErr)
	}

	log.Printf("UpdateExistingChildEntity: Updating child entity: parent=(id=%v,kind=%v) child=(id=%v,kind=%v: updated entity = %+v)",
		uniqueID.parentID, entityRel.ParentEntityKind, uniqueID.childID, entityRel.ChildEntityKind, entityToUpdate)

	parentKey, getErr := GetExistingRootEntityKey(appEngContext, entityRel.ParentEntityKind, uniqueID.parentID)
	if getErr != nil {
		return fmt.Errorf("UpdateExistingChildEntity: Unable to retrieve parent entity: %v", getErr)
	}

	if updateErr := updateExistingEntity(appEngContext,
		uniqueID.childID, entityRel.ChildEntityKind, parentKey, entityToUpdate); updateErr != nil {
		return fmt.Errorf("UpdateExistingChildEntity: Unable to update child entity: %v", updateErr)
	}

	return nil
}
