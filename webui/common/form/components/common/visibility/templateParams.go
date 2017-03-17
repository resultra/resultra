package visibility

import (
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type VisibilityPropertyTemplateParams struct {
	PanelParams                          propertiesSidebar.PanelTemplateParams
	ElemPrefix                           string
	VisibilityFilterConditionPanelParams recordFilter.FilterPanelTemplateParams
}

func NewComponentVisibilityTempalteParams(elemPrefix string, panelID string) VisibilityPropertyTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Conditional Visibility",
		PanelID:          panelID}

	params := VisibilityPropertyTemplateParams{ElemPrefix: elemPrefix,
		PanelParams:                          panelParams,
		VisibilityFilterConditionPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

	return params
}
