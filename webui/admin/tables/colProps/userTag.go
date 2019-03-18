package colProps

import (
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type UserTagColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newUserTagTemplateParams() UserTagColPropsTemplateParams {

	elemPrefix := "userTag_"

	templParams := UserTagColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "userTagLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "userTagPerms")}

	return templParams

}
