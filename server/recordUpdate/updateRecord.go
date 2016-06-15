package recordUpdate

import (
	"fmt"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/record"
)

func updateRecordValue(recUpdater record.RecordUpdater) (*record.Record, error) {

	// Perform the low-level datastore write of the record value update.
	recordForUpdate, writeErr := record.UpdateRecordValue(recUpdater)
	if writeErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set record value: err = %v", writeErr)
	}

	// Changing this value may have caused the values for calculated fields to also change.
	// Clients of this function need a fully up to date record reference, so the calculated
	// field's values must also be recalculated when a value changes.
	if calcErr := calcField.UpdateCalcFieldValues(recordForUpdate.ParentTableID, &recordForUpdate.FieldValues); calcErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set value: Error calculating fields to reflect update: err = %v", calcErr)
	}

	// Write the record back to the datastore (including the calculated values)
	updatedRecord, updateErr := record.UpdateExistingRecord(recordForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set value: Error retrieving existing record for update: err = %v", updateErr)
	}

	return updatedRecord, nil

}
