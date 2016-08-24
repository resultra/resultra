package common

import (
	"resultra/datasheet/webui/dashboard/components/common/componentTitle"
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
	"resultra/datasheet/webui/dashboard/components/common/valueGrouping"
	"resultra/datasheet/webui/dashboard/components/common/valueSummary"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, componentTitle.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newComponentDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueGrouping.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueSummary.TemplateFileList...)

}
