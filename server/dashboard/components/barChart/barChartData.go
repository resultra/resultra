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
	BarChartID  string            `json:"barChartID"`
	BarChartRef BarChartRef       `json:"barChartRef"`
	Title       string            `json:"title"`
	XAxisTitle  string            `json:"xAxisTitle"`
	YAxisTitle  string            `json:"yAxisTitle"`
	DataRows    []BarChartDataRow `json:"dataRows"`
}

func getOneBarChartData(appEngContext appengine.Context, params BarChartUniqueID, barChart *BarChart) (*BarChartData, error) {

	barChartRef, refErr := getBarChartRef(appEngContext, params, barChart)
	if refErr != nil {
		return nil, fmt.Errorf("getDashboardBarChartsData: Error getting bar chart reference: %v", refErr)

	}

	tableID := barChart.DataSrcTable.Encode()

	recordRefs, getRecErr := recordFilter.GetFilteredRecords(appEngContext, recordFilter.GetFilteredRecordsParams{TableID: tableID})
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult, groupingErr := barChart.XAxisVals.GroupRecords(appEngContext, recordRefs)
	if groupingErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error grouping records for bar chart: %v", groupingErr)
	}

	dataRows := []BarChartDataRow{}
	for _, valGroup := range valGroupingResult.ValGroups {
		dataRows = append(dataRows,
			BarChartDataRow{valGroup.GroupLabel, float64(len(valGroup.RecordsInGroup))})
	}

	barChartData := BarChartData{
		BarChartID:  barChartRef.BarChartID,
		BarChartRef: *barChartRef,
		Title:       barChartRef.Title,
		XAxisTitle:  valGroupingResult.GroupingLabel,
		YAxisTitle:  "Count",
		DataRows:    dataRows}

	return &barChartData, nil

}

func GetBarChartData(appEngContext appengine.Context, params BarChartUniqueID) (*BarChartData, error) {

	barChart, getBarChartErr := getBarChart(appEngContext, params)
	if getBarChartErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart: %v", getBarChartErr)
	}

	barChartData, dataErr := getOneBarChartData(appEngContext, params, barChart)
	if dataErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
	}

	return barChartData, nil

}

func GetDashboardBarChartsData(appEngContext appengine.Context, parentDashboardID string) ([]BarChartData, error) {

	var barCharts []BarChart
	barChartIDs, getBarChartsErr := datastoreWrapper.GetAllChildEntities(appEngContext,
		parentDashboardID, barChartEntityKind, &barCharts)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("getDashboardBarChartsData: unable to retrieve bar charts: %v", getBarChartsErr)
	}

	var barChartsData []BarChartData
	for barChartIndex, barChart := range barCharts {

		barChartID := barChartIDs[barChartIndex]

		params := BarChartUniqueID{
			ParentDashboardID: parentDashboardID,
			BarChartID:        barChartID}

		barChartData, dataErr := getOneBarChartData(appEngContext, params, &barChart)
		if dataErr != nil {
			return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart data: %v", dataErr)
		}
		barChartsData = append(barChartsData, *barChartData)
	}

	return barChartsData, nil
}
