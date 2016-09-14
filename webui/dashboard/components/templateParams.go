package components

import (
	"resultra/datasheet/webui/dashboard/components/barChart"
	"resultra/datasheet/webui/dashboard/components/summaryTable"
)

type ComponentTemplateParams struct {
	BarChartParams     barChart.BarChartTemplateParams
	SummaryTableParams summaryTable.SummaryTableTemplateParams
}

var TemplateParams ComponentTemplateParams

func init() {
	TemplateParams = ComponentTemplateParams{
		BarChartParams:     barChart.TemplateParams,
		SummaryTableParams: summaryTable.TemplateParams}
}
