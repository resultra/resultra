package dashboard

import (
	"appengine"
	"fmt"
	"resultra/datasheet/datamodel"
)

const barChartEntityKind string = "DashboardBarChart"

const xAxisSortAsc string = "asc"
const xAxisSortDesc string = "desc"

// DashboardBarChart is the datastore object for dashboard bar charts.
type BarChart struct {
	// TODO: Include table selection, filter settings

	// XAxisVals is a grouping of field values displayed along the x axis of the bar chart.
	XAxisVals ValGrouping

	// XAxisSortValues configures how values (bars) along the x axis are sorted. Options include
	// yAxisValAsc, yAxisValDesc, xAxisValAsc, xAxisValDesc. The default is xAxisValAsc.
	XAxisSortValues string

	Geometry datamodel.LayoutGeometry

	YAxisVals ValSummary
}

// BarChartRef is the opaque API reference object for dashboard bar charts.
type BarChartRef struct {
	BarChartID string                   `json:"barChartID"`
	Geometry   datamodel.LayoutGeometry `json:"geometry"`

	XAxisValsRef    ValGroupingRef `json:"xAxisValsRef"`
	XAxisSortValues string         `json:"xAxisSortValues"`

	YAxisValRef ValSummaryRef `json:"yAxisValSummary"`
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

	if !((params.XAxisSortValues == xAxisSortAsc) || (params.XAxisSortValues == xAxisSortDesc)) {
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
		Geometry:        params.Geometry}

	barChartID, insertErr := datamodel.InsertNewEntity(appEngContext, barChartEntityKind,
		parentDashboardKey, &newBarChart)
	if insertErr != nil {
		return nil, fmt.Errorf("NewBarChart: Unable to create new bar chart: %v", insertErr)
	}

	barChartRef := BarChartRef{
		BarChartID:      barChartID,
		Geometry:        params.Geometry,
		XAxisValsRef:    *valGroupingRef,
		XAxisSortValues: params.XAxisSortValues,
		YAxisValRef:     *valSummaryRef,
	}

	return &barChartRef, fmt.Errorf("NewBarChart: Not implemented yet")

}
