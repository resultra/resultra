package generic

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
	"resultra/datasheet/webui/generic/userAuth"
	"resultra/datasheet/webui/generic/wizardDialog"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/generic/include.html"}

	TemplateFileList = append(TemplateFileList, propertiesSidebar.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, wizardDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, userAuth.TemplateFileList...)

}
