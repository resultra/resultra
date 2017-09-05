package summaryTable

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

type SummaryTableProps struct {

	// XAxisVals is a grouping of field values displayed along the x axis of the bar chart.
	RowGroupingVals values.ValGrouping `json:"rowGroupingVals"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`

	Title string `json:"title"`

	ColumnValSummaries []values.ValSummary `json:"columnValSummaries"`

	DefaultFilterRules recordFilter.RecordFilterRuleSet `json:"defaultFilterRules"`
	PreFilterRules     recordFilter.RecordFilterRuleSet `json:"preFilterRules"`

	HelpPopupMsg string `json:"helpPopupMsg"`
}

func (srcProps SummaryTableProps) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*SummaryTableProps, error) {

	destProps := srcProps

	rowGroupingVals, err := srcProps.RowGroupingVals.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.RowGroupingVals = *rowGroupingVals

	destColSummaries := []values.ValSummary{}
	for _, srcColSummary := range srcProps.ColumnValSummaries {
		destColSummary, err := srcColSummary.Clone(remappedIDs)
		if err != nil {
			return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
		}
		destColSummaries = append(destColSummaries, *destColSummary)
	}
	destProps.ColumnValSummaries = destColSummaries

	clonedFilterRules, err := srcProps.DefaultFilterRules.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.DefaultFilterRules = *clonedFilterRules

	clonedPreFilterRules, err := srcProps.PreFilterRules.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.PreFilterRules = *clonedPreFilterRules

	return &destProps, nil

}

func newDefaultSummaryTableProps() SummaryTableProps {
	props := SummaryTableProps{
		Title:              "",
		DefaultFilterRules: recordFilter.NewDefaultRecordFilterRuleSet(),
		PreFilterRules:     recordFilter.NewDefaultRecordFilterRuleSet(),
		HelpPopupMsg:       ""}
	return props
}
