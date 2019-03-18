package attachment

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/form/components/common/delete"
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/newFormElemDialog"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type ImageDesignTemplateParams struct {
	ElemPrefix               string
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type ImageViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams ImageDesignTemplateParams
var ViewTemplateParams ImageViewTemplateParams

func init() {

	elemPrefix := "attachment_"

	DesignTemplateParams = ImageDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "attachmentLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "attachmentPerms"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "attachmentHelp"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "attachDelete", "Delete Attachment Box"),
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Validation", PanelID: "attachmentValidation"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Attachment Area",
			FieldInfoPrompt: `Attachments are stored in fields. Either a new field can be created to store the attachments, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the attachments.`}}

	ViewTemplateParams = ImageViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "imageTimeline"}}

}
