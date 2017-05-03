package dashboardController

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/gauge"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordReadController"
)

type GaugeData struct {
	GaugeID               string                `json:"gaugeID"`
	Gauge                 gauge.Gauge           `json:"gauge"`
	Title                 string                `json:"title"`
	GroupedSummarizedVals GroupedSummarizedVals `json:"groupedSummarizedVals"`
}

func getOneGaugeData(gauge *gauge.Gauge, filterRules recordFilter.RecordFilterRuleSet) (*GaugeData, error) {

	parentDashboard, err := dashboard.GetDashboard(gauge.ParentDashboardID)
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
	recordRefs, getRecErr := recordReadController.GetFilteredSortedRecords(getRecordParams)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult := groupRecordsIntoSingleGroup(recordRefs)

	summaries := []values.ValSummary{}
	summaries = append(summaries, gauge.Properties.ValSummary)
	groupedSummarizedVals, summarizeErr := summarizeGroupedRecords(valGroupingResult, summaries)
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

func getGaugeData(params GetGaugeDataParams) (*GaugeData, error) {

	gauge, getErr := gauge.GetGauge(params.ParentDashboardID, params.GaugeID)
	if getErr != nil {
		return nil, fmt.Errorf("getGaugeData: Error retrieving gauge with params = %+v: error= %v",
			params, getErr)
	}

	gaugeData, dataErr := getOneGaugeData(gauge, params.FilterRules)
	if dataErr != nil {
		return nil, fmt.Errorf("getGaugeData: Error retrieving gauge data: %v", dataErr)
	}

	return gaugeData, nil

}

func getDefaultDashboardGaugesData(parentDashboardID string) ([]GaugeData, error) {

	gauges, err := gauge.GetGauges(parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("getDefaultDashboardGaugesData: Error retrieving gauges: %v", err)
	}

	var gaugesData []GaugeData
	for _, gauge := range gauges {

		gaugeData, dataErr := getOneGaugeData(&gauge, gauge.Properties.DefaultFilterRules)
		if dataErr != nil {
			return nil, fmt.Errorf("getDefaultDashboardGaugesData: Error retrieving gauge data: %v", dataErr)
		}
		gaugesData = append(gaugesData, *gaugeData)
	}

	return gaugesData, nil
}
