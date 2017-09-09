package file

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type FileDesignTemplateParams struct {
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

type FileViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams FileDesignTemplateParams
var ViewTemplateParams FileViewTemplateParams

func init() {

	elemPrefix := "file_"

	DesignTemplateParams = FileDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "fileLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "fileDelete", "Delete File Attachment"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "fileHelp"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "fileFormat"},
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "fileClearValue"},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "fileValidation"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "fileVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "filePerms"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Email Address Input",
			FieldInfoPrompt:    `File attachments are stored in fields. Either a new field can be created, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the file.'`}}

	ViewTemplateParams = FileViewTemplateParams{
		ElemPrefix: elemPrefix}

}
