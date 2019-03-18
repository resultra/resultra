package colProps

import (
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type UserSelectionColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newUserSelectionTemplateParams() UserSelectionColPropsTemplateParams {

	elemPrefix := "userSelection_"

	templParams := UserSelectionColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "userSelectionLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "userSelectionPerms")}

	return templParams

}
