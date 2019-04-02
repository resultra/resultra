// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package recordUpdate

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/alert"
	"github.com/resultra/resultra/server/record"
	"github.com/resultra/resultra/server/recordValue"
	"github.com/resultra/resultra/server/recordValueMappingController"
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
