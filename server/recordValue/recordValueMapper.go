package recordValue

import (
	"fmt"
	"resultra/datasheet/server/record"
)

func MapOneRecordUpdatesToFieldValues(parentTableID string, recordID string) (*RecordValueResults, error) {
	return nil, fmt.Errorf("MapRecordUpdatesToFieldValues: Not implemented")

	//	cellUpdateFieldValIndex, indexErr := NewUpdateFieldValueIndex(parentTableID, recordID)
	_, indexErr := record.NewUpdateFieldValueIndex(parentTableID, recordID)
	if indexErr != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: error mapping updates to field values for "+
			" parent table = %v, record = %v: error = %v", parentTableID, recordID, indexErr)
	}

	//	for fieldID := range *cellUpdateFieldValIndex {
	//	for _ := range *cellUpdateFieldValIndex {

	//	}

	// TODO - Iterate through all the fields with value updates and populate the RecordValues structure with
	// the latest values for non-calculated field. These latest values represent the "current value" for a given field.
	// This current value is the basis for calculated fields.

	// TODO - Now that all the non-calculated fields have been populated into RecordValues, all the calculated
	// fields also need to be populated. The formulas for calculated field by refer to the latest value of non-calculated
	// fields.
	//if calcErr := calcField.UpdateCalcFieldValues(appEngContext, recordForUpdate); calcErr != nil {
	//	return nil, fmt.Errorf("updateRecordValue: Can't set value: Error calculating fields to reflect update: err = %v", calcErr)
	//}

	// TODO - Write the complete RecordValue to the datastore. The intent is to keep the RecordValue results up to date
	// in the datastore.  When individual records change, a single RecordValue can be updated. However, if a calculated field
	// formula changes, a new field is added, or some other change impacting all the records, all the RecordValue's must
	// be updated.

	// TODO - Return the individual RecordValue. Although the RecordValue has been written to the datastore, the updated
	// RecordValue may need to be returned immediately to the client in cases where an individual record has been updated
	// and the results need to be displayed to the user immediately.
	return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: not implemented yet")
}
