package summaryTable

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/values"
)

// The BarChartPropertyUpdater interface along with UpdateBarChartProps() implement a harness for
// property updates. All property updates consiste of: (1) Retrieve the entity from the datastore,
// (2) Do the update on the entity itself, (3) Save the updated entity back to the datastore.
// Steps (1) and (3) can be done in a wrapper function UpdateBarChartProps, while only (2) needs
// be defined for each different property update. The goal is to minimize code bloat of property
// setting code and also make property updating code more uniform and less error prone.
type SummaryTablePropertyUpdater interface {
	uniqueSummaryTableID() string
	parentDashboardID() string
	updateSummaryTableProps(summaryTable *SummaryTable) error
}

type SummaryTableUniqueIDHeader struct {
	ParentDashboardID string `json:"parentDashboardID"`
	SummaryTableID    string `json:"summaryTableID"`
}

func (idHeader SummaryTableUniqueIDHeader) parentDashboardID() string {
	return idHeader.ParentDashboardID
}

func (idHeader SummaryTableUniqueIDHeader) uniqueSummaryTableID() string {
	return idHeader.SummaryTableID
}

func updateSummaryTableProps(propUpdater SummaryTablePropertyUpdater) (*SummaryTable, error) {

	// Retrieve the bar chart from the data store
	summaryTableForUpdate, getErr := GetSummaryTable(propUpdater.parentDashboardID(), propUpdater.uniqueSummaryTableID())
	if getErr != nil {
		return nil, fmt.Errorf("updateSummaryTableProps: Unable to get existing summary table: %v", getErr)
	}

	// Do the actual update
	propUpdateErr := propUpdater.updateSummaryTableProps(summaryTableForUpdate)
	if propUpdateErr != nil {
		return nil, fmt.Errorf("updateSummaryTableProps: Unable to update existing summary table: %v", propUpdateErr)
	}

	// Save the updated bar chart back to the data store
	updatedSummaryTable, updateErr := updateExistingSummaryTable(summaryTableForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateSummaryTableProps: Unable to update existing bar chart: %v", updateErr)
	}

	return updatedSummaryTable, nil

}

// Title Property

type SetSummaryTableTitleParams struct {
	// Embed a common header to reference the BarChart in the datastore. This header also supports
	// the niqueBarChartID() method to retrieve the unique ID. So, once decoded, the struct can be passed as an
	// BarChartPropertyUpdater interface to a generic/reusable function to process the property update.
	SummaryTableUniqueIDHeader
	NewTitle string `json:"newTitle"`
}

func (titleParam SetSummaryTableTitleParams) updateSummaryTableProps(summaryTable *SummaryTable) error {

	log.Printf("Updating summary table title: %v", titleParam.NewTitle)

	summaryTable.Properties.Title = titleParam.NewTitle

	return nil
}

// Dimensions Property

type SetSummaryTableDimensionsParams struct {
	SummaryTableUniqueIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (params SetSummaryTableDimensionsParams) updateSummaryTableProps(summaryTable *SummaryTable) error {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return fmt.Errorf("setBarChartDimensions: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	summaryTable.Properties.Geometry = params.Geometry

	return nil
}

type SetSummaryTableAvailableFilterParams struct {
	SummaryTableUniqueIDHeader
	AvailableFilterIDs []string `json:"availableFilterIDs"`
}

func (params SetSummaryTableAvailableFilterParams) updateSummaryTableProps(summaryTable *SummaryTable) error {

	summaryTable.Properties.AvailableFilterIDs = params.AvailableFilterIDs

	return nil
}

type SetSummaryTableDefaultFilterParams struct {
	SummaryTableUniqueIDHeader
	DefaultFilterIDs []string `json:"defaultFilterIDs"`
}

func (params SetSummaryTableDefaultFilterParams) updateSummaryTableProps(summaryTable *SummaryTable) error {

	summaryTable.Properties.DefaultFilterIDs = params.DefaultFilterIDs

	return nil
}

type SetSummaryTableSummaryColumns struct {
	SummaryTableUniqueIDHeader
	ColumnValSummaries []values.ValSummary `json:"columnValSummaries"`
}

func (params SetSummaryTableSummaryColumns) updateSummaryTableProps(summaryTable *SummaryTable) error {

	summaryTable.Properties.ColumnValSummaries = params.ColumnValSummaries

	return nil
}

type SetRowGroupingParams struct {
	SummaryTableUniqueIDHeader
	RowValueGrouping values.ValGrouping `json:"rowValueGrouping"`
}

func (params SetRowGroupingParams) updateSummaryTableProps(summaryTable *SummaryTable) error {

	summaryTable.Properties.RowGroupingVals = params.RowValueGrouping

	return nil
}
