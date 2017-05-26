package colProps

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type TextInputColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newTextInputTemplateParams() TextInputColPropsTemplateParams {

	elemPrefix := "textInput_"

	templParams := TextInputColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "numberInputLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "numberInputPerms")}

	return templParams

}
