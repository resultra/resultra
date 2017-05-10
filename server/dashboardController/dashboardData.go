package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/header"
	"resultra/datasheet/server/recordValueMappingController"
)

type GetDashboardDataParams struct {
	DashboardID string `json:"dashboardID"`
}

type DashboardDataRef struct {
	Dashboard         dashboard.Dashboard `json:"dashboard"`
	BarChartsData     []BarChartData      `json:"barChartsData"`
	SummaryTablesData []SummaryTableData  `json:"summaryTablesData"`
	GaugesData        []GaugeData         `json:"gaugesData"`
	Headers           []header.Header     `json:"headers"`
}

func getDefaultDashboardData(params GetDashboardDataParams) (*DashboardDataRef, error) {

	dashboard, err := dashboard.GetDashboard(params.DashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard: error = %v", err)
	}

	mapErr := recordValueMappingController.MapAllRecordUpdatesToFieldValues(dashboard.ParentDatabaseID)
	if mapErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error updating records: %v", mapErr)
	}

	barChartData, getBarChartsErr := getDefaultDashboardBarChartsData(params.DashboardID)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard barchart data: error = %v", getBarChartsErr)
	}

	summaryTablesData, getTableErr := getDefaultDashboardSummaryTablesData(params.DashboardID)
	if getTableErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard summary tables data: error = %v", getTableErr)
	}

	gaugesData, getGaugeErr := getDefaultDashboardGaugesData(params.DashboardID)
	if getGaugeErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard gauges data: error = %v", getGaugeErr)
	}

	headers, getHeaderErr := header.GetHeaders(params.DashboardID)
	if getHeaderErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve headers: error = %v", getTableErr)
	}

	dashboardDataRef := DashboardDataRef{
		Dashboard:         *dashboard,
		BarChartsData:     barChartData,
		SummaryTablesData: summaryTablesData,
		GaugesData:        gaugesData,
		Headers:           headers}

	return &dashboardDataRef, nil
}
