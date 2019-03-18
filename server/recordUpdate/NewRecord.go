package recordUpdate

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/alert"
	"resultra/tracker/server/record"
	"resultra/tracker/server/recordValue"
	"resultra/tracker/server/recordValueMappingController"
)

func newRecord(trackerDBHandle *sql.DB, currUserID string, params record.NewRecordParams) (*recordValue.RecordValueResults, error) {

	// Perform the low-level datastore record creation
	newRecord, newErr := record.NewRecord(trackerDBHandle, params)
	if newErr != nil {
		return nil, fmt.Errorf("newRecord: Can't create record: err = %v", newErr)
	}

	recCellUpdates := record.NewRecordCellUpdates(newRecord.RecordID)

	// Create an initial set of mapped record values. The mapped values are in the format
	// needed by clients to record creation. Although no values have been set yet, some of the
	// calculated fields may also have fixed values which don't depend on any values being set
	// in the record.
	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToLatestFieldValues(
		trackerDBHandle, currUserID, newRecord.ParentDatabaseID, recCellUpdates, record.FullyCommittedCellUpdatesChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"newRecord: Error mapping field values: err = %v", mapErr)
	}

	// Force a recalculation of results the next time results are loaded.
	recordValue.ResultsCache.Remove(params.ParentDatabaseID)
	alert.RemoveTrackerDatabaseCacheEntries(params.ParentDatabaseID)

	return updateRecordValResult, nil

}
