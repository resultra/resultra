package components

import (
	"resultra/datasheet/webui/dashboard/components/barChart"
)

type ComponentTemplateParams struct {
	BarChartParams barChart.BarChartTemplateParams
}

var TemplateParams ComponentTemplateParams

func init() {
	TemplateParams = ComponentTemplateParams{BarChartParams: barChart.TemplateParams}
}
