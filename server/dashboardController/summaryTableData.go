package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/summaryTable"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordReadController"
)

type SummaryTableData struct {
	SummaryTableID        string                    `json:"summaryTableID"`
	SummaryTable          summaryTable.SummaryTable `json:"summaryTable"`
	GroupedSummarizedVals GroupedSummarizedVals     `json:"groupedSummarizedVals"`
}

func getOneSummaryTableData(currUserID string,
	summaryTable *summaryTable.SummaryTable, filterRules recordFilter.RecordFilterRuleSet) (*SummaryTableData, error) {

	parentDashboard, err := dashboard.GetDashboard(summaryTable.ParentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: %v", err)
	}

	sortRules := []recordSortDataModel.RecordSortRule{}
	getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
		DatabaseID:     parentDashboard.ParentDatabaseID,
		PreFilterRules: summaryTable.Properties.PreFilterRules,
		FilterRules:    filterRules,
		SortRules:      sortRules}
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(currUserID, getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: Error retrieving records for summary table: %v", getRecErr)
	}

	valGroupingResult, groupingErr := groupRecords(summaryTable.Properties.RowGroupingVals, recordRefs)
	if groupingErr != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: Error grouping records for summary table: %v", groupingErr)
	}

	groupedSummarizedVals, summarizeErr := summarizeGroupedRecords(valGroupingResult,
		summaryTable.Properties.ColumnValSummaries)
	if summarizeErr != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: Error grouping records for summary table: %v", summarizeErr)
	}

	summaryTableData := SummaryTableData{
		SummaryTableID:        summaryTable.SummaryTableID,
		SummaryTable:          *summaryTable,
		GroupedSummarizedVals: *groupedSummarizedVals}

	return &summaryTableData, nil
}

type GetSummaryTableDataParams struct {
	ParentDashboardID string                           `json:"parentDashboardID"`
	SummaryTableID    string                           `json:"summaryTableID"`
	FilterRules       recordFilter.RecordFilterRuleSet `json:"filterRules"`
}

func getSummaryTableData(currUserID string, params GetSummaryTableDataParams) (*SummaryTableData, error) {

	if len(params.SummaryTableID) <= 0 {
		return nil, fmt.Errorf("GetSummaryTableData: missing summary table ID")
	}

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("GetSummaryTableData: missing dashboard ID")
	}

	summaryTable, getSummaryTableErr := summaryTable.GetSummaryTable(params.ParentDashboardID, params.SummaryTableID)
	if getSummaryTableErr != nil {
		return nil, fmt.Errorf("GetSummaryTableData: Error retrieving summary table with params=%+v: error= %v",
			params, getSummaryTableErr)
	}

	summaryTableData, dataErr := getOneSummaryTableData(currUserID, summaryTable, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("GetSummaryTableData: Error retrieving bar chart data: %v", dataErr)
	}

	return summaryTableData, nil

}

func getDefaultDashboardSummaryTablesData(currUserID string, parentDashboardID string) ([]SummaryTableData, error) {

	summaryTables, err := summaryTable.GetSummaryTables(parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardBarChartsData: Error retrieving bar charts: %v", err)
	}

	summaryTablesData := []SummaryTableData{}
	for _, summaryTable := range summaryTables {

		summaryTableData, dataErr := getOneSummaryTableData(currUserID, &summaryTable,
			summaryTable.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("GetData: Error retrieving summary table data: %v", dataErr)
		}
		summaryTablesData = append(summaryTablesData, *summaryTableData)
	}

	return summaryTablesData, nil
}
