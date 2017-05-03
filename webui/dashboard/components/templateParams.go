package components

import (
	"resultra/datasheet/webui/dashboard/components/barChart"
	"resultra/datasheet/webui/dashboard/components/gauge"
	"resultra/datasheet/webui/dashboard/components/header"
	"resultra/datasheet/webui/dashboard/components/summaryTable"
)

type ComponentDesignTemplateParams struct {
	BarChartParams     barChart.BarChartDesignTemplateParams
	SummaryTableParams summaryTable.SummaryTableDesignTemplateParams
	HeaderParams       header.HeaderDesignTemplateParams
	GaugeParams        gauge.GaugeDesignTemplateParams
}

type ComponentViewTemplateParams struct {
	SummaryTableParams summaryTable.SummaryTableViewTemplateParams
	BarChartParams     barChart.BarChartViewTemplateParams
}

var DesignTemplateParams ComponentDesignTemplateParams
var ViewTemplateParams ComponentViewTemplateParams

func init() {
	DesignTemplateParams = ComponentDesignTemplateParams{
		BarChartParams:     barChart.DesignTemplateParams,
		SummaryTableParams: summaryTable.DesignTemplateParams,
		HeaderParams:       header.DesignTemplateParams,
		GaugeParams:        gauge.DesignTemplateParams}

	ViewTemplateParams = ComponentViewTemplateParams{
		SummaryTableParams: summaryTable.ViewTemplateParams,
		BarChartParams:     barChart.ViewTemplateParams}
}
