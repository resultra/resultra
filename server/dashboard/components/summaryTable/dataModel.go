package summaryTable

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/components/common"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

const summaryTableEntityKind string = "SummaryTable"

// DashboardBarChart is the datastore object for dashboard bar charts.
type SummaryTable struct {
	ParentDashboardID string `json:"parentDashboardID"`

	SummaryTableID string `json:"summaryTableID"`

	// DataSrcTable is the table the bar chart gets its data from
	Properties SummaryTableProps `json:"properties"`
}

type NewSummaryTableParams struct {
	ParentDashboardID string `json:"parentDashboardID"`

	DataSrcTableID string `json:"dataSrcTableID"`

	RowGroupingVals values.NewValGroupingParams `json:"rowGroupingVals"`

	ColumnValSummaries []values.NewValSummaryParams `json:"columnValSummaries"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func saveSummaryTable(newSummaryTable SummaryTable) error {

	if saveErr := common.SaveNewDashboardComponent(summaryTableEntityKind,
		newSummaryTable.ParentDashboardID, newSummaryTable.SummaryTableID, newSummaryTable.Properties); saveErr != nil {
		return fmt.Errorf("newSummaryTable: Unable to save summary table component: error = %v", saveErr)
	}
	return nil

}

func newSummaryTable(params NewSummaryTableParams) (*SummaryTable, error) {

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("newSummaryTable: Error creating summary table: missing parent dashboard ID")
	}

	if len(params.DataSrcTableID) <= 0 {
		return nil, fmt.Errorf("newSummaryTable: Error creating summary table: missing table ID")
	}

	rowGrouping, rowGroupingErr := values.NewValGrouping(params.RowGroupingVals)
	if rowGroupingErr != nil {
		return nil, fmt.Errorf("newSummaryTable: Error creating new value grouping for bar chart: error = %v", rowGroupingErr)
	}

	colSummaries := []values.ValSummary{}
	for _, currColSummary := range params.ColumnValSummaries {
		colSummary, colSummaryErr := values.NewValSummary(currColSummary)
		if colSummaryErr != nil {
			return nil, fmt.Errorf("newSummaryTable: Error creating summary values for bar chart: error = %v", colSummaryErr)
		}
		colSummaries = append(colSummaries, *colSummary)
	}

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("newSummaryTable: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	summaryTableProps := SummaryTableProps{
		RowGroupingVals:    *rowGrouping,
		ColumnValSummaries: colSummaries,
		DataSrcTableID:     params.DataSrcTableID,
		Geometry:           params.Geometry,
		Title:              "",
		DefaultFilterRules: []recordFilter.RecordFilterRule{}}

	newSummaryTable := SummaryTable{
		ParentDashboardID: params.ParentDashboardID,
		SummaryTableID:    uniqueID.GenerateSnowflakeID(),
		Properties:        summaryTableProps}

	if saveErr := saveSummaryTable(newSummaryTable); saveErr != nil {
		return nil, fmt.Errorf("newSummaryTable: Unable to save summary component with params=%+v: error = %v", params, saveErr)
	}

	return &newSummaryTable, nil
}

func GetSummaryTable(parentDashboardID string, summaryTableID string) (*SummaryTable, error) {

	summaryTableProps := SummaryTableProps{}
	if getErr := common.GetDashboardComponent(summaryTableEntityKind, parentDashboardID, summaryTableID, &summaryTableProps); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to retrieve bar chart component: %v", getErr)
	}

	summaryTable := SummaryTable{
		ParentDashboardID: parentDashboardID,
		SummaryTableID:    summaryTableID,
		Properties:        summaryTableProps}

	return &summaryTable, nil

}

func GetSummaryTables(parentDashboardID string) ([]SummaryTable, error) {

	summaryTables := []SummaryTable{}
	addSummaryTable := func(summaryTableID string, encodedProps string) error {

		var summaryTableProps SummaryTableProps
		if decodeErr := generic.DecodeJSONString(encodedProps, &summaryTableProps); decodeErr != nil {
			return fmt.Errorf("GetSummaryTables: can't decode properties: %v", encodedProps)
		}

		currSummaryTable := SummaryTable{
			ParentDashboardID: parentDashboardID,
			SummaryTableID:    summaryTableID,
			Properties:        summaryTableProps}

		summaryTables = append(summaryTables, currSummaryTable)

		return nil
	}
	if getErr := common.GetDashboardComponents(summaryTableEntityKind, parentDashboardID, addSummaryTable); getErr != nil {
		return nil, fmt.Errorf("getBarCharts: Can't get bar chart components: %v")
	}

	return summaryTables, nil
}

func CloneSummaryTables(remappedIDs uniqueID.UniqueIDRemapper, srcParentDashboardID string) error {

	remappedDashboardID, err := remappedIDs.GetExistingRemappedID(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneSummaryTables: %v", err)
	}

	summaryTables, err := GetSummaryTables(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneSummaryTables: %v", err)
	}

	for _, srcSummaryTable := range summaryTables {

		remappedSummaryTableID, err := remappedIDs.AllocNewRemappedID(srcSummaryTable.SummaryTableID)
		if err != nil {
			return fmt.Errorf("CloneSummaryTables: %v", err)
		}

		clonedProps, err := srcSummaryTable.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneSummaryTables: %v", err)
		}

		destSummaryTable := SummaryTable{
			ParentDashboardID: remappedDashboardID,
			SummaryTableID:    remappedSummaryTableID,
			Properties:        *clonedProps}

		if err := saveSummaryTable(destSummaryTable); err != nil {
			return fmt.Errorf("CloneSummaryTables: %v", err)
		}
	}

	return nil
}

func updateExistingSummaryTable(updatedSummaryTable *SummaryTable) (*SummaryTable, error) {

	if updateErr := common.UpdateDashboardComponent(summaryTableEntityKind, updatedSummaryTable.ParentDashboardID,
		updatedSummaryTable.SummaryTableID, updatedSummaryTable.Properties); updateErr != nil {
		return nil, fmt.Errorf("Error updating summary table %+v: %v", updatedSummaryTable, updateErr)
	}

	return updatedSummaryTable, nil

}
