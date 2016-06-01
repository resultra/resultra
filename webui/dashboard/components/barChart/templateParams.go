package barChart

import (
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type BarChartTemplateParams struct {
	ElemPrefix          string
	SelectTableParams   newComponentDialog.SelectTableTemplateParams
	ValueGroupingParams newComponentDialog.ValueGroupingTemplateParams
	ValueSummaryParams  newComponentDialog.ValueSummaryTemplateParams
	TitlePanelParams    propertiesSidebar.PanelTemplateParams
	XAxisPanelParams    propertiesSidebar.PanelTemplateParams
}

var TemplateParams BarChartTemplateParams

func init() {

	elemPrefix := "barChart_"

	tableSelectionParams := newComponentDialog.SelectTableTemplateParams{elemPrefix,
		"Select a table as the source of data for this bar chart."}

	valueGroupingParams := newComponentDialog.ValueGroupingTemplateParams{
		elemPrefix, "Configure which values are shown along the X axis and how these values are grouped",
		"Field for X axis' values", "Group Values By"}

	valueSummaryParams := newComponentDialog.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized along the Y axis.",
		"Field to summarize with", "Summarize values by"}

	TemplateParams = BarChartTemplateParams{
		ElemPrefix:          elemPrefix,
		SelectTableParams:   tableSelectionParams,
		ValueGroupingParams: valueGroupingParams,
		ValueSummaryParams:  valueSummaryParams,
		TitlePanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "barChartTitle"},
		XAxisPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "X Axis", PanelID: "barChartXAxis"}}
}
