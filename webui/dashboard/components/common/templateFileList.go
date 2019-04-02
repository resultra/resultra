// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"github.com/resultra/resultra/webui/dashboard/components/common/componentTitle"
	"github.com/resultra/resultra/webui/dashboard/components/common/delete"
	"github.com/resultra/resultra/webui/dashboard/components/common/newComponentDialog"
	"github.com/resultra/resultra/webui/dashboard/components/common/valueGrouping"
	"github.com/resultra/resultra/webui/dashboard/components/common/valueSummary"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, componentTitle.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, newComponentDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueGrouping.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueSummary.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, delete.TemplateFileList...)

}
