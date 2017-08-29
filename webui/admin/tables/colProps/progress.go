package colProps

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/valueThreshold"
	"resultra/datasheet/webui/generic/propertiesSidebar"
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
