// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"resultra/tracker/webui/common/alert"
	"resultra/tracker/webui/common/attachment"
	"resultra/tracker/webui/common/conditionalFormat"
	"resultra/tracker/webui/common/defaultValues"
	"resultra/tracker/webui/common/field"
	"resultra/tracker/webui/common/form"
	"resultra/tracker/webui/common/formulaEditor"
	"resultra/tracker/webui/common/helpMenu"
	"resultra/tracker/webui/common/helpPopup"
	"resultra/tracker/webui/common/itemList"
	"resultra/tracker/webui/common/recordFilter"
	"resultra/tracker/webui/common/recordSort"
	"resultra/tracker/webui/common/singlePageApp"
	"resultra/tracker/webui/common/timeline"
	"resultra/tracker/webui/common/trackerPageContent"
	"resultra/tracker/webui/common/userAuth"
	"resultra/tracker/webui/common/valueList"
	"resultra/tracker/webui/common/valueThreshold"
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
