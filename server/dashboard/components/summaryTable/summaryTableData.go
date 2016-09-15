package summaryTable

import (
	"fmt"
)

type SummaryTableData struct {
	SummaryTableID string       `json:"summaryTableID"`
	SummaryTable   SummaryTable `json:"summaryTable"`
}

func getOneSummaryTableData(summaryTable *SummaryTable) (*SummaryTableData, error) {

	summaryTableData := SummaryTableData{
		SummaryTableID: summaryTable.SummaryTableID,
		SummaryTable:   *summaryTable}

	// TODO - Retrieve data for table

	return &summaryTableData, nil
}

type GetSummaryTableDataParams struct {
	ParentDashboardID string `json:"parentDashboardID"`
	SummaryTableID    string `json:"summaryTableID"`
}

func GetSummaryTableData(params GetSummaryTableDataParams) (*SummaryTableData, error) {

	if len(params.SummaryTableID) <= 0 {
		return nil, fmt.Errorf("GetSummaryTableData: missing summary table ID")
	}

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("GetSummaryTableData: missing dashboard ID")
	}

	summaryTable, getSummaryTableErr := getSummaryTable(params.ParentDashboardID, params.SummaryTableID)
	if getSummaryTableErr != nil {
		return nil, fmt.Errorf("GetSummaryTableData: Error retrieving summary table with params=%+v: error= %v",
			params, getSummaryTableErr)
	}

	summaryTableData, dataErr := getOneSummaryTableData(summaryTable)
	if dataErr != nil {
		return nil, fmt.Errorf("GetSummaryTableData: Error retrieving bar chart data: %v", dataErr)
	}

	return summaryTableData, nil

}

func GetDashboardSummaryTablesData(parentDashboardID string) ([]SummaryTableData, error) {

	summaryTables, err := getSummaryTables(parentDashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardBarChartsData: Error retrieving bar charts: %v", err)
	}

	summaryTablesData := []SummaryTableData{}
	for _, summaryTable := range summaryTables {

		summaryTableData, dataErr := getOneSummaryTableData(&summaryTable)
		if dataErr != nil {
			return nil, fmt.Errorf("GetData: Error retrieving summary table data: %v", dataErr)
		}
		summaryTablesData = append(summaryTablesData, *summaryTableData)
	}

	return summaryTablesData, nil
}
