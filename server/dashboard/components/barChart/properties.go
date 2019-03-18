// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package barChart

import (
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/dashboard/values"
	"resultra/tracker/server/recordFilter"
	"resultra/tracker/server/trackerDatabase"
)

const xAxisSortAsc string = "asc"
const xAxisSortDesc string = "desc"

type BarChartProps struct {

	// XAxisVals is a grouping of field values displayed along the x axis of the bar chart.
	XAxisVals values.ValGrouping `json:"xAxisVals"`

	// XAxisSortValues configures how values (bars) along the x axis are sorted. Options include
	// yAxisValAsc, yAxisValDesc, xAxisValAsc, xAxisValDesc. The default is xAxisValAsc.
	XAxisSortValues string `json:"xAxisSortValues"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`

	Title string `json:"title"`

	YAxisVals values.ValSummary `json:"yAxisValSummary"`

	DefaultFilterRules recordFilter.RecordFilterRuleSet `json:"defaultFilterRules"`
	PreFilterRules     recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
	HelpPopupMsg       string                           `json:"helpPopupMsg"`
}

func (srcProps BarChartProps) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*BarChartProps, error) {

	destProps := srcProps

	xAxisVals, err := srcProps.XAxisVals.Clone(cloneParams.IDRemapper)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.XAxisVals = *xAxisVals

	yAxisVals, err := srcProps.YAxisVals.Clone(cloneParams.IDRemapper)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.YAxisVals = *yAxisVals

	clonedFilterRules, err := srcProps.DefaultFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.DefaultFilterRules = *clonedFilterRules

	clonedPreFilterRules, err := srcProps.PreFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.PreFilterRules = *clonedPreFilterRules

	return &destProps, nil

}

func newDefaultBarChartProps() BarChartProps {

	props := BarChartProps{
		Title:              "",
		DefaultFilterRules: recordFilter.NewDefaultRecordFilterRuleSet(),
		PreFilterRules:     recordFilter.NewDefaultRecordFilterRuleSet(),
		HelpPopupMsg:       ""}

	return props
}
