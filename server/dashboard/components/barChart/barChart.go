package barChart

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const barChartEntityKind string = "BarChart"

const xAxisSortAsc string = "asc"
const xAxisSortDesc string = "desc"

type BarChartUniqueIDHeader struct {
	BarChartID string `json:"barChartID"`
}

func (idHeader BarChartUniqueIDHeader) uniqueBarChartID() string {
	return idHeader.BarChartID
}

// DashboardBarChart is the datastore object for dashboard bar charts.
type BarChart struct {
	ParentDashboardID string `json:"parentDashboardID"`

	BarChartID string `json:"barChartID"`

	// DataSrcTable is the table the bar chart gets its data from
	DataSrcTableID string `json:"dataSrcTableID"`

	// XAxisVals is a grouping of field values displayed along the x axis of the bar chart.
	XAxisVals values.ValGrouping `json"xAxisVals"`

	// XAxisSortValues configures how values (bars) along the x axis are sorted. Options include
	// yAxisValAsc, yAxisValDesc, xAxisValAsc, xAxisValDesc. The default is xAxisValAsc.
	XAxisSortValues string `json:"xAxisSortValues"`

	Geometry common.LayoutGeometry `json:"geometry"`

	Title string `json:"title"`

	YAxisVals values.ValSummary `json:"yAxisValSummary"`
}

const barChartIDFieldName string = "BarChartID"
const barChartParentDashboardIDFieldName string = "ParentDashboardID"

type NewBarChartParams struct {
	ParentDashboardID string `json:"parentDashboardID"`

	DataSrcTableID string `json:"dataSrcTableID"`

	XAxisVals       values.NewValGroupingParams `json:"xAxisVals"`
	XAxisSortValues string                      `json:"xAxisSortValues"`

	YAxisVals values.NewValSummaryParams `json:"yAxisVals"`

	Geometry common.LayoutGeometry `json:"geometry"`
}

func updateExistingBarChart(appEngContext appengine.Context, barChartID string, updatedBarChart *BarChart) (*BarChart, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		barChartID, barChartEntityKind, barChartIDFieldName, updatedBarChart); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating text box: error = %v", updateErr)
	}

	return updatedBarChart, nil

}

func validBarChartSortXAxisProp(xAxisSortVal string) bool {
	if (xAxisSortVal == xAxisSortAsc) || (xAxisSortVal == xAxisSortDesc) {
		return true
	} else {
		return false
	}
}

func NewBarChart(appEngContext appengine.Context, params NewBarChartParams) (*BarChart, error) {

	valGrouping, valGroupingErr := values.NewValGrouping(appEngContext, params.XAxisVals)
	if valGroupingErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating new value grouping for bar chart: error = %v", valGroupingErr)
	}

	valSummary, valSummaryErr := values.NewValSummary(appEngContext, params.YAxisVals)
	if valSummaryErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating summary values for bar chart: error = %v", valSummaryErr)
	}

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("NewBarChart: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	if !validBarChartSortXAxisProp(params.XAxisSortValues) {
		return nil, fmt.Errorf("NewBarChart: Invalid X axis sort order: %v", params.XAxisSortValues)
	}

	newBarChart := BarChart{
		ParentDashboardID: params.ParentDashboardID,
		BarChartID:        uniqueID.GenerateUniqueID(),
		XAxisVals:         *valGrouping,
		XAxisSortValues:   params.XAxisSortValues,
		DataSrcTableID:    params.DataSrcTableID,
		YAxisVals:         *valSummary,
		Geometry:          params.Geometry,
		Title:             ""} // default to empty title

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, barChartEntityKind, &newBarChart)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new text box component: error inserting into datastore: %v", insertErr)
	}

	return &newBarChart, nil

}

func getBarChart(appEngContext appengine.Context, barChartID string) (*BarChart, error) {

	var barChart BarChart

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, barChartEntityKind,
		barChartIDFieldName, barChartID, &barChart); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to retrieve bar chart from datastore: error = %v", getErr)
	}

	return &barChart, nil

}
