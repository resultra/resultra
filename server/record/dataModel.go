package record

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

type Record struct {
	ParentTableID string `json:"parentTableID"`
	RecordID      string `json:"recordID"`
}

type NewRecordParams struct {
	ParentTableID string `json:"parentTableID"`
}

func NewRecord(params NewRecordParams) (*Record, error) {

	newRecord := Record{ParentTableID: params.ParentTableID,
		RecordID: uniqueID.GenerateSnowflakeID()}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO records (table_id, record_id) VALUES ($1,$2)`, newRecord.ParentTableID, newRecord.RecordID); insertErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", insertErr)
	}

	return &newRecord, nil

}

func GetRecord(parentTableID string, recordID string) (*Record, error) {

	getRecord := Record{}

	if getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT table_id,record_id FROM records WHERE table_id=$1 AND record_id=$2 LIMIT 1`, parentTableID, recordID).Scan(&getRecord.ParentTableID,
		&getRecord.RecordID); getErr != nil {
		return nil, fmt.Errorf("GetRecord: Unabled to get record: id = %v: datastore err=%v", recordID, getErr)
	}

	return &getRecord, nil

}

type GetRecordsParams struct {
	TableID string `json:"tableID"`
}

func GetRecords(params GetRecordsParams) ([]Record, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT table_id,record_id FROM records WHERE table_id=$1`, params.TableID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecords: Failure querying database: %v", queryErr)
	}
	records := []Record{}
	for rows.Next() {
		var currRecord Record
		if scanErr := rows.Scan(&currRecord.ParentTableID,
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
