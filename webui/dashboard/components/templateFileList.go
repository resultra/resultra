package components

import (
	"resultra/tracker/webui/dashboard/components/barChart"
	"resultra/tracker/webui/dashboard/components/common"
	"resultra/tracker/webui/dashboard/components/gauge"
	"resultra/tracker/webui/dashboard/components/header"
	"resultra/tracker/webui/dashboard/components/summaryTable"
	"resultra/tracker/webui/dashboard/components/summaryValue"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, header.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, barChart.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, summaryTable.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, gauge.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, common.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, summaryValue.TemplateFileList...)

}
