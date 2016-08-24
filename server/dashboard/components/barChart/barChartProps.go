package barChart

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/dashboard/values"
)

// The BarChartPropertyUpdater interface along with UpdateBarChartProps() implement a harness for
// property updates. All property updates consiste of: (1) Retrieve the entity from the datastore,
// (2) Do the update on the entity itself, (3) Save the updated entity back to the datastore.
// Steps (1) and (3) can be done in a wrapper function UpdateBarChartProps, while only (2) needs
// be defined for each different property update. The goal is to minimize code bloat of property
// setting code and also make property updating code more uniform and less error prone.
type BarChartPropertyUpdater interface {
	uniqueBarChartID() string
	parentDashboardID() string
	updateBarChartProps(barChart *BarChart) error
}

type BarChartUniqueIDHeader struct {
	ParentDashboardID string `json:"parentDashboardID"`
	BarChartID        string `json:"barChartID"`
}

func (idHeader BarChartUniqueIDHeader) parentDashboardID() string {
	return idHeader.ParentDashboardID
}

func (idHeader BarChartUniqueIDHeader) uniqueBarChartID() string {
	return idHeader.BarChartID
}

func UpdateBarChartProps(propUpdater BarChartPropertyUpdater) (*BarChart, error) {

	// Retrieve the bar chart from the data store
	barChartForUpdate, getBarChartErr := getBarChart(propUpdater.parentDashboardID(), propUpdater.uniqueBarChartID())
	if getBarChartErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to get existing bar chart: %v", getBarChartErr)
	}

	// Do the actual update
	propUpdateErr := propUpdater.updateBarChartProps(barChartForUpdate)
	if propUpdateErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to update existing bar chart: %v", propUpdateErr)
	}

	// Save the updated bar chart back to the data store
	updatedBarChart, updateErr := updateExistingBarChart(barChartForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to update existing bar chart: %v", updateErr)
	}

	return updatedBarChart, nil

}

// Title Property

type SetBarChartTitleParams struct {
	// Embed a common header to reference the BarChart in the datastore. This header also supports
	// the niqueBarChartID() method to retrieve the unique ID. So, once decoded, the struct can be passed as an
	// BarChartPropertyUpdater interface to a generic/reusable function to process the property update.
	BarChartUniqueIDHeader
	NewTitle string `json:"newTitle"`
}

func (titleParam SetBarChartTitleParams) updateBarChartProps(barChart *BarChart) error {

	log.Printf("Updating bar chart title: %v", titleParam.NewTitle)

	barChart.Properties.Title = titleParam.NewTitle

	return nil
}

// Dimensions Property

type SetBarChartDimensionsParams struct {
	BarChartUniqueIDHeader
	Geometry common.LayoutGeometry `json:"geometry"`
}

func (params SetBarChartDimensionsParams) updateBarChartProps(barChart *BarChart) error {

	if !common.ValidGeometry(params.Geometry) {
		return fmt.Errorf("setBarChartDimensions: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	barChart.Properties.Geometry = params.Geometry

	return nil
}

type SetBarChartAvailableFilterParams struct {
	BarChartUniqueIDHeader
	AvailableFilterIDs []string `json:"availableFilterIDs"`
}

func (params SetBarChartAvailableFilterParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.AvailableFilterIDs = params.AvailableFilterIDs

	return nil
}

type SetBarChartDefaultFilterParams struct {
	BarChartUniqueIDHeader
	DefaultFilterIDs []string `json:"defaultFilterIDs"`
}

func (params SetBarChartDefaultFilterParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.DefaultFilterIDs = params.DefaultFilterIDs

	return nil
}

type SetXAxisValuesParams struct {
	BarChartUniqueIDHeader
	XAxisValueGrouping values.ValGrouping `json:"xAxisValueGrouping"`
}

func (params SetXAxisValuesParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.XAxisVals = params.XAxisValueGrouping

	return nil
}
