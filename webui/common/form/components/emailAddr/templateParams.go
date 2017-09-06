package emailAddr

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type EmailAddrDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	ClearValuePanelParams    propertiesSidebar.PanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type EmailAddrViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams EmailAddrDesignTemplateParams
var ViewTemplateParams EmailAddrViewTemplateParams

func init() {

	elemPrefix := "emailAddr_"

	DesignTemplateParams = EmailAddrDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "emailAddrLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "emailAddrDelete", "Delete Email Address Input"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "emailAddrHelp"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "emailAddrFormat"},
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "emailAddrClearValue"},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "emailAddrValidation"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "emailAddrVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "emailAddrPerms"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Email Address Input",
			FieldInfoPrompt:    `Email addresses are stored in fields. Either a new field can be created, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the email addresses.'`}}

	ViewTemplateParams = EmailAddrViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "emailAddrTimeline"}}

}
