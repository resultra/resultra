package colProps

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type CheckBoxColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newCheckBoxTemplateParams() CheckBoxColPropsTemplateParams {

	elemPrefix := "checkBox_"

	templParams := CheckBoxColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "checkBoxLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "checkBoxPerms")}

	return templParams

}
