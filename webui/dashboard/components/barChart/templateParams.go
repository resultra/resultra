package barChart

import (
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
)

type BarChartTemplateParams struct {
	ElemPrefix          string
	SelectTableParams   newComponentDialog.SelectTableTemplateParams
	ValueGroupingParams newComponentDialog.ValueGroupingTemplateParams
}

var TemplateParams BarChartTemplateParams

func init() {

	elemPrefix := "barChart_"

	TemplateParams = BarChartTemplateParams{
		ElemPrefix: elemPrefix,
		SelectTableParams: newComponentDialog.SelectTableTemplateParams{elemPrefix,
			"Select a table as the source of data for this bar chart."},
		ValueGroupingParams: newComponentDialog.ValueGroupingTemplateParams{
			elemPrefix, "Configure which values are shown on the X axis and how these values are grouped",
			"Field for X axis' values", "Group Values By"}}
}
