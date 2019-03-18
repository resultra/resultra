// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package dashboardController

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/dashboard/values"
	"resultra/tracker/server/recordValue"
)

type GroupDataRow struct {
	GroupLabel  string    `json:"groupLabel"`
	SummaryVals []float64 `json:"summaryVals"`
}

type GroupedSummarizedVals struct {
	GroupedDataRows      []GroupDataRow `json:"groupedDataRows"`
	OverallDataRow       GroupDataRow   `json:"overallDataRow"`
	GroupingLabel        string         `json:"groupingLabel"`
	SummaryLabels        []string       `json:"summaryLabels"`
	SummaryNumberFormats []string       `json:"summaryNumberFormats"`
}

func computeSummarySum(recordsInGroup []recordValue.RecordValueResults, fieldID string) (float64, error) {
	sum := 0.0
	for _, currRecordVal := range recordsInGroup {
		numberVal, valFound := currRecordVal.FieldValues.GetNumberFieldValue(fieldID)
		if valFound {
			sum += numberVal
		}
	}
	return sum, nil
}

func computeSummaryAvg(recordsInGroup []recordValue.RecordValueResults, fieldID string) (float64, error) {
	sum := 0.0
	numRecords := 0.0
	for _, currRecordVal := range recordsInGroup {
		numberVal, valFound := currRecordVal.FieldValues.GetNumberFieldValue(fieldID)
		if valFound {
			sum += numberVal
			numRecords += 1.0
		}
	}
	if numRecords > 0.0 {
		avg := sum / numRecords
		return avg, nil
	} else {
		return 0.0, nil
	}
	return sum, nil
}

func computePercTrue(recordsInGroup []recordValue.RecordValueResults, fieldID string) (float64, error) {
	numTrue := 0.0
	numRecords := 0.0
	for _, currRecordVal := range recordsInGroup {
		boolVal, valFound := currRecordVal.FieldValues.GetBoolFieldValue(fieldID)
		if valFound {
			numRecords += 1.0
			if boolVal == true {
				numTrue += 1.0
			}
		}
	}
	if numRecords > 0.0 {
		percTrue := numTrue / numRecords
		return percTrue, nil
	} else {
		return 0.0, nil
	}
}

func computeCountTrue(recordsInGroup []recordValue.RecordValueResults, fieldID string) (float64, error) {
	numTrue := 0.0
	numRecords := 0.0
	for _, currRecordVal := range recordsInGroup {
		boolVal, valFound := currRecordVal.FieldValues.GetBoolFieldValue(fieldID)
		if valFound {
			numRecords += 1.0
			if boolVal == true {
				numTrue += 1.0
			}
		}
	}
	if numRecords > 0.0 {
		return numTrue, nil
	} else {
		return 0.0, nil
	}
}

func computeCountFalse(recordsInGroup []recordValue.RecordValueResults, fieldID string) (float64, error) {
	numFalse := 0.0
	numRecords := 0.0
	for _, currRecordVal := range recordsInGroup {
		boolVal, valFound := currRecordVal.FieldValues.GetBoolFieldValue(fieldID)
		if valFound {
			numRecords += 1.0
			if boolVal == false {
				numFalse += 1.0
			}
		}
	}
	if numRecords > 0.0 {
		return numFalse, nil
	} else {
		return 0.0, nil
	}
}

func summarizeOneGroupedVal(recordsInGroup []recordValue.RecordValueResults, summary values.ValSummary) (float64, error) {
	// TODO - Replace dummied up summary value with one computed for the
	// specific ValSummary.
	switch summary.SummarizeValsWith {
	case values.ValSummaryCount:
		countOfRecords := float64(len(recordsInGroup))
		summaryVal := countOfRecords
		return summaryVal, nil
	case values.ValSummarySum:
		return computeSummarySum(recordsInGroup, summary.SummarizeByFieldID)
	case values.ValSummaryAvg:
		return computeSummaryAvg(recordsInGroup, summary.SummarizeByFieldID)
	case values.ValSummaryCountTrue:
		return computeCountTrue(recordsInGroup, summary.SummarizeByFieldID)
	case values.ValSummaryCountFalse:
		return computeCountFalse(recordsInGroup, summary.SummarizeByFieldID)
	case values.ValSummaryPercTrue:
		return computePercTrue(recordsInGroup, summary.SummarizeByFieldID)
	case values.ValSummaryPercFalse:
		percTrue, err := computePercTrue(recordsInGroup, summary.SummarizeByFieldID)
		if err != nil {
			return 0.0, err
		} else {
			return (1.0 - percTrue), nil
		}
	default:
		return 0.0, fmt.Errorf("Unsupported summary type = %v", summary.SummarizeValsWith)
	}

}

func computeOneGroupSummarizedVals(valGroup ValGroup, summaries []values.ValSummary) (*GroupDataRow, error) {
	summaryVals := []float64{}
	for _, currValSummary := range summaries {

		summaryVal, summaryErr := summarizeOneGroupedVal(valGroup.RecordsInGroup, currValSummary)
		if summaryErr != nil {
			return nil, fmt.Errorf("summarizeGroupedRecords: Error summarizing value: %v", summaryErr)
		}

		summaryVals = append(summaryVals, summaryVal)
	}

	groupDataRow := GroupDataRow{
		GroupLabel:  valGroup.GroupLabel,
		SummaryVals: summaryVals}

	return &groupDataRow, nil

}

func summarizeGroupedRecords(trackerDBHandle *sql.DB,
	valGroupingResult *ValGroupingResult, summaries []values.ValSummary) (*GroupedSummarizedVals, error) {

	groupDataRows := []GroupDataRow{}

	for _, currValGroup := range valGroupingResult.ValGroups {
		groupDataRow, err := computeOneGroupSummarizedVals(currValGroup, summaries)
		if err != nil {
			return nil, fmt.Errorf("summarizeGroupedRecords: Error summarizing value: %v", err)
		}
		groupDataRows = append(groupDataRows, *groupDataRow)
	}

	overallDataRow, overallErr := computeOneGroupSummarizedVals(valGroupingResult.OverallGroup, summaries)
	if overallErr != nil {
		return nil, fmt.Errorf("summarizeGroupedRecords: Error summarizing value: %v", overallErr)
	}

	summaryLabels := []string{}
	summaryFormats := []string{}
	for _, currValSummary := range summaries {

		summaryLabel, err := currValSummary.SummaryLabel(trackerDBHandle)
		if err != nil {
			return nil, fmt.Errorf("summarizeGroupedRecords: can't generate summary: %v", err)
		}

		summaryLabels = append(summaryLabels, summaryLabel)

		summaryFormats = append(summaryFormats, currValSummary.NumberFormat)
	}

	groupedSummarizedVals := GroupedSummarizedVals{
		GroupedDataRows:      groupDataRows,
		GroupingLabel:        valGroupingResult.GroupingLabel,
		SummaryLabels:        summaryLabels,
		OverallDataRow:       *overallDataRow,
		SummaryNumberFormats: summaryFormats}

	return &groupedSummarizedVals, nil
}
