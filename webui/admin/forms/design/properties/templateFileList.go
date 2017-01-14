package properties

import (
	"resultra/datasheet/webui/admin/forms/design/properties/formName"
	"resultra/datasheet/webui/admin/forms/design/properties/newItem"
	"resultra/datasheet/webui/admin/forms/design/properties/userRole"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/admin/forms/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, userRole.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, formName.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newItem.TemplateFileList...)

}
