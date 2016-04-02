package datamodel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
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
	recordID, insertErr := InsertNewEntity(appEngContext, recordEntityKind, nil, &newRecord)
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
	getErr := GetRootEntityByID(appEngContext, recordEntityKind, recordParams.RecordID, &getRecord)
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
		recordID, encodeErr := EncodeUniqueEntityIDToStr(recKey)
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
func validateFieldForRecordValue(appEngContext appengine.Context, fieldID string, expectedFieldType string,
	allowCalcField bool) error {
	fieldRef, fieldGetErr := GetField(appEngContext, GetFieldParams{fieldID})
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

type SetRecordTextValueParams struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
	Value    string `json:"value"`
}

func SetRecordTextValue(appEngContext appengine.Context, setValParams SetRecordTextValueParams) (*RecordRef, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, setValParams.FieldID, FieldTypeText, false); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordTextValue(params=%+v):"+
			" Error validating record's field for update: %v", setValParams, fieldValidateErr)
	}

	recordForUpdate := Record{}
	getErr := GetRootEntityByID(appEngContext, recordEntityKind, setValParams.RecordID, &recordForUpdate)
	if getErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordValue(params=%+v):"+
			" Error retrieving existing record for update: err = %v", setValParams, getErr)
	}

	updatedRecordRef := RecordRef{setValParams.RecordID, recordForUpdate}
	// TODO - Consider putting a prefix on the field ID before saving it to the datastore.
	//	recordForUpdate.FieldValues[setValParams.FieldID] = setValParams.Value
	log.Printf("SetRecordTextValue: Setting value for field = %v, val = %v", setValParams.FieldID, setValParams.Value)
	updatedRecordRef.FieldValues[setValParams.FieldID] = setValParams.Value

	// Changing this value may have caused the values for calculated fields to also change.
	// Clients of this function need a fully up to date record reference, so the calculated
	// field's values must also be recalculated when a value changes.
	if calcErr := updateCalcFieldValues(appEngContext, &updatedRecordRef); calcErr != nil {
		return nil, fmt.Errorf("Can't set value: Error calculating fields to reflect update: params=%+v, err = %v",
			setValParams, calcErr)
	}

	// Write the record back to the datastore (including the calculated values)
	if updateErr := updateExistingRootEntity(appEngContext, recordEntityKind,
		setValParams.RecordID, &updatedRecordRef.FieldValues); updateErr != nil {
		return nil, fmt.Errorf("Can't set value: Error retrieving existing record for update: params=%+v, err = %v",
			setValParams, updateErr)
	}

	// Return the updated record
	// TODO - Depending upon how calculated values are implemented,
	// this is where calculated field values may also be updated.
	return &updatedRecordRef, nil

}

type SetRecordNumberValueParams struct {
	RecordID string  `json:"recordID"`
	FieldID  string  `json:"fieldID"`
	Value    float64 `json:"value"`
}

func SetRecordNumberValue(appEngContext appengine.Context, setValParams SetRecordNumberValueParams) (*RecordRef, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, setValParams.FieldID, FieldTypeNumber, false); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordTextValue(params=%+v):"+
			" Error validating record's field for update: %v", setValParams, fieldValidateErr)
	}

	// TODO - Check field is not a calculated field.

	recordForUpdate := Record{}
	getErr := GetRootEntityByID(appEngContext, recordEntityKind, setValParams.RecordID, &recordForUpdate)
	if getErr != nil {
		return nil, fmt.Errorf("Can't set value in SetRecordValue(params=%+v):"+
			" Error retrieving existing record for update: err = %v", setValParams, getErr)
	}

	updatedRecordRef := RecordRef{setValParams.RecordID, recordForUpdate}
	/* TODO - The evaluation of calculated field equations is currently returning an error for undefined fields.
		 Instead, there needs to be support for an undefined results, which will cascade up through the
	     recursion. The lack of support for undefined values is preventing the code below from
	     working.*/

	// TODO - Consider putting a prefix on the field ID before saving it to the datastore.
	log.Printf("SetRecordNumberValue: Setting value for field = %v, val = %v", setValParams.FieldID, setValParams.Value)
	updatedRecordRef.FieldValues[setValParams.FieldID] = setValParams.Value

	// Now that the value has been set, update the calculated field values to reflect changes made to the
	// field: Changing this value may have caused the values for calculated fields to also change.
	// Clients of this function need a fully up to date record reference, so the calculated
	// field's values must also be recalculated when a value changes.
	if calcErr := updateCalcFieldValues(appEngContext, &updatedRecordRef); calcErr != nil {
		return nil, fmt.Errorf("Can't set value: Error calculating fields to reflect update: params=%+v, err = %v",
			setValParams, calcErr)
	}

	// TODO - Since a value has changed, update any calculated field values. Should this happen before saving the record?

	if updateErr := updateExistingRootEntity(appEngContext, recordEntityKind,
		setValParams.RecordID, &updatedRecordRef.FieldValues); updateErr != nil {
		return nil, fmt.Errorf("Can't set value: Error saving record for update: params=%+v, err = %v", setValParams, updateErr)
	}

	// Return the updated record
	// TODO - Depending upon how calculated values are implemented,
	// this is where calculated field values may also be updated.
	return &updatedRecordRef, nil

}
