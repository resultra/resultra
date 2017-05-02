package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/header"
)

type GetDashboardDataParams struct {
	DashboardID string `json:"dashboardID"`
}

type DashboardDataRef struct {
	Dashboard         dashboard.Dashboard `json:"dashboard"`
	BarChartsData     []BarChartData      `json:"barChartsData"`
	SummaryTablesData []SummaryTableData  `json:"summaryTablesData"`
	Headers           []header.Header     `json:"headers"`
}

func getDefaultDashboardData(params GetDashboardDataParams) (*DashboardDataRef, error) {

	dashboard, err := dashboard.GetDashboard(params.DashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard: error = %v", err)
	}

	barChartData, getBarChartsErr := getDefaultDashboardBarChartsData(params.DashboardID)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard barchart data: error = %v", getBarChartsErr)
	}

	summaryTablesData, getTableErr := getDefaultDashboardSummaryTablesData(params.DashboardID)
	if getTableErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard summary tables data: error = %v", getTableErr)
	}

	headers, getHeaderErr := header.GetHeaders(params.DashboardID)
	if getHeaderErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve headers: error = %v", getTableErr)
	}

	dashboardDataRef := DashboardDataRef{
		Dashboard:         *dashboard,
		BarChartsData:     barChartData,
		SummaryTablesData: summaryTablesData,
		Headers:           headers}

	return &dashboardDataRef, nil
}
