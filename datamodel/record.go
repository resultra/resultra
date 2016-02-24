package datamodel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"math"
)

const recordEntityKind string = "Record"

//const recordCreateDateReservedPropName = "__CreateDate__"

type Record map[string]interface{}

type RecordRef struct {
	RecordID    string `json:"recordID"`
	FieldValues Record `json:"fieldValues"`
}

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

func NewRecord(appEngContext appengine.Context) (*RecordRef, error) {

	newRecord := Record{}

	// TODO - Replace nil with database parent
	recordID, insertErr := insertNewEntity(appEngContext, recordEntityKind, nil, &newRecord)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new field: error inserting into datastore: %v", insertErr)
	}

	return &RecordRef{recordID, newRecord}, nil

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

// TODO - The GetRecord function initially returns all records. However, more fields will be included for:
// - maximum record to retrieve, along with
// - sort and filter criteria.
// - parent table ID (once tables are implemented)
// - cursor indicating where to start the query (for retrieving results in batches)

func GetRecords(appEngContext appengine.Context) ([]RecordRef, error) {

	var records []Record
	recordQuery := datastore.NewQuery(recordEntityKind)
	keys, err := recordQuery.GetAll(appEngContext, &records)
	if err != nil {
		return nil, fmt.Errorf("GetRecords: Unable to retrieve records from datastore: datastore error = %v", err)
	}

	recordRefs := make([]RecordRef, len(records))
	for recIter, currRec := range records {
		recKey := keys[recIter]
		recordID, encodeErr := encodeUniqueEntityIDToStr(recKey)
		if encodeErr != nil {
			return nil, fmt.Errorf("Failed to encode unique ID for record: key=%+v, encode err=%v", recKey, encodeErr)
		}
		recordRefs[recIter] = RecordRef{recordID, currRec}
	}
	return recordRefs, nil
}

type SetRecordTextValueParams struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
	Value    string `json:"value"`
}

func validateFieldForRecordValue(appEngContext appengine.Context, fieldID string, expectedFieldType string) error {
	fieldRef, fieldGetErr := GetField(appEngContext, GetFieldParams{fieldID})
	if fieldGetErr != nil {
		return fmt.Errorf(" Error retrieving field for updating/setting value: err = %v", fieldGetErr)
	}
	if fieldRef.FieldInfo.Type != expectedFieldType {
		return fmt.Errorf("Can't update/set value:"+
			" Type mismatch with field: expecting %v: got %v", expectedFieldType, fieldRef.FieldInfo.Type)

	}
	return nil
}

func (recordRef RecordRef) GetTextRecordValue(appEngContext appengine.Context, fieldID string) (string, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, fieldID, fieldTypeText); fieldValidateErr != nil {
		return "", fmt.Errorf("Can't get value from record = %+v and fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", recordRef, fieldID, fieldValidateErr)
	}

	val, foundVal := recordRef.FieldValues[fieldID]
	if !foundVal {
		return "", fmt.Errorf("Undefined value for record with ID = %v, field = %v",
			recordRef.RecordID, fieldID)
	} else {
		if textVal, foundText := val.(string); !foundText {
			return "", fmt.Errorf("Type mismatch retrieving value from record with ID = %v, field = %v:"+
				" expecting string, got %v", recordRef.RecordID, fieldID, val)

		} else {
			return textVal, nil
		}
	} // else (if found a value for the given field ID)
}

func (recordRef RecordRef) GetNumberRecordValue(appEngContext appengine.Context, fieldID string) (float64, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, fieldID, fieldTypeNumber); fieldValidateErr != nil {
		return math.NaN(), fmt.Errorf("Can't get value from record = %+v and fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", recordRef, fieldID, fieldValidateErr)
	}

	val, foundVal := recordRef.FieldValues[fieldID]
	if !foundVal {
		return math.NaN(), fmt.Errorf("Undefined value for record with ID = %v, field = %v",
			recordRef.RecordID, fieldID)
	} else {
		if numberVal, foundNumber := val.(float64); !foundNumber {
			return math.NaN(), fmt.Errorf("Type mismatch retrieving value from record with ID = %v, field = %v:"+
				" expecting number, got %v", recordRef.RecordID, fieldID, val)
		} else {
			return numberVal, nil
		}
	} // else (if found a value for the given field ID)
}

func SetRecordTextValue(appEngContext appengine.Context, setValParams SetRecordTextValueParams) (*RecordRef, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, setValParams.FieldID, fieldTypeText); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordTextValue(params=%+v):"+
			" Error validating record's field for update: %v", setValParams, fieldValidateErr)
	}

	recordForUpdate := Record{}
	getErr := getRootEntityByID(appEngContext, recordEntityKind, setValParams.RecordID, &recordForUpdate)
	if getErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordValue(params=%+v):"+
			" Error retrieving existing record for update: err = %v", setValParams, getErr)
	}

	// TODO - Consider putting a prefix on the field ID before saving it to the datastore.
	//	recordForUpdate.FieldValues[setValParams.FieldID] = setValParams.Value
	recordForUpdate[setValParams.FieldID] = setValParams.Value

	if updateErr := updateExistingRootEntity(appEngContext, recordEntityKind,
		setValParams.RecordID, &recordForUpdate); updateErr != nil {
		return nil, fmt.Errorf("Can't set value: Error retrieving existing record for update: params=%+v, err = %v",
			setValParams, updateErr)
	}

	// Return the updated record
	// TODO - Depending upon how calculated values are implemented,
	// this is where calculated field values may also be updated.
	return &RecordRef{setValParams.RecordID, recordForUpdate}, nil

}

type SetRecordNumberValueParams struct {
	RecordID string  `json:"recordID"`
	FieldID  string  `json:"fieldID"`
	Value    float64 `json:"value"`
}

func SetRecordNumberValue(appEngContext appengine.Context, setValParams SetRecordNumberValueParams) (*RecordRef, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, setValParams.FieldID, fieldTypeNumber); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordTextValue(params=%+v):"+
			" Error validating record's field for update: %v", setValParams, fieldValidateErr)
	}

	recordForUpdate := Record{}
	getErr := getRootEntityByID(appEngContext, recordEntityKind, setValParams.RecordID, &recordForUpdate)
	if getErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordValue(params=%+v):"+
			" Error retrieving existing record for update: err = %v", setValParams, getErr)
	}

	// TODO - Consider putting a prefix on the field ID before saving it to the datastore.
	//	recordForUpdate.FieldValues[setValParams.FieldID] = setValParams.Value
	recordForUpdate[setValParams.FieldID] = setValParams.Value

	if updateErr := updateExistingRootEntity(appEngContext, recordEntityKind,
		setValParams.RecordID, &recordForUpdate); updateErr != nil {
		return nil, fmt.Errorf("Can't set value: Error retrieving existing record for update: params=%+v, err = %v", setValParams, updateErr)
	}

	// Return the updated record
	// TODO - Depending upon how calculated values are implemented,
	// this is where calculated field values may also be updated.
	return &RecordRef{setValParams.RecordID, recordForUpdate}, nil

}
