package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/dashboard/components/summaryTable"
	"resultra/datasheet/server/recordReadController"
	"resultra/datasheet/server/recordSort"
)

type SummaryTableData struct {
	SummaryTableID        string                    `json:"summaryTableID"`
	SummaryTable          summaryTable.SummaryTable `json:"summaryTable"`
	GroupedSummarizedVals GroupedSummarizedVals     `json:"groupedSummarizedVals"`
}

func getOneSummaryTableData(summaryTable *summaryTable.SummaryTable) (*SummaryTableData, error) {

	tableID := summaryTable.Properties.DataSrcTableID

	// TODO - Store the list of filters with the bar chart and include it in the query.
	filterIDs := []string{}
	sortRules := []recordSort.RecordSortRule{}
	getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
		TableID:   tableID,
		FilterIDs: filterIDs,
		SortRules: sortRules}
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: Error retrieving records for summary table: %v", getRecErr)
	}

	valGroupingResult, groupingErr := groupRecords(summaryTable.Properties.RowGroupingVals, tableID, recordRefs)
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
	ParentDashboardID string `json:"parentDashboardID"`
	SummaryTableID    string `json:"summaryTableID"`
}

func getSummaryTableData(params GetSummaryTableDataParams) (*SummaryTableData, error) {

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

	summaryTableData, dataErr := getOneSummaryTableData(summaryTable)
	if dataErr != nil {
		return nil, fmt.Errorf("GetSummaryTableData: Error retrieving bar chart data: %v", dataErr)
	}

	return summaryTableData, nil

}

func getDashboardSummaryTablesData(parentDashboardID string) ([]SummaryTableData, error) {

	summaryTables, err := summaryTable.GetSummaryTables(parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardBarChartsData: Error retrieving bar charts: %v", err)
	}

	summaryTablesData := []SummaryTableData{}
	for _, summaryTable := range summaryTables {

		summaryTableData, dataErr := getOneSummaryTableData(&summaryTable)
		if dataErr != nil {
			return nil, fmt.Errorf("GetData: Error retrieving summary table data: %v", dataErr)
		}
		summaryTablesData = append(summaryTablesData, *summaryTableData)
	}

	return summaryTablesData, nil
}