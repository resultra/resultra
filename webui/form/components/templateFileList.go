package components

import (
	"resultra/datasheet/webui/form/components/checkBox"
	"resultra/datasheet/webui/form/components/textBox"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = []string{"static/form/components/jsInclude.html"}

	TemplateFileList = append(TemplateFileList, checkBox.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, textBox.TemplateFileList...)

}
