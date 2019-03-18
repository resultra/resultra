package record

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic/timestamp"
	"resultra/tracker/server/generic/uniqueID"
)

type RecordUpdater interface {
	fieldType() string
	fieldID() string
	recordID() string
	parentDatabaseID() string
	generateCellValue() (string, error)
	GetChangeSetID() string
	doCollapseRecentValues() bool
}

// RecordUpdateHeader is a common header for all record value updates. It also implements
// part of the RecorddUpdater interface. This struct should be embedded in other structs
// used to update values of specific types.
type RecordUpdateHeader struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	RecordID         string `json:"recordID"`
	FieldID          string `json:"fieldID"`
	ChangeSetID      string `json:"changeSetID"`
}

func (recUpdateHeader RecordUpdateHeader) fieldID() string {
	return recUpdateHeader.FieldID
}

func (recUpdateHeader RecordUpdateHeader) recordID() string {
	return recUpdateHeader.RecordID
}

func (recUpdateHeader RecordUpdateHeader) parentDatabaseID() string {
	return recUpdateHeader.ParentDatabaseID
}

func (recUpdateHeader RecordUpdateHeader) GetChangeSetID() string {
	return recUpdateHeader.ChangeSetID
}

// updateRecordValue implements a generic algorithm (strategy design pattern) which wrapp the updating of records.
// It leaves the low-level updating of values to implementers of the RecordUpdater interface. Different RecordUpdaters
// are needed for different value types, while the code to (1) retrieve the record, (2) validate the field type,
// (3) re-calculate calculated fields, then (4) save the updated record is made common.
func UpdateRecordValue(trackingDBHandle *sql.DB, currUserID string, recUpdater RecordUpdater) (*Record, error) {

	recordID := recUpdater.recordID()

	field, fieldGetErr := field.GetField(trackingDBHandle, recUpdater.fieldID())
	if fieldGetErr != nil {
		return nil, fmt.Errorf("UpdateRecordValue: Error retrieving field for updating/setting value: err = %v", fieldGetErr)
	}

	if fieldValidateErr := ValidateFieldForRecordValue(
		*field, recUpdater.fieldType(), false); fieldValidateErr != nil {
		return nil, fmt.Errorf("UpdateRecordValue: Can't set record value:"+
			" Error validating record's field for update: %v", fieldValidateErr)
	}

	recordForUpdate, getErr := GetRecord(trackingDBHandle, recordID)
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

	uniqueUpdateID := uniqueID.GenerateUniqueID()
	updateTimestamp := timestamp.CurrentTimestampUTC()

	cellUpdate := CellUpdate{
		UpdateID:         uniqueUpdateID,
		UserID:           currUserID,
		ParentDatabaseID: recUpdater.parentDatabaseID(),
		FieldID:          recUpdater.fieldID(),
		RecordID:         recUpdater.recordID(),
		CellValue:        cellValue,
		UpdateTimeStamp:  updateTimestamp,
		ChangeSetID:      recUpdater.GetChangeSetID()}

	if saveErr := SaveCellUpdate(trackingDBHandle, cellUpdate, recUpdater.doCollapseRecentValues()); saveErr != nil {
		return nil, fmt.Errorf("UpdateRecordValue: Error saving cell update: %v", saveErr)
	}

	// Return the updated record
	// TODO - Depending upon how calculated values are implemented,
	// this is where calculated field values may also be updated.
	return recordForUpdate, nil

}
