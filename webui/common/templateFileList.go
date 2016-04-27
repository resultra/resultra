package common

import (
	"resultra/datasheet/webui/common/objectSelection"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/common/include.html"}

	TemplateFileList = append(TemplateFileList, objectSelection.TemplateFileList...)

}
