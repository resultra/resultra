package generic

import (
	"resultra/tracker/webui/generic/confirmDelete"
	"resultra/tracker/webui/generic/gauge"
	"resultra/tracker/webui/generic/propertiesSidebar"
	"resultra/tracker/webui/generic/valueFormat"
	"resultra/tracker/webui/generic/wizardDialog"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, propertiesSidebar.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, wizardDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, gauge.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueFormat.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, confirmDelete.TemplateFileList...)

}
