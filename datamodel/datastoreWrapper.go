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

func NewRootEntityKey(appEngContext appengine.Context,
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

func newChildEntityKey(appEngContext appengine.Context,
	entityKind string, encodedID string, parentKey *datastore.Key) (*datastore.Key, error) {

	if parentKey == nil {
		return nil, fmt.Errorf("nil parent key used to create chiild key: child kind=%v, id=%v", entityKind, encodedID)
	}

	decodedID, err := decodeUniqueEntityIDStrToInt(encodedID)
	if err != nil {
		return nil, err
	}
	if len(entityKind) == 0 {
		return nil, errors.New("Invalid entity kind used to create key: empty entity kind name")
	}
	childKey := datastore.NewKey(appEngContext, entityKind, "", decodedID, parentKey) // nil for no parent entity

	return childKey, nil
}

// Verify the existance of the entity with the given key. When inserting child
// entities, the datastore doesn't verify the key to the parent entity actually
// exists in the datastore. Using this function makes these types of operations
// more robust.
//
// Pass a dummy value to the query - it is ignored when the KeysOnly query is used
func verifyEntityExists(appEngContext appengine.Context, entityKind string, existingEntityKey *datastore.Key) error {

	var dummyDest interface{}

	entityExistsQuery := datastore.NewQuery(entityKind).Filter("__key__=", existingEntityKey).KeysOnly()
	foundKeys, existanceErr := entityExistsQuery.GetAll(appEngContext, dummyDest)

	if existanceErr != nil {
		return fmt.Errorf("Failure verifying existing of datastore entity with existing key (entity kind =%v key=%+v): datastore error=%v",
			entityKind, existingEntityKey, existanceErr)
	} else if len(foundKeys) != 1 {
		return fmt.Errorf("Can't find datastore entity for entity: entity kind=%v key=%+v",
			entityKind, existingEntityKey)
	}
	return nil
}

// Get an entity key for an existing entity - verify the entity exits
func GetExistingRootEntityKey(appEngContext appengine.Context,
	entityKind string, encodedID string) (*datastore.Key, error) {

	rootKey, keyErr := NewRootEntityKey(appEngContext, entityKind, encodedID)
	if keyErr != nil {
		return nil, keyErr
	}

	if err := verifyEntityExists(appEngContext, entityKind, rootKey); err != nil {
		return nil, err
	}

	return rootKey, nil

}

func InsertNewEntity(appEngContext appengine.Context, entityKind string,
	parentKey *datastore.Key, src interface{}) (string, error) {

	// nil argument OK for parentKey (meaning no parent)
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

func UpdateExistingEntity(appEngContext appengine.Context,
	encodedID string, entityKind string,
	parentKey *datastore.Key, src interface{}) error {

	childKey, keyErr := newChildEntityKey(appEngContext, entityKind, encodedID, parentKey)
	if keyErr != nil {
		return fmt.Errorf("updateExistingEntity failed: err = %v", keyErr)
	}

	_, putErr := datastore.Put(appEngContext, childKey, src)
	if putErr != nil {
		return fmt.Errorf("updateExistingEntity failed: entity kind=%v,child key=%+v, parent key=%+v, datastore error=%v",
			entityKind, childKey, parentKey, putErr)
	}

	return nil

}

func updateExistingRootEntity(appEngContext appengine.Context, entityKind string,
	encodedID string, src interface{}) error {

	rootKey, keyErr := NewRootEntityKey(appEngContext, entityKind, encodedID)
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

func GetChildEntityByID(encodedID string, appEngContext appengine.Context, entityKind string,
	parentKey *datastore.Key, dest interface{}) error {

	decodedID, err := decodeUniqueEntityIDStrToInt(encodedID)
	if err != nil {
		return err
	}

	// nil argument for parentKey (no parent in this case)
	getKey := datastore.NewKey(appEngContext, entityKind, "", decodedID, parentKey)

	if getErr := datastore.Get(appEngContext, getKey, dest); getErr != nil {
		return getErr
	}

	return nil
}

func GetRootEntityByID(appEngContext appengine.Context, entityKind string, encodedID string, dest interface{}) error {

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
