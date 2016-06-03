package recordUpdate

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/record"
)

type RecordUpdater interface {
	fieldType() string
	fieldID() string
	recordID() string
	updateRecordValue(rec *record.Record)
}

// RecordUpdateHeader is a common header for all record value updates. It also implements
// part of the RecorddUpdater interface. This struct should be embedded in other structs
// used to update values of specific types.
type RecordUpdateHeader struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
}

func (recUpdateHeader RecordUpdateHeader) fieldID() string {
	return recUpdateHeader.FieldID
}

func (recUpdateHeader RecordUpdateHeader) recordID() string {
	return recUpdateHeader.RecordID
}

// updateRecordValue implements a generic algorithm (strategy design pattern) which wrapp the updating of records.
// It leaves the low-level updating of values to implementers of the RecordUpdater interface. Different RecordUpdaters
// are needed for different value types, while the code to (1) retrieve the record, (2) validate the field type,
// (3) re-calculate calculated fields, then (4) save the updated record is made common.
func UpdateRecordValue(appEngContext appengine.Context, recUpdater RecordUpdater) (*record.Record, error) {

	recordID := recUpdater.recordID()

	if fieldValidateErr := record.ValidateFieldForRecordValue(appEngContext, recUpdater.fieldID(), recUpdater.fieldType(), false); fieldValidateErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set record value:"+
			" Error validating record's field for update: %v", fieldValidateErr)
	}

	recordForUpdate, getErr := record.GetRecord(appEngContext, recordID)
	if getErr != nil {
		return nil, fmt.Errorf("SetRecordTextValue: Can't set value:"+
			" Error retrieving existing record for update: err = %v", getErr)
	}

	// TODO - Consider putting a prefix on the field ID before saving it to the datastore.
	//	recordForUpdate.FieldValues[setValParams.FieldID] = setValParams.Value
	log.Printf("updateRecordValue: Setting value for field = %v", recUpdater.fieldID())
	recUpdater.updateRecordValue(recordForUpdate)

	// Changing this value may have caused the values for calculated fields to also change.
	// Clients of this function need a fully up to date record reference, so the calculated
	// field's values must also be recalculated when a value changes.
	if calcErr := calcField.UpdateCalcFieldValues(appEngContext, recordForUpdate); calcErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set value: Error calculating fields to reflect update: err = %v", calcErr)
	}

	// Write the record back to the datastore (including the calculated values)
	updatedRecord, updateErr := record.UpdateExistingRecord(appEngContext, recordID, recordForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't set value: Error retrieving existing record for update: err = %v", updateErr)
	}

	// Return the updated record
	// TODO - Depending upon how calculated values are implemented,
	// this is where calculated field values may also be updated.
	return updatedRecord, nil

}
