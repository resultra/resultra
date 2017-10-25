package summaryValue

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic/threshold"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/trackerDatabase"
)

type SummaryValProps struct {
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
	Title    string                         `json:"title"`

	ValSummary    values.ValSummary           `json:"valSummary"`
	ThresholdVals []threshold.ThresholdValues `json:"thresholdVals"`

	DefaultFilterRules recordFilter.RecordFilterRuleSet `json:"defaultFilterRules"`
	PreFilterRules     recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
	HelpPopupMsg       string                           `json:"helpPopupMsg"`
}

func (srcProps SummaryValProps) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*SummaryValProps, error) {

	destProps := srcProps

	valSummary, err := srcProps.ValSummary.Clone(cloneParams.IDRemapper)
	if err != nil {
		return nil, fmt.Errorf("SummaryValProps.Clone: %v", err)
	}
	destProps.ValSummary = *valSummary

	clonedFilterRules, err := srcProps.DefaultFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("SummaryValProps.Clone: %v", err)
	}
	destProps.DefaultFilterRules = *clonedFilterRules

	clonedPreFilterRules, err := srcProps.PreFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("SummaryValProps.Clone: %v", err)
	}
	destProps.PreFilterRules = *clonedPreFilterRules

	return &destProps, nil

}

func newDefaultSummaryValProps() SummaryValProps {
	props := SummaryValProps{
		Title:              "SummaryVal",
		ThresholdVals:      []threshold.ThresholdValues{},
		DefaultFilterRules: recordFilter.NewDefaultRecordFilterRuleSet(),
		PreFilterRules:     recordFilter.NewDefaultRecordFilterRuleSet(),
		HelpPopupMsg:       ""}
	return props
}
