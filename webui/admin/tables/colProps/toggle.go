package colProps

import (
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type ToggleColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newToggleTemplateParams() ToggleColPropsTemplateParams {

	elemPrefix := "toggle_"

	templParams := ToggleColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "toggleLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "togglePerms")}

	return templParams

}
