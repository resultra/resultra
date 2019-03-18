// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package generic

import (
	"resultra/tracker/webui/generic/confirmDelete"
	"resultra/tracker/webui/generic/gauge"
	"resultra/tracker/webui/generic/propertiesSidebar"
	"resultra/tracker/webui/generic/valueFormat"
	"resultra/tracker/webui/generic/wizardDialog"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, propertiesSidebar.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, wizardDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, gauge.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueFormat.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, confirmDelete.TemplateFileList...)

}
