package datastoreWrapper

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"resultra/datasheet/server/generic/uniqueID"
)

// This file contains the public functions for the datastoreWrapper package.
// Basically a standardized insert, get, update (and delete) wrapper function is
// included for both child and root entity types. These functions encapsulate access
// to the GAE datastore by using opaque/encoded IDs instead of raw datastore keys.
// The only place where a raw datastore key is needed is when one entity refers to
// another and needs to create a pointer to it using a *datastore.Key.

func InsertNewRootEntity(appEngContext appengine.Context, entityKind string,
	src interface{}) error {

	// nil argument is for no parent
	newKey := datastore.NewIncompleteKey(appEngContext, entityKind, nil)

	_, err := datastore.Put(appEngContext, newKey, src)
	if err != nil {
		return fmt.Errorf("InsertNewRootEntity: Error inserting entity into datastore: error = %v", err)
	}

	return nil
}

// START New code for using generated UUIDs instead of Googles internal key format.

func getEntityKeyByUUID(appEngContext appengine.Context, entityKind, entityIDFieldName, entityID string) (*datastore.Key, error) {

	matchingEntityIDFilter := fmt.Sprintf("%v =", entityIDFieldName)
	getEntityKeyQuery := datastore.NewQuery(entityKind).Filter(matchingEntityIDFilter, entityID).KeysOnly()
	entityKeys, getKeyErr := getEntityKeyQuery.GetAll(appEngContext, nil)
	if getKeyErr != nil {
		return nil, fmt.Errorf("getEntityKeyByUUID: Unable to find entity to update: no entity found with ID = %v,error=%v", entityID, getKeyErr)
	}
	if len(entityKeys) != 1 {
		return nil, fmt.Errorf("getEntityKeyByUUID: Unable to find entity to update: no entity found with ID = %v", entityID)
	}
	entityKey := entityKeys[0]
	return entityKey, nil
}

func GetEntityByUUID(appEngContext appengine.Context, entityKind string, entityIDFieldName string, entityID string, dest interface{}) error {

	if err := uniqueID.ValidatedWellFormedID(entityID); err != nil {
		return fmt.Errorf("GetEntityByUUID: Invalid entity ID = %v: %v", entityID, err)
	}

	entityKey, getKeyErr := getEntityKeyByUUID(appEngContext, entityKind, entityIDFieldName, entityID)
	if getKeyErr != nil {
		return getKeyErr
	}

	if getErr := datastore.Get(appEngContext, entityKey, dest); getErr != nil {
		return getErr
	}
	return nil

}

// GetAllChildEntities wraps a call to a datastore GetAll() query, given datastore IDs and their entity kinds.
// It also converts the keys to opaque IDs before returning the results.
func GetAllChildEntitiesWithParentUUID(appEngContext appengine.Context, parentID string,
	childEntityKind string, parentIDFieldName string, destSlice interface{}) error {

	if err := uniqueID.ValidatedWellFormedID(parentID); err != nil {
		return fmt.Errorf("GetAllChildEntities: Invalid parent ID: %v", err)
	}

	matchingParentIDFilter := fmt.Sprintf("%v =", parentIDFieldName)
	getChildrenQuery := datastore.NewQuery(childEntityKind).Filter(matchingParentIDFilter, parentID)
	_, getErr := getChildrenQuery.GetAll(appEngContext, destSlice)

	if getErr != nil {
		return fmt.Errorf("GetAllChildEntities: Unable to get all child entities:  parent id=%+v, error=%v",
			parentID, getErr)
	}

	return nil
}

func UpdateExistingEntityByUUID(appEngContext appengine.Context,
	entityID string, entityKind string, entityIDFieldName string, entityToUpdate interface{}) error {

	if err := uniqueID.ValidatedWellFormedID(entityID); err != nil {
		return fmt.Errorf("GetAllChildEntities: Invalid parent ID: %v", err)
	}

	entityKey, getKeyErr := getEntityKeyByUUID(appEngContext, entityKind, entityIDFieldName, entityID)
	if getKeyErr != nil {
		return getKeyErr
	}

	_, putErr := datastore.Put(appEngContext, entityKey, entityToUpdate)
	if putErr != nil {
		return fmt.Errorf("UpdateExistingEntity failed: entity key=%+v, datastore error=%v", entityKey, putErr)
	}

	return nil
}
