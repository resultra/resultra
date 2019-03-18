package common

import (
	"resultra/tracker/webui/dashboard/components/common/componentTitle"
	"resultra/tracker/webui/dashboard/components/common/delete"
	"resultra/tracker/webui/dashboard/components/common/newComponentDialog"
	"resultra/tracker/webui/dashboard/components/common/valueGrouping"
	"resultra/tracker/webui/dashboard/components/common/valueSummary"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, componentTitle.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newComponentDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueGrouping.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueSummary.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, delete.TemplateFileList...)

}
