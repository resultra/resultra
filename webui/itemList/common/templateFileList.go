package common

import (
	"resultra/datasheet/webui/itemList/common/timeline"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, timeline.TemplateFileList...)

}
