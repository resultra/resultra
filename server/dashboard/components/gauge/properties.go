package gauge

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic/threshold"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/trackerDatabase"
)

type GaugeProps struct {
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
	Title    string                         `json:"title"`

	ValSummary    values.ValSummary           `json:"valSummary"`
	MinVal        float64                     `json:"minVal"`
	MaxVal        float64                     `json:"maxVal"`
	ThresholdVals []threshold.ThresholdValues `json:"thresholdVals"`

	DefaultFilterRules recordFilter.RecordFilterRuleSet `json:"defaultFilterRules"`
	PreFilterRules     recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
	HelpPopupMsg       string                           `json:"helpPopupMsg"`
}

func (srcProps GaugeProps) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*GaugeProps, error) {

	destProps := srcProps

	valSummary, err := srcProps.ValSummary.Clone(cloneParams.IDRemapper)
	if err != nil {
		return nil, fmt.Errorf("GaugeProps.Clone: %v", err)
	}
	destProps.ValSummary = *valSummary

	clonedFilterRules, err := srcProps.DefaultFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("GaugeProps.Clone: %v", err)
	}
	destProps.DefaultFilterRules = *clonedFilterRules

	clonedPreFilterRules, err := srcProps.PreFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("GaugeProps.Clone: %v", err)
	}
	destProps.PreFilterRules = *clonedPreFilterRules

	return &destProps, nil

}

func newDefaultGaugeProps() GaugeProps {
	props := GaugeProps{
		Title:              "Gauge",
		MinVal:             0.0,
		MaxVal:             100.0,
		ThresholdVals:      []threshold.ThresholdValues{},
		DefaultFilterRules: recordFilter.NewDefaultRecordFilterRuleSet(),
		PreFilterRules:     recordFilter.NewDefaultRecordFilterRuleSet(),
		HelpPopupMsg:       ""}
	return props
}
