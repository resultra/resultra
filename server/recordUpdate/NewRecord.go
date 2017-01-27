package recordUpdate

import (
	"fmt"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordValue"
)

func newRecord(params record.NewRecordParams) (*recordValue.RecordValueResults, error) {

	// Perform the low-level datastore record creation
	newRecord, newErr := record.NewRecord(params)
	if newErr != nil {
		return nil, fmt.Errorf("newRecord: Can't create record: err = %v", newErr)
	}

	// Create an initial set of mapped record values. The mapped values are in the format
	// needed by clients to record creation. Although no values have been set yet, some of the
	// calculated fields may also have fixed values which don't depend on any values being set
	// in the record.
	updateRecordValResult, mapErr := recordValue.MapOneRecordUpdatesToFieldValues(
		newRecord.ParentDatabaseID, newRecord.RecordID, record.FullyCommittedCellUpdatesChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"newRecord: Error mapping field values: err = %v", mapErr)
	}

	return updateRecordValResult, nil

}
