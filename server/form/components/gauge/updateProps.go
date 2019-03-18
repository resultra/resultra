package gauge

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/generic/numberFormat"
	"resultra/tracker/server/generic/threshold"
)

type GaugeIDInterface interface {
	getGaugeID() string
	getParentFormID() string
}

type GaugeIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	GaugeID      string `json:"gaugeID"`
}

func (idHeader GaugeIDHeader) getGaugeID() string {
	return idHeader.GaugeID
}

func (idHeader GaugeIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type GaugePropUpdater interface {
	GaugeIDInterface
	updateProps(gauge *Gauge) error
}

func updateGaugeProps(trackerDBHandle *sql.DB, propUpdater GaugePropUpdater) (*Gauge, error) {

	// Retrieve the bar chart from the data store
	gaugeForUpdate, getErr := getGauge(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getGaugeID())
	if getErr != nil {
		return nil, fmt.Errorf("updateGaugeProps: Unable to get existing gauge indicator: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(gaugeForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateGaugeProps: Unable to update existing gauge indicator properties: %v", propUpdateErr)
	}

	gauge, updateErr := updateExistingGauge(trackerDBHandle, gaugeForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateGaugeProps: Unable to update existing gauge indicator properties: datastore update error =  %v", updateErr)
	}

	return gauge, nil
}

type GaugeResizeParams struct {
	GaugeIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams GaugeResizeParams) updateProps(gauge *Gauge) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set gauge indicator dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	gauge.Properties.Geometry = updateParams.Geometry

	return nil
}

type SetRangeParams struct {
	GaugeIDHeader
	MinVal float64 `json:"minVal"`
	MaxVal float64 `json:"maxVal"`
}

func (updateParams SetRangeParams) updateProps(gauge *Gauge) error {

	if updateParams.MaxVal <= updateParams.MinVal {
		return fmt.Errorf("invalid gauge indicator range: %v %v", updateParams.MinVal, updateParams.MaxVal)
	}

	gauge.Properties.MinVal = updateParams.MinVal
	gauge.Properties.MaxVal = updateParams.MaxVal

	return nil
}

type SetThresholdsParams struct {
	GaugeIDHeader
	ThresholdVals []threshold.ThresholdValues `json:"thresholdVals"`
}

func (updateParams SetThresholdsParams) updateProps(gauge *Gauge) error {

	gauge.Properties.ThresholdVals = updateParams.ThresholdVals

	return nil
}

type GaugeLabelFormatParams struct {
	GaugeIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams GaugeLabelFormatParams) updateProps(gauge *Gauge) error {

	// TODO - Validate format is well-formed.

	gauge.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type GaugeVisibilityParams struct {
	GaugeIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams GaugeVisibilityParams) updateProps(gauge *Gauge) error {

	// TODO - Validate conditions

	gauge.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type HelpPopupMsgParams struct {
	GaugeIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(gauge *Gauge) error {

	gauge.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}

type GaugeValueFormatParams struct {
	GaugeIDHeader
	ValueFormat numberFormat.NumberFormatProperties `json:"valueFormat"`
}

func (updateParams GaugeValueFormatParams) updateProps(gauge *Gauge) error {

	gauge.Properties.ValueFormat = updateParams.ValueFormat

	return nil
}
