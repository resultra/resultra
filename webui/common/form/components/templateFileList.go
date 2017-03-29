package components

import (
	"resultra/datasheet/webui/common/form/components/caption"
	"resultra/datasheet/webui/common/form/components/checkBox"
	"resultra/datasheet/webui/common/form/components/comment"
	"resultra/datasheet/webui/common/form/components/common"
	"resultra/datasheet/webui/common/form/components/datePicker"
	"resultra/datasheet/webui/common/form/components/formButton"
	"resultra/datasheet/webui/common/form/components/gauge"
	"resultra/datasheet/webui/common/form/components/header"
	"resultra/datasheet/webui/common/form/components/htmlEditor"
	"resultra/datasheet/webui/common/form/components/image"
	"resultra/datasheet/webui/common/form/components/progress"
	"resultra/datasheet/webui/common/form/components/rating"
	"resultra/datasheet/webui/common/form/components/selection"
	"resultra/datasheet/webui/common/form/components/textBox"
	"resultra/datasheet/webui/common/form/components/userSelection"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = []string{"static/common/form/components/include.html",
		"static/common/form/components/properties.html"}

	TemplateFileList = append(TemplateFileList, common.TemplateFileList...)

	TemplateFileList = append(TemplateFileList, checkBox.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, datePicker.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, htmlEditor.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, image.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, textBox.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, header.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, rating.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, selection.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, userSelection.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, comment.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, formButton.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, progress.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, caption.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, gauge.TemplateFileList...)

}
