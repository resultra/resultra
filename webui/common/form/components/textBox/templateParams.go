package textBox

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type TextboxDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
}

type TextboxViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams TextboxDesignTemplateParams
var ViewTemplateParams TextboxViewTemplateParams

func init() {

	elemPrefix := "textBox_"

	DesignTemplateParams = TextboxDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "textboxLabel"}},
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "textboxFormat"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "textBoxVisibility"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Text Box",
			FieldInfoPrompt: `Values from text boxes are stored in fields. Either a new field can be created for this
					text box, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this text box's values.'`}}

	ViewTemplateParams = TextboxViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "textBoxTimeline"}}

}
