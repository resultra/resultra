package valueThreshold

type ThresholdValuesPanelTemplateParams struct {
	ElemPrefix string
}

func NewThresholdValuesTemplateParams(elemPrefix string) ThresholdValuesPanelTemplateParams {

	panelParams := ThresholdValuesPanelTemplateParams{
		ElemPrefix: elemPrefix}

	return panelParams
}
