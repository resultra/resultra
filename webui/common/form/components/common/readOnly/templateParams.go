package readOnly

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ReadOnlyPropertyTemplateParams struct {
	PanelParams propertiesSidebar.PanelTemplateParams
	ElemPrefix  string
}

func NewReadOnlyTemplateParams(elemPrefix string, panelID string) ReadOnlyPropertyTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Read Only",
		PanelID:          panelID}

	params := ReadOnlyPropertyTemplateParams{ElemPrefix: elemPrefix,
		PanelParams: panelParams}

	return params
}
