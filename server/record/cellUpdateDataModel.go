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
	ChangeSetID      string
	CellValue        string               // Value encoded as JSON
	Properties       CellUpdateProperties // Properties encoded as JSON
}

const FullyCommittedCellUpdatesChangeSetID string = ""

func SaveCellUpdate(cellUpdate CellUpdate) error {

	encodedProps, err := generic.EncodeJSONString(cellUpdate.Properties)
	if err != nil {
		return fmt.Errorf("SaveCellUpdate: Error encoding properties: %v", err)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO cell_updates (update_id, user_id, database_id, record_id, field_id,update_timestamp_utc,value,properties,change_set_id) 
			 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		cellUpdate.UpdateID,
		cellUpdate.UserID,
		cellUpdate.ParentDatabaseID,
		cellUpdate.RecordID,
		cellUpdate.FieldID,
		cellUpdate.UpdateTimeStamp,
		cellUpdate.CellValue,
		encodedProps,
		cellUpdate.ChangeSetID); insertErr != nil {
		return fmt.Errorf("saveCellUpdate: insert failed: update = %+v, error = %v", cellUpdate, insertErr)
	}
	return nil

}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordCellUpdates(recordID string, changeSetID string) ([]CellUpdate, error) {

	selectFields := `SELECT update_id,user_id,database_id,record_id,field_id,update_timestamp_utc,value,properties
							FROM cell_updates`
	matchRecordQuery := ` record_id = $1 `

	// Build up the query depending on whether or not the changeSetID is empty or not.
	changeSetIDMatch := ""
	changeIDQuery := ` (change_set_id is null OR change_set_id = $2)`
	if len(changeSetID) > 0 {
		// match a specific change_set_id
		// TODO - Get the cell updates for changes without a changeSetID (i.e., baseline/main values)
		// and with the given changeSetID
		changeIDQuery = ` (change_set_id is null OR change_set_id = '' OR change_set_id = $2)`
		changeSetIDMatch = changeSetID
	}

	cellUpdatesQuery := selectFields + ` WHERE ` + matchRecordQuery + ` AND ` + changeIDQuery

	rows, queryErr := databaseWrapper.DBHandle().Query(cellUpdatesQuery, recordID, changeSetIDMatch)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordCellUpdates: Failure querying database for record ID = %v: %v", recordID, queryErr)
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
			return nil, fmt.Errorf("GetRecordCellUpdates: Failure scanning database row: %v", scanErr)

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
func GetRecordFieldCellUpdates(recordID string, fieldID string, changeSetID string) ([]CellUpdate, error) {

	selectFields := `SELECT update_id,user_id,database_id, record_id,field_id, update_timestamp_utc, value,properties
			FROM cell_updates`
	matchRecordAndFieldQuery := ` record_id = $1 AND field_id = $2 `

	// Build up the query depending on whether or not the changeSetID is empty or not.
	changeSetIDMatch := ""
	changeIDQuery := ` (change_set_id is null or change_set_id = $3)`
	if len(changeSetID) > 0 {
		// match a specific change_set_id or value without a changeSetID (i.e., main/baseline values)
		changeIDQuery = `(change_set_id is null OR change_set_id = '' OR change_set_id = $3)`
		changeSetIDMatch = changeSetID
	}

	cellUpdatesQuery := selectFields + ` WHERE ` + matchRecordAndFieldQuery + ` AND ` + changeIDQuery

	rows, queryErr := databaseWrapper.DBHandle().Query(cellUpdatesQuery, recordID, fieldID, changeSetIDMatch)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecordFieldCellUpdates: Failure querying database: %v", queryErr)
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
			return nil, fmt.Errorf("GetRecordFieldCellUpdates: Failure querying database: %v", scanErr)

		}
		cellUpdateProps := newDefaultCellUpdateProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &cellUpdateProps); decodeErr != nil {
			return nil, fmt.Errorf("GetRecordFieldCellUpdates: can't decode properties: %v", encodedProps)
		}
		currCellUpdate.Properties = cellUpdateProps

		cellUpdates = append(cellUpdates, currCellUpdate)
	}

	return cellUpdates, nil
}

func CommitChangeSet(recordID string, changeSetID string) error {

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE cell_updates 
				SET change_set_id=$1
				WHERE record_id=$2 AND change_set_id=$3`,
		FullyCommittedCellUpdatesChangeSetID, recordID, changeSetID); updateErr != nil {
		return fmt.Errorf("CommitChangeSet: Can't commit changes to record=%v, change set=%v: error = %v",
			recordID, changeSetID, updateErr)
	}

	return nil

}
