package components

import (
	"resultra/datasheet/webui/dashboard/components/barChart"
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/dashboard/components/include.html"}

	TemplateFileList = append(TemplateFileList, barChart.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newComponentDialog.TemplateFileList...)

}
