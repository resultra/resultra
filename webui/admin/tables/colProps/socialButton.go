package colProps

import (
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type SocialButtonColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newSocialButtonTemplateParams() SocialButtonColPropsTemplateParams {

	elemPrefix := "socialButton_"

	templParams := SocialButtonColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "socialButtonLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "socialButtonPerms")}

	return templParams

}
