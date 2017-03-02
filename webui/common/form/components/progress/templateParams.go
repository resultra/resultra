package progress

import (
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ProgressDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	RangePanelParams         propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

type ProgressViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams ProgressDesignTemplateParams
var ViewTemplateParams ProgressViewTemplateParams

func init() {

	elemPrefix := "progress_"

	DesignTemplateParams = ProgressDesignTemplateParams{
		ElemPrefix:        elemPrefix,
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "progressFormat"},
		RangePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Range", PanelID: "progressRange"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Progress Indicator",
			FieldInfoPrompt:    `Progress indicators use a field value to determine the progress level.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this progress indicator's values.`}}

	ViewTemplateParams = ProgressViewTemplateParams{
		ElemPrefix: elemPrefix}

}