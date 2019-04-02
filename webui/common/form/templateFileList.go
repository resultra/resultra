// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package form

import (
	"github.com/resultra/resultra/webui/common/form/components"
	"github.com/resultra/resultra/webui/common/form/submit"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, components.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, submit.TemplateFileList...)

}
