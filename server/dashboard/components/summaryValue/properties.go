// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package summaryValue

import (
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/dashboard/values"
	"github.com/resultra/resultra/server/generic/threshold"
	"github.com/resultra/resultra/server/recordFilter"
	"github.com/resultra/resultra/server/trackerDatabase"
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
