package summaryTable

import (
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
	"resultra/datasheet/webui/dashboard/components/common/valueSummary"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

// Template parameters when the summary table is in design mode
type SummaryTableDesignTemplateParams struct {
	ElemPrefix             string
	RowValueGroupingParams newComponentDialog.ValueGroupingTemplateParams
	ColValueSummaryParams  valueSummary.ValueSummaryTemplateParams
	TitlePanelParams       propertiesSidebar.PanelTemplateParams
	RowPanelParams         propertiesSidebar.PanelTemplateParams
	ColPanelParams         propertiesSidebar.PanelTemplateParams
	FilteringPanelParams   propertiesSidebar.PanelTemplateParams
	FilterPropPanelParams  recordFilter.FilterPanelTemplateParams
}

// Template parameters when the summary table is in view mode
type SummaryTableViewTemplateParams struct {
	ElemPrefix           string
	FilteringPanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams SummaryTableDesignTemplateParams
var ViewTemplateParams SummaryTableViewTemplateParams

func init() {

	elemPrefix := "summaryTable_"

	rowGroupingParams := newComponentDialog.ValueGroupingTemplateParams{
		elemPrefix, "Configure which field is used to group values into rows",
		"Field for grouping values into rows", "Group Values By"}

	valueSummaryParams := valueSummary.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized along the Y axis.",
		"Field to summarize with", "Summarize values by"}

	DesignTemplateParams = SummaryTableDesignTemplateParams{
		ElemPrefix:             elemPrefix,
		RowValueGroupingParams: rowGroupingParams,
		ColValueSummaryParams:  valueSummaryParams,
		TitlePanelParams:       propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "summaryTableTitle"},
		RowPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Rows", PanelID: "summaryTableRows"},
		ColPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Columns", PanelID: "summaryTableCols"},
		FilteringPanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "summaryTableFiltering"},
		FilterPropPanelParams:  recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

	ViewTemplateParams = SummaryTableViewTemplateParams{
		ElemPrefix:           elemPrefix,
		FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "summaryTableFiltering"}}
}
