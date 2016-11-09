package properties

import (
	"resultra/datasheet/webui/form/design/properties/formName"
	"resultra/datasheet/webui/form/design/properties/recordSorting"
	"resultra/datasheet/webui/form/design/properties/userRole"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/form/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, userRole.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, formName.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordSorting.TemplateFileList...)

}
