package components

import (
	"resultra/datasheet/webui/dashboard/components/barChart"
	"resultra/datasheet/webui/dashboard/components/gauge"
	"resultra/datasheet/webui/dashboard/components/header"
	"resultra/datasheet/webui/dashboard/components/summaryTable"
	"resultra/datasheet/webui/dashboard/components/summaryValue"
)

type ComponentDesignTemplateParams struct {
	BarChartParams     barChart.BarChartDesignTemplateParams
	SummaryTableParams summaryTable.SummaryTableDesignTemplateParams
	HeaderParams       header.HeaderDesignTemplateParams
	GaugeParams        gauge.GaugeDesignTemplateParams
	SummaryValParams   summaryValue.SummaryValDesignTemplateParams
}

type ComponentViewTemplateParams struct {
	SummaryTableParams summaryTable.SummaryTableViewTemplateParams
	BarChartParams     barChart.BarChartViewTemplateParams
	SummaryValParams   summaryValue.SummaryValViewTemplateParams
}

var DesignTemplateParams ComponentDesignTemplateParams
var ViewTemplateParams ComponentViewTemplateParams

func init() {
	DesignTemplateParams = ComponentDesignTemplateParams{
		BarChartParams:     barChart.DesignTemplateParams,
		SummaryTableParams: summaryTable.DesignTemplateParams,
		HeaderParams:       header.DesignTemplateParams,
		GaugeParams:        gauge.DesignTemplateParams,
		SummaryValParams:   summaryValue.DesignTemplateParams}

	ViewTemplateParams = ComponentViewTemplateParams{
		SummaryTableParams: summaryTable.ViewTemplateParams,
		BarChartParams:     barChart.ViewTemplateParams,
		SummaryValParams:   summaryValue.ViewTemplateParams}
}
