package inputProperties

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HelpPopupPropertyTemplateParams struct {
	PanelParams propertiesSidebar.PanelTemplateParams
	ElemPrefix  string
}

func NewHelpPopupTemplateParams(elemPrefix string, panelID string) HelpPopupPropertyTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Popup Help Message",
		PanelID:          panelID}

	params := HelpPopupPropertyTemplateParams{ElemPrefix: elemPrefix,
		PanelParams: panelParams}

	return params
}
