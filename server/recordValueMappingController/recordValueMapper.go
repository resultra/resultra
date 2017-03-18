package recordValueMappingController

import (
	"fmt"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordValue"
)

// Re-map the series of value updates to "flattened" current (most recent) values for both calculated
// and non-calculated fields.
func MapOneRecordUpdatesToFieldValues(parentDatabaseID string, recordID string, changeSetID string) (*recordValue.RecordValueResults, error) {

	cellUpdateFieldValIndex, indexErr := record.NewUpdateFieldValueIndex(parentDatabaseID, recordID, changeSetID)
	if indexErr != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", indexErr)
	}

	// For non-calculated fields, get the latest (most recent) field values.
	latestFieldValues := cellUpdateFieldValIndex.LatestNonCalcFieldValues()

	// Now that all the non-calculated fields have been populated into latestFieldValues, all the calculated
	// fields also need to be populated. The formulas for calculated field by refer to the latest value of non-calculated
	// fields, so this set of values needs to be passed into UpdateCalcFieldValues as a starting point.
	if calcErr := calcField.UpdateCalcFieldValues(parentDatabaseID, latestFieldValues); calcErr != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: Can't set value: Error calculating fields to reflect update: err = %v", calcErr)
	}

	recValResults := recordValue.RecordValueResults{
		ParentDatabaseID: parentDatabaseID,
		RecordID:         recordID,
		FieldValues:      *latestFieldValues}

	// Write the complete RecordValue to the datastore. The intent is to keep the RecordValue results up to date
	// in the datastore.  When individual records change, a single RecordValue can be updated. However, if a calculated field
	// formula changes, a new field is added, or some other change impacting all the records, all the RecordValue's must
	// be updated.
	if changeSetID == record.FullyCommittedCellUpdatesChangeSetID {
		// Only permanently save mapped results values if the changeSetID is for a fully committed set of values. Otherwise,
		// the values being mapped are for a temporary set of changes being made in a modal dialog and subject to being
		// cancelled.
		if saveErr := recordValue.SaveRecordValueResults(recValResults); saveErr != nil {
			return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: Error saving mapped record value results: err = %v", saveErr)
		}

	}

	return &recValResults, nil
}
