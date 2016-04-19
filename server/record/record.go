package record

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/field"
)

const recordEntityKind string = "Record"

//const recordCreateDateReservedPropName = "__CreateDate__"

type Record map[string]interface{}

func (rec Record) ValueIsSet(fieldID string) bool {
	_, valueExists := rec[fieldID]
	if valueExists {
		return true
	} else {
		return false
	}
}

func (rec Record) GetTextFieldValue(fieldID string) (string, error) {
	rawVal := rec[fieldID]
	if theStr, validType := rawVal.(string); validType {
		return theStr, nil
	} else {
		return "", fmt.Errorf("Type mismatch retrieving text field value from record: field ID = %v, raw value = %v", fieldID, rawVal)
	}
}

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
	recordID, insertErr := datastoreWrapper.InsertNewEntity(appEngContext, recordEntityKind, nil, &newRecord)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new field: error inserting into datastore: %v", insertErr)
	}

	return &RecordRef{recordID, newRecord}, nil

}

type RecordID struct {
	// TODO - More fields will go here once a record is
	// tied to a database table
	RecordID string `json:"recordID"`
}

func GetRecord(appEngContext appengine.Context, recordParams RecordID) (*RecordRef, error) {

	getRecord := Record{}
	getErr := datastoreWrapper.GetRootEntityByID(appEngContext, recordEntityKind, recordParams.RecordID, &getRecord)
	if getErr != nil {
		return nil, fmt.Errorf("Can't get record: Error retrieving existing record: record params=%+v, err = %v", recordParams, getErr)
	}

	return &RecordRef{RecordID: recordParams.RecordID, FieldValues: getRecord}, nil

}

func UpdateExistingRecord(appEngContext appengine.Context, recordID RecordID, rec Record) (*RecordRef, error) {

	if updateErr := datastoreWrapper.UpdateExistingRootEntity(appEngContext, recordEntityKind,
		recordID.RecordID, &rec); updateErr != nil {
		return nil, fmt.Errorf("UpdateExistingRecord: Can't set value: Error updating existing record: params=%+v, err = %v",
			recordID, updateErr)
	}
	return &RecordRef{RecordID: recordID.RecordID, FieldValues: rec}, nil

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
		recordID, encodeErr := datastoreWrapper.EncodeUniqueEntityIDToStr(recKey)
		if encodeErr != nil {
			return nil, fmt.Errorf("Failed to encode unique ID for record: key=%+v, encode err=%v", recKey, encodeErr)
		}
		recordRefs[recIter] = RecordRef{recordID, currRec}
	}
	return recordRefs, nil
}

// Validate the field is of the correct type and not a calculated field (if allowCalcField not true). This is for validating
// the field when setting/getting values from regular "literal" fields which store values entered by end-users (as opposed to
// calculated fields)
func ValidateFieldForRecordValue(appEngContext appengine.Context, fieldID string, expectedFieldType string,
	allowCalcField bool) error {
	fieldRef, fieldGetErr := field.GetField(appEngContext, field.GetFieldParams{fieldID})
	if fieldGetErr != nil {
		return fmt.Errorf(" Error retrieving field for updating/setting value: err = %v", fieldGetErr)
	}
	if fieldRef.FieldInfo.Type != expectedFieldType {
		return fmt.Errorf("Can't update/set value:"+
			" Type mismatch with field: expecting %v: got %v", expectedFieldType, fieldRef.FieldInfo.Type)

	} else if (!allowCalcField) && fieldRef.FieldInfo.IsCalcField {
		return fmt.Errorf("Field is a calculated field, setting values directly not supported: field=%v",
			fieldRef.FieldInfo.RefName)
	}
	return nil
}
