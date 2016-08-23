package components

import (
	"resultra/datasheet/webui/dashboard/components/barChart"
	"resultra/datasheet/webui/dashboard/components/common"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, barChart.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, common.TemplateFileList...)

}
