package dashboardController

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/recordSortDataModel"
	"resultra/tracker/server/dashboard"
	"resultra/tracker/server/dashboard/components/summaryValue"
	"resultra/tracker/server/dashboard/values"
	"resultra/tracker/server/recordFilter"
	"resultra/tracker/server/recordReadController"
)

type SummaryValData struct {
	SummaryValID          string                  `json:"summaryValID"`
	SummaryVal            summaryValue.SummaryVal `json:"summaryVal"`
	Title                 string                  `json:"title"`
	GroupedSummarizedVals GroupedSummarizedVals   `json:"groupedSummarizedVals"`
}

func getOneSummaryValData(trackerDBHandle *sql.DB, currUserID string, summaryVal *summaryValue.SummaryVal,
	filterRules recordFilter.RecordFilterRuleSet) (*SummaryValData, error) {

	parentDashboard, err := dashboard.GetDashboard(trackerDBHandle, summaryVal.ParentDashboardID)
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
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(trackerDBHandle, currUserID, getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult := groupRecordsIntoSingleGroup(recordRefs)

	summaries := []values.ValSummary{}
	summaries = append(summaries, summaryVal.Properties.ValSummary)
	groupedSummarizedVals, summarizeErr := summarizeGroupedRecords(trackerDBHandle, valGroupingResult, summaries)
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

func getSummaryValData(trackerDBHandle *sql.DB, currUserID string, params GetSummaryValDataParams) (*SummaryValData, error) {

	summaryVal, getErr := summaryValue.GetSummaryVal(trackerDBHandle, params.ParentDashboardID, params.SummaryValID)
	if getErr != nil {
		return nil, fmt.Errorf("getSummaryValData: Error retrieving summaryVal with params = %+v: error= %v",
			params, getErr)
	}

	summaryValData, dataErr := getOneSummaryValData(trackerDBHandle, currUserID, summaryVal, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("getSummaryValData: Error retrieving summaryVal data: %v", dataErr)
	}

	return summaryValData, nil

}

func getDefaultDashboardSummaryValsData(trackerDBHandle *sql.DB, currUserID string, parentDashboardID string) ([]SummaryValData, error) {

	summaryVals, err := summaryValue.GetSummaryVals(trackerDBHandle, parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getDefaultDashboardSummaryValsData: Error retrieving summaryVals: %v", err)
	}

	var summaryValsData []SummaryValData
	for _, summaryVal := range summaryVals {

		summaryValData, dataErr := getOneSummaryValData(trackerDBHandle, currUserID, &summaryVal, summaryVal.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("getDefaultDashboardSummaryValsData: Error retrieving summaryVal data: %v", dataErr)
		}
		summaryValsData = append(summaryValsData, *summaryValData)
	}

	return summaryValsData, nil
}
