// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package values

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic/uniqueID"
)

const ValGroupByNone string = "none"
const ValGroupByDay string = "day"
const ValGroupByWeek string = "week"
const ValGroupByMonthYear string = "monthYear"
const ValGroupByBucket string = "bucket"

const ValGroupIntervalEndOfDay string = "days"
const ValGroupIntervalIntervalEndOfWeek string = "weeks"
const ValGroupIntervalEndOfMonth string = "months"

// ValGrouping represents a grouping of field values for purposes of summarizing
// in bar charts, lines charts, pie charts, and summary tables.
type ValGrouping struct {

	// The field or time incremnent used to group values.
	// One of the following needs to be set.
	GroupValsByFieldID       *string `json:"groupValsByFieldID,omitempty"`
	GroupValsByTimeIncrement *string `json:"groupValsByTimeIncrement,omitempty"`

	// Time range when the GroupValsByTimeIncrement is set.
	TimeRange *string `json:"timeRange,omitempty"`

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
	GroupValsBy *string `json:"groupValsBy,omitempty"`

	// GroupByValBucketWidth is used with the GroupValsBy "bucket" property to configure a threshold for
	// grouping values.
	GroupByValBucketWidth *float64 `json:"groupValsByBucketWidth,omitempty"`

	BucketStart *float64 `json:"bucketStart,omitempty"`
	BucketEnd   *float64 `json:"bucketEnd,omitempty"`

	NumberFormat *string `json:"numberFormat,omitempty"`

	IncludeBlank bool `json:"includeBlank"`
}

func (srcGrouping ValGrouping) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ValGrouping, error) {
	destGrouping := srcGrouping

	if srcGrouping.GroupValsByFieldID != nil {
		remappedFieldID, err := remappedIDs.GetExistingRemappedID(*srcGrouping.GroupValsByFieldID)
		if err != nil {
			return nil, fmt.Errorf("ValGrouping.Clone: %v", err)
		}
		destGrouping.GroupValsByFieldID = &remappedFieldID
	}

	return &destGrouping, nil
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
	case ValGroupByDay, ValGroupByMonthYear, ValGroupByWeek:
		if fieldType != field.FieldTypeTime {
			return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
		}
	default:
		return fmt.Errorf("Invalid grouping = %v for field type = %v", groupValsBy, fieldType)
	} // switch groupValsBy
	return nil
}

func ValidateValGrouping(trackingDBHandle *sql.DB, valGrouping ValGrouping) error {

	if valGrouping.GroupValsByFieldID != nil {
		groupingField, fieldErr := field.GetField(trackingDBHandle, *valGrouping.GroupValsByFieldID)
		if fieldErr != nil {
			return fmt.Errorf("NewValGrouping: Can't create value grouping with field ID = '%v': datastore error=%v",
				*valGrouping.GroupValsByFieldID, fieldErr)
		}

		if valGrouping.GroupValsBy == nil {
			return fmt.Errorf("NewValGrouping: Can't create value grouping with field ID = '%v', missing grouping",
				*valGrouping.GroupValsByFieldID)
		}

		if groupByErr := validateFieldTypeWithGrouping(groupingField.Type, *valGrouping.GroupValsBy,
			valGrouping.GroupByValBucketWidth, valGrouping.BucketStart, valGrouping.BucketEnd); groupByErr != nil {
			return fmt.Errorf("NewValGrouping: Invalid value grouping: %v", groupByErr)
		}
	}

	return nil

}

func (valGrouping ValGrouping) GroupingLabel(trackingDBHandle *sql.DB) (string, error) {

	if valGrouping.GroupValsByFieldID != nil {

		groupingField, fieldErr := field.GetField(trackingDBHandle, *valGrouping.GroupValsByFieldID)
		if fieldErr != nil {
			return "", fmt.Errorf("GroupingLabel: Can't create grouping label: %v", fieldErr)
		}

		if valGrouping.GroupValsBy == nil {
			return "", fmt.Errorf("GroupingLabel: Can't create grouping label: missing grouping")
		}
		groupValsBy := *valGrouping.GroupValsBy

		switch groupValsBy {
		case ValGroupByNone:
			return groupingField.Name, nil
		case ValGroupByBucket:
			return groupingField.Name, nil
		case ValGroupByDay:
			return "Date", nil
		case ValGroupByMonthYear:
			return "Month and Year", nil
		case ValGroupByWeek:
			return "Week starting", nil
		default:
			return "", fmt.Errorf("GroupingLabel: unsupported grouping type: %v", valGrouping.GroupValsBy)
		} // switch groupValsBy
	} else {

		if valGrouping.GroupValsByTimeIncrement == nil {
			return "", fmt.Errorf("GroupingLabel: missing time increment")
		}
		timeInterval := *valGrouping.GroupValsByTimeIncrement

		switch timeInterval {
		case ValGroupIntervalEndOfDay:
			return "End of Day", nil
		case ValGroupIntervalIntervalEndOfWeek:
			return "End of Week", nil
		case ValGroupIntervalEndOfMonth:
			return "End of Month", nil
		default:
			return "", fmt.Errorf("GroupingLabel: unsupported time increment: %v", timeInterval)
		}
	}

}
