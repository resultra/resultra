package properties

import (
	"resultra/tracker/webui/admin/forms/design/properties/formName"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/admin/forms/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, formName.TemplateFileList...)

}
