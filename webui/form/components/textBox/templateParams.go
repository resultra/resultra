package textBox

import (
	"resultra/datasheet/webui/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type TextboxTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

var TemplateParams TextboxTemplateParams

func init() {
	TemplateParams = TextboxTemplateParams{
		ElemPrefix:        "textBox_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "textboxFormat"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  "textBox_",
			DialogTitle: "New Text Box",
			FieldInfoPrompt: `Values from text boxes are stored in fields. Either a new field can be created for this
					text box, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this text box's values.'`}}
}
