package colProps

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type TextSelectionColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newTextSelectionTemplateParams() TextSelectionColPropsTemplateParams {

	elemPrefix := "textSelection_"

	templParams := TextSelectionColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "textSelectionLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "textSelectionPerms")}

	return templParams

}
