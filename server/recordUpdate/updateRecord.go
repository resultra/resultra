package recordUpdate

import (
	"fmt"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordValue"
)

func updateRecordValue(recUpdater record.RecordUpdater) (*recordValue.RecordValueResults, error) {

	// Perform the low-level datastore write of the record value update.
	recordForUpdate, writeErr := record.UpdateRecordValue(recUpdater)
	if writeErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set record value: err = %v", writeErr)
	}

	// Since a change has occored to one of the record's values, a new set of mapped record
	// values needs to be created.
	updateRecordValResult, mapErr := recordValue.MapOneRecordUpdatesToFieldValues(
		recordForUpdate.ParentTableID, recordForUpdate.RecordID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	return updateRecordValResult, nil

}
