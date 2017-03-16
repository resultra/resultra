package checkBox

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type CheckboxDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
}

type CheckboxViewTemplateParams struct {
	ElemPrefix         string
	CommentPanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams CheckboxDesignTemplateParams
var ViewTemplateParams CheckboxViewTemplateParams

func init() {

	elemPrefix := "checkbox_"

	DesignTemplateParams = CheckboxDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix, HideNoLabelOption: true,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "checkBoxLabel"}},
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "checkboxFormat"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Check Box",
			FieldInfoPrompt: `Check box values are stored in fields. Either a new field can be created to store the values, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this check box's values.`}}

	ViewTemplateParams = CheckboxViewTemplateParams{
		ElemPrefix:         elemPrefix,
		CommentPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Comments", PanelID: "checkboxComments"}}

}
