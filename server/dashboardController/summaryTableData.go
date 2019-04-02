// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package dashboardController

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/recordSortDataModel"
	"github.com/resultra/resultra/server/dashboard"
	"github.com/resultra/resultra/server/dashboard/components/summaryTable"
	"github.com/resultra/resultra/server/recordFilter"
	"github.com/resultra/resultra/server/recordReadController"
)

type SummaryTableData struct {
	SummaryTableID        string                    `json:"summaryTableID"`
	SummaryTable          summaryTable.SummaryTable `json:"summaryTable"`
	GroupedSummarizedVals GroupedSummarizedVals     `json:"groupedSummarizedVals"`
}

func getOneSummaryTableData(trackerDBHandle *sql.DB, currUserID string,
	summaryTable *summaryTable.SummaryTable, filterRules recordFilter.RecordFilterRuleSet) (*SummaryTableData, error) {

	parentDashboard, err := dashboard.GetDashboard(trackerDBHandle, summaryTable.ParentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: %v", err)
	}

	var valGroupingResult *ValGroupingResult
	if summaryTable.Properties.RowGroupingVals.GroupValsByFieldID != nil {
		sortRules := []recordSortDataModel.RecordSortRule{}
		getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
			DatabaseID:     parentDashboard.ParentDatabaseID,
			PreFilterRules: summaryTable.Properties.PreFilterRules,
			FilterRules:    filterRules,
			SortRules:      sortRules}
		recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(trackerDBHandle, currUserID, getRecordParams)
		if getRecErr != nil {
			return nil, fmt.Errorf("getOneSummaryTableData: Error retrieving records for summary table: %v", getRecErr)
		}

		groupingResult, groupingErr := groupRecordsByFieldValue(trackerDBHandle, summaryTable.Properties.RowGroupingVals, recordRefs)
		if groupingErr != nil {
			return nil, fmt.Errorf("getOneSummaryTableData: Error grouping records for summary table: %v", groupingErr)
		}
		valGroupingResult = groupingResult

	} else {
		timeIncrementGroupingParams := GroupByTimeIntervalParams{
			trackerDBHandle: trackerDBHandle,
			databaseID:      parentDashboard.ParentDatabaseID,
			currUserID:      currUserID,
			preFilterRules:  summaryTable.Properties.PreFilterRules,
			filterRules:     filterRules,
			valGrouping:     summaryTable.Properties.RowGroupingVals}
		groupingResult, groupingErr := groupRecordsByTimeInterval(timeIncrementGroupingParams)
		if groupingErr != nil {
			return nil, fmt.Errorf("getOneSummaryTableData: Error grouping records for summary table: %v", groupingErr)
		}
		valGroupingResult = groupingResult
	}

	groupedSummarizedVals, summarizeErr := summarizeGroupedRecords(trackerDBHandle, valGroupingResult,
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

func getSummaryTableData(trackerDBHandle *sql.DB, currUserID string, params GetSummaryTableDataParams) (*SummaryTableData, error) {

	if len(params.SummaryTableID) <= 0 {
		return nil, fmt.Errorf("GetSummaryTableData: missing summary table ID")
	}

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("GetSummaryTableData: missing dashboard ID")
	}

	summaryTable, getSummaryTableErr := summaryTable.GetSummaryTable(trackerDBHandle, params.ParentDashboardID, params.SummaryTableID)
	if getSummaryTableErr != nil {
		return nil, fmt.Errorf("GetSummaryTableData: Error retrieving summary table with params=%+v: error= %v",
			params, getSummaryTableErr)
	}

	summaryTableData, dataErr := getOneSummaryTableData(trackerDBHandle, currUserID, summaryTable, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("GetSummaryTableData: Error retrieving bar chart data: %v", dataErr)
	}

	return summaryTableData, nil

}

func getDefaultDashboardSummaryTablesData(trackerDBHandle *sql.DB, currUserID string, parentDashboardID string) ([]SummaryTableData, error) {

	summaryTables, err := summaryTable.GetSummaryTables(trackerDBHandle, parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardBarChartsData: Error retrieving bar charts: %v", err)
	}

	summaryTablesData := []SummaryTableData{}
	for _, summaryTable := range summaryTables {

		summaryTableData, dataErr := getOneSummaryTableData(trackerDBHandle, currUserID, &summaryTable,
			summaryTable.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("GetData: Error retrieving summary table data: %v", dataErr)
		}
		summaryTablesData = append(summaryTablesData, *summaryTableData)
	}

	return summaryTablesData, nil
}
