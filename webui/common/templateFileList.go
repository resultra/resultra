// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"github.com/resultra/resultra/webui/common/alert"
	"github.com/resultra/resultra/webui/common/attachment"
	"github.com/resultra/resultra/webui/common/conditionalFormat"
	"github.com/resultra/resultra/webui/common/defaultValues"
	"github.com/resultra/resultra/webui/common/field"
	"github.com/resultra/resultra/webui/common/form"
	"github.com/resultra/resultra/webui/common/formulaEditor"
	"github.com/resultra/resultra/webui/common/helpMenu"
	"github.com/resultra/resultra/webui/common/helpPopup"
	"github.com/resultra/resultra/webui/common/itemList"
	"github.com/resultra/resultra/webui/common/recordFilter"
	"github.com/resultra/resultra/webui/common/recordSort"
	"github.com/resultra/resultra/webui/common/singlePageApp"
	"github.com/resultra/resultra/webui/common/timeline"
	"github.com/resultra/resultra/webui/common/trackerPageContent"
	"github.com/resultra/resultra/webui/common/userAuth"
	"github.com/resultra/resultra/webui/common/valueList"
	"github.com/resultra/resultra/webui/common/valueThreshold"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, attachment.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, formulaEditor.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordFilter.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, trackerPageContent.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordSort.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, field.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, helpMenu.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, form.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, defaultValues.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueThreshold.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueList.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, itemList.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, conditionalFormat.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, helpPopup.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, alert.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, timeline.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, userAuth.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, singlePageApp.TemplateFileList...)

}
