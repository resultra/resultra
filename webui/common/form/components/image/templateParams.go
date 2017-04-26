package image

import (
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ImageDesignTemplateParams struct {
	ElemPrefix               string
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
}

type ImageViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams ImageDesignTemplateParams
var ViewTemplateParams ImageViewTemplateParams

func init() {

	elemPrefix := "image_"

	DesignTemplateParams = ImageDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "attachmentLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "attachmentPerms"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "attachDelete", "Delete Attachment Box"),
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Validation", PanelID: "attachmentValidation"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Image Area",
			FieldInfoPrompt: `Images are stored in fields. Either a new field can be created to store the images, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the images.`}}

	ViewTemplateParams = ImageViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "imageTimeline"}}

}
