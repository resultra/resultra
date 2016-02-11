package datamodel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	//	"log"
)

const recordEntityKind string = "Record"

//const recordCreateDateReservedPropName = "__CreateDate__"

type Record map[string]interface{}

func (rec *Record) Load(ch <-chan datastore.Property) error {
	// Note: you might want to clear current values from the map or create a new map
	for p := range ch { // Read until channel is closed
		(*rec)[p.Name] = p.Value
	}
	return nil
}

func (rec *Record) Save(ch chan<- datastore.Property) error {
	defer close(ch) // Channel must be closed
	for k, v := range *rec {
		ch <- datastore.Property{Name: k, Value: v}
	}
	return nil
}

func SaveNewRecord(appEngContext appengine.Context, newRecord Record) (string, error) {

	// TODO - Replace nil with database parent
	recordID, insertErr := insertNewEntity(appEngContext, recordEntityKind, nil, &newRecord)
	if insertErr != nil {
		return "", fmt.Errorf("Can't create new field: error inserting into datastore: %v", insertErr)
	}

	return recordID, nil

}
