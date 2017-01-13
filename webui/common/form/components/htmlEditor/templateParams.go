package htmlEditor

import (
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HTMLEditorDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

type HTMLEditorViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams HTMLEditorDesignTemplateParams
var ViewTemplateParams HTMLEditorViewTemplateParams

func init() {

	elemPrefix := "htmlEditor_"

	DesignTemplateParams = HTMLEditorDesignTemplateParams{
		ElemPrefix:        elemPrefix,
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "htmlEditorFormat"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New HTML Editor",
			FieldInfoPrompt: `Content provided in an editor is stored in fields. Either a new field can be created for this
					editor, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this editor's content.`}}

	ViewTemplateParams = HTMLEditorViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "htmlEditorTimeline"}}
}
