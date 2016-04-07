package dashboard

import (
	"appengine"
	"fmt"
	"resultra/datasheet/datamodel"
)

const barChartEntityKind string = "DashboardBarChart"

const xAxisSortAsc string = "asc"
const xAxisSortDesc string = "desc"

type BarChartUniqueID struct {
	ParentDashboardID string `json:"parentDashboardID"`
	BarChartID        string `json:"barChartID"`
}

// This BarChartUniqueIDHeader and its uniqueBarChartID() method are intended to
// be embedded in other structs for purposes of having a standard header on structs
// which are received as parameters to update properties or to return information or
// data related to the bar chart.
type BarChartUniqueIDHeader struct {
	UniqueID BarChartUniqueID `json:"uniqueID"`
}

func (idHeader BarChartUniqueIDHeader) uniqueBarChartID() BarChartUniqueID {
	return idHeader.UniqueID
}

// DashboardBarChart is the datastore object for dashboard bar charts.
type BarChart struct {
	// TODO: Include table selection, filter settings

	// XAxisVals is a grouping of field values displayed along the x axis of the bar chart.
	XAxisVals ValGrouping

	// XAxisSortValues configures how values (bars) along the x axis are sorted. Options include
	// yAxisValAsc, yAxisValDesc, xAxisValAsc, xAxisValDesc. The default is xAxisValAsc.
	XAxisSortValues string

	Geometry datamodel.LayoutGeometry

	Title string

	YAxisVals ValSummary
}

// BarChartRef is the opaque API reference object for dashboard bar charts.
type BarChartRef struct {
	ParentDashboardID string                   `json:"parentDashboardID"`
	BarChartID        string                   `json:"barChartID"`
	Geometry          datamodel.LayoutGeometry `json:"geometry"`

	XAxisValsRef    ValGroupingRef `json:"xAxisValsRef"`
	XAxisSortValues string         `json:"xAxisSortValues"`

	YAxisValRef ValSummaryRef `json:"yAxisValSummary"`
	Title       string        `json:"title"`
}

type NewBarChartParams struct {
	ParentDashboardID string `json:"parentDashboardID"`

	// PlaceholderID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	PlaceholderID string `json:"placeholderID"`

	XAxisVals       NewValGroupingParams `json:"xAxisVals"`
	XAxisSortValues string               `json:"xAxisSortValues"`

	YAxisVals NewValSummaryParams `json:"yAxisVals"`

	Geometry datamodel.LayoutGeometry `json:"geometry"`
}

func updateExistingBarChart(appEngContext appengine.Context, uniqueID BarChartUniqueID, updatedBarChart *BarChart) (*BarChartRef, error) {

	parentDashboardKey, getDashboardErr := datamodel.GetExistingRootEntityKey(appEngContext, dashboardEntityKind,
		uniqueID.ParentDashboardID)
	if getDashboardErr != nil {
		return nil, fmt.Errorf("updateExistingBarChart: Invalid dashboard for bar chart: %v", getDashboardErr)
	}

	if updateErr := datamodel.UpdateExistingEntity(appEngContext,
		uniqueID.BarChartID, barChartEntityKind,
		parentDashboardKey, updatedBarChart); updateErr != nil {
		return nil, updateErr
	}

	updatedBarChartRef, refErr := getBarChartRef(appEngContext, uniqueID, updatedBarChart)
	if refErr != nil {
		return nil, fmt.Errorf("updateExistingBarChart: Invalid dashboard for bar chart: %v", refErr)
	}

	return updatedBarChartRef, nil

}

func validBarChartSortXAxisProp(xAxisSortVal string) bool {
	if (xAxisSortVal == xAxisSortAsc) || (xAxisSortVal == xAxisSortDesc) {
		return true
	} else {
		return false
	}
}

func NewBarChart(appEngContext appengine.Context, params NewBarChartParams) (*BarChartRef, error) {

	valGrouping, valGroupingRef, valGroupingErr := NewValGrouping(appEngContext, params.XAxisVals)
	if valGroupingErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating new value grouping for bar chart: error = %v", valGroupingErr)
	}

	valSummary, valSummaryRef, valSummaryErr := NewValSummary(appEngContext, params.YAxisVals)
	if valSummaryErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating summary values for bar chart: error = %v", valSummaryErr)
	}

	if !datamodel.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("NewBarChart: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	if !validBarChartSortXAxisProp(params.XAxisSortValues) {
		return nil, fmt.Errorf("NewBarChart: Invalid X axis sort order: %v", params.XAxisSortValues)
	}

	parentDashboardKey, getDashboardErr := datamodel.GetExistingRootEntityKey(appEngContext, dashboardEntityKind,
		params.ParentDashboardID)
	if getDashboardErr != nil {
		return nil, fmt.Errorf("NewBarChart: Invalid dashboard for bar chart: %v", getDashboardErr)
	}

	newBarChart := BarChart{
		XAxisVals:       *valGrouping,
		XAxisSortValues: params.XAxisSortValues,
		YAxisVals:       *valSummary,
		Geometry:        params.Geometry,
		Title:           ""} // default to empty title

	barChartID, insertErr := datamodel.InsertNewEntity(appEngContext, barChartEntityKind,
		parentDashboardKey, &newBarChart)
	if insertErr != nil {
		return nil, fmt.Errorf("NewBarChart: Unable to create new bar chart: %v", insertErr)
	}

	barChartRef := BarChartRef{
		ParentDashboardID: params.ParentDashboardID,
		BarChartID:        barChartID,
		Geometry:          params.Geometry,
		XAxisValsRef:      *valGroupingRef,
		XAxisSortValues:   params.XAxisSortValues,
		YAxisValRef:       *valSummaryRef,
	}

	return &barChartRef, nil

}

func getBarChart(appEngContext appengine.Context, params BarChartUniqueID) (*BarChart, error) {

	parentDashboardKey, dashboardKeyErr := datamodel.NewRootEntityKey(appEngContext,
		dashboardEntityKind, params.ParentDashboardID)
	if dashboardKeyErr != nil {
		return nil, fmt.Errorf("getBarChart: unable to retrieve parent dashboard key for dashboard = %v", params.ParentDashboardID)
	}

	var barChart BarChart
	getErr := datamodel.GetChildEntityByID(params.BarChartID, appEngContext, barChartEntityKind,
		parentDashboardKey, &barChart)
	if getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to get bar chart from datastore: error = %v", getErr)
	}

	return &barChart, nil

}

func getBarChartRef(appEngContext appengine.Context, params BarChartUniqueID, barChart *BarChart) (*BarChartRef, error) {

	xAxisValsRef, groupingErr := barChart.XAxisVals.GetValGroupingRef(appEngContext)
	if groupingErr != nil {
		return nil, fmt.Errorf("GetBarChartRef: Error getting bar chart reference: %v", groupingErr)
	}

	yAxisValsRef, summaryRef := barChart.YAxisVals.GetValSummaryRef(appEngContext)
	if summaryRef != nil {
		return nil, fmt.Errorf("GetBarChartRef: Error getting bar chart reference: %v", summaryRef)
	}

	barChartRef := BarChartRef{
		ParentDashboardID: params.ParentDashboardID,
		BarChartID:        params.BarChartID,
		Geometry:          barChart.Geometry,
		XAxisValsRef:      *xAxisValsRef,
		XAxisSortValues:   barChart.XAxisSortValues,
		YAxisValRef:       *yAxisValsRef,
		Title:             barChart.Title}

	return &barChartRef, nil
}
