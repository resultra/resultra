package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/dashboard/values"
	//	"resultra/datasheet/server/field"
	"resultra/datasheet/server/recordValue"
)

type GroupDataRow struct {
	GroupLabel  string    `json:"groupLabel"`
	SummaryVals []float64 `json:"summaryVals"`
}

type GroupedSummarizedVals struct {
	GroupedDataRows []GroupDataRow `json:"groupedDataRows"`
	GroupingLabel   string         `json:"groupingLabel"`
	SummaryLabels   []string       `json:"summaryLabels"`
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
	default:
		return 0.0, fmt.Errorf("Unsupported summary type = %v", summary.SummarizeValsWith)
	}

}

func summarizeGroupedRecords(valGroupingResult *ValGroupingResult, summaries []values.ValSummary) (*GroupedSummarizedVals, error) {

	groupDataRows := []GroupDataRow{}

	for _, currValGroup := range valGroupingResult.ValGroups {

		summaryVals := []float64{}
		for _, currValSummary := range summaries {

			summaryVal, summaryErr := summarizeOneGroupedVal(currValGroup.RecordsInGroup, currValSummary)
			if summaryErr != nil {
				return nil, fmt.Errorf("summarizeGroupedRecords: Error summarizing value: %v", summaryErr)
			}

			summaryVals = append(summaryVals, summaryVal)
		}

		groupDataRow := GroupDataRow{
			GroupLabel:  currValGroup.GroupLabel,
			SummaryVals: summaryVals}

		groupDataRows = append(groupDataRows, groupDataRow)

	}

	summaryLabels := []string{}
	for _, currValSummary := range summaries {

		summaryLabel, err := currValSummary.SummaryLabel()
		if err != nil {
			return nil, fmt.Errorf("summarizeGroupedRecords: can't generate summary: %v", err)
		}

		summaryLabels = append(summaryLabels, summaryLabel)
	}

	groupedSummarizedVals := GroupedSummarizedVals{
		GroupedDataRows: groupDataRows,
		GroupingLabel:   valGroupingResult.GroupingLabel,
		SummaryLabels:   summaryLabels}

	return &groupedSummarizedVals, nil
}
