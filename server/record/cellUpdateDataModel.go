package record

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"time"
)

type CellUpdate struct {
	ParentTableID string
	RecordID      string
	FieldID       string
	// userID string TODO: Add user ID
	UpdateTimeStamp time.Time
	CellValue       string // Value encoded as JSON
}

func NewCellUpdate(parentTableID string, fieldID string, recordID string, cellValue string) CellUpdate {
	return CellUpdate{
		ParentTableID: parentTableID,
		FieldID:       fieldID,
		RecordID:      recordID,
		CellValue:     cellValue}
}

func SaveCellUpdate(cellUpdate CellUpdate) error {

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO cell_updates (table_id, record_id, field_id,update_timestamp_utc,value) 
			 VALUES ($1,$2,$3,$4,$5)`, cellUpdate.ParentTableID,
		cellUpdate.RecordID,
		cellUpdate.FieldID,
		time.Now().UTC(),
		cellUpdate.CellValue); insertErr != nil {
		return fmt.Errorf("saveCellUpdate: insert failed: error = %v", insertErr)
	}
	return nil

}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordCellUpdates(parentTableID string, recordID string) ([]CellUpdate, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT field_id, update_timestamp_utc, value
			FROM cell_updates
			WHERE table_id=$1 and record_id=$2`,
		parentTableID, recordID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordCellUpdates: Failure querying database: %v", queryErr)
	}
	cellUpdates := []CellUpdate{}
	for rows.Next() {
		var currCellUpdate CellUpdate
		if scanErr := rows.Scan(&currCellUpdate.FieldID,
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

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT field_id, update_timestamp_utc, value
			FROM cell_updates
			WHERE table_id=$1 and record_id=$2 and field_id=$3`,
		parentTableID, recordID, fieldID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordCellUpdates: Failure querying database: %v", queryErr)
	}
	cellUpdates := []CellUpdate{}
	for rows.Next() {
		var currCellUpdate CellUpdate
		if scanErr := rows.Scan(&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)

		}
		cellUpdates = append(cellUpdates, currCellUpdate)
	}

	return cellUpdates, nil
}
