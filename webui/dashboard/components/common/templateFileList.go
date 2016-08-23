package common

import (
	"resultra/datasheet/webui/dashboard/components/common/componentTitle"
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, componentTitle.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newComponentDialog.TemplateFileList...)

}
