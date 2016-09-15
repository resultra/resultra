package barChart

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/components/common"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const barChartEntityKind string = "BarChart"

const xAxisSortAsc string = "asc"
const xAxisSortDesc string = "desc"

type BarChartProps struct {

	// DataSrcTable is the table the bar chart gets its data from
	DataSrcTableID string `json:"dataSrcTableID"`

	// XAxisVals is a grouping of field values displayed along the x axis of the bar chart.
	XAxisVals values.ValGrouping `json:"xAxisVals"`

	// XAxisSortValues configures how values (bars) along the x axis are sorted. Options include
	// yAxisValAsc, yAxisValDesc, xAxisValAsc, xAxisValDesc. The default is xAxisValAsc.
	XAxisSortValues string `json:"xAxisSortValues"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`

	Title string `json:"title"`

	YAxisVals values.ValSummary `json:"yAxisValSummary"`

	AvailableFilterIDs []string `json:"availableFilterIDs"`
	DefaultFilterIDs   []string `json:"defaultFilterIDs"`
}

// DashboardBarChart is the datastore object for dashboard bar charts.
type BarChart struct {
	ParentDashboardID string `json:"parentDashboardID"`

	BarChartID string `json:"barChartID"`

	// DataSrcTable is the table the bar chart gets its data from
	Properties BarChartProps `json:"properties"`
}

const barChartIDFieldName string = "BarChartID"
const barChartParentDashboardIDFieldName string = "ParentDashboardID"

type NewBarChartParams struct {
	ParentDashboardID string `json:"parentDashboardID"`

	DataSrcTableID string `json:"dataSrcTableID"`

	XAxisVals       values.NewValGroupingParams `json:"xAxisVals"`
	XAxisSortValues string                      `json:"xAxisSortValues"`

	YAxisVals values.NewValSummaryParams `json:"yAxisVals"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func updateExistingBarChart(updatedBarChart *BarChart) (*BarChart, error) {

	if updateErr := common.UpdateDashboardComponent(barChartEntityKind, updatedBarChart.ParentDashboardID,
		updatedBarChart.BarChartID, updatedBarChart.Properties); updateErr != nil {
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

func NewBarChart(params NewBarChartParams) (*BarChart, error) {

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("newSummaryTable: Error creating bar chart: missing parent dashboard ID")
	}

	if len(params.DataSrcTableID) <= 0 {
		return nil, fmt.Errorf("newSummaryTable: Error creating bar chart: missing table ID")
	}

	valGrouping, valGroupingErr := values.NewValGrouping(params.XAxisVals)
	if valGroupingErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating new value grouping for bar chart: error = %v", valGroupingErr)
	}

	valSummary, valSummaryErr := values.NewValSummary(params.YAxisVals)
	if valSummaryErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating summary values for bar chart: error = %v", valSummaryErr)
	}

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("NewBarChart: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	if !validBarChartSortXAxisProp(params.XAxisSortValues) {
		return nil, fmt.Errorf("NewBarChart: Invalid X axis sort order: %v", params.XAxisSortValues)
	}

	barChartProps := BarChartProps{
		XAxisVals:          *valGrouping,
		XAxisSortValues:    params.XAxisSortValues,
		DataSrcTableID:     params.DataSrcTableID,
		YAxisVals:          *valSummary,
		Geometry:           params.Geometry,
		Title:              "", // optional
		AvailableFilterIDs: []string{},
		DefaultFilterIDs:   []string{}}

	newBarChart := BarChart{
		ParentDashboardID: params.ParentDashboardID,
		BarChartID:        uniqueID.GenerateSnowflakeID(),
		Properties:        barChartProps}

	if saveErr := common.SaveNewDashboardComponent(barChartEntityKind,
		newBarChart.ParentDashboardID, newBarChart.BarChartID, newBarChart.Properties); saveErr != nil {
		return nil, fmt.Errorf("NewBarChart: Unable to save bar chart component with params=%+v: error = %v", params, saveErr)
	}

	return &newBarChart, nil

}

func getBarChart(parentDashboardID string, barChartID string) (*BarChart, error) {

	barChartProps := BarChartProps{}
	if getErr := common.GetDashboardComponent(barChartEntityKind, parentDashboardID, barChartID, &barChartProps); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to retrieve bar chart component: %v", getErr)
	}

	barChart := BarChart{
		ParentDashboardID: parentDashboardID,
		BarChartID:        barChartID,
		Properties:        barChartProps}

	return &barChart, nil

}

func getBarCharts(parentDashboardID string) ([]BarChart, error) {

	barCharts := []BarChart{}
	addBarChart := func(barChartID string, encodedProps string) error {

		var barChartProps BarChartProps
		if decodeErr := generic.DecodeJSONString(encodedProps, &barChartProps); decodeErr != nil {
			return fmt.Errorf("getBarCharts: can't decode properties: %v", encodedProps)
		}

		currBarChart := BarChart{
			ParentDashboardID: parentDashboardID,
			BarChartID:        barChartID,
			Properties:        barChartProps}
		barCharts = append(barCharts, currBarChart)

		return nil
	}
	if getErr := common.GetDashboardComponents(barChartEntityKind, parentDashboardID, addBarChart); getErr != nil {
		return nil, fmt.Errorf("getBarCharts: Can't get bar chart components: %v")
	}

	return barCharts, nil
}
