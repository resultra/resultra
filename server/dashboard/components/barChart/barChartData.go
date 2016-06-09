package barChart

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/recordFilter"
)

type BarChartDataRow struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type BarChartData struct {
	BarChartID string            `json:"barChartID"`
	BarChart   BarChart          `json:"barChart"`
	Title      string            `json:"title"`
	XAxisTitle string            `json:"xAxisTitle"`
	YAxisTitle string            `json:"yAxisTitle"`
	DataRows   []BarChartDataRow `json:"dataRows"`
}

func getOneBarChartData(appEngContext appengine.Context, barChart *BarChart) (*BarChartData, error) {

	tableID := barChart.Properties.DataSrcTableID

	// TODO - Store the list of filters with the bar chart and include it in the query.
	filterIDs := []string{}
	getRecordParams := recordFilter.GetFilteredRecordsParams{
		TableID:   tableID,
		FilterIDs: filterIDs}
	recordRefs, getRecErr := recordFilter.GetFilteredRecords(appEngContext, getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult, groupingErr := barChart.Properties.XAxisVals.GroupRecords(appEngContext, tableID, recordRefs)
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

func GetBarChartData(appEngContext appengine.Context, parentDashboardID string, barChartID string) (*BarChartData, error) {

	barChart, getBarChartErr := getBarChart(appEngContext, parentDashboardID, barChartID)
	if getBarChartErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart with id=%v, parent dashboard = %v: error= %v",
			barChartID, parentDashboardID, getBarChartErr)
	}

	barChartData, dataErr := getOneBarChartData(appEngContext, barChart)
	if dataErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
	}

	return barChartData, nil

}

func GetDashboardBarChartsData(appEngContext appengine.Context, parentDashboardID string) ([]BarChartData, error) {

	var barCharts []BarChart

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentDashboardID,
		barChartEntityKind, barChartParentDashboardIDFieldName, &barCharts)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve text box components: form id=%v", parentDashboardID)
	}

	var barChartsData []BarChartData
	for _, barChart := range barCharts {

		barChartData, dataErr := getOneBarChartData(appEngContext, &barChart)
		if dataErr != nil {
			return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
		}
		barChartsData = append(barChartsData, *barChartData)
	}

	return barChartsData, nil
}
