package dashboard

import (
	"appengine"
	"fmt"
	"resultra/datasheet/datamodel"
)

// The BarChartPropertyUpdater interface along with UpdateBarChartProps() implement a harness for
// property updates. All property updates consiste of: (1) Retrieve the entity from the datastore,
// (2) Do the update on the entity itself, (3) Save the updated entity back to the datastore.
// Steps (1) and (3) can be done in a wrapper function UpdateBarChartProps, while only (2) needs
// be defined for each different property update. The goal is to minimize code bloat of property
// setting code and also make property updating code more uniform and less error prone.
type BarChartPropertyUpdater interface {
	uniqueBarChartID() BarChartUniqueID
	updateBarChartProps(barChart *BarChart) error
}

func UpdateBarChartProps(appEngContext appengine.Context, propUpdater BarChartPropertyUpdater) (*BarChartRef, error) {

	// Retrieve the bar chart from the data store
	barChartForUpdate, getBarChartErr := getBarChart(appEngContext, propUpdater.uniqueBarChartID())
	if getBarChartErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to get existing bar chart: %v", getBarChartErr)
	}

	// Do the actual update
	propUpdateErr := propUpdater.updateBarChartProps(barChartForUpdate)
	if propUpdateErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to update existing bar chart: %v", propUpdateErr)
	}

	// Save the updated bar chart back to the data store
	barChartRef, updateErr := updateExistingBarChart(appEngContext, propUpdater.uniqueBarChartID(), barChartForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to update existing bar chart: %v", updateErr)
	}

	return barChartRef, nil

}

// Title Property

type SetBarChartTitleParams struct {
	// Embed a common header to reference the BarChart in the datastore. This header also supports
	// the niqueBarChartID() method to retrieve the unique ID. So, once decoded, the struct can be passed as an
	// BarChartPropertyUpdater interface to a generic/reusable function to process the property update.
	BarChartUniqueIDHeader
	Title string `json:"title"`
}

func (titleParam SetBarChartTitleParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Title = titleParam.Title

	return nil
}

// Dimensions Property

type SetBarChartDimensionsParams struct {
	BarChartUniqueIDHeader
	Geometry datamodel.LayoutGeometry `json:"geometry"`
}

func (params SetBarChartDimensionsParams) updateBarChartProps(barChart *BarChart) error {

	if !datamodel.ValidGeometry(params.Geometry) {
		return fmt.Errorf("setBarChartDimensions: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	barChart.Geometry = params.Geometry

	return nil
}
