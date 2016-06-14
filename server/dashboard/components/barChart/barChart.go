package barChart

import (
	"fmt"
	"github.com/gocql/gocql"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

const barChartEntityKind string = "BarChart"

const xAxisSortAsc string = "asc"
const xAxisSortDesc string = "desc"

type BarChartProps struct {

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

	Geometry common.LayoutGeometry `json:"geometry"`
}

func updateExistingBarChart(updatedBarChart *BarChart) (*BarChart, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedProps, encodeErr := generic.EncodeJSONString(updatedBarChart.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingBarChart: Unable to update bar chart: error = %v", encodeErr)
	}

	if updateErr := dbSession.Query(`UPDATE bar_charts SET properties=? WHERE dashboard_id=? and barchart_id=?`,
		encodedProps,
		updatedBarChart.ParentDashboardID,
		updatedBarChart.BarChartID).Exec(); updateErr != nil {
		fmt.Errorf("updateExistingBarChart: Can't create bar chart: unable to update bar chart: error = %v", updateErr)
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

	valGrouping, valGroupingErr := values.NewValGrouping(params.XAxisVals)
	if valGroupingErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating new value grouping for bar chart: error = %v", valGroupingErr)
	}

	valSummary, valSummaryErr := values.NewValSummary(params.YAxisVals)
	if valSummaryErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating summary values for bar chart: error = %v", valSummaryErr)
	}

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("NewBarChart: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	if !validBarChartSortXAxisProp(params.XAxisSortValues) {
		return nil, fmt.Errorf("NewBarChart: Invalid X axis sort order: %v", params.XAxisSortValues)
	}

	barChartProps := BarChartProps{
		XAxisVals:       *valGrouping,
		XAxisSortValues: params.XAxisSortValues,
		DataSrcTableID:  params.DataSrcTableID,
		YAxisVals:       *valSummary,
		Geometry:        params.Geometry,
		Title:           ""}

	newBarChart := BarChart{
		ParentDashboardID: params.ParentDashboardID,
		BarChartID:        gocql.TimeUUID().String(),
		Properties:        barChartProps}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedProps, encodeErr := generic.EncodeJSONString(newBarChart.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("NewDashboard: Unable to create dashboard: error = %v", encodeErr)
	}

	if insertErr := dbSession.Query(
		`INSERT INTO bar_charts (dashboard_id, barchart_id, properties) 
			VALUES (?,?,?)`,
		newBarChart.ParentDashboardID,
		newBarChart.BarChartID,
		encodedProps).Exec(); insertErr != nil {
		fmt.Errorf("NewBarChart: Can't create bar chart: unable to create bar chart: error = %v", insertErr)
	}

	return &newBarChart, nil

}

func getBarChart(parentDashboardID string, barChartID string) (*BarChart, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedProperties := ""

	var barChart BarChart
	getErr := dbSession.Query(
		`SELECT dashboard_id, barchart_id, properties 
			FROM bar_charts 
			WHERE dashboard_id=? AND barchart_id=? LIMIT 1`,
		parentDashboardID, barChartID).Scan(&barChart.ParentDashboardID,
		&barChart.BarChartID,
		&encodedProperties)
	if getErr != nil {
		return nil, fmt.Errorf("Unabled to get bar chart: id = %+v: datastore err=%v", barChartID, getErr)
	}

	decodeErr := generic.DecodeJSONString(encodedProperties, &barChart.Properties)
	if decodeErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to get bar chart: error = %v", decodeErr)
	}

	return &barChart, nil

}

func getBarCharts(parentDashboardID string) ([]BarChart, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedProperties := ""

	barChartIter := dbSession.Query(`SELECT dashboard_id, barchart_id, properties
			FROM bar_charts
			WHERE dashboard_id=?`,
		parentDashboardID).Iter()

	var currBarChart BarChart
	barCharts := []BarChart{}
	for barChartIter.Scan(&currBarChart.ParentDashboardID,
		&currBarChart.BarChartID,
		&encodedProperties) {

		decodeErr := generic.DecodeJSONString(encodedProperties, &currBarChart.Properties)
		if decodeErr != nil {
			return nil, fmt.Errorf("getBarChart: Unable to get bar chart: error = %v", decodeErr)
		}

		barCharts = append(barCharts, currBarChart)

		encodedProperties = ""
	}
	if closeErr := barChartIter.Close(); closeErr != nil {
		return nil, fmt.Errorf("getTableList: Failure querying database: %v", closeErr)
	}

	return barCharts, nil

}
