package form

import (
	"resultra/datasheet/webui/common/form/components"
	"resultra/datasheet/webui/common/form/submit"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, components.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, submit.TemplateFileList...)

}
