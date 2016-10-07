package summaryTable

import (
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
	"resultra/datasheet/webui/dashboard/components/common/valueSummary"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

// Template parameters when the summary table is in design mode
type SummaryTableDesignTemplateParams struct {
	ElemPrefix             string
	SelectTableParams      newComponentDialog.SelectTableTemplateParams
	RowValueGroupingParams newComponentDialog.ValueGroupingTemplateParams
	ColValueSummaryParams  valueSummary.ColumnsValueSummaryTemplateParams
	TitlePanelParams       propertiesSidebar.PanelTemplateParams
	RowPanelParams         propertiesSidebar.PanelTemplateParams
	ColPanelParams         propertiesSidebar.PanelTemplateParams
	FilteringPanelParams   propertiesSidebar.PanelTemplateParams
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

	tableSelectionParams := newComponentDialog.SelectTableTemplateParams{elemPrefix,
		"Select a table as the source of data for this summary table."}

	rowGroupingParams := newComponentDialog.ValueGroupingTemplateParams{
		elemPrefix, "Configure which field is used to group values into rows",
		"Field for grouping values into rows", "Group Values By"}

	colSummaryParams := valueSummary.ColumnsValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized in the 1st column (more to come later)"}

	DesignTemplateParams = SummaryTableDesignTemplateParams{
		ElemPrefix:             elemPrefix,
		SelectTableParams:      tableSelectionParams,
		RowValueGroupingParams: rowGroupingParams,
		ColValueSummaryParams:  colSummaryParams,
		TitlePanelParams:       propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "summaryTableTitle"},
		RowPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Rows", PanelID: "summaryTableRows"},
		ColPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Columns", PanelID: "summaryTableCols"},
		FilteringPanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "summaryTableFiltering"}}

	ViewTemplateParams = SummaryTableViewTemplateParams{
		ElemPrefix:           elemPrefix,
		FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "summaryTableFiltering"}}
}
