package properties

import (
	"resultra/datasheet/webui/admin/forms/design/properties/formName"
	"resultra/datasheet/webui/admin/forms/design/properties/newItem"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/admin/forms/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, formName.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newItem.TemplateFileList...)

}
