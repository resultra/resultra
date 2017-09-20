package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/summaryValue"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordReadController"
)

type SummaryValData struct {
	SummaryValID          string                  `json:"summaryValID"`
	SummaryVal            summaryValue.SummaryVal `json:"summaryVal"`
	Title                 string                  `json:"title"`
	GroupedSummarizedVals GroupedSummarizedVals   `json:"groupedSummarizedVals"`
}

func getOneSummaryValData(currUserID string, summaryVal *summaryValue.SummaryVal,
	filterRules recordFilter.RecordFilterRuleSet) (*SummaryValData, error) {

	parentDashboard, err := dashboard.GetDashboard(summaryVal.ParentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: %v", err)
	}

	// TODO - Store the list of filters with the bar chart and include it in the query.
	sortRules := []recordSortDataModel.RecordSortRule{}
	getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
		DatabaseID:     parentDashboard.ParentDatabaseID,
		PreFilterRules: summaryVal.Properties.PreFilterRules,
		FilterRules:    filterRules,
		SortRules:      sortRules}
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(currUserID, getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult := groupRecordsIntoSingleGroup(recordRefs)

	summaries := []values.ValSummary{}
	summaries = append(summaries, summaryVal.Properties.ValSummary)
	groupedSummarizedVals, summarizeErr := summarizeGroupedRecords(valGroupingResult, summaries)
	if summarizeErr != nil {
		return nil, fmt.Errorf("getOneBarSummaryValData: Error grouping records for summaryVal: %v", summarizeErr)
	}

	summaryValData := SummaryValData{
		SummaryValID: summaryVal.SummaryValID,
		SummaryVal:   *summaryVal,
		Title:        summaryVal.Properties.Title,
		GroupedSummarizedVals: *groupedSummarizedVals}

	return &summaryValData, nil

}

type GetSummaryValDataParams struct {
	ParentDashboardID string                           `json:"parentDashboardID"`
	SummaryValID      string                           `json:"summaryValID"`
	FilterRules       recordFilter.RecordFilterRuleSet `json:"filterRules"`
}

func getSummaryValData(currUserID string, params GetSummaryValDataParams) (*SummaryValData, error) {

	summaryVal, getErr := summaryValue.GetSummaryVal(params.ParentDashboardID, params.SummaryValID)
	if getErr != nil {
		return nil, fmt.Errorf("getSummaryValData: Error retrieving summaryVal with params = %+v: error= %v",
			params, getErr)
	}

	summaryValData, dataErr := getOneSummaryValData(currUserID, summaryVal, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("getSummaryValData: Error retrieving summaryVal data: %v", dataErr)
	}

	return summaryValData, nil

}

func getDefaultDashboardSummaryValsData(currUserID string, parentDashboardID string) ([]SummaryValData, error) {

	summaryVals, err := summaryValue.GetSummaryVals(parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getDefaultDashboardSummaryValsData: Error retrieving summaryVals: %v", err)
	}

	var summaryValsData []SummaryValData
	for _, summaryVal := range summaryVals {

		summaryValData, dataErr := getOneSummaryValData(currUserID, &summaryVal, summaryVal.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("getDefaultDashboardSummaryValsData: Error retrieving summaryVal data: %v", dataErr)
		}
		summaryValsData = append(summaryValsData, *summaryValData)
	}

	return summaryValsData, nil
}
