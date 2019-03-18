package permissions

import (
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type PermissionsPropertyTemplateParams struct {
	PanelParams propertiesSidebar.PanelTemplateParams
	ElemPrefix  string
}

func NewPermissionTemplateParams(elemPrefix string, panelID string) PermissionsPropertyTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Permissions",
		PanelID:          panelID}

	params := PermissionsPropertyTemplateParams{ElemPrefix: elemPrefix,
		PanelParams: panelParams}

	return params
}
