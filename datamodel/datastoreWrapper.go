package datamodel

import (
	"appengine"
	"appengine/datastore"
)

func insertNewEntity(appEngContext appengine.Context, entityKind string, src interface{}) (string, error) {

	newKey := datastore.NewIncompleteKey(appEngContext, entityKind, nil)

	putKey, err := datastore.Put(appEngContext, newKey, src)
	if err != nil {
		return "", err
	}

	encodedID, encodeErr := encodeUniqueEntityIDToStr(putKey)
	if encodeErr != nil {
		return "", encodeErr
	}

	return encodedID, nil

}

func getEntityByID(encodedID string, appEngContext appengine.Context, entityKind string, dest interface{}) error {

	decodedID, err := decodeUniqueEntityIDStrToInt(encodedID)
	if err != nil {
		return err
	}

	getKey := datastore.NewKey(appEngContext, entityKind, "", decodedID, nil)

	getErr := datastore.Get(appEngContext, getKey, dest)
	if getErr != nil {
		return getErr
	}

	return nil
}
