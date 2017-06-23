package colProps

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type FormButtonColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newFormButtonTemplateParams() FormButtonColPropsTemplateParams {

	elemPrefix := "button_"

	templParams := FormButtonColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "formButtonLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "formButtonPerms")}

	return templParams

}
