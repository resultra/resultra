package dashboardController

import (
	"database/sql"
	"fmt"
	"math"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/numberFormat"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordValue"
	"resultra/datasheet/server/recordValueMappingController"
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

func groupRecordsIntoSingleGroup(recValResults []recordValue.RecordValueResults) *ValGroupingResult {

	overallGroup := ValGroup{
		GroupLabel:     "Overall",
		RecordsInGroup: recValResults}

	valGroups := []ValGroup{}
	valGroups = append(valGroups, overallGroup)

	return &ValGroupingResult{
		ValGroups:     valGroups,
		GroupingLabel: "Overall",
		OverallGroup:  overallGroup}

}

type GroupByTimeIntervalParams struct {
	trackerDBHandle *sql.DB
	databaseID      string
	currUserID      string
	preFilterRules  recordFilter.RecordFilterRuleSet
	filterRules     recordFilter.RecordFilterRuleSet
	valGrouping     values.ValGrouping
}

func groupRecordsByTimeInterval(params GroupByTimeIntervalParams) (*ValGroupingResult, error) {

	computeOneTimeIntervalVals := func(asOfTime time.Time) (*ValGroup, error) {

		unfilteredRecordValues, mapErr := recordValueMappingController.MapAllRecordUpdatesToFieldValues(
			params.trackerDBHandle, params.currUserID, params.databaseID, asOfTime)
		if mapErr != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Error updating records: %v", mapErr)
		}

		preFilteredRecords, preFilterErr := recordFilter.FilterRecordValues(
			params.trackerDBHandle, params.currUserID, params.preFilterRules, unfilteredRecordValues)
		if preFilterErr != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Error pre-filtering records: %v", preFilterErr)
		}

		filteredRecords, err := recordFilter.FilterRecordValues(
			params.trackerDBHandle, params.currUserID, params.filterRules, preFilteredRecords)
		if err != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Error filtering records: %v", err)
		}

		groupLabel := asOfTime.Format("2006-01-02")

		valGroup := ValGroup{
			GroupLabel:     groupLabel,
			RecordsInGroup: filteredRecords}

		return &valGroup, nil

	}

	currDate := time.Now().UTC()
	currValGroup, err := computeOneTimeIntervalVals(currDate)
	if err != nil {
		return nil, fmt.Errorf("groupRecordsByTimeInterval: %v", err)
	}

	currGroup := ValGroup{
		GroupLabel:     "Current",
		RecordsInGroup: currValGroup.RecordsInGroup}

	valGroups := []ValGroup{}
	timeInrements := generateTimeIncrementsForValGrouping(params.valGrouping)
	for _, currIncrTime := range timeInrements {
		intervalValGroup, err := computeOneTimeIntervalVals(currIncrTime)
		if err != nil {
			return nil, fmt.Errorf("groupRecordsByTimeInterval: %v", err)
		}
		valGroups = append(valGroups, *intervalValGroup)

	}

	valGroupingResult := ValGroupingResult{
		ValGroups:     valGroups,
		OverallGroup:  currGroup,
		GroupingLabel: "Date"}

	return &valGroupingResult, nil

}

