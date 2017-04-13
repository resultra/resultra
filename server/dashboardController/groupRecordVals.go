package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/recordValue"
)

type ValGroup struct {
	GroupLabel     string
	RecordsInGroup []recordValue.RecordValueResults
}

type ValGroupingRecordVal struct {
	groupLabel string
}

type ValGroupingResult struct {
	ValGroups     []ValGroup
	GroupingLabel string
}

func groupRecords(valGrouping values.ValGrouping,
	recValResults []recordValue.RecordValueResults) (*ValGroupingResult, error) {

	groupingField, fieldErr := field.GetField(valGrouping.GroupValsByFieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("groupRecords: Can't get field to group records: error = %v", fieldErr)
	}

	// Use a map to group the values. Values are added to the same GroupVal if they have the same
	// group label.
	groupLabelValGroupMap := map[string]*ValGroup{}
	for _, currRecValResults := range recValResults {
		groupLabel, lblErr := recordGroupLabel(valGrouping, *groupingField, currRecValResults)
		if lblErr != nil {
			return nil, fmt.Errorf("groupRecords: Error getting label to group records: error = %v", lblErr)
		}
		_, groupExists := groupLabelValGroupMap[groupLabel]
		if !groupExists {
			groupLabelValGroupMap[groupLabel] = &ValGroup{groupLabel, []recordValue.RecordValueResults{}}
		}
		valGroup := groupLabelValGroupMap[groupLabel]
		valGroup.RecordsInGroup = append(valGroup.RecordsInGroup, currRecValResults)
	}

	// Flatten the group values into an array
	var valGroups []ValGroup
	for _, currValGroup := range groupLabelValGroupMap {
		valGroups = append(valGroups, *currValGroup)
	}

	groupingLabel, groupingLabelErr := valGrouping.GroupingLabel()
	if groupingLabelErr != nil {
		return nil, fmt.Errorf("groupRecords: Error getting grouping label: error = %v", groupingLabelErr)
	}

	return &ValGroupingResult{ValGroups: valGroups,
		GroupingLabel: groupingLabel}, nil
}

func groupTimeFieldRecordVal(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (string, error) {

	if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
		timeVal, valFound := recValResults.FieldValues.GetTimeFieldValue(fieldGroup.FieldID)
		if !valFound {
			return "", fmt.Errorf("groupTimeFieldRecordVal: Unabled to retrieve value for grouping label")
		} else {
			switch valGrouping.GroupValsBy {
			case values.ValGroupByNone:
				return "All Dates", nil
			case values.ValGroupByDay:
				return timeVal.Format("2006-01-02"), nil
			case values.ValGroupByMonthYear:
				return timeVal.Format("Jan 2006"), nil
			default:
				return "", fmt.Errorf("Invalid grouping = %v for time field type", valGrouping.GroupValsBy)
			} // switch groupValsBy
		}
	} else {
		return "BLANK", nil
	}
}

func recordGroupLabel(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (string, error) {
	switch fieldGroup.Type {
	case field.FieldTypeText:
		if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
			textVal, valErr := recValResults.FieldValues.GetTextFieldValue(fieldGroup.FieldID)
			if valErr != nil {
				return "", fmt.Errorf("recordGroupLabel: Unabled to retrieve value for grouping label: error = %v", valErr)
			} else {
				return textVal, nil
			}
		} else {
			return "BLANK", nil
		}
	case field.FieldTypeNumber:
		return "All Numbers", nil // TODO - Group by number and/or bucket the values
	case field.FieldTypeTime:
		return groupTimeFieldRecordVal(valGrouping, fieldGroup, recValResults)
	}
	return "", fmt.Errorf("recordGroupLabel: unsupported grouping: fieldRef = %+v", fieldGroup)
}
