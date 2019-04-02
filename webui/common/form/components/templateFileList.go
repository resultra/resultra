// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package components

import (
	"github.com/resultra/resultra/webui/common/form/components/attachment"
	"github.com/resultra/resultra/webui/common/form/components/caption"
	"github.com/resultra/resultra/webui/common/form/components/checkBox"
	"github.com/resultra/resultra/webui/common/form/components/comment"
	"github.com/resultra/resultra/webui/common/form/components/common"
	"github.com/resultra/resultra/webui/common/form/components/datePicker"
	"github.com/resultra/resultra/webui/common/form/components/emailAddr"
	"github.com/resultra/resultra/webui/common/form/components/file"
	"github.com/resultra/resultra/webui/common/form/components/formButton"
	"github.com/resultra/resultra/webui/common/form/components/gauge"
	"github.com/resultra/resultra/webui/common/form/components/header"
	"github.com/resultra/resultra/webui/common/form/components/htmlEditor"
	"github.com/resultra/resultra/webui/common/form/components/image"
	"github.com/resultra/resultra/webui/common/form/components/label"
	"github.com/resultra/resultra/webui/common/form/components/numberInput"
	"github.com/resultra/resultra/webui/common/form/components/progress"
	"github.com/resultra/resultra/webui/common/form/components/rating"
	"github.com/resultra/resultra/webui/common/form/components/selection"
	"github.com/resultra/resultra/webui/common/form/components/socialButton"
	"github.com/resultra/resultra/webui/common/form/components/textBox"
	"github.com/resultra/resultra/webui/common/form/components/toggle"
	"github.com/resultra/resultra/webui/common/form/components/urlLink"
	"github.com/resultra/resultra/webui/common/form/components/userSelection"
	"github.com/resultra/resultra/webui/common/form/components/userTag"
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
