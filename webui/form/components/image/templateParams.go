package image

import (
	"resultra/datasheet/webui/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ImageDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
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
		ElemPrefix:        elemPrefix,
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "imageFormat"},
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
