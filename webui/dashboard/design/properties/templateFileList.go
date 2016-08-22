package properties

import (
	"resultra/datasheet/webui/dashboard/design/properties/dashboardName"
	"resultra/datasheet/webui/dashboard/design/properties/userRole"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/dashboard/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, dashboardName.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, userRole.TemplateFileList...)

}
