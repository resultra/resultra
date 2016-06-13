package recordUpdate

import (
	"fmt"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

type CellUpdate struct {
	ParentTableID string
	RecordID      string
	FieldID       string
	// userID string TODO: Add user ID
	CellValue string // Value encoded as JSON
}

func newCellUpdate(parentTableID string, fieldID string, recordID string, cellValue string) CellUpdate {
	return CellUpdate{
		ParentTableID: parentTableID,
		FieldID:       fieldID,
		RecordID:      recordID,
		CellValue:     cellValue}
}

func saveCellUpdate(cellUpdate CellUpdate) error {

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
