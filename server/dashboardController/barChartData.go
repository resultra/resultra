package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/dashboard/components/barChart"
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

func getOneBarChartData(barChart *barChart.BarChart, filterIDs []string) (*BarChartData, error) {

	tableID := barChart.Properties.DataSrcTableID

	// TODO - Store the list of filters with the bar chart and include it in the query.
	sortRules := []recordSortDataModel.RecordSortRule{}
	getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
		TableID:   tableID,
		FilterIDs: filterIDs,
		SortRules: sortRules}
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult, groupingErr := groupRecords(barChart.Properties.XAxisVals, tableID, recordRefs)
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
	ParentDashboardID string   `json:"parentDashboardID"`
	BarChartID        string   `json:"barChartID"`
	FilterIDs         []string `json:"filterIDs"`
}

func getBarChartData(params GetBarChartDataParams) (*BarChartData, error) {

	barChart, getBarChartErr := barChart.GetBarChart(params.ParentDashboardID, params.BarChartID)
	if getBarChartErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart with params = %+v: error= %v",
			params, getBarChartErr)
	}

	barChartData, dataErr := getOneBarChartData(barChart, params.FilterIDs)
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

		filterIDs := barChart.Properties.DefaultFilterIDs
		barChartData, dataErr := getOneBarChartData(&barChart, filterIDs)
		if dataErr != nil {
			return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
		}
		barChartsData = append(barChartsData, *barChartData)
	}

	return barChartsData, nil
}
