package common

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/visibility"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, newFormElemDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, label.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, visibility.TemplateFileList...)

}
