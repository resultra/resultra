package gauge

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/dashboard/values"
	"resultra/tracker/server/generic/threshold"
	"resultra/tracker/server/recordFilter"
)

// The BarChartPropertyUpdater interface along with UpdateBarChartProps() implement a harness for
// property updates. All property updates consiste of: (1) Retrieve the entity from the datastore,
// (2) Do the update on the entity itself, (3) Save the updated entity back to the datastore.
// Steps (1) and (3) can be done in a wrapper function UpdateBarChartProps, while only (2) needs
// be defined for each different property update. The goal is to minimize code bloat of property
// setting code and also make property updating code more uniform and less error prone.
type GaugePropertyUpdater interface {
	uniqueGaugeID() string
	parentDashboardID() string
	updateGaugeProps(gauge *Gauge) error
}

type GaugeUniqueIDGauge struct {
	ParentDashboardID string `json:"parentDashboardID"`
	GaugeID           string `json:"gaugeID"`
}

func (idGauge GaugeUniqueIDGauge) parentDashboardID() string {
	return idGauge.ParentDashboardID
}

func (idGauge GaugeUniqueIDGauge) uniqueGaugeID() string {
	return idGauge.GaugeID
}

func updateGaugeProps(trackerDBHandle *sql.DB, propUpdater GaugePropertyUpdater) (*Gauge, error) {

	// Retrieve the bar chart from the data store
	gaugeForUpdate, getErr := GetGauge(trackerDBHandle,
		propUpdater.parentDashboardID(), propUpdater.uniqueGaugeID())
	if getErr != nil {
		return nil, fmt.Errorf("updateGaugeProps: Unable to get existing gauge: %v", getErr)
	}

	// Do the actual update
	propUpdateErr := propUpdater.updateGaugeProps(gaugeForUpdate)
	if propUpdateErr != nil {
		return nil, fmt.Errorf("updateGaugeProps: Unable to update existing gauge: %v", propUpdateErr)
	}

	// Save the updated bar chart back to the data store
	updatedGauge, updateErr := updateExistingGauge(trackerDBHandle, gaugeForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateGaugeProps: Unable to update existing gauge: %v", updateErr)
	}

	return updatedGauge, nil

}

// Title Property

type SetGaugeTitleParams struct {
	// Embed a common gauge to reference the BarChart in the datastore. This gauge also supports
	// the niqueBarChartID() method to retrieve the unique ID. So, once decoded, the struct can be passed as an
	// BarChartPropertyUpdater interface to a generic/reusable function to process the property update.
	GaugeUniqueIDGauge
	NewTitle string `json:"newTitle"`
}

func (titleParam SetGaugeTitleParams) updateGaugeProps(gauge *Gauge) error {

	log.Printf("Updating gauge title: %v", titleParam.NewTitle)

	gauge.Properties.Title = titleParam.NewTitle

	return nil
}

// Dimensions Property

type SetGaugeDimensionsParams struct {
	GaugeUniqueIDGauge
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (params SetGaugeDimensionsParams) updateGaugeProps(gauge *Gauge) error {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return fmt.Errorf("setBarChartDimensions: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	gauge.Properties.Geometry = params.Geometry

	return nil
}

type SetValSummaryParams struct {
	GaugeUniqueIDGauge
	ValSummary values.ValSummary `json:"valSummary"`
}

func (params SetValSummaryParams) updateGaugeProps(gauge *Gauge) error {

	gauge.Properties.ValSummary = params.ValSummary

	return nil
}

type SetDefaultFilterRulesParams struct {
	GaugeUniqueIDGauge
	DefaultFilterRules recordFilter.RecordFilterRuleSet `json:"defaultFilterRules"`
}

func (params SetDefaultFilterRulesParams) updateGaugeProps(gauge *Gauge) error {

	gauge.Properties.DefaultFilterRules = params.DefaultFilterRules

	return nil
}

type SetPreFilterRulesParams struct {
	GaugeUniqueIDGauge
	PreFilterRules recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
}

func (params SetPreFilterRulesParams) updateGaugeProps(gauge *Gauge) error {

	gauge.Properties.PreFilterRules = params.PreFilterRules

	return nil
}

type SetRangeParams struct {
	GaugeUniqueIDGauge
	MinVal float64 `json:"minVal"`
	MaxVal float64 `json:"maxVal"`
}

func (updateParams SetRangeParams) updateGaugeProps(gauge *Gauge) error {

	if updateParams.MaxVal <= updateParams.MinVal {
		return fmt.Errorf("invalid gauge indicator range: %v %v", updateParams.MinVal, updateParams.MaxVal)
	}

	gauge.Properties.MinVal = updateParams.MinVal
	gauge.Properties.MaxVal = updateParams.MaxVal

	return nil
}

type SetThresholdsParams struct {
	GaugeUniqueIDGauge
	ThresholdVals []threshold.ThresholdValues `json:"thresholdVals"`
}

func (updateParams SetThresholdsParams) updateGaugeProps(gauge *Gauge) error {

	gauge.Properties.ThresholdVals = updateParams.ThresholdVals

	return nil
}

type SetHelpPopupMsgParams struct {
	GaugeUniqueIDGauge
	PopupMsg string `json:"popupMsg"`
}

func (params SetHelpPopupMsgParams) updateGaugeProps(gauge *Gauge) error {

	gauge.Properties.HelpPopupMsg = params.PopupMsg

	return nil
}
