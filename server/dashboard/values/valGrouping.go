package values

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
)

const ValGroupByNone string = "none"
const ValGroupByDay string = "day"
const ValGroupByMonthYear string = "monthYear"
const ValGroupByBucket string = "bucket"

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
	GroupByValBucketWidth *float64 `json:"groupValsByBucketWidth,omitempty"`

	BucketStart *float64 `json:"bucketStart,omitempty"`
	BucketEnd   *float64 `json:"bucketEnd,omitempty"`

	NumberFormat *string `json:"numberFormat,omitempty"`
}

func (srcGrouping ValGrouping) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ValGrouping, error) {
	destGrouping := srcGrouping

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcGrouping.GroupValsByFieldID)
	if err != nil {
		return nil, fmt.Errorf("ValGrouping.Clone: %v", err)
	}
	destGrouping.GroupValsByFieldID = remappedFieldID

	return &destGrouping, nil
}

type NewValGroupingParams struct {
	FieldID               string   `json:"fieldID"`
	GroupValsBy           string   `json:"groupValsBy"`
	GroupByValBucketWidth *float64 `json:"groupByValBucketWidth,omitempty"`
	BucketStart           *float64 `json:"bucketStart,omitempty"`
	BucketEnd             *float64 `json:"bucketEnd,omitempty"`
}

func validateFieldTypeWithGrouping(fieldType string, groupValsBy string,
	bucketWidth *float64, bucketStart *float64, bucketEnd *float64) error {
	switch groupValsBy {
	case ValGroupByNone:
		return nil
	case ValGroupByBucket:
		if fieldType != field.FieldTypeNumber {
			return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
		}
		if bucketWidth == nil {
			return fmt.Errorf("Invalid grouping = %v for field type = %v, bucket width missing", groupValsBy, fieldType)
		}
		if *bucketWidth <= 0.0 {
			return fmt.Errorf("Invalid grouping = %v for field type = %v, bucket width must be > 0.0", groupValsBy, fieldType)
		}
		if bucketStart != nil && bucketEnd != nil {
			if *bucketEnd < *bucketStart {
				return fmt.Errorf("Invalid grouping = %v for field type = %v, bucket end must be greater than bucket start",
					groupValsBy, fieldType)

			}
		}
	case ValGroupByDay, ValGroupByMonthYear:
		if fieldType != field.FieldTypeTime {
			return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
		}
	default:
		return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
	} // switch groupValsBy
	return nil
}

func NewValGrouping(params NewValGroupingParams) (*ValGrouping, error) {

	groupingField, fieldErr := field.GetField(params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Can't create value grouping with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if groupByErr := validateFieldTypeWithGrouping(groupingField.Type, params.GroupValsBy,
		params.GroupByValBucketWidth, params.BucketStart, params.BucketEnd); groupByErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Invalid value grouping: %v", groupByErr)
	}

	valGrouping := ValGrouping{
		GroupValsByFieldID:    params.FieldID,
		GroupValsBy:           params.GroupValsBy,
		GroupByValBucketWidth: params.GroupByValBucketWidth,
		BucketStart:           params.BucketStart,
		BucketEnd:             params.BucketEnd}

	return &valGrouping, nil

}

func (valGrouping ValGrouping) GroupingLabel() (string, error) {

	groupingField, fieldErr := field.GetField(valGrouping.GroupValsByFieldID)
	if fieldErr != nil {
		return "", fmt.Errorf("GroupingLabel: Can't create grouping label: %v", fieldErr)
	}

	switch valGrouping.GroupValsBy {
	case ValGroupByNone:
		return groupingField.Name, nil
	case ValGroupByBucket:
		return "TBD", nil
	case ValGroupByDay:
		return "Date", nil
	case ValGroupByMonthYear:
		return "Month and Year", nil
	default:
		return "", fmt.Errorf("GroupingLabel: unsupported grouping type: %v", valGrouping.GroupValsBy)
	} // switch groupValsBy

}
