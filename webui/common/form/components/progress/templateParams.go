package progress

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/valueThreshold"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ProgressDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	RangePanelParams         propertiesSidebar.PanelTemplateParams
	ThresholdPanelParams     propertiesSidebar.PanelTemplateParams
	ThresholdValueParams     valueThreshold.ThresholdValuesPanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
}

type ProgressViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams ProgressDesignTemplateParams
var ViewTemplateParams ProgressViewTemplateParams

func init() {

	elemPrefix := "progress_"

	DesignTemplateParams = ProgressDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "progressLabel"}},
		FormatPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "progressFormat"},
		RangePanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Range", PanelID: "progressRange"},
		ThresholdPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Thresholds", PanelID: "progressThreshold"},
		ThresholdValueParams: valueThreshold.NewThresholdValuesTemplateParams(elemPrefix),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Progress Indicator",
			FieldInfoPrompt:    `Progress indicators use a field value to determine the progress level.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this progress indicator's values.`}}

	ViewTemplateParams = ProgressViewTemplateParams{
		ElemPrefix: elemPrefix}

}
