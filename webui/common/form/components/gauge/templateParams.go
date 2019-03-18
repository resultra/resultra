package gauge

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/form/components/common/delete"
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/newFormElemDialog"
	"resultra/tracker/webui/common/form/components/common/visibility"
	"resultra/tracker/webui/common/valueThreshold"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type GaugeDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	RangePanelParams         propertiesSidebar.PanelTemplateParams
	ThresholdPanelParams     propertiesSidebar.PanelTemplateParams
	ThresholdValueParams     valueThreshold.ThresholdValuesPanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
}

type GaugeViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams GaugeDesignTemplateParams
var ViewTemplateParams GaugeViewTemplateParams

func init() {

	elemPrefix := "gauge_"

	DesignTemplateParams = GaugeDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "gaugeLabel"}},
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "gaugeHelp"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "gaugeDelete", "Delete Gauge"),
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "progressVisibility"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "gaugeFormat"},
		RangePanelParams:      propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Range", PanelID: "gaugeRange"},
		ThresholdPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Thresholds", PanelID: "gaugeThreshold"},
		ThresholdValueParams:  valueThreshold.NewThresholdValuesTemplateParams(elemPrefix),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Gauge",
			FieldInfoPrompt:    `Gauges use a field value to determine the progress level.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this gauge's values.`}}

	ViewTemplateParams = GaugeViewTemplateParams{
		ElemPrefix: elemPrefix}

}
