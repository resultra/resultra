package dashboardController

import (
	"database/sql"
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
	GaugesData        []GaugeData         `json:"gaugesData"`
	SummaryValsData   []SummaryValData    `json:"summaryValsData"`
	Headers           []header.Header     `json:"headers"`
}

func getDefaultDashboardData(trackerDBHandle *sql.DB, currUserID string, params GetDashboardDataParams) (*DashboardDataRef, error) {

	dashboard, err := dashboard.GetDashboard(trackerDBHandle, params.DashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard: error = %v", err)
	}

	barChartData, getBarChartsErr := getDefaultDashboardBarChartsData(trackerDBHandle, currUserID, params.DashboardID)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard barchart data: error = %v", getBarChartsErr)
	}

	summaryTablesData, getTableErr := getDefaultDashboardSummaryTablesData(trackerDBHandle, currUserID, params.DashboardID)
	if getTableErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard summary tables data: error = %v", getTableErr)
	}

	gaugesData, getGaugeErr := getDefaultDashboardGaugesData(trackerDBHandle, currUserID, params.DashboardID)
	if getGaugeErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard gauges data: error = %v", getGaugeErr)
	}

	summaryValsData, getSummaryValsErr := getDefaultDashboardSummaryValsData(trackerDBHandle, currUserID, params.DashboardID)
	if getSummaryValsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard gauges data: error = %v", getSummaryValsErr)
	}

	headers, getHeaderErr := header.GetHeaders(trackerDBHandle, params.DashboardID)
	if getHeaderErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve headers: error = %v", getTableErr)
	}

	dashboardDataRef := DashboardDataRef{
		Dashboard:         *dashboard,
		BarChartsData:     barChartData,
		SummaryTablesData: summaryTablesData,
		GaugesData:        gaugesData,
		SummaryValsData:   summaryValsData,
		Headers:           headers}

	return &dashboardDataRef, nil
}
