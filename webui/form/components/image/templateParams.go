package image

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ImageTemplateParams struct {
	ElemPrefix        string
	FormatPanelParams propertiesSidebar.PanelTemplateParams
}

var TemplateParams ImageTemplateParams

func init() {
	TemplateParams = ImageTemplateParams{
		ElemPrefix:        "image_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "imageFormat"}}
}
