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
	ParentDashboardID string                   `json:"parentDashboardID"`
	BarChartID        string                   `json:"barChartID"`
	Geometry          datamodel.LayoutGeometry `json:"geometry"`

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

// Update the properties to match those found in updatedProps
func UpdateBarChartProps(appEngContext appengine.Context, updatedProps BarChartRef) (*BarChartRef, error) {

	valGrouping, valGroupingRef, valGroupingErr := NewValGroupingFromRef(appEngContext, updatedProps.XAxisValsRef)
	if valGroupingErr != nil {
		return nil, fmt.Errorf("UpdateBarChartProperties: Error creating new value grouping for bar chart: error = %v", valGroupingErr)
	}

	valSummary, valSummaryRef, valSummaryErr := NewValSummaryFromRef(appEngContext, updatedProps.YAxisValRef)
	if valSummaryErr != nil {
		return nil, fmt.Errorf("UpdateBarChartProperties: Error creating summary values for bar chart: error = %v", valSummaryErr)
	}

	if !datamodel.ValidGeometry(updatedProps.Geometry) {
		return nil, fmt.Errorf("UpdateBarChartProperties: Invalid geometry for bar chart: %+v", updatedProps.Geometry)
	}

	if !validBarChartSortXAxisProp(updatedProps.XAxisSortValues) {
		return nil, fmt.Errorf("UpdateBarChartProperties: Invalid X axis sort order: %v", updatedProps.XAxisSortValues)
	}

	updatedBarChart := BarChart{
		XAxisVals:       *valGrouping,
		XAxisSortValues: updatedProps.XAxisSortValues,
		YAxisVals:       *valSummary,
		Geometry:        updatedProps.Geometry}

	parentDashboardKey, getDashboardErr := datamodel.GetExistingRootEntityKey(appEngContext, dashboardEntityKind,
		updatedProps.ParentDashboardID)
	if getDashboardErr != nil {
		return nil, fmt.Errorf("UpdateBarChartProperties: Invalid dashboard for bar chart: %v", getDashboardErr)
	}

	if updateErr := datamodel.UpdateExistingEntity(appEngContext,
		updatedProps.BarChartID, barChartEntityKind,
		parentDashboardKey, &updatedBarChart); updateErr != nil {
		return nil, updateErr
	}

	barChartRef := BarChartRef{
		ParentDashboardID: updatedProps.ParentDashboardID,
		BarChartID:        updatedProps.BarChartID,
		Geometry:          updatedBarChart.Geometry,
		XAxisValsRef:      *valGroupingRef,
		XAxisSortValues:   updatedBarChart.XAxisSortValues,
		YAxisValRef:       *valSummaryRef,
	}

	return &barChartRef, nil
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
		Geometry:        params.Geometry}

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

type GetBarChartParams struct {
	ParentDashboardID string `json:"parentDashboardID"`
	BarChartID        string `json:"barChartID"`
}

func getBarChart(appEngContext appengine.Context, params GetBarChartParams) (*BarChart, error) {

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

func getBarChartAndRef(appEngContext appengine.Context, params GetBarChartParams) (*BarChart, *BarChartRef, error) {

	barChart, getBarChartErr := getBarChart(appEngContext, params)
	if getBarChartErr != nil {
		return nil, nil, fmt.Errorf("GetBarChartRef: Error getting bar chart reference: %v", getBarChartErr)
	}

	xAxisValsRef, groupingErr := barChart.XAxisVals.GetValGroupingRef(appEngContext)
	if groupingErr != nil {
		return nil, nil, fmt.Errorf("GetBarChartRef: Error getting bar chart reference: %v", groupingErr)
	}

	yAxisValsRef, summaryRef := barChart.YAxisVals.GetValSummaryRef(appEngContext)
	if summaryRef != nil {
		return nil, nil, fmt.Errorf("GetBarChartRef: Error getting bar chart reference: %v", summaryRef)
	}

	barChartRef := BarChartRef{
		ParentDashboardID: params.ParentDashboardID,
		BarChartID:        params.BarChartID,
		Geometry:          barChart.Geometry,
		XAxisValsRef:      *xAxisValsRef,
		XAxisSortValues:   barChart.XAxisSortValues,
		YAxisValRef:       *yAxisValsRef}

	return barChart, &barChartRef, nil
}

type BarChartDataRow struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type BarChartData struct {
	BarChartID  string            `json:"barChartID"`
	BarChartRef BarChartRef       `json:"barChartRef"`
	Title       string            `json:"title"`
	XAxisTitle  string            `json:"xAxisTitle"`
	YAxisTitle  string            `json:"yAxisTitle"`
	DataRows    []BarChartDataRow `json:"dataRows"`
}

func GetBarChartData(appEngContext appengine.Context, params GetBarChartParams) (*BarChartData, error) {

	barChart, barChartRef, getBarChartErr := getBarChartAndRef(appEngContext, params)
	if getBarChartErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving bar chart: %v", getBarChartErr)
	}

	recordRefs, getRecErr := datamodel.GetFilteredRecords(appEngContext)
	if getRecErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error retrieving records for bar chart: %v", getRecErr)
	}

	valGroupingResult, groupingErr := barChart.XAxisVals.groupRecords(appEngContext, recordRefs)
	if groupingErr != nil {
		return nil, fmt.Errorf("GetBarChartData: Error grouping records for bar chart: %v", groupingErr)
	}

	dataRows := []BarChartDataRow{}
	for _, valGroup := range valGroupingResult.valGroups {
		dataRows = append(dataRows,
			BarChartDataRow{valGroup.groupLabel, float64(len(valGroup.recordsInGroup))})
	}

	barChartData := BarChartData{
		BarChartID:  params.BarChartID,
		BarChartRef: *barChartRef,
		Title:       "Chart Title TBD", // not implemented yet
		XAxisTitle:  valGroupingResult.groupingLabel,
		YAxisTitle:  "Count",
		DataRows:    dataRows}

	return &barChartData, nil

}
