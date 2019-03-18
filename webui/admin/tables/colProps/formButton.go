package colProps

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/defaultValues"
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type FormButtonColPropsTemplateParams struct {
	ElemPrefix              string
	LabelPanelParams        label.LabelPropertyTemplateParams
	ButtonLabelPanelParams  inputProperties.FormButtonLabelPropertyTemplateParams
	PermissionPanelParams   permissions.PermissionsPropertyTemplateParams
	DefaultValuePanelParams defaultValues.DefaultValuesPanelTemplateParams
}

func newFormButtonTemplateParams() FormButtonColPropsTemplateParams {

	elemPrefix := "button_"

	templParams := FormButtonColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "formButtonLabel"}},

		ButtonLabelPanelParams: inputProperties.FormButtonLabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Button Label", PanelID: "formButtonButtonLabel"}},

		PermissionPanelParams:   permissions.NewPermissionTemplateParams(elemPrefix, "formButtonPerms"),
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	return templParams

}
