package htmlEditor

import (
	"resultra/datasheet/webui/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HTMLEditorTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

var TemplateParams HTMLEditorTemplateParams

func init() {
	TemplateParams = HTMLEditorTemplateParams{
		ElemPrefix:        "htmlEditor_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "htmlEditorFormat"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  "htmlEditor_",
			DialogTitle: "New HTML Editor",
			FieldInfoPrompt: `Content provided in an editor is stored in fields. Either a new field can be created for this
					editor, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this editor's content.`}}
}
