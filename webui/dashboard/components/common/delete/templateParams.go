package delete

import (
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type DeletePropertyPanelTemplateParams struct {
	PanelParams       propertiesSidebar.PanelTemplateParams
	ElemPrefix        string
	DeleteButtonLabel string
}

func NewDeletePropertyPanelTemplateParams(elemPrefix string, panelID string, deleteButtonLabel string) DeletePropertyPanelTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Delete",
		PanelID:          panelID}

	params := DeletePropertyPanelTemplateParams{ElemPrefix: elemPrefix,
		PanelParams:       panelParams,
		DeleteButtonLabel: deleteButtonLabel}

	return params
}
