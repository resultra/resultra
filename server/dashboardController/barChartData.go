package dashboardController

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordReadController"
)

type BarChartDataRow struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type BarChartData struct {
	BarChartID            string                `json:"barChartID"`
	BarChart              barChart.BarChart     `json:"barChart"`
	Title                 string                `json:"title"`
	GroupedSummarizedVals GroupedSummarizedVals `json:"groupedSummarizedVals"`
}

func getOneBarChartData(trackerDBHandle *sql.DB, currUserID string, barChart *barChart.BarChart,
	filterRules recordFilter.RecordFilterRuleSet) (*BarChartData, error) {

	parentDashboard, err := dashboard.GetDashboard(trackerDBHandle, barChart.ParentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: %v", err)
	}

	// TODO - Store the list of filters with the bar chart and include it in the query.
	sortRules := []recordSortDataModel.RecordSortRule{}
	getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
		DatabaseID:     parentDashboard.ParentDatabaseID,
		PreFilterRules: barChart.Properties.PreFilterRules,
		FilterRules:    filterRules,
		SortRules:      sortRules}
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(trackerDBHandle, currUserID, getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult, groupingErr := groupRecords(trackerDBHandle, barChart.Properties.XAxisVals, recordRefs)
	if groupingErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error grouping records for bar chart: %v", groupingErr)
	}

	summaries := []values.ValSummary{}
	summaries = append(summaries, barChart.Properties.YAxisVals)
	groupedSummarizedVals, summarizeErr := summarizeGroupedRecords(trackerDBHandle, valGroupingResult, summaries)
	if summarizeErr != nil {
		return nil, fmt.Errorf("getOneBarChartData: Error grouping records for bar chart: %v", summarizeErr)
	}

	barChartData := BarChartData{
		BarChartID: barChart.BarChartID,
		BarChart:   *barChart,
		Title:      barChart.Properties.Title,
		GroupedSummarizedVals: *groupedSummarizedVals}

	return &barChartData, nil

}

type GetBarChartDataParams struct {
	ParentDashboardID string                           `json:"parentDashboardID"`
	BarChartID        string                           `json:"barChartID"`
	FilterRules       recordFilter.RecordFilterRuleSet `json:"filterRules"`
}

func getBarChartData(trackerDBHandle *sql.DB, currUserID string, params GetBarChartDataParams) (*BarChartData, error) {

	barChart, getBarChartErr := barChart.GetBarChart(trackerDBHandle, params.ParentDashboardID, params.BarChartID)
	if getBarChartErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart with params = %+v: error= %v",
			params, getBarChartErr)
	}

	barChartData, dataErr := getOneBarChartData(trackerDBHandle, currUserID, barChart, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
	}

	return barChartData, nil

}

func getDefaultDashboardBarChartsData(trackerDBHandle *sql.DB, currUserID string, parentDashboardID string) ([]BarChartData, error) {

	barCharts, err := barChart.GetBarCharts(trackerDBHandle, parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getDefaultDashboardBarChartsData: Error retrieving bar charts: %v", err)
	}

	var barChartsData []BarChartData
	for _, barChart := range barCharts {

		barChartData, dataErr := getOneBarChartData(trackerDBHandle, currUserID, &barChart, barChart.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
		}
		barChartsData = append(barChartsData, *barChartData)
	}

	return barChartsData, nil
}