func groupRecordsByFieldValue(trackerDBHandle *sql.DB, valGrouping values.ValGrouping,
	recValResults []recordValue.RecordValueResults) (*ValGroupingResult, error) {

	if (valGrouping.GroupValsByFieldID == nil) || (valGrouping.GroupValsBy == nil) {
		return nil, fmt.Errorf("Malformed grouping - missing field ID or grouping")
	}

	groupingField, fieldErr := field.GetField(trackerDBHandle, *valGrouping.GroupValsByFieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("groupRecords: Can't get field to group records: error = %v", fieldErr)
	}

	// Use a map to group the values. Values are added to the same GroupVal if they have the same
	// group label.
	groupLabelValGroupMap := map[string]*IntermediateValGroup{}
	overallRecordValResults := []recordValue.RecordValueResults{}
	appendRecordToResults := func(recValResults recordValue.RecordValueResults, groupLabelInfo valGroupLabelInfo) {
		_, groupExists := groupLabelValGroupMap[groupLabelInfo.label]
		if !groupExists {
			groupLabelValGroupMap[groupLabelInfo.label] = &IntermediateValGroup{groupLabelInfo, []recordValue.RecordValueResults{}}
		}
		valGroup := groupLabelValGroupMap[groupLabelInfo.label]
		valGroup.RecordsInGroup = append(valGroup.RecordsInGroup, recValResults)

		overallRecordValResults = append(overallRecordValResults, recValResults)
	}

	for _, currRecValResults := range recValResults {
		groupLabelInfo, lblErr := recordGroupLabelInfo(valGrouping, *groupingField, currRecValResults)
		if lblErr != nil {
			return nil, fmt.Errorf("groupRecords: Error getting label to group records: error = %v", lblErr)
		}

		if groupLabelInfo.isBlank {
			if valGrouping.IncludeBlank {
				appendRecordToResults(currRecValResults, *groupLabelInfo)
			}
		} else {
			appendRecordToResults(currRecValResults, *groupLabelInfo)
		}

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

	groupingLabel, groupingLabelErr := valGrouping.GroupingLabel(trackerDBHandle)
	if groupingLabelErr != nil {
		return nil, fmt.Errorf("groupRecords: Error getting grouping label: error = %v", groupingLabelErr)
	}

	overallGroup := ValGroup{
		GroupLabel:     "Overall",
		RecordsInGroup: overallRecordValResults}

	return &ValGroupingResult{
		ValGroups:     valGroups,
		GroupingLabel: groupingLabel,
		OverallGroup:  overallGroup}, nil
}

// valGroupLabelInfo is used as intermediate data to produce a "normalized" sort value
// for use on the grouping column of data results.
type valGroupLabelInfo struct {
	label       string
	isBlank     bool
	textSortVal *string
	numSortVal  *float64
}

func textGroupLabelInfo(label string) *valGroupLabelInfo {
	sortVal := label
	return &valGroupLabelInfo{
		label:       label,
		textSortVal: &sortVal,
		isBlank:     false}
}

func blankGroupLabelInfo() *valGroupLabelInfo {
	textSortVal := ""
	label := ""
	numSortVal := -1.0 * math.MaxFloat64
	return &valGroupLabelInfo{
		label:       label,
		textSortVal: &textSortVal,
		numSortVal:  &numSortVal,
		isBlank:     true}
}

func numberGroupLabelInfo(label string, numberSortVal float64) *valGroupLabelInfo {
	sortVal := numberSortVal
	return &valGroupLabelInfo{
		label:      label,
		numSortVal: &sortVal,
		isBlank:    false}
}

func timeGroupLabelInfo(label string, timeVal time.Time) *valGroupLabelInfo {
	sortVal := float64(timeVal.UnixNano())
	return &valGroupLabelInfo{
		label:      label,
		numSortVal: &sortVal,
		isBlank:    false}
}

func groupTimeFieldRecordVal(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (*valGroupLabelInfo, error) {

	if valGrouping.GroupValsBy == nil {
		return nil, fmt.Errorf("Malformed value grouping: missing grouping by value")
	}
	groupValsBy := *valGrouping.GroupValsBy

	if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
		timeVal, valFound := recValResults.FieldValues.GetTimeFieldValue(fieldGroup.FieldID)
		if !valFound {
			return blankGroupLabelInfo(), nil
		} else {
			switch groupValsBy {
			case values.ValGroupByNone:
				return numberGroupLabelInfo("All Dates", 0.0), nil
			case values.ValGroupByDay:
				return timeGroupLabelInfo(timeVal.Format("2006-01-02"), timeVal), nil
			case values.ValGroupByWeek:
				weekStart := beginningOfWeek(timeVal)
				return timeGroupLabelInfo(weekStart.Format("2006-01-02"), weekStart), nil
			case values.ValGroupByMonthYear:
				return timeGroupLabelInfo(timeVal.Format("Jan 2006"), timeVal), nil
			default:
				return nil, fmt.Errorf("Invalid grouping = %v for time field type", valGrouping.GroupValsBy)
			} // switch groupValsBy
		}
	} else {
		return blankGroupLabelInfo(), nil
	}
}

func groupBoolFieldRecordVal(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (*valGroupLabelInfo, error) {

	if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
		boolVal, valFound := recValResults.FieldValues.GetBoolFieldValue(fieldGroup.FieldID)
		if !valFound {
			return blankGroupLabelInfo(), nil
		} else {
			if boolVal == true {
				return numberGroupLabelInfo("True", 1.0), nil
			} else {
				return numberGroupLabelInfo("False", 0.0), nil
			}
		}
	} else {
		return blankGroupLabelInfo(), nil
	}
}

func formattedValGroupNumber(val float64, valGrouping values.ValGrouping) string {
	if valGrouping.NumberFormat != nil {
		return numberFormat.FormatNumber(val, *valGrouping.NumberFormat)
	} else {
		return numberFormat.FormatNumber(val, numberFormat.NumberFormatGeneral)
	}
}

func bucketedNumberGroupLabelInfo(numberVal float64, valGrouping values.ValGrouping) *valGroupLabelInfo {

	if (valGrouping.BucketStart) != nil && numberVal < (*valGrouping.BucketStart) {
		formattedStart := formattedValGroupNumber((*valGrouping.BucketStart), valGrouping)
		return numberGroupLabelInfo(fmt.Sprintf("< %v", formattedStart), numberVal)
	}
	if (valGrouping.BucketEnd) != nil && numberVal > (*valGrouping.BucketEnd) {
		formattedEnd := formattedValGroupNumber((*valGrouping.BucketEnd), valGrouping)
		return numberGroupLabelInfo(fmt.Sprintf("> %v", formattedEnd), numberVal)
	}

	var bucketWidth = 1.0
	if valGrouping.GroupByValBucketWidth != nil && *valGrouping.GroupByValBucketWidth > 0.0 {
		bucketWidth = *valGrouping.GroupByValBucketWidth
	}

	numBuckets := math.Trunc(numberVal / bucketWidth)
	rem := math.Remainder(numberVal, bucketWidth)
	start := bucketWidth * numBuckets
	if numberVal < 0.0 && rem != 0.0 {
		start = start - bucketWidth
	}
	end := start + bucketWidth

	return numberGroupLabelInfo(fmt.Sprintf("%v to %v",
		formattedValGroupNumber(start, valGrouping), formattedValGroupNumber(end, valGrouping)), numberVal)
}

func groupNumberFieldRecordVal(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (*valGroupLabelInfo, error) {

	if valGrouping.GroupValsBy == nil {
		return nil, fmt.Errorf("Malformed value grouping: missing grouping by value")
	}
	groupValsBy := *valGrouping.GroupValsBy

	if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
		numberVal, valFound := recValResults.FieldValues.GetNumberFieldValue(fieldGroup.FieldID)
		if !valFound {
			return nil, fmt.Errorf("groupNumberFieldRecordVal: Unabled to retrieve value for grouping label")
		} else {
			switch groupValsBy {
			case values.ValGroupByNone:
				formattedVal := formattedValGroupNumber(numberVal, valGrouping)
				return numberGroupLabelInfo(formattedVal, numberVal), nil
			case values.ValGroupByBucket:
				return bucketedNumberGroupLabelInfo(numberVal, valGrouping), nil
			default:
				return numberGroupLabelInfo("All Numbers", 0.0), nil
			} // switch groupValsBy
		}
	} else {
		return blankGroupLabelInfo(), nil
	}

}

func recordGroupLabelInfo(valGrouping values.ValGrouping, fieldGroup field.Field,
	recValResults recordValue.RecordValueResults) (*valGroupLabelInfo, error) {
	switch fieldGroup.Type {
	case field.FieldTypeText:
		if recValResults.FieldValues.ValueIsSet(fieldGroup.FieldID) {
			textVal, foundVal := recValResults.FieldValues.GetTextFieldValue(fieldGroup.FieldID)
			if !foundVal {
				return nil, fmt.Errorf("recordGroupLabel: Unabled to retrieve value for grouping label")
			} else {
				return textGroupLabelInfo(textVal), nil
			}
		} else {
			return blankGroupLabelInfo(), nil
		}
	case field.FieldTypeBool:
		return groupBoolFieldRecordVal(valGrouping, fieldGroup, recValResults)
	case field.FieldTypeNumber:
		return groupNumberFieldRecordVal(valGrouping, fieldGroup, recValResults)
	case field.FieldTypeTime:
		return groupTimeFieldRecordVal(valGrouping, fieldGroup, recValResults)
	}
	return nil, fmt.Errorf("recordGroupLabel: unsupported grouping: fieldRef = %+v", fieldGroup)
}
