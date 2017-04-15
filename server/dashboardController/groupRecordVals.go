package dashboardController

import (
	"fmt"
	"math"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/recordValue"
	"sort"
	"strings"
	"time"
)

type IntermediateValGroup struct {
	GroupLabelInfo valGroupLabelInfo
	RecordsInGroup []recordValue.RecordValueResults
}

type ByGroupLabelSortOrder []IntermediateValGroup

func (s ByGroupLabelSortOrder) Less(j, k int) bool {
	jInfo := s[j].GroupLabelInfo
	kInfo := s[k].GroupLabelInfo

	if (jInfo.textSortVal != nil) && (kInfo.textSortVal != nil) {
		compareVal := strings.Compare(*jInfo.textSortVal, *kInfo.textSortVal)
		if compareVal > 0 {
			return false
		} else {
			return true
		}
	} else if (jInfo.numSortVal != nil) && (kInfo.numSortVal != nil) {
		return (*jInfo.numSortVal) < (*kInfo.numSortVal)
	} else {
		return false
	}
}

func (s ByGroupLabelSortOrder) Len() int {
	return len(s)
}

func (s ByGroupLabelSortOrder) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ValGroup struct {
	GroupLabel     string
	RecordsInGroup []recordValue.RecordValueResults
}

type ValGroupingRecordVal struct {
	groupLabel string
}

type ValGroupingResult struct {
	ValGroups     []ValGroup
	OverallGroup  ValGroup
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
	groupLabelValGroupMap := map[string]*IntermediateValGroup{}
	for _, currRecValResults := range recValResults {
		groupLabelInfo, lblErr := recordGroupLabelInfo(valGrouping, *groupingField, currRecValResults)
		if lblErr != nil {
			return nil, fmt.Errorf("groupRecords: Error getting label to group records: error = %v", lblErr)
		}
		_, groupExists := groupLabelValGroupMap[groupLabelInfo.label]
		if !groupExists {
			groupLabelValGroupMap[groupLabelInfo.label] = &IntermediateValGroup{*groupLabelInfo, []recordValue.RecordValueResults{}}
		}
		valGroup := groupLabelValGroupMap[groupLabelInfo.label]
		valGroup.RecordsInGroup = append(valGroup.RecordsInGroup, currRecValResults)
	}

	// Flatten the intermediate value groups into an array
	intermValGroups := []IntermediateValGroup{}
	for _, currValGroup := range groupLabelValGroupMap {
		intermValGroups = append(intermValGroups, *currValGroup)
	}

	// Sort the intermediate value groups
	sort.Sort(ByGroupLabelSortOrder(intermValGroups))

	// Flatten the intermediate group values into an array finalized ValGroup(s)
	var valGroups []ValGroup
	for _, currValGroup := range intermValGroups {
		valGroup := ValGroup{currValGroup.GroupLabelInfo.label, currValGroup.RecordsInGroup}
		valGroups = append(valGroups, valGroup)
	}

	groupingLabel, groupingLabelErr := valGrouping.GroupingLabel()
	if groupingLabelErr != nil {
		return nil, fmt.Errorf("groupRecords: Error getting grouping label: error = %v", groupingLabelErr)
	}

	overallGroup := ValGroup{
		GroupLabel:     "Overall",
		RecordsInGroup: recValResults}

	return &ValGroupingResult{
		ValGroups:     valGroups,
		GroupingLabel: groupingLabel,
		OverallGroup:  overallGroup}, nil
}

// valGroupLabelInfo is used as intermediate data to produce a "normalized" sort value
// for use on the grouping column of data results.
type valGroupLabelInfo struct {
	label       string
	textSortVal *string
	numSortVal  *float64
}

func textGroupLabelInfo(label string) *valGroupLabelInfo {
	sortVal := label
	return &valGroupLabelInfo{
		label:       label,
		textSortVal: &sortVal}
}

func blankTextGroupLabelInfo(label string) *valGroupLabelInfo {
	sortVal := ""
	return &valGroupLabelInfo{
		label:       label,
		textSortVal: &sortVal}
}

func numberGroupLabelInfo(label string, numberSortVal float64) *valGroupLabelInfo {
	sortVal := numberSortVal
	return &valGroupLabelInfo{
		label:      label,
		numSortVal: &sortVal}
}

func timeGroupLabelInfo(label string, timeVal time.Time) *valGroupLabelInfo {
	sortVal := float64(timeVal.UnixNano())
	return &valGroupLabelInfo{
		label:      label,
		numSortVal: &sortVal}
}

func groupTimeFieldRecordVal(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (*valGroupLabelInfo, error) {

	if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
		timeVal, valFound := recValResults.FieldValues.GetTimeFieldValue(fieldGroup.FieldID)
		if !valFound {
			return nil, fmt.Errorf("groupTimeFieldRecordVal: Unabled to retrieve value for grouping label")
		} else {
			switch valGrouping.GroupValsBy {
			case values.ValGroupByNone:
				return numberGroupLabelInfo("All Dates", 0.0), nil
			case values.ValGroupByDay:
				return timeGroupLabelInfo(timeVal.Format("2006-01-02"), timeVal), nil
			case values.ValGroupByMonthYear:
				return timeGroupLabelInfo(timeVal.Format("Jan 2006"), timeVal), nil
			default:
				return nil, fmt.Errorf("Invalid grouping = %v for time field type", valGrouping.GroupValsBy)
			} // switch groupValsBy
		}
	} else {
		return numberGroupLabelInfo("BLANK", -1.0*math.MaxFloat64), nil
	}
}

func groupBoolFieldRecordVal(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (*valGroupLabelInfo, error) {

	if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
		boolVal, valFound := recValResults.FieldValues.GetBoolFieldValue(fieldGroup.FieldID)
		if !valFound {
			return nil, fmt.Errorf("groupBoolFieldRecordVal: Unabled to retrieve value for grouping label")
		} else {
			if boolVal == true {
				return numberGroupLabelInfo("True", 1.0), nil
			} else {
				return numberGroupLabelInfo("False", 0.0), nil
			}
		}
	} else {
		return numberGroupLabelInfo("BLANK", -1.0), nil
	}
}

func recordGroupLabelInfo(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (*valGroupLabelInfo, error) {
	switch fieldGroup.Type {
	case field.FieldTypeText:
		if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
			textVal, valErr := recValResults.FieldValues.GetTextFieldValue(fieldGroup.FieldID)
			if valErr != nil {
				return nil, fmt.Errorf("recordGroupLabel: Unabled to retrieve value for grouping label: error = %v", valErr)
			} else {
				return textGroupLabelInfo(textVal), nil
			}
		} else {
			return blankTextGroupLabelInfo("BLANK"), nil
		}
	case field.FieldTypeBool:
		return groupBoolFieldRecordVal(valGrouping, fieldGroup, recValResults)
	case field.FieldTypeNumber:
		return numberGroupLabelInfo("All Numbers", 0.0), nil // TODO - Group by number and/or bucket the values
	case field.FieldTypeTime:
		return groupTimeFieldRecordVal(valGrouping, fieldGroup, recValResults)
	}
	return nil, fmt.Errorf("recordGroupLabel: unsupported grouping: fieldRef = %+v", fieldGroup)
}
