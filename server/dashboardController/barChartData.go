package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordReadController"
)

type BarChartDataRow struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type BarChartData struct {
	BarChartID string            `json:"barChartID"`
	BarChart   barChart.BarChart `json:"barChart"`
	Title      string            `json:"title"`
	XAxisTitle string            `json:"xAxisTitle"`
	YAxisTitle string            `json:"yAxisTitle"`
	DataRows   []BarChartDataRow `json:"dataRows"`
}

func getOneBarChartData(barChart *barChart.BarChart, filterRules []recordFilter.RecordFilterRule) (*BarChartData, error) {

	parentDashboard, err := dashboard.GetDashboard(barChart.ParentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: %v", err)
	}

	// TODO - Store the list of filters with the bar chart and include it in the query.
	sortRules := []recordSortDataModel.RecordSortRule{}
	// TODO - Fully support pre-filtering for bar charts instead of no prefiltering.
	preFilterRules := []recordFilter.RecordFilterRule{}
	getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
		DatabaseID:     parentDashboard.ParentDatabaseID,
		PreFilterRules: preFilterRules,
		FilterRules:    filterRules,
		SortRules:      sortRules}
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult, groupingErr := groupRecords(barChart.Properties.XAxisVals, recordRefs)
	if groupingErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error grouping records for bar chart: %v", groupingErr)
	}

	dataRows := []BarChartDataRow{}
	for _, valGroup := range valGroupingResult.ValGroups {
		dataRows = append(dataRows,
			BarChartDataRow{valGroup.GroupLabel, float64(len(valGroup.RecordsInGroup))})
	}

	barChartData := BarChartData{
		BarChartID: barChart.BarChartID,
		BarChart:   *barChart,
		Title:      barChart.Properties.Title,
		XAxisTitle: valGroupingResult.GroupingLabel,
		YAxisTitle: "Count",
		DataRows:   dataRows}

	return &barChartData, nil

}

type GetBarChartDataParams struct {
	ParentDashboardID string                          `json:"parentDashboardID"`
	BarChartID        string                          `json:"barChartID"`
	FilterRules       []recordFilter.RecordFilterRule `json:"filterRules"`
}

func getBarChartData(params GetBarChartDataParams) (*BarChartData, error) {

	barChart, getBarChartErr := barChart.GetBarChart(params.ParentDashboardID, params.BarChartID)
	if getBarChartErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart with params = %+v: error= %v",
			params, getBarChartErr)
	}

	barChartData, dataErr := getOneBarChartData(barChart, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
	}

	return barChartData, nil

}

func getDefaultDashboardBarChartsData(parentDashboardID string) ([]BarChartData, error) {

	barCharts, err := barChart.GetBarCharts(parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getDefaultDashboardBarChartsData: Error retrieving bar charts: %v", err)
	}

	var barChartsData []BarChartData
	for _, barChart := range barCharts {

		barChartData, dataErr := getOneBarChartData(&barChart, barChart.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
		}
		barChartsData = append(barChartsData, *barChartData)
	}

	return barChartsData, nil
}
