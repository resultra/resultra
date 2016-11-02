package record

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

type CellUpdate struct {
	UpdateID        string
	ParentTableID   string
	RecordID        string
	FieldID         string
	UserID          string
	UpdateTimeStamp time.Time
	CellValue       string // Value encoded as JSON
}

func NewCellUpdate(userID string, parentTableID string, fieldID string, recordID string, cellValue string) CellUpdate {
	return CellUpdate{
		UserID:        userID,
		ParentTableID: parentTableID,
		FieldID:       fieldID,
		RecordID:      recordID,
		CellValue:     cellValue}
}

func SaveCellUpdate(cellUpdate CellUpdate) error {

	uniqueUpdateID := uniqueID.GenerateSnowflakeID()

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO cell_updates (update_id, user_id, table_id, record_id, field_id,update_timestamp_utc,value) 
			 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		uniqueUpdateID,
		cellUpdate.UserID,
		cellUpdate.ParentTableID,
		cellUpdate.RecordID,
		cellUpdate.FieldID,
		time.Now().UTC(),
		cellUpdate.CellValue); insertErr != nil {
		return fmt.Errorf("saveCellUpdate: insert failed: update = %+v, error = %v", cellUpdate, insertErr)
	}
	return nil

}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordCellUpdates(parentTableID string, recordID string) ([]CellUpdate, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT update_id, user_id, table_id, record_id, field_id, update_timestamp_utc, value
			FROM cell_updates
			WHERE table_id=$1 and record_id=$2`,
		parentTableID, recordID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordCellUpdates: Failure querying database: %v", queryErr)
	}
	cellUpdates := []CellUpdate{}
	for rows.Next() {
		var currCellUpdate CellUpdate
		if scanErr := rows.Scan(
			&currCellUpdate.UpdateID,
			&currCellUpdate.UserID,
			&currCellUpdate.ParentTableID,
			&currCellUpdate.RecordID,
			&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)

		}
		cellUpdates = append(cellUpdates, currCellUpdate)
	}

	return cellUpdates, nil
}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordFieldCellUpdates(parentTableID string, recordID string, fieldID string) ([]CellUpdate, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT update_id,user_id,table_id, record_id,field_id, update_timestamp_utc, value
			FROM cell_updates
			WHERE table_id=$1 and record_id=$2 and field_id=$3`,
		parentTableID, recordID, fieldID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordCellUpdates: Failure querying database: %v", queryErr)
	}
	cellUpdates := []CellUpdate{}
	for rows.Next() {
		var currCellUpdate CellUpdate
		if scanErr := rows.Scan(
			&currCellUpdate.UpdateID,
			&currCellUpdate.UserID,
			&currCellUpdate.ParentTableID,
			&currCellUpdate.RecordID,
			&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)

		}
		cellUpdates = append(cellUpdates, currCellUpdate)
	}

	return cellUpdates, nil
}
