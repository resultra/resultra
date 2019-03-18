// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package recordUpdate

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/tracker/server/alert"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/record"
	"resultra/tracker/server/recordValue"
	"resultra/tracker/server/recordValueMappingController"
	"resultra/tracker/server/userRole"
)

func updateRecordValue(req *http.Request, recUpdater record.RecordUpdater) (*recordValue.RecordValueResults, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get current user: err = %v", userErr)
	}

	// Perform the low-level datastore write of the record value update.
	recordForUpdate, writeErr := record.UpdateRecordValue(trackerDBHandle, currUserID, recUpdater)
	if writeErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set record value: err = %v", writeErr)
	}

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(trackerDBHandle, recordForUpdate.RecordID, recUpdater.GetChangeSetID())
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	// Since a change has occored to one of the record's values, a new set of mapped record
	// values needs to be created.
	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToLatestFieldValues(
		trackerDBHandle, currUserID, recordForUpdate.ParentDatabaseID, recCellUpdates, recUpdater.GetChangeSetID())
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	userIsAdmin := userRole.CurrUserIsDatabaseAdmin(req, recordForUpdate.ParentDatabaseID)
	// (re)generate any alerts which may have been triggered by the current update
	alert.GenerateOneRecordAlerts(trackerDBHandle, currUserID, recordForUpdate.ParentDatabaseID,
		recordForUpdate.RecordID, currUserID, userIsAdmin)

	// Force a recalculation of results the next time results are loaded.
	recordValue.ResultsCache.Remove(recordForUpdate.ParentDatabaseID)
	alert.RemoveTrackerDatabaseCacheEntries(recordForUpdate.ParentDatabaseID)

	return updateRecordValResult, nil

}

type CommitChangeSetParams struct {
	RecordID    string `json:"recordID"`
	ChangeSetID string `json:"changeSetID"`
}

func commitChangeSet(trackerDBHandle *sql.DB, currUserID string, params CommitChangeSetParams) (*recordValue.RecordValueResults, error) {

	commitRecord, err := record.GetRecord(trackerDBHandle, params.RecordID)
	if err != nil {
		return nil, fmt.Errorf("commitChangeSet: error getting record: %v", err)
	}

	if commitErr := record.CommitChangeSet(trackerDBHandle, params.RecordID, params.ChangeSetID); commitErr != nil {
		return nil, fmt.Errorf("commitChangeSet: error committing changes : %v", commitErr)
	}

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(trackerDBHandle, params.RecordID, record.FullyCommittedCellUpdatesChangeSetID)
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	// Temporary changes made under the given changeSetID have been made permanent, so
	// a new set of mapped record values needs to be created.
	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToLatestFieldValues(
		trackerDBHandle, currUserID, commitRecord.ParentDatabaseID, recCellUpdates, record.FullyCommittedCellUpdatesChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	// Force a recalculation of results the next time results are loaded.
	recordValue.ResultsCache.Remove(commitRecord.ParentDatabaseID)
	alert.RemoveTrackerDatabaseCacheEntries(commitRecord.ParentDatabaseID)

	return updateRecordValResult, nil

}

func setDefaultValues(req *http.Request, params record.SetDefaultValsParams) (*recordValue.RecordValueResults, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	currUserID, err := userAuth.GetCurrentUserID(req)
	if err != nil {
		return nil, fmt.Errorf("setDefaultValues: %v", err)
	}

	if setDefaultErr := record.SetDefaultValues(trackerDBHandle, currUserID, params); setDefaultErr != nil {
		return nil, fmt.Errorf("setDefaultValues: %v", setDefaultErr)
	}

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(trackerDBHandle, params.RecordID, params.ChangeSetID)
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToLatestFieldValues(
		trackerDBHandle, currUserID, params.ParentDatabaseID, recCellUpdates, params.ChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	// Force a recalculation of results the next time results are loaded.
	recordValue.ResultsCache.Remove(params.ParentDatabaseID)
	alert.RemoveTrackerDatabaseCacheEntries(params.ParentDatabaseID)

	return updateRecordValResult, nil

}
