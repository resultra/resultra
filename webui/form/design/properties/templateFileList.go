package properties

import (
	"resultra/datasheet/webui/form/design/properties/recordFiltering"
	"resultra/datasheet/webui/form/design/properties/userRole"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/form/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, userRole.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordFiltering.TemplateFileList...)

}
