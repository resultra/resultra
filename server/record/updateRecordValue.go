package record

import (
	"fmt"
)

type RecordUpdater interface {
	fieldType() string
	fieldID() string
	recordID() string
	parentTableID() string
	generateCellValue() (string, error)
}

// RecordUpdateHeader is a common header for all record value updates. It also implements
// part of the RecorddUpdater interface. This struct should be embedded in other structs
// used to update values of specific types.
type RecordUpdateHeader struct {
	ParentTableID string `json:"parentTableID"`
	RecordID      string `json:"recordID"`
	FieldID       string `json:"fieldID"`
}

func (recUpdateHeader RecordUpdateHeader) fieldID() string {
	return recUpdateHeader.FieldID
}

func (recUpdateHeader RecordUpdateHeader) recordID() string {
	return recUpdateHeader.RecordID
}

func (recUpdateHeader RecordUpdateHeader) parentTableID() string {
	return recUpdateHeader.ParentTableID
}

// updateRecordValue implements a generic algorithm (strategy design pattern) which wrapp the updating of records.
// It leaves the low-level updating of values to implementers of the RecordUpdater interface. Different RecordUpdaters
// are needed for different value types, while the code to (1) retrieve the record, (2) validate the field type,
// (3) re-calculate calculated fields, then (4) save the updated record is made common.
func UpdateRecordValue(recUpdater RecordUpdater) (*Record, error) {

	recordID := recUpdater.recordID()

	if fieldValidateErr := ValidateFieldForRecordValue(recUpdater.parentTableID(),
		recUpdater.fieldID(), recUpdater.fieldType(), false); fieldValidateErr != nil {
		return nil, fmt.Errorf("UpdateRecordValue: Can't set record value:"+
			" Error validating record's field for update: %v", fieldValidateErr)
	}

	recordForUpdate, getErr := GetRecord(recUpdater.parentTableID(), recordID)
	if getErr != nil {
		return nil, fmt.Errorf("UpdateRecordValue: Can't set value:"+
			" Error retrieving existing record for update: err = %v", getErr)
	}

	// Changes to records are stored as a series of updates, which are then rolled up into a simpler
	// structure which has all the calculated values and filter results.
	cellValue, cellErr := recUpdater.generateCellValue()
	if cellErr != nil {
		return nil, fmt.Errorf("UpdateRecordValue: Error generating value for cell update: %v", cellErr)
	}
	cellUpdate := NewCellUpdate(recUpdater.parentTableID(), recUpdater.fieldID(), recUpdater.recordID(), cellValue)
	if saveErr := SaveCellUpdate(cellUpdate); saveErr != nil {
		return nil, fmt.Errorf("UpdateRecordValue: Error saving cell update: %v", saveErr)
	}

	// Return the updated record
	// TODO - Depending upon how calculated values are implemented,
	// this is where calculated field values may also be updated.
	return recordForUpdate, nil

}
