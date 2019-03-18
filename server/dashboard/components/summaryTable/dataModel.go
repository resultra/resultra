// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package summaryTable

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/dashboard/components/common"
	"resultra/tracker/server/dashboard/values"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
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

	RowGroupingVals values.ValGrouping `json:"rowGroupingVals"`

	ColumnValSummaries []values.NewValSummaryParams `json:"columnValSummaries"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func saveSummaryTable(destDBHandle *sql.DB, newSummaryTable SummaryTable) error {

	if saveErr := common.SaveNewDashboardComponent(destDBHandle, summaryTableEntityKind,
		newSummaryTable.ParentDashboardID, newSummaryTable.SummaryTableID, newSummaryTable.Properties); saveErr != nil {
		return fmt.Errorf("newSummaryTable: Unable to save summary table component: error = %v", saveErr)
	}
	return nil

}

func newSummaryTable(trackerDBHandle *sql.DB, params NewSummaryTableParams) (*SummaryTable, error) {

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("newSummaryTable: Error creating summary table: missing parent dashboard ID")
	}

	rowGroupingErr := values.ValidateValGrouping(trackerDBHandle, params.RowGroupingVals)
	if rowGroupingErr != nil {
		return nil, fmt.Errorf("newSummaryTable: Error creating new value grouping for bar chart: error = %v", rowGroupingErr)
	}

	colSummaries := []values.ValSummary{}
	for _, currColSummary := range params.ColumnValSummaries {
		colSummary, colSummaryErr := values.NewValSummary(trackerDBHandle, currColSummary)
		if colSummaryErr != nil {
			return nil, fmt.Errorf("newSummaryTable: Error creating summary values for bar chart: error = %v", colSummaryErr)
		}
		colSummaries = append(colSummaries, *colSummary)
	}

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("newSummaryTable: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	summaryTableProps := newDefaultSummaryTableProps()
	summaryTableProps.RowGroupingVals = params.RowGroupingVals
	summaryTableProps.ColumnValSummaries = colSummaries
	summaryTableProps.Geometry = params.Geometry

	newSummaryTable := SummaryTable{
		ParentDashboardID: params.ParentDashboardID,
		SummaryTableID:    uniqueID.GenerateUniqueID(),
		Properties:        summaryTableProps}

	if saveErr := saveSummaryTable(trackerDBHandle, newSummaryTable); saveErr != nil {
		return nil, fmt.Errorf("newSummaryTable: Unable to save summary component with params=%+v: error = %v", params, saveErr)
	}

	return &newSummaryTable, nil
}

func GetSummaryTable(trackerDBHandle *sql.DB, parentDashboardID string, summaryTableID string) (*SummaryTable, error) {

	summaryTableProps := newDefaultSummaryTableProps()
	if getErr := common.GetDashboardComponent(trackerDBHandle,
		summaryTableEntityKind, parentDashboardID, summaryTableID, &summaryTableProps); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to retrieve bar chart component: %v", getErr)
	}

	summaryTable := SummaryTable{
		ParentDashboardID: parentDashboardID,
		SummaryTableID:    summaryTableID,
		Properties:        summaryTableProps}

	return &summaryTable, nil

}

func getSummaryTablesFromSrc(srcDBHandle *sql.DB, parentDashboardID string) ([]SummaryTable, error) {

	summaryTables := []SummaryTable{}
	addSummaryTable := func(summaryTableID string, encodedProps string) error {

		summaryTableProps := newDefaultSummaryTableProps()
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
	if getErr := common.GetDashboardComponents(srcDBHandle, summaryTableEntityKind, parentDashboardID, addSummaryTable); getErr != nil {
		return nil, fmt.Errorf("getBarCharts: Can't get bar chart components: %v")
	}

	return summaryTables, nil
}

func GetSummaryTables(trackerDBHandle *sql.DB, parentDashboardID string) ([]SummaryTable, error) {
	return getSummaryTablesFromSrc(trackerDBHandle, parentDashboardID)
}

func CloneSummaryTables(cloneParams *trackerDatabase.CloneDatabaseParams, srcParentDashboardID string) error {

	remappedDashboardID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneSummaryTables: %v", err)
	}

	summaryTables, err := getSummaryTablesFromSrc(cloneParams.SrcDBHandle, srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneSummaryTables: %v", err)
	}

	for _, srcSummaryTable := range summaryTables {

		remappedSummaryTableID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcSummaryTable.SummaryTableID)

		clonedProps, err := srcSummaryTable.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneSummaryTables: %v", err)
		}

		destSummaryTable := SummaryTable{
			ParentDashboardID: remappedDashboardID,
			SummaryTableID:    remappedSummaryTableID,
			Properties:        *clonedProps}

		if err := saveSummaryTable(cloneParams.DestDBHandle, destSummaryTable); err != nil {
			return fmt.Errorf("CloneSummaryTables: %v", err)
		}
	}

	return nil
}

func updateExistingSummaryTable(trackerDBHandle *sql.DB, updatedSummaryTable *SummaryTable) (*SummaryTable, error) {

	if updateErr := common.UpdateDashboardComponent(trackerDBHandle,
		summaryTableEntityKind, updatedSummaryTable.ParentDashboardID,
		updatedSummaryTable.SummaryTableID, updatedSummaryTable.Properties); updateErr != nil {
		return nil, fmt.Errorf("Error updating summary table %+v: %v", updatedSummaryTable, updateErr)
	}

	return updatedSummaryTable, nil

}
