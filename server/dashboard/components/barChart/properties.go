package barChart

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

const xAxisSortAsc string = "asc"
const xAxisSortDesc string = "desc"

type BarChartProps struct {

	// DataSrcTable is the table the bar chart gets its data from
	DataSrcTableID string `json:"dataSrcTableID"`

	// XAxisVals is a grouping of field values displayed along the x axis of the bar chart.
	XAxisVals values.ValGrouping `json:"xAxisVals"`

	// XAxisSortValues configures how values (bars) along the x axis are sorted. Options include
	// yAxisValAsc, yAxisValDesc, xAxisValAsc, xAxisValDesc. The default is xAxisValAsc.
	XAxisSortValues string `json:"xAxisSortValues"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`

	Title string `json:"title"`

	YAxisVals values.ValSummary `json:"yAxisValSummary"`

	DefaultFilterRules []recordFilter.RecordFilterRule `json:"defaultFilterRules"`
}

func (srcProps BarChartProps) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*BarChartProps, error) {

	destProps := srcProps

	remappedTableID, err := remappedIDs.GetExistingRemappedID(srcProps.DataSrcTableID)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.DataSrcTableID = remappedTableID

	xAxisVals, err := srcProps.XAxisVals.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.XAxisVals = *xAxisVals

	yAxisVals, err := srcProps.YAxisVals.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.YAxisVals = *yAxisVals

	clonedFilterRules, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.DefaultFilterRules)
	if err != nil {
		return nil, fmt.Errorf("BarChartProps.Clone: %v", err)
	}
	destProps.DefaultFilterRules = clonedFilterRules

	return &destProps, nil

}
