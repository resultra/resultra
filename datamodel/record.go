package datamodel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
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

type RecordRef struct {
	RecordID    string `json:"recordID"`
	FieldValues Record `json:"fieldValues"`
}

type GetRecordParams struct {
	// TODO - More fields will go here once a record is
	// tied to a database table
	RecordID string `json:"recordID"`
}

func GetRecord(appEngContext appengine.Context, recordParams GetRecordParams) (*RecordRef, error) {

	getRecord := Record{}
	getErr := getRootEntityByID(appEngContext, recordEntityKind, recordParams.RecordID, &getRecord)
	if getErr != nil {
		return nil, fmt.Errorf("Can't get record: Error retrieving existing record: record params=%+v, err = %v", recordParams, getErr)
	}

	return &RecordRef{RecordID: recordParams.RecordID, FieldValues: getRecord}, nil

}

type SetRecordValueParams struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
	Value    string `json:"value"`
}

func SetRecordValue(appEngContext appengine.Context, setValParams SetRecordValueParams) error {

	_, fieldGetErr := GetField(appEngContext, GetFieldParams{setValParams.FieldID})
	if fieldGetErr != nil {
		return fmt.Errorf("Can't set value in SetRecordValue(params=%+v):"+
			" Error retrieving value's field for update: err = %v", setValParams, fieldGetErr)
	}

	recordForUpdate := Record{}
	getErr := getRootEntityByID(appEngContext, recordEntityKind, setValParams.RecordID, &recordForUpdate)
	if getErr != nil {
		return fmt.Errorf("Can't set value in SetRecordValue(params=%+v):"+
			" Error retrieving existing record for update: err = %v", setValParams, getErr)
	}

	//	recordForUpdate.FieldValues[setValParams.FieldID] = setValParams.Value
	recordForUpdate[setValParams.FieldID] = setValParams.Value

	if updateErr := updateExistingRootEntity(appEngContext, recordEntityKind,
		setValParams.RecordID, &recordForUpdate); updateErr != nil {
		return fmt.Errorf("Can't set value: Error retrieving existing record for update: params=%+v, err = %v", setValParams, updateErr)
	}

	return nil

}
