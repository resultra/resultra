package colProps

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/valueThreshold"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type ProgressColPropsTemplateParams struct {
	ElemPrefix           string
	ThresholdValueParams valueThreshold.ThresholdValuesPanelTemplateParams
	LabelPanelParams     label.LabelPropertyTemplateParams
	HelpPopupParams      inputProperties.HelpPopupPropertyTemplateParams
}

func newProgressTemplateParams() ProgressColPropsTemplateParams {

	elemPrefix := "progress_"

	templParams := ProgressColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "progressLabel"}},
		HelpPopupParams:      inputProperties.NewHelpPopupTemplateParams(elemPrefix, "progressHelp"),
		ThresholdValueParams: valueThreshold.NewThresholdValuesTemplateParams(elemPrefix)}

	return templParams

}
