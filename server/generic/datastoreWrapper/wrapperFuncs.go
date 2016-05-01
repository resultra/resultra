package datastoreWrapper

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	//	"log"
)

// This file contains the public functions for the datastoreWrapper package.
// Basically a standardized insert, get, update (and delete) wrapper function is
// included for both child and root entity types. These functions encapsulate access
// to the GAE datastore by using opaque/encoded IDs instead of raw datastore keys.
// The only place where a raw datastore key is needed is when one entity refers to
// another and needs to create a pointer to it using a *datastore.Key.

func InsertNewRootEntity(appEngContext appengine.Context, entityKind string,
	src interface{}) (string, error) {

	// nil argument is for no parent
	newKey := datastore.NewIncompleteKey(appEngContext, entityKind, nil)

	putKey, err := datastore.Put(appEngContext, newKey, src)
	if err != nil {
		return "", err
	}

	return putKey.Encode(), nil
}

// Get an entity key for an existing entity,but first verify the entity exits
func GetExistingRootEntityKey(appEngContext appengine.Context,
	entityKind string, encodedID string) (*datastore.Key, error) {

	rootKey, decodeErr := decodeUniqueEntityIDStrToKey(encodedID)
	if decodeErr != nil {
		return nil, fmt.Errorf("GetRootEntity: Unable to decode entity key: %v", decodeErr)
	}

	return rootKey, nil

}

func GetRootEntity(appEngContext appengine.Context, entityKind string, encodedID string, dest interface{}) error {

	rootKey, decodeErr := decodeUniqueEntityIDStrToKey(encodedID)
	if decodeErr != nil {
		return fmt.Errorf("GetRootEntity: Unable to decode entity key: %v", decodeErr)
	}

	getErr := datastore.Get(appEngContext, rootKey, dest)
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
func GetRootEntityFromKey(appEngContext appengine.Context, entityKind string,
	rootKey *datastore.Key, dest interface{}) (string, error) {

	if getErr := datastore.Get(appEngContext, rootKey, dest); getErr != nil {
		return "", getErr
	}

	return rootKey.Encode(), nil

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

func UpdateExistingRootEntity(appEngContext appengine.Context,
	encodedID string, src interface{}) error {

	rootKey, decodeErr := decodeUniqueEntityIDStrToKey(encodedID)
	if decodeErr != nil {
		return fmt.Errorf("UpdateExistingRootEntity: Unable to decode entity key: %v", decodeErr)
	}

	_, putErr := datastore.Put(appEngContext, rootKey, src)
	if putErr != nil {
		return fmt.Errorf("UpdateExistingRootEntity Put() failed: root key=%+v, datastore error=%v",
			rootKey, putErr)
	}

	return nil

}

func InsertNewChildEntity(appEngContext appengine.Context,
	parentID string, childEntityKind string, newEntity interface{}) (string, error) {

	parentKey, decodeErr := decodeUniqueEntityIDStrToKey(parentID)
	if decodeErr != nil {
		return "", fmt.Errorf("InsertNewChildEntity: Unable to decode parent key: %v", decodeErr)
	}

	newKey := datastore.NewIncompleteKey(appEngContext, childEntityKind, parentKey)

	putKey, err := datastore.Put(appEngContext, newKey, newEntity)
	if err != nil {
		return "", err
	}

	return putKey.Encode(), nil

}

// GetChildEntity retrieves a child entity for the given unique ID and associated entity kind for both the child
// and parent entitiy.
func GetChildEntity(appEngContext appengine.Context,
	encodedChildID string, getDest interface{}) error {

	childKey, decodeErr := decodeUniqueEntityIDStrToKey(encodedChildID)
	if decodeErr != nil {
		return fmt.Errorf("GetChildEntity: Unable to decode child key: %v", decodeErr)
	}

	if getErr := datastore.Get(appEngContext, childKey, getDest); getErr != nil {
		return getErr
	}

	return nil
}

// GetAllChildEntities wraps a call to a datastore GetAll() query, given datastore IDs and their entity kinds.
// It also converts the keys to opaque IDs before returning the results.
func GetAllChildEntities(appEngContext appengine.Context, parentID string,
	childEntityKind string, destSlice interface{}) ([]string, error) {

	parentKey, decodeErr := decodeUniqueEntityIDStrToKey(parentID)
	if decodeErr != nil {
		return nil, fmt.Errorf("InsertNewChildEntity: Unable to decode parent key: %v", decodeErr)
	}

	getAllQuery := datastore.NewQuery(childEntityKind).Ancestor(parentKey)
	keys, getErr := getAllQuery.GetAll(appEngContext, destSlice)

	if getErr != nil {
		return nil, fmt.Errorf("GetAllChildEntities: Unable to get all child entities:  parent id=%+v, error=%v",
			parentID, getErr)
	}

	childIDs := make([]string, len(keys))
	for keyIter, currKey := range keys {
		childIDs[keyIter] = currKey.Encode()
	}

	return childIDs, nil
}

func UpdateExistingChildEntity(appEngContext appengine.Context, childID string, entityToUpdate interface{}) error {

	childKey, decodeErr := decodeUniqueEntityIDStrToKey(childID)
	if decodeErr != nil {
		return fmt.Errorf("UpdateExistingChildEntity: Unable to decode entity key: %v", decodeErr)
	}

	_, putErr := datastore.Put(appEngContext, childKey, entityToUpdate)
	if putErr != nil {
		return fmt.Errorf("updateExistingEntity failed: child key=%+v, datastore error=%v", childKey, putErr)
	}

	return nil
}
