package dashboardController

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/recordSortDataModel"
	"resultra/tracker/server/dashboard"
	"resultra/tracker/server/dashboard/components/gauge"
	"resultra/tracker/server/dashboard/values"
	"resultra/tracker/server/recordFilter"
	"resultra/tracker/server/recordReadController"
)

type GaugeData struct {
	GaugeID               string                `json:"gaugeID"`
	Gauge                 gauge.Gauge           `json:"gauge"`
	Title                 string                `json:"title"`
	GroupedSummarizedVals GroupedSummarizedVals `json:"groupedSummarizedVals"`
}

func getOneGaugeData(trackerDBHandle *sql.DB,
	currUserID string, gauge *gauge.Gauge, filterRules recordFilter.RecordFilterRuleSet) (*GaugeData, error) {

	parentDashboard, err := dashboard.GetDashboard(trackerDBHandle, gauge.ParentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getOneSummaryTableData: %v", err)
	}

	// TODO - Store the list of filters with the bar chart and include it in the query.
	sortRules := []recordSortDataModel.RecordSortRule{}
	getRecordParams := recordReadController.GetFilteredSortedRecordsParams{
		DatabaseID:     parentDashboard.ParentDatabaseID,
		PreFilterRules: gauge.Properties.PreFilterRules,
		FilterRules:    filterRules,
		SortRules:      sortRules}
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(trackerDBHandle, currUserID, getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult := groupRecordsIntoSingleGroup(recordRefs)

	summaries := []values.ValSummary{}
	summaries = append(summaries, gauge.Properties.ValSummary)
	groupedSummarizedVals, summarizeErr := summarizeGroupedRecords(trackerDBHandle, valGroupingResult, summaries)
	if summarizeErr != nil {
		return nil, fmt.Errorf("getOneBarGaugeData: Error grouping records for gauge: %v", summarizeErr)
	}

	gaugeData := GaugeData{
		GaugeID: gauge.GaugeID,
		Gauge:   *gauge,
		Title:   gauge.Properties.Title,
		GroupedSummarizedVals: *groupedSummarizedVals}

	return &gaugeData, nil

}

type GetGaugeDataParams struct {
	ParentDashboardID string                           `json:"parentDashboardID"`
	GaugeID           string                           `json:"gaugeID"`
	FilterRules       recordFilter.RecordFilterRuleSet `json:"filterRules"`
}

func getGaugeData(trackerDBHandle *sql.DB, currUserID string, params GetGaugeDataParams) (*GaugeData, error) {

	gauge, getErr := gauge.GetGauge(trackerDBHandle, params.ParentDashboardID, params.GaugeID)
	if getErr != nil {
		return nil, fmt.Errorf("getGaugeData: Error retrieving gauge with params = %+v: error= %v",
			params, getErr)
	}

	gaugeData, dataErr := getOneGaugeData(trackerDBHandle, currUserID, gauge, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("getGaugeData: Error retrieving gauge data: %v", dataErr)
	}

	return gaugeData, nil

}

func getDefaultDashboardGaugesData(trackerDBHandle *sql.DB, currUserID string, parentDashboardID string) ([]GaugeData, error) {

	gauges, err := gauge.GetGauges(trackerDBHandle, parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getDefaultDashboardGaugesData: Error retrieving gauges: %v", err)
	}

	var gaugesData []GaugeData
	for _, gauge := range gauges {

		gaugeData, dataErr := getOneGaugeData(trackerDBHandle, currUserID, &gauge, gauge.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("getDefaultDashboardGaugesData: Error retrieving gauge data: %v", dataErr)
		}
		gaugesData = append(gaugesData, *gaugeData)
	}

	return gaugesData, nil
}
