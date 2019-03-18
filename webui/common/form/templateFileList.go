package form

import (
	"resultra/tracker/webui/common/form/components"
	"resultra/tracker/webui/common/form/submit"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, components.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, submit.TemplateFileList...)

}
