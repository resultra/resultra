package barChart

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/components/common"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

const barChartEntityKind string = "BarChart"

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

	XAxisVals       values.NewValGroupingParams `json:"xAxisVals"`
	XAxisSortValues string                      `json:"xAxisSortValues"`

	YAxisVals values.NewValSummaryParams `json:"yAxisVals"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func updateExistingBarChart(updatedBarChart *BarChart) (*BarChart, error) {

	if updateErr := common.UpdateDashboardComponent(barChartEntityKind, updatedBarChart.ParentDashboardID,
		updatedBarChart.BarChartID, updatedBarChart.Properties); updateErr != nil {
		return nil, fmt.Errorf("Error updating bar chart %+v: %v", updatedBarChart, updateErr)

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

func saveBarChart(newBarChart BarChart) error {
	if saveErr := common.SaveNewDashboardComponent(barChartEntityKind,
		newBarChart.ParentDashboardID, newBarChart.BarChartID, newBarChart.Properties); saveErr != nil {
		return fmt.Errorf("NewBarChart: Unable to save bar chart component: error = %v", saveErr)
	}
	return nil
}

func NewBarChart(params NewBarChartParams) (*BarChart, error) {

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("newSummaryTable: Error creating bar chart: missing parent dashboard ID")
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
		YAxisVals:          *valSummary,
		Geometry:           params.Geometry,
		Title:              "",
		DefaultFilterRules: []recordFilter.RecordFilterRule{}}

	newBarChart := BarChart{
		ParentDashboardID: params.ParentDashboardID,
		BarChartID:        uniqueID.GenerateSnowflakeID(),
		Properties:        barChartProps}

	if saveErr := saveBarChart(newBarChart); saveErr != nil {
		return nil, fmt.Errorf("NewBarChart: Unable to save bar chart component with params=%+v: error = %v", params, saveErr)
	}

	return &newBarChart, nil

}

func GetBarChart(parentDashboardID string, barChartID string) (*BarChart, error) {

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

func GetBarCharts(parentDashboardID string) ([]BarChart, error) {

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

func CloneBarCharts(remappedIDs uniqueID.UniqueIDRemapper, srcParentDashboardID string) error {

	remappedDashboardID, err := remappedIDs.GetExistingRemappedID(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneBarCharts: %v", err)
	}

	barCharts, err := GetBarCharts(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneBarCharts: %v", err)
	}

	for _, srcBarChart := range barCharts {

		remappedBarChartID, err := remappedIDs.AllocNewRemappedID(srcBarChart.BarChartID)
		if err != nil {
			return fmt.Errorf("CloneBarCharts: %v", err)
		}

		clonedProps, err := srcBarChart.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneBarCharts: %v", err)
		}

		destBarChart := BarChart{
			ParentDashboardID: remappedDashboardID,
			BarChartID:        remappedBarChartID,
			Properties:        *clonedProps}

		if err := saveBarChart(destBarChart); err != nil {
			return fmt.Errorf("CloneBarCharts: %v", err)
		}
	}

	return nil
}
