package recordUpdate

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordValue"
	"resultra/datasheet/server/recordValueMappingController"
)

func updateRecordValue(req *http.Request, recUpdater record.RecordUpdater) (*recordValue.RecordValueResults, error) {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get current user: err = %v", userErr)
	}

	// Perform the low-level datastore write of the record value update.
	recordForUpdate, writeErr := record.UpdateRecordValue(currUserID, recUpdater)
	if writeErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set record value: err = %v", writeErr)
	}

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(recordForUpdate.RecordID, recUpdater.GetChangeSetID())
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	// Since a change has occored to one of the record's values, a new set of mapped record
	// values needs to be created.
	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToFieldValues(
		recordForUpdate.ParentDatabaseID, recCellUpdates, recUpdater.GetChangeSetID())
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	return updateRecordValResult, nil

}

type CommitChangeSetParams struct {
	RecordID    string `json:"recordID"`
	ChangeSetID string `json:"changeSetID"`
}

func commitChangeSet(params CommitChangeSetParams) (*recordValue.RecordValueResults, error) {

	commitRecord, err := record.GetRecord(params.RecordID)
	if err != nil {
		return nil, fmt.Errorf("commitChangeSet: error getting record: %v", err)
	}

	if commitErr := record.CommitChangeSet(params.RecordID, params.ChangeSetID); commitErr != nil {
		return nil, fmt.Errorf("commitChangeSet: error committing changes : %v", commitErr)
	}

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(params.RecordID, record.FullyCommittedCellUpdatesChangeSetID)
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	// Temporary changes made under the given changeSetID have been made permanent, so
	// a new set of mapped record values needs to be created.
	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToFieldValues(
		commitRecord.ParentDatabaseID, recCellUpdates, record.FullyCommittedCellUpdatesChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	return updateRecordValResult, nil

}

func setDefaultValues(req *http.Request, params record.SetDefaultValsParams) (*recordValue.RecordValueResults, error) {

	currUserID, err := userAuth.GetCurrentUserID(req)
	if err != nil {
		return nil, fmt.Errorf("setDefaultValues: %v", err)
	}

	if setDefaultErr := record.SetDefaultValues(currUserID, params); setDefaultErr != nil {
		return nil, fmt.Errorf("setDefaultValues: %v", setDefaultErr)
	}

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(params.RecordID, params.ChangeSetID)
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToFieldValues(
		params.ParentDatabaseID, recCellUpdates, params.ChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	return updateRecordValResult, nil

}
