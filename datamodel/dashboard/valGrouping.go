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
