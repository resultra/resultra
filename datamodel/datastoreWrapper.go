package datamodel

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"fmt"
	"log"
)

type DummyGetDest struct {
	dummy string
}

func newRootEntityKey(appEngContext appengine.Context,
	entityKind string, encodedID string) (*datastore.Key, error) {

	decodedID, err := decodeUniqueEntityIDStrToInt(encodedID)
	if err != nil {
		return nil, err
	}
	if len(entityKind) == 0 {
		return nil, errors.New("Invalid entity kind used to create key: empty entity kind name")
	}
	rootKey := datastore.NewKey(appEngContext, entityKind, "", decodedID, nil) // nil for no parent entity
	return rootKey, nil
}

// Get an entity key for an existing entity - verify the entity exits
func getExistingRootEntityKey(appEngContext appengine.Context,
	entityKind string, encodedID string) (*datastore.Key, error) {

	rootKey, keyErr := newRootEntityKey(appEngContext, entityKind, encodedID)
	if keyErr != nil {
		return nil, keyErr
	}

	// Verify the existance of the entity with the given key. When inserting child
	// entities, the datastore doesn't verify the key to the parent entity actually
	// exists in the datastore. Using this function makes these types of operations
	// more robust.
	//
	// Pass a dummy value to the query - it is ignored when the KeysOnly query is used
	var dummyDest interface{}
	entityExistsQuery := datastore.NewQuery(entityKind).Filter("__key__=", rootKey).KeysOnly()
	foundKeys, existanceErr := entityExistsQuery.GetAll(appEngContext, dummyDest)
	if existanceErr != nil {
		return nil, fmt.Errorf("Failure verifying existing of datastore entity with existing key (entity kind =%v key=%v): datastore error=%v",
			entityKind, encodedID, existanceErr)
	} else if len(foundKeys) != 1 {
		return nil, fmt.Errorf("Can't find datastore entity for entity: entity kind=%v id=%v",
			entityKind, encodedID)
	}

	return rootKey, nil

}

func newChildEntityKey(appEngContext appengine.Context, encodedChildID string,
	childEntityKind string, parentKey *datastore.Key) (*datastore.Key, error) {

	decodedChildID, err := decodeUniqueEntityIDStrToInt(encodedChildID)
	if err != nil {
		return nil, err
	}
	if parentKey == nil {
		return nil, errors.New("Invalid parent key used to create child key: parent key = nil")
	}
	if len(childEntityKind) == 0 {
		return nil, errors.New("Invalid entity kind used to create key: empty entity kind name")
	}

	childKey := datastore.NewKey(appEngContext, childEntityKind, "", decodedChildID, parentKey)

	return childKey, nil

}

func insertNewEntity(appEngContext appengine.Context, entityKind string,
	parentKey *datastore.Key, src interface{}) (string, error) {

	// nil argument for parentKey (no parent in this case)
	newKey := datastore.NewIncompleteKey(appEngContext, entityKind, parentKey)

	putKey, err := datastore.Put(appEngContext, newKey, src)
	if err != nil {
		return "", err
	}

	encodedID, encodeErr := encodeUniqueEntityIDToStr(putKey)
	if encodeErr != nil {
		return "", encodeErr
	}

	log.Printf("INSERT new entity: kind=%v, id (base36)=%v id(base10)=%v",
		entityKind, encodedID, putKey.IntID())

	return encodedID, nil

}

func getEntityByID(encodedID string, appEngContext appengine.Context, entityKind string, dest interface{}) error {

	decodedID, err := decodeUniqueEntityIDStrToInt(encodedID)
	if err != nil {
		return err
	}

	// nil argument for parentKey (no parent in this case)
	getKey := datastore.NewKey(appEngContext, entityKind, "", decodedID, nil)

	getErr := datastore.Get(appEngContext, getKey, dest)
	if getErr != nil {
		return getErr
	}

	return nil
}
