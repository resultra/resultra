package dashboard

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"resultra/datasheet/datamodel"
)

const valGroupByNone string = "none"
const valGroupByDay string = "day"
const valGroupByBucket string = "bucket"

// ValGrouping represents a grouping of field values for purposes of summarizing
// in bar charts, lines charts, pie charts, and summary tables.
type ValGrouping struct {

	// XAxisField is the field used to group values along the x axis of the bar chart.
	GroupValsByField *datastore.Key

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
	GroupValsBy string

	// GroupByValBucketWidth is used with the GroupValsBy "bucket" property to configure a threshold for
	// grouping values.
	GroupByValBucketWidth float64
}

// DashboardValGroupingRef is an opaque reference to the dashboard value grouping parameters.
type ValGroupingRef struct {
	GroupValsByFieldRef   datamodel.FieldRef `json:"groupValsByFieldRef"`
	GroupValsBy           string             `json:"groupValsBy"`
	GroupByValBucketWidth float64            `json:"groupByValBucketWidth"`
}

type NewValGroupingParams struct {
	FieldID               string  `json:"fieldID"`
	GroupValsBy           string  `json:"groupValsBy"`
	GroupByValBucketWidth float64 `json:"groupByValBucketWidth"`
}

func (valGroupingRef ValGroupingRef) toNewValGroupingParams() NewValGroupingParams {
	return NewValGroupingParams{
		FieldID:               valGroupingRef.GroupValsByFieldRef.FieldID,
		GroupValsBy:           valGroupingRef.GroupValsBy,
		GroupByValBucketWidth: valGroupingRef.GroupByValBucketWidth}
}

func validateFieldTypeWithGrouping(fieldType string, groupValsBy string, bucketWidth float64) error {
	switch groupValsBy {
	case valGroupByNone:
		return nil
	case valGroupByBucket:
		if fieldType != datamodel.FieldTypeNumber {
			return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
		}
		if bucketWidth <= 0.0 {
			return fmt.Errorf("Invalid grouping = %v for field type = %v, bucket width must be > 0.0", groupValsBy, fieldType)
		}
	case valGroupByDay:
		if fieldType != datamodel.FieldTypeDate {
			return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
		}
	default:
		return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
	} // switch groupValsBy
	return nil
}

func NewValGrouping(appEngContext appengine.Context, params NewValGroupingParams) (*ValGrouping, *ValGroupingRef, error) {

	fieldKey, fieldRef, fieldErr := datamodel.GetExistingFieldRefAndKey(appEngContext, datamodel.GetFieldParams{params.FieldID})
	if fieldErr != nil {
		return nil, nil, fmt.Errorf("NewValGrouping: Can't get field value grouping: datastore error = %v", fieldErr)
	}

	if groupByErr := validateFieldTypeWithGrouping(fieldRef.FieldInfo.Type, params.GroupValsBy,
		params.GroupByValBucketWidth); groupByErr != nil {
		return nil, nil, fmt.Errorf("NewValGrouping: Invalid value grouping: %v", groupByErr)
	}

	valGroupingRef := ValGroupingRef{*fieldRef, params.GroupValsBy, params.GroupByValBucketWidth}
	valGrouping := ValGrouping{fieldKey, params.GroupValsBy, params.GroupByValBucketWidth}

	return &valGrouping, &valGroupingRef, nil

}

func NewValGroupingFromRef(appEngContext appengine.Context, groupingRef ValGroupingRef) (*ValGrouping, *ValGroupingRef, error) {
	newValGroupingParams := groupingRef.toNewValGroupingParams()
	return NewValGrouping(appEngContext, newValGroupingParams)
}

func (valGrouping ValGrouping) GetValGroupingRef(appEngContext appengine.Context) (*ValGroupingRef, error) {

	fieldRef, fieldErr := datamodel.GetFieldFromKey(appEngContext, valGrouping.GroupValsByField)
	if fieldErr != nil {
		return nil, fmt.Errorf("GetValGroupingRef: Can't get field  for value grouping: datastore error = %v", fieldErr)
	}

	valGroupingRef := ValGroupingRef{*fieldRef, valGrouping.GroupValsBy, valGrouping.GroupByValBucketWidth}

	return &valGroupingRef, nil

}

type ValGroup struct {
	groupLabel     string
	recordsInGroup []datamodel.RecordRef
}

type ValGroupingRecordVal struct {
	groupLabel string
}

func recordGroupLabel(fieldGroupRef datamodel.FieldRef, recordRef datamodel.RecordRef) (string, error) {
	switch fieldGroupRef.FieldInfo.Type {
	case datamodel.FieldTypeText:
		if recordRef.FieldValues.ValueIsSet(fieldGroupRef.FieldID) {
			textVal, valErr := recordRef.FieldValues.GetTextFieldValue(fieldGroupRef.FieldID)
			if valErr != nil {
				return "", fmt.Errorf("recordGroupLabel: Unabled to retrieve value for grouping label: error = %v", valErr)
			} else {
				return textVal, nil
			}
		} else {
			return "BLANK", nil
		}
	case datamodel.FieldTypeNumber:
		return "All Numbers", nil // TODO - Group by number and/or bucket the values
	case datamodel.FieldTypeDate:
		return "All Dates", nil // TODO - Group by date and/or bucket the values by day, month, etc.
	}
	return "", fmt.Errorf("recordGroupLabel: unsupported grouping: fieldRef = %+v", fieldGroupRef)
}

type ValGroupingResult struct {
	valGroups     []ValGroup
	groupingLabel string
}

func (valGrouping ValGrouping) groupRecords(appEngContext appengine.Context, recordRefs []datamodel.RecordRef) (*ValGroupingResult, error) {
	fieldRef, fieldErr := datamodel.GetFieldFromKey(appEngContext, valGrouping.GroupValsByField)
	if fieldErr != nil {
		return nil, fmt.Errorf("groupRecords: Can't get field to group records: error = %v", fieldErr)
	}

	// Use a map to group the values. Values are added to the same GroupVal if they have the same
	// group label.
	groupLabelValGroupMap := map[string]*ValGroup{}
	for _, currRecordRef := range recordRefs {
		groupLabel, lblErr := recordGroupLabel(*fieldRef, currRecordRef)
		if lblErr != nil {
			return nil, fmt.Errorf("groupRecords: Error getting label to group records: error = %v", lblErr)
		}
		_, groupExists := groupLabelValGroupMap[groupLabel]
		if !groupExists {
			groupLabelValGroupMap[groupLabel] = &ValGroup{groupLabel, []datamodel.RecordRef{}}
		}
		valGroup := groupLabelValGroupMap[groupLabel]
		valGroup.recordsInGroup = append(valGroup.recordsInGroup, currRecordRef)
	}

	// Flatten the group values into an array
	var valGroups []ValGroup
	for _, currValGroup := range groupLabelValGroupMap {
		valGroups = append(valGroups, *currValGroup)
	}

	return &ValGroupingResult{valGroups, fieldRef.FieldInfo.Name}, nil
}
