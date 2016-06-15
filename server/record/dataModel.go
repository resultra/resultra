package record

import (
	"fmt"
	"github.com/gocql/gocql"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

type RecFieldValues map[string]interface{}

type Record struct {
	ParentTableID string         `json:"parentTableID"`
	RecordID      string         `json:"recordID"`
	FieldValues   RecFieldValues `json:"fieldValues"`
}

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

type NewRecordParams struct {
	ParentTableID string `json:"parentTableID"`
}

func NewRecord(params NewRecordParams) (*Record, error) {

	newRecord := Record{ParentTableID: params.ParentTableID,
		RecordID:    gocql.TimeUUID().String(),
		FieldValues: RecFieldValues{}}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedFieldVals, encodeErr := generic.EncodeJSONString(newRecord.FieldValues)
	if encodeErr != nil {
		return nil, fmt.Errorf("NewRecord: failure encoding field values: error = %v", encodeErr)
	}

	if insertErr := dbSession.Query(`INSERT INTO records (tableID, record_id, field_values) VALUES (?,?,?)`,
		newRecord.ParentTableID, newRecord.RecordID, encodedFieldVals).Exec(); insertErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", insertErr)
	}

	return &newRecord, nil

}

func GetRecord(parentTableID string, recordID string) (*Record, error) {

	getRecord := Record{"", "", RecFieldValues{}}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("GetRecord: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedFieldVals := ""
	getErr := dbSession.Query(`SELECT tableID,record_id,field_values FROM records WHERE tableID=? AND record_id=? LIMIT 1`,
		parentTableID, recordID).Scan(&getRecord.ParentTableID,
		&getRecord.RecordID,
		&encodedFieldVals)
	if getErr != nil {
		return nil, fmt.Errorf("GetRecord: Unabled to get record: id = %v: datastore err=%v", recordID, getErr)
	}

	if decodeErr := generic.DecodeJSONString(encodedFieldVals, &getRecord.FieldValues); decodeErr != nil {
		return nil, fmt.Errorf("GetRecord: Unabled to get record: id = %v: datastore err=%v", recordID, decodeErr)
	}

	return &getRecord, nil

}

func UpdateExistingRecord(rec *Record) (*Record, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedFieldVals, encodeErr := generic.EncodeJSONString(rec.FieldValues)
	if encodeErr != nil {
		return nil, fmt.Errorf("NewRecord: failure encoding field values: error = %v", encodeErr)
	}

	if updateErr := dbSession.Query(`UPDATE records set field_values=? WHERE tableID=? and record_id=?`,
		encodedFieldVals, rec.ParentTableID, rec.RecordID).Exec(); updateErr != nil {
		return nil, fmt.Errorf("UpdateExistingRecord: Can't update record: unable to update record: error = %v", updateErr)
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

func GetRecords(params GetRecordsParams) ([]Record, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	recordIter := dbSession.Query(`SELECT tableID,record_id,field_values FROM records WHERE tableID=?`,
		params.TableID).Iter()

	var currRecord Record
	encodedFieldVals := ""
	records := []Record{}
	for recordIter.Scan(&currRecord.ParentTableID,
		&currRecord.RecordID,
		&encodedFieldVals) {

		if decodeErr := generic.DecodeJSONString(encodedFieldVals, &currRecord.FieldValues); decodeErr != nil {
			return nil, fmt.Errorf("GetRecord: Unabled to get record: id = %v: datastore err=%v", currRecord.RecordID, decodeErr)
		}
		records = append(records, currRecord)
	}
	if closeErr := recordIter.Close(); closeErr != nil {
		return nil, fmt.Errorf("GetRecords: Failure querying database: %v", closeErr)
	}

	return records, nil
}

// Validate the field is of the correct type and not a calculated field (if allowCalcField not true). This is for validating
// the field when setting/getting values from regular "literal" fields which store values entered by end-users (as opposed to
// calculated fields)
func ValidateFieldForRecordValue(fieldParentTableID string, fieldID string, expectedFieldType string,
	allowCalcField bool) error {

	field, fieldGetErr := field.GetField(fieldParentTableID, fieldID)
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
