package recordSort

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/recordValue"
)

func SortRecordValues(trackerDBHandle *sql.DB, parentDatabaseID string,
	recordVals []recordValue.RecordValueResults,
	sortRules []recordSortDataModel.RecordSortRule) error {

	if len(sortRules) == 0 {
		return nil // no sorting necessary
	}

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(trackerDBHandle,
		field.GetFieldListParams{ParentDatabaseID: parentDatabaseID})
	if indexErr != nil {
		return fmt.Errorf("SortRecordValues: %v", indexErr)
	}

	sortFuncs := []ByRecordValueLessFunc{}
	for _, currSortRule := range sortRules {

		if !recordSortDataModel.ValidSortDirection(currSortRule.SortDirection) {
			return fmt.Errorf("SortRecordValues: invalid sort direction for sort rule = %+v", currSortRule)
		}

		sortField, fieldErr := fieldRefIndex.GetFieldRefByID(currSortRule.SortFieldID)
		if fieldErr != nil {
			return fmt.Errorf("SortRecordValues: invalid sort field for sort rule = %+v: err = %v",
				currSortRule, fieldErr)
		}

		switch sortField.Type {
		case field.FieldTypeText:
			sortFuncs = append(sortFuncs, SortByTextField(currSortRule.SortFieldID, currSortRule.SortDirection))
		case field.FieldTypeNumber:
			sortFuncs = append(sortFuncs, SortByNumberField(currSortRule.SortFieldID, currSortRule.SortDirection))
		case field.FieldTypeTime:
			sortFuncs = append(sortFuncs, SortByTimeField(currSortRule.SortFieldID, currSortRule.SortDirection))
		case field.FieldTypeBool:
			sortFuncs = append(sortFuncs, SortByBoolField(currSortRule.SortFieldID, currSortRule.SortDirection))
		case field.FieldTypeUser:
			// TODO - Implement sort function for users
			return fmt.Errorf("Sorting by user not yet implemented")
			//			sortFuncs = append(sortFuncs, SortByUserField(currSortRule.SortFieldID, currSortRule.SortDirection))
		default:
			return fmt.Errorf("SortRecordValues: invalid sort field type %v for sort rule = %+v",
				sortField.Type, sortField)

		}
		OrderedBy(sortFuncs...).Sort(recordVals)

	}
	return nil

}
