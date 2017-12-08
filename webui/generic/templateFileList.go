package generic

import (
	"resultra/datasheet/webui/generic/confirmDelete"
	"resultra/datasheet/webui/generic/gauge"
	"resultra/datasheet/webui/generic/propertiesSidebar"
	"resultra/datasheet/webui/generic/valueFormat"
	"resultra/datasheet/webui/generic/wizardDialog"
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
