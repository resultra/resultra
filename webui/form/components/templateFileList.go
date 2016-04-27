package components

import (
	"resultra/datasheet/webui/form/components/checkBox"
	"resultra/datasheet/webui/form/components/common/componentEdit"
	"resultra/datasheet/webui/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/form/components/textBox"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = []string{"static/form/components/include.html",
		"static/form/components/properties.html"}

	TemplateFileList = append(TemplateFileList, componentEdit.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newFormElemDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, checkBox.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, textBox.TemplateFileList...)

}
