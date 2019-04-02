// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package generic

import (
	"github.com/resultra/resultra/webui/generic/confirmDelete"
	"github.com/resultra/resultra/webui/generic/gauge"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
	"github.com/resultra/resultra/webui/generic/valueFormat"
	"github.com/resultra/resultra/webui/generic/wizardDialog"
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
