package values

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
)

const valGroupByNone string = "none"
const valGroupByDay string = "day"
const valGroupByBucket string = "bucket"

// ValGrouping represents a grouping of field values for purposes of summarizing
// in bar charts, lines charts, pie charts, and summary tables.
type ValGrouping struct {

	// XAxisField is the field used to group values along the x axis of the bar chart.
	GroupValsByFieldID string `json:"groupValsByFieldID"`

	// GroupValsBy configures how values from GroupValsByField are grouped.
	// Especially for date and number fields, the values will typically be grouped (bucketed), rather
	// than shown in their raw/ungrouped format.
	//
	// Depending on the data type of the field, different options are
	// available to group the values, including:
	//
	// Number: none, bucket
	// Date: none, hour, day, week, month, quarter, year
	// Text: none
	// Bool: none
	GroupValsBy string `json:"groupValsBy"`

	// GroupByValBucketWidth is used with the GroupValsBy "bucket" property to configure a threshold for
	// grouping values.
	GroupByValBucketWidth float64 `json:"groupValsByBucketWidth"`
}

type NewValGroupingParams struct {
	FieldParentTableID    string  `json:"fieldParentTableID"`
	FieldID               string  `json:"fieldID"`
	GroupValsBy           string  `json:"groupValsBy"`
	GroupByValBucketWidth float64 `json:"groupByValBucketWidth"`
}

func validateFieldTypeWithGrouping(fieldType string, groupValsBy string, bucketWidth float64) error {
	switch groupValsBy {
	case valGroupByNone:
		return nil
	case valGroupByBucket:
		if fieldType != field.FieldTypeNumber {
			return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
		}
		if bucketWidth <= 0.0 {
			return fmt.Errorf("Invalid grouping = %v for field type = %v, bucket width must be > 0.0", groupValsBy, fieldType)
		}
	case valGroupByDay:
		if fieldType != field.FieldTypeTime {
			return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
		}
	default:
		return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
	} // switch groupValsBy
	return nil
}

func NewValGrouping(params NewValGroupingParams) (*ValGrouping, error) {

	groupingField, fieldErr := field.GetField(params.FieldParentTableID, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Can't create value grouping with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if groupByErr := validateFieldTypeWithGrouping(groupingField.Type, params.GroupValsBy,
		params.GroupByValBucketWidth); groupByErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Invalid value grouping: %v", groupByErr)
	}

	valGrouping := ValGrouping{params.FieldID, params.GroupValsBy, params.GroupByValBucketWidth}

	return &valGrouping, nil

}

type ValGroup struct {
	GroupLabel     string
	RecordsInGroup []record.Record
}

type ValGroupingRecordVal struct {
	groupLabel string
}

func recordGroupLabel(fieldGroup field.Field, rec record.Record) (string, error) {
	switch fieldGroup.Type {
	case field.FieldTypeText:
		if rec.ValueIsSet(fieldGroup.FieldID) {
			textVal, valErr := rec.GetTextFieldValue(fieldGroup.FieldID)
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
		return "All Dates", nil // TODO - Group by date and/or bucket the values by day, month, etc.
	}
	return "", fmt.Errorf("recordGroupLabel: unsupported grouping: fieldRef = %+v", fieldGroup)
}

type ValGroupingResult struct {
	ValGroups     []ValGroup
	GroupingLabel string
}

func (valGrouping ValGrouping) GroupRecords(parentFieldID string, records []record.Record) (*ValGroupingResult, error) {

	groupingField, fieldErr := field.GetField(parentFieldID, valGrouping.GroupValsByFieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("groupRecords: Can't get field to group records: error = %v", fieldErr)
	}

	// Use a map to group the values. Values are added to the same GroupVal if they have the same
	// group label.
	groupLabelValGroupMap := map[string]*ValGroup{}
	for _, currRecord := range records {
		groupLabel, lblErr := recordGroupLabel(*groupingField, currRecord)
		if lblErr != nil {
			return nil, fmt.Errorf("groupRecords: Error getting label to group records: error = %v", lblErr)
		}
		_, groupExists := groupLabelValGroupMap[groupLabel]
		if !groupExists {
			groupLabelValGroupMap[groupLabel] = &ValGroup{groupLabel, []record.Record{}}
		}
		valGroup := groupLabelValGroupMap[groupLabel]
		valGroup.RecordsInGroup = append(valGroup.RecordsInGroup, currRecord)
	}

	// Flatten the group values into an array
	var valGroups []ValGroup
	for _, currValGroup := range groupLabelValGroupMap {
		valGroups = append(valGroups, *currValGroup)
	}

	return &ValGroupingResult{valGroups, groupingField.Name}, nil
}
