package components

import (
	"resultra/tracker/webui/common/form/components/attachment"
	"resultra/tracker/webui/common/form/components/caption"
	"resultra/tracker/webui/common/form/components/checkBox"
	"resultra/tracker/webui/common/form/components/comment"
	"resultra/tracker/webui/common/form/components/common"
	"resultra/tracker/webui/common/form/components/datePicker"
	"resultra/tracker/webui/common/form/components/emailAddr"
	"resultra/tracker/webui/common/form/components/file"
	"resultra/tracker/webui/common/form/components/formButton"
	"resultra/tracker/webui/common/form/components/gauge"
	"resultra/tracker/webui/common/form/components/header"
	"resultra/tracker/webui/common/form/components/htmlEditor"
	"resultra/tracker/webui/common/form/components/image"
	"resultra/tracker/webui/common/form/components/label"
	"resultra/tracker/webui/common/form/components/numberInput"
	"resultra/tracker/webui/common/form/components/progress"
	"resultra/tracker/webui/common/form/components/rating"
	"resultra/tracker/webui/common/form/components/selection"
	"resultra/tracker/webui/common/form/components/socialButton"
	"resultra/tracker/webui/common/form/components/textBox"
	"resultra/tracker/webui/common/form/components/toggle"
	"resultra/tracker/webui/common/form/components/urlLink"
	"resultra/tracker/webui/common/form/components/userSelection"
	"resultra/tracker/webui/common/form/components/userTag"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = []string{"static/common/form/components/include.html",
		"static/common/form/components/properties.html"}

	TemplateFileList = append(TemplateFileList, common.TemplateFileList...)

	TemplateFileList = append(TemplateFileList, checkBox.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, toggle.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, datePicker.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, htmlEditor.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, attachment.TemplateFileList...)
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
	TemplateFileList = append(TemplateFileList, numberInput.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, socialButton.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, label.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, emailAddr.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, urlLink.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, file.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, image.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, userTag.TemplateFileList...)

}
