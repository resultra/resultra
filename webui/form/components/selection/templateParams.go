package selection

import (
	"resultra/datasheet/webui/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type SelectionDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ValuesPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

type SelectionViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams SelectionDesignTemplateParams
var ViewTemplateParams SelectionViewTemplateParams

func init() {

	elemPrefix := "selection_"

	DesignTemplateParams = SelectionDesignTemplateParams{
		ElemPrefix: elemPrefix,

		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "selectionFormat"},
		ValuesPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Selectable Values", PanelID: "selectionValues"},

		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Selection",
			FieldInfoPrompt: `Values from selections are stored in fields. Either a new field can be created for this
					selection, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this selection's values.'`}}

	ViewTemplateParams = SelectionViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "selectionTimeline"}}

}
