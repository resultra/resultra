package record

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"time"
)

type CellUpdateValueFormat struct {
	Context string `json:"context"`
	Format  string `json:"format"`
}

type CellUpdateProperties struct {
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
}

func newDefaultCellUpdateProperties() CellUpdateProperties {
	valFormat := CellUpdateValueFormat{"", ""}
	props := CellUpdateProperties{valFormat}
	return props
}

type CellUpdate struct {
	UpdateID         string
	ParentDatabaseID string
	RecordID         string
	FieldID          string
	UserID           string
	UpdateTimeStamp  time.Time
	CellValue        string               // Value encoded as JSON
	Properties       CellUpdateProperties // Properties encoded as JSON
}

func SaveCellUpdate(cellUpdate CellUpdate) error {

	encodedProps, err := generic.EncodeJSONString(cellUpdate.Properties)
	if err != nil {
		return fmt.Errorf("SaveCellUpdate: Error encoding properties: %v", err)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO cell_updates (update_id, user_id, database_id, record_id, field_id,update_timestamp_utc,value,properties) 
			 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		cellUpdate.UpdateID,
		cellUpdate.UserID,
		cellUpdate.ParentDatabaseID,
		cellUpdate.RecordID,
		cellUpdate.FieldID,
		cellUpdate.UpdateTimeStamp,
		cellUpdate.CellValue,
		encodedProps); insertErr != nil {
		return fmt.Errorf("saveCellUpdate: insert failed: update = %+v, error = %v", cellUpdate, insertErr)
	}
	return nil

}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordCellUpdates(parentDatabaseID string, recordID string) ([]CellUpdate, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT update_id, user_id, database_id, record_id, field_id, update_timestamp_utc, value,properties
			FROM cell_updates
			WHERE database_id=$1 and record_id=$2`,
		parentDatabaseID, recordID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordCellUpdates: Failure querying database: %v", queryErr)
	}
	cellUpdates := []CellUpdate{}
	for rows.Next() {
		var currCellUpdate CellUpdate
		var encodedProps string
		if scanErr := rows.Scan(
			&currCellUpdate.UpdateID,
			&currCellUpdate.UserID,
			&currCellUpdate.ParentDatabaseID,
			&currCellUpdate.RecordID,
			&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue,
			&encodedProps); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)

		}
		cellUpdateProps := newDefaultCellUpdateProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &cellUpdateProps); decodeErr != nil {
			return nil, fmt.Errorf("GetRecordCellUpdates: can't decode properties: %v", encodedProps)
		}
		currCellUpdate.Properties = cellUpdateProps

		cellUpdates = append(cellUpdates, currCellUpdate)
	}

	return cellUpdates, nil
}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordFieldCellUpdates(recordID string, fieldID string) ([]CellUpdate, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT update_id,user_id,database_id, record_id,field_id, update_timestamp_utc, value,properties
			FROM cell_updates
			WHERE record_id=$2 and field_id=$3`,
		recordID, fieldID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordCellUpdates: Failure querying database: %v", queryErr)
	}
	cellUpdates := []CellUpdate{}
	for rows.Next() {
		var currCellUpdate CellUpdate
		var encodedProps string
		if scanErr := rows.Scan(
			&currCellUpdate.UpdateID,
			&currCellUpdate.UserID,
			&currCellUpdate.ParentDatabaseID,
			&currCellUpdate.RecordID,
			&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue,
			&encodedProps); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)

		}
		cellUpdateProps := newDefaultCellUpdateProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &cellUpdateProps); decodeErr != nil {
			return nil, fmt.Errorf("GetRecordCellUpdates: can't decode properties: %v", encodedProps)
		}
		currCellUpdate.Properties = cellUpdateProps

		cellUpdates = append(cellUpdates, currCellUpdate)
	}

	return cellUpdates, nil
}
