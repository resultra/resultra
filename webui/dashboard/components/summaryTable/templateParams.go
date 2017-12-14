package summaryTable

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/dashboard/components/common/delete"
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
	"resultra/datasheet/webui/dashboard/components/common/valueSummary"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

// Template parameters when the summary table is in design mode
type SummaryTableDesignTemplateParams struct {
	ElemPrefix               string
	RowValueGroupingParams   newComponentDialog.ValueGroupingTemplateParams
	ColValueSummaryParams    valueSummary.ValueSummaryTemplateParams
	TitlePanelParams         propertiesSidebar.PanelTemplateParams
	RowPanelParams           propertiesSidebar.PanelTemplateParams
	ColPanelParams           propertiesSidebar.PanelTemplateParams
	FilteringPanelParams     propertiesSidebar.PanelTemplateParams
	PreFilteringPanelParams  propertiesSidebar.PanelTemplateParams
	FilterPropPanelParams    recordFilter.FilterPanelTemplateParams
	PreFilterPropPanelParams recordFilter.FilterPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
}

// Template parameters when the summary table is in view mode
type SummaryTableViewTemplateParams struct {
	ElemPrefix              string
	FilteringPanelParams    propertiesSidebar.PanelTemplateParams
	FilterConfigPanelParams recordFilter.FilterPanelTemplateParams
}

var DesignTemplateParams SummaryTableDesignTemplateParams
var ViewTemplateParams SummaryTableViewTemplateParams

func init() {

	elemPrefix := "summaryTable_"
	preFilterElemPrefix := "summaryTablePreFilter_"

	rowGroupingParams := newComponentDialog.ValueGroupingTemplateParams{
		elemPrefix, "Configure which field is used to group values into rows",
		"Field for grouping values into rows", "Group Values By"}

	valueSummaryParams := valueSummary.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized in the first column (more columns can be added later).",
		"Field to summarize with", "Summarize values by"}

	DesignTemplateParams = SummaryTableDesignTemplateParams{
		ElemPrefix:               elemPrefix,
		RowValueGroupingParams:   rowGroupingParams,
		ColValueSummaryParams:    valueSummaryParams,
		TitlePanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "summaryTableTitle"},
		HelpPopupParams:          inputProperties.NewHelpPopupTemplateParams(elemPrefix, "summaryTableHelp"),
		RowPanelParams:           propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Rows", PanelID: "summaryTableRows"},
		ColPanelParams:           propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Columns", PanelID: "summaryTableCols"},
		PreFilteringPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Pre-Filtering", PanelID: "summaryTablePreFiltering"},
		FilteringPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "summaryTableFiltering"},
		DeletePanelParams:        delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "summaryTableDelete", "Delete Summary Table"),
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix)}

	ViewTemplateParams = SummaryTableViewTemplateParams{
		ElemPrefix:              elemPrefix,
		FilteringPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "summaryTableFiltering"},
		FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}
}
