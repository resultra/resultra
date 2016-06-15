package record

import (
	"fmt"
	"resultra/datasheet/server/generic/cassandraWrapper"
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

	// TODO - Verify all IDs are well-formed and cell value is non empty.

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if insertErr := dbSession.Query(
		`INSERT INTO cell_updates (table_id, record_id, field_id, update_timestamp_utc,value) 
			 VALUES (?,?,?,toTimestamp(now()),?)`,
		cellUpdate.ParentTableID,
		cellUpdate.RecordID,
		cellUpdate.FieldID,
		cellUpdate.CellValue).Exec(); insertErr != nil {
		return fmt.Errorf("saveCellUpdate: Can't save cell update: error = %v", insertErr)
	}

	return nil

}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordCellUpdates(parentTableID string, recordID string) ([]CellUpdate, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("GetCellUpdates: Can't get cell updaets: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	// TODO - Verify parentTableID and recordID are well-formed IDs

	cellUpdateIter := dbSession.Query(`SELECT field_id, update_timestamp_utc, value
			FROM cell_updates
			WHERE table_id=? and record_id=?`,
		parentTableID, recordID).Iter()

	var currCellUpdate CellUpdate
	cellUpdates := []CellUpdate{}
	for cellUpdateIter.Scan(&currCellUpdate.FieldID,
		&currCellUpdate.UpdateTimeStamp,
		&currCellUpdate.CellValue) {

		currCellUpdate.ParentTableID = parentTableID
		currCellUpdate.RecordID = recordID

		cellUpdates = append(cellUpdates, currCellUpdate)

	}
	if closeErr := cellUpdateIter.Close(); closeErr != nil {
		return nil, fmt.Errorf("GetCellUpdates: Failure querying database: %v", closeErr)
	}

	return cellUpdates, nil
}
