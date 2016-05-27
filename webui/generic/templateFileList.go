package generic

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/generic/include.html"}

	TemplateFileList = append(TemplateFileList, propertiesSidebar.TemplateFileList...)

}
