package record

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"time"
)

type CellUpdate struct {
	UpdateID         string
	ParentDatabaseID string
	RecordID         string
	FieldID          string
	UserID           string
	UpdateTimeStamp  time.Time
	ChangeSetID      string
	CellValue        string // Value encoded as JSON
}

// Custom sort function for the FieldValueUpate
type CellUpdateByUpdateTime []CellUpdate

func (s CellUpdateByUpdateTime) Len() int {
	return len(s)
}
func (s CellUpdateByUpdateTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort in  chronological order; i.e. the most recent dates come first.
func (s CellUpdateByUpdateTime) Less(i, j int) bool {
	return s[i].UpdateTimeStamp.Before(s[j].UpdateTimeStamp)
}

const FullyCommittedCellUpdatesChangeSetID string = ""

func SaveCellUpdate(cellUpdate CellUpdate, doCollapseRecentValues bool) error {

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO cell_updates (update_id, user_id, database_id, record_id, field_id,update_timestamp_utc,value,change_set_id) 
			 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		cellUpdate.UpdateID,
		cellUpdate.UserID,
		cellUpdate.ParentDatabaseID,
		cellUpdate.RecordID,
		cellUpdate.FieldID,
		cellUpdate.UpdateTimeStamp,
		cellUpdate.CellValue,
		cellUpdate.ChangeSetID); insertErr != nil {
		return fmt.Errorf("saveCellUpdate: insert failed: update = %+v, error = %v", cellUpdate, insertErr)
	}

	if doCollapseRecentValues {
		// Delete any cell updates which were committed within the last minute. This is intended to collapse the editing
		// of values, so there isn't superfulous updates. For example, if a spinner is used to select a number, then
		// there is no need to keep all the intermediate values, only the last value.
		collapseUpdatesStart := cellUpdate.UpdateTimeStamp.Add(-1 * time.Minute)

		if _, deleteRecentErr := databaseWrapper.DBHandle().Exec(
			`DELETE FROM cell_updates
				 WHERE record_id=$1 AND field_id=$2 AND user_id=$3 AND change_set_id=$4
				 AND update_timestamp_utc<$5 AND update_timestamp_utc>$6 AND update_id<>$7`,
			cellUpdate.RecordID,
			cellUpdate.FieldID,
			cellUpdate.UserID,
			cellUpdate.ChangeSetID,
			cellUpdate.UpdateTimeStamp,
			collapseUpdatesStart,
			cellUpdate.UpdateID); deleteRecentErr != nil {
			return fmt.Errorf("saveCellUpdate: deletion of recent records failed: update = %+v, error = %v", cellUpdate, deleteRecentErr)
		}

	}

	return nil

}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordCellUpdates(recordID string, changeSetID string) (*RecordCellUpdates, error) {

	selectFields := `SELECT update_id,user_id,database_id,record_id,field_id,update_timestamp_utc,value
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
		if scanErr := rows.Scan(
			&currCellUpdate.UpdateID,
			&currCellUpdate.UserID,
			&currCellUpdate.ParentDatabaseID,
			&currCellUpdate.RecordID,
			&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue); scanErr != nil {
			return nil, fmt.Errorf("GetRecordCellUpdates: Failure scanning database row: %v", scanErr)

		}

		cellUpdates = append(cellUpdates, currCellUpdate)
	}

	recCellUpdates := RecordCellUpdates{
		RecordID:    recordID,
		CellUpdates: cellUpdates}

	return &recCellUpdates, nil
}

type RecordCellUpdates struct {
	RecordID    string
	CellUpdates []CellUpdate
}

func NewRecordCellUpdates(recordID string) *RecordCellUpdates {
	recCellUpdates := RecordCellUpdates{
		RecordID:    recordID,
		CellUpdates: []CellUpdate{}}
	return &recCellUpdates
}

type RecordCellUpdateMap map[string]*RecordCellUpdates

// Get all cell updates for records which have been fully committed. This excludes records for forms which were
// opened up in a form, but the results were never saved.
func GetAllNonDraftCellUpdates(databaseID string, changeSetID string) (RecordCellUpdateMap, error) {

	// Pre-populate the map of record ID to the cell updates structure. This ensures the structure
	// is populated, even for record IDs without any cell updates yet.
	records, err := GetNonDraftRecords(databaseID)
	if err != nil {
		return nil, fmt.Errorf("GetAllCellUpdates: %v", err)
	}
	recordCellUpdateMap := RecordCellUpdateMap{}
	for _, currRecord := range records {
		cellUpdates := RecordCellUpdates{RecordID: currRecord.RecordID,
			CellUpdates: []CellUpdate{}}
		recordCellUpdateMap[currRecord.RecordID] = &cellUpdates
	}

	selectFields := `SELECT 
			cell_updates.update_id,cell_updates.user_id,cell_updates.database_id,
			cell_updates.record_id,cell_updates.field_id,cell_updates.update_timestamp_utc,
			cell_updates.value
		FROM records, cell_updates `
	matchDatabaseQuery := ` records.is_draft_record=false AND cell_updates.database_id=$1 AND cell_updates.record_id=records.record_id  `

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

	cellUpdatesQuery := selectFields + ` WHERE ` + matchDatabaseQuery + ` AND ` + changeIDQuery

	rows, queryErr := databaseWrapper.DBHandle().Query(cellUpdatesQuery, databaseID, changeSetIDMatch)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllCellUpdates: Failure querying database = %v for cell updates: %v", databaseID, queryErr)
	}

	for rows.Next() {
		var currCellUpdate CellUpdate
		if scanErr := rows.Scan(
			&currCellUpdate.UpdateID,
			&currCellUpdate.UserID,
			&currCellUpdate.ParentDatabaseID,
			&currCellUpdate.RecordID,
			&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue); scanErr != nil {
			return nil, fmt.Errorf("GetRecordCellUpdates: Failure scanning database row: %v", scanErr)

		}

		var recCellUpdates *RecordCellUpdates
		recCellUpdates, found := recordCellUpdateMap[currCellUpdate.RecordID]
		if found {
			recCellUpdates.CellUpdates = append(recCellUpdates.CellUpdates, currCellUpdate)
		}
	}

	return recordCellUpdateMap, nil
}

// GetCellUpdates retrieves a list of cell updates for all the fields in the given record.
func GetRecordFieldCellUpdates(recordID string, fieldID string, changeSetID string) ([]CellUpdate, error) {

	selectFields := `SELECT update_id,user_id,database_id, record_id,field_id, update_timestamp_utc, value
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
		if scanErr := rows.Scan(
			&currCellUpdate.UpdateID,
			&currCellUpdate.UserID,
			&currCellUpdate.ParentDatabaseID,
			&currCellUpdate.RecordID,
			&currCellUpdate.FieldID,
			&currCellUpdate.UpdateTimeStamp,
			&currCellUpdate.CellValue); scanErr != nil {
			return nil, fmt.Errorf("GetRecordFieldCellUpdates: Failure querying database: %v", scanErr)

		}

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
