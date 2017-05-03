package components

import (
	"resultra/datasheet/webui/dashboard/components/barChart"
	"resultra/datasheet/webui/dashboard/components/common"
	"resultra/datasheet/webui/dashboard/components/gauge"
	"resultra/datasheet/webui/dashboard/components/header"
	"resultra/datasheet/webui/dashboard/components/summaryTable"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, header.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, barChart.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, summaryTable.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, gauge.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, common.TemplateFileList...)

}
