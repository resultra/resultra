// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, newFormElemDialog.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, label.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, visibility.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, permissions.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, delete.TemplateFileList...)
}
