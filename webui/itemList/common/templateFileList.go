package common

import (
	"resultra/datasheet/webui/itemList/common/display"
	"resultra/datasheet/webui/itemList/common/timeline"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, display.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, timeline.TemplateFileList...)

}
