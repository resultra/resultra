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

func summarizeOneGroupedVal(recordsInGroup []recordValue.RecordValueResults, summary values.ValSummary) (float64, error) {
	// TODO - Replace dummied up summary value with one computed for the
	// specific ValSummary.
	countOfRecords := float64(len(recordsInGroup))

	summaryVal := countOfRecords

	return summaryVal, nil

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
