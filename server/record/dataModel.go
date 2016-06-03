package record

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const recordEntityKind string = "Record"

//const recordCreateDateReservedPropName = "__CreateDate__"

type RecFieldValues map[string]interface{}

type Record struct {
	ParentTableID string         `jsaon:"parentTableID"`
	RecordID      string         `json:"recordID"`
	FieldValues   RecFieldValues `json:"fieldValues"`
}

const recordRecordIDFieldName string = "RecordID"
const recordParentTableIDFieldName string = "ParentTableID"

func (rec Record) ValueIsSet(fieldID string) bool {
	_, valueExists := rec.FieldValues[fieldID]
	if valueExists {
		return true
	} else {
		return false
	}
}

func (rec Record) GetTextFieldValue(fieldID string) (string, error) {
	rawVal := rec.FieldValues[fieldID]
	if theStr, validType := rawVal.(string); validType {
		return theStr, nil
	} else {
		return "", fmt.Errorf("Type mismatch retrieving text field value from record: field ID = %v, raw value = %v", fieldID, rawVal)
	}
}

func (fieldVals *RecFieldValues) Load(ch <-chan datastore.Property) error {
	// Note: you might want to clear current values from the map or create a new map
	for p := range ch { // Read until channel is closed
		(*fieldVals)[p.Name] = p.Value
	}
	return nil
}

func (fieldVals *RecFieldValues) Save(ch chan<- datastore.Property) error {
	defer close(ch) // Channel must be closed
	for k, v := range *fieldVals {
		ch <- datastore.Property{Name: k, Value: v}
	}
	return nil
}

type NewRecordParams struct {
	TableID string `json:"tableID"`
}

func NewRecord(appEngContext appengine.Context, params NewRecordParams) (*Record, error) {

	newRecord := Record{ParentTableID: params.TableID,
		RecordID: uniqueID.GenerateUniqueID()}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, recordEntityKind, &newRecord)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new record: error inserting into datastore: %v", insertErr)
	}

	return &newRecord, nil

}

func GetRecord(appEngContext appengine.Context, recordID string) (*Record, error) {

	var getRecord Record

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, recordEntityKind,
		recordRecordIDFieldName, recordID, &getRecord); getErr != nil {
		return nil, fmt.Errorf("GetRecord: Unable to get record from datastore: error = %v", getErr)
	}

	return &getRecord, nil

}

func UpdateExistingRecord(appEngContext appengine.Context, recordID string, rec *Record) (*Record, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		recordID, recordEntityKind, recordRecordIDFieldName, rec); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating record: error = %v", updateErr)
	}

	return rec, nil

}

// TODO - The GetRecord function initially returns all records. However, more fields will be included for:
// - maximum record to retrieve, along with
// - sort and filter criteria.
// - cursor indicating where to start the query (for retrieving results in batches)

type GetRecordsParams struct {
	TableID string `json:"tableID"`
}

func GetRecords(appEngContext appengine.Context, params GetRecordsParams) ([]Record, error) {

	var records []Record

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, params.TableID,
		recordEntityKind, recordParentTableIDFieldName, &records)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve records: form id=%v", params.TableID)
	}

	return records, nil
}

// Validate the field is of the correct type and not a calculated field (if allowCalcField not true). This is for validating
// the field when setting/getting values from regular "literal" fields which store values entered by end-users (as opposed to
// calculated fields)
func ValidateFieldForRecordValue(appEngContext appengine.Context, fieldID string, expectedFieldType string,
	allowCalcField bool) error {

	field, fieldGetErr := field.GetField(appEngContext, fieldID)
	if fieldGetErr != nil {
		return fmt.Errorf(" Error retrieving field for updating/setting value: err = %v", fieldGetErr)
	}
	if field.Type != expectedFieldType {
		return fmt.Errorf("Can't update/set value:"+
			" Type mismatch with field: expecting %v: got %v: field=%+v", expectedFieldType, field.Type, field)

	} else if (!allowCalcField) && field.IsCalcField {
		return fmt.Errorf("Field is a calculated field, setting values directly not supported: field=%v",
			field.RefName)
	}
	return nil
}
