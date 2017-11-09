package gauge

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/components/common"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const gaugeEntityKind string = "gauge"

// DashboardBarChart is the datastore object for dashboard bar charts.
type Gauge struct {
	ParentDashboardID string `json:"parentDashboardID"`

	GaugeID string `json:"gaugeID"`

	// DataSrcTable is the table the bar chart gets its data from
	Properties GaugeProps `json:"properties"`
}

type NewGaugeParams struct {
	ParentDashboardID string `json:"parentDashboardID"`

	ValSummary values.NewValSummaryParams `json:"valSummary"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func saveGauge(destDBHandle *sql.DB, newGauge Gauge) error {

	if saveErr := common.SaveNewDashboardComponent(destDBHandle, gaugeEntityKind,
		newGauge.ParentDashboardID, newGauge.GaugeID, newGauge.Properties); saveErr != nil {
		return fmt.Errorf("newGauge: Unable to save gauge component: error = %v", saveErr)
	}
	return nil

}

func newGauge(trackerDBHandle *sql.DB, params NewGaugeParams) (*Gauge, error) {

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("newGauge: Error creating summary table: missing parent dashboard ID")
	}

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("newGauge: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	valSummary, valSummaryErr := values.NewValSummary(trackerDBHandle, params.ValSummary)
	if valSummaryErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating summary values for bar chart: error = %v", valSummaryErr)
	}

	gaugeProps := newDefaultGaugeProps()
	gaugeProps.Geometry = params.Geometry
	gaugeProps.ValSummary = *valSummary

	newGauge := Gauge{
		ParentDashboardID: params.ParentDashboardID,
		GaugeID:           uniqueID.GenerateSnowflakeID(),
		Properties:        gaugeProps}

	if saveErr := saveGauge(trackerDBHandle, newGauge); saveErr != nil {
		return nil, fmt.Errorf("newGauge: Unable to save summary component with params=%+v: error = %v", params, saveErr)
	}

	return &newGauge, nil
}

func GetGauge(trackerDBHandle *sql.DB, parentDashboardID string, gaugeID string) (*Gauge, error) {

	gaugeProps := newDefaultGaugeProps()
	if getErr := common.GetDashboardComponent(trackerDBHandle,
		gaugeEntityKind, parentDashboardID, gaugeID, &gaugeProps); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to retrieve bar chart component: %v", getErr)
	}

	gauge := Gauge{
		ParentDashboardID: parentDashboardID,
		GaugeID:           gaugeID,
		Properties:        gaugeProps}

	return &gauge, nil

}

func getGaugesFromSrc(srcDBHandle *sql.DB, parentDashboardID string) ([]Gauge, error) {

	gauges := []Gauge{}
	addGauge := func(gaugeID string, encodedProps string) error {

		gaugeProps := newDefaultGaugeProps()
		if decodeErr := generic.DecodeJSONString(encodedProps, &gaugeProps); decodeErr != nil {
			return fmt.Errorf("GetGauges: can't decode properties: %v", encodedProps)
		}

		currGauge := Gauge{
			ParentDashboardID: parentDashboardID,
			GaugeID:           gaugeID,
			Properties:        gaugeProps}

		gauges = append(gauges, currGauge)

		return nil
	}
	if getErr := common.GetDashboardComponents(srcDBHandle, gaugeEntityKind, parentDashboardID, addGauge); getErr != nil {
		return nil, fmt.Errorf("getBarCharts: Can't get bar chart components: %v")
	}

	return gauges, nil
}

func GetGauges(trackerDBHandle *sql.DB, parentDashboardID string) ([]Gauge, error) {
	return getGaugesFromSrc(trackerDBHandle, parentDashboardID)
}

func CloneGauges(cloneParams *trackerDatabase.CloneDatabaseParams, srcParentDashboardID string) error {

	remappedDashboardID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneGauges: %v", err)
	}

	gauges, err := getGaugesFromSrc(cloneParams.SrcDBHandle, srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneGauges: %v", err)
	}

	for _, srcGauge := range gauges {

		remappedGaugeID, err := cloneParams.IDRemapper.AllocNewRemappedID(srcGauge.GaugeID)
		if err != nil {
			return fmt.Errorf("CloneGauges: %v", err)
		}

		clonedProps, err := srcGauge.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneGauges: %v", err)
		}

		destGauge := Gauge{
			ParentDashboardID: remappedDashboardID,
			GaugeID:           remappedGaugeID,
			Properties:        *clonedProps}

		if err := saveGauge(cloneParams.DestDBHandle, destGauge); err != nil {
			return fmt.Errorf("CloneGauges: %v", err)
		}
	}

	return nil
}

func updateExistingGauge(trackerDBHandle *sql.DB, updatedGauge *Gauge) (*Gauge, error) {

	if updateErr := common.UpdateDashboardComponent(trackerDBHandle,
		gaugeEntityKind, updatedGauge.ParentDashboardID,
		updatedGauge.GaugeID, updatedGauge.Properties); updateErr != nil {
		return nil, fmt.Errorf("Error updating summary table %+v: %v", updatedGauge, updateErr)
	}

	return updatedGauge, nil

}
