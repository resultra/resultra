package checkBox

import (
	"resultra/datasheet/webui/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type CheckboxDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

type CheckboxViewTemplateParams struct {
	ElemPrefix         string
	CommentPanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams CheckboxDesignTemplateParams
var ViewTemplateParams CheckboxViewTemplateParams

func init() {
	DesignTemplateParams = CheckboxDesignTemplateParams{
		ElemPrefix:        "checkbox_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "checkboxFormat"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  "checkbox_",
			DialogTitle: "New Check Box",
			FieldInfoPrompt: `Check box values are stored in fields. Either a new field can be created to store the values, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this check box's values.`}}

	ViewTemplateParams = CheckboxViewTemplateParams{
		ElemPrefix:         "checkbox_",
		CommentPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Comments", PanelID: "checkboxComments"}}

}
