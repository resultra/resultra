package record

import (
	"fmt"
	"github.com/gocql/gocql"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

type RecFieldValues map[string]interface{}

type Record struct {
	ParentTableID string `json:"parentTableID"`
	RecordID      string `json:"recordID"`
}

type NewRecordParams struct {
	ParentTableID string `json:"parentTableID"`
}

func NewRecord(params NewRecordParams) (*Record, error) {

	newRecord := Record{ParentTableID: params.ParentTableID,
		RecordID: gocql.TimeUUID().String()}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if insertErr := dbSession.Query(`INSERT INTO records (table_id, record_id) VALUES (?,?)`,
		newRecord.ParentTableID, newRecord.RecordID).Exec(); insertErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", insertErr)
	}

	return &newRecord, nil

}

func GetRecord(parentTableID string, recordID string) (*Record, error) {

	getRecord := Record{}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("GetRecord: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	getErr := dbSession.Query(`SELECT table_id,record_id FROM records WHERE table_id=? AND record_id=? LIMIT 1`,
		parentTableID, recordID).Scan(&getRecord.ParentTableID,
		&getRecord.RecordID)
	if getErr != nil {
		return nil, fmt.Errorf("GetRecord: Unabled to get record: id = %v: datastore err=%v", recordID, getErr)
	}

	return &getRecord, nil

}

type GetRecordsParams struct {
	TableID string `json:"tableID"`
}

func GetRecords(params GetRecordsParams) ([]Record, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	recordIter := dbSession.Query(`SELECT table_id,record_id FROM records WHERE table_id=?`,
		params.TableID).Iter()

	var currRecord Record
	records := []Record{}
	for recordIter.Scan(&currRecord.ParentTableID,
		&currRecord.RecordID) {

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
