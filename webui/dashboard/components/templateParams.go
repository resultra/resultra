package components

import (
	"resultra/tracker/webui/dashboard/components/barChart"
	"resultra/tracker/webui/dashboard/components/gauge"
	"resultra/tracker/webui/dashboard/components/header"
	"resultra/tracker/webui/dashboard/components/summaryTable"
	"resultra/tracker/webui/dashboard/components/summaryValue"
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
	GaugeParams        gauge.GaugeViewTemplateParams
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
		GaugeParams:        gauge.ViewTemplateParams,
		SummaryValParams:   summaryValue.ViewTemplateParams}
}
