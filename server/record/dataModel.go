package record

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

type Record struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	RecordID         string `json:"recordID"`
}

type NewRecordParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func NewRecord(params NewRecordParams) (*Record, error) {

	newRecord := Record{ParentDatabaseID: params.ParentDatabaseID,
		RecordID: uniqueID.GenerateSnowflakeID()}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO records (database_id, record_id) VALUES ($1,$2)`, newRecord.ParentDatabaseID, newRecord.RecordID); insertErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", insertErr)
	}

	return &newRecord, nil

}

func GetRecord(recordID string) (*Record, error) {

	getRecord := Record{}

	if getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id,record_id FROM records WHERE record_id=$1 LIMIT 1`, recordID).Scan(
		&getRecord.ParentDatabaseID,
		&getRecord.RecordID); getErr != nil {
		return nil, fmt.Errorf("GetRecord: Unabled to get record: id = %v: datastore err=%v", recordID, getErr)
	}

	return &getRecord, nil

}

type GetRecordsParams struct {
	DatabaseID string `json:"databaseID"`
}

func GetRecords(params GetRecordsParams) ([]Record, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT database_id,record_id FROM records WHERE database_id=$1`,
		params.DatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecords: Failure querying database: %v", queryErr)
	}
	records := []Record{}
	for rows.Next() {
		var currRecord Record
		if scanErr := rows.Scan(&currRecord.ParentDatabaseID,
			&currRecord.RecordID); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)

		}
		records = append(records, currRecord)
	}

	return records, nil
}

// Validate the field is of the correct type and not a calculated field (if allowCalcField not true). This is for validating
// the field when setting/getting values from regular "literal" fields which store values entered by end-users (as opposed to
// calculated fields)
func ValidateFieldForRecordValue(fieldID string, expectedFieldType string,
	allowCalcField bool) error {

	field, fieldGetErr := field.GetField(fieldID)
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
