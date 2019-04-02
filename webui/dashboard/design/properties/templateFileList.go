// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package properties

import (
	"github.com/resultra/resultra/webui/dashboard/design/properties/dashboardName"
	"github.com/resultra/resultra/webui/dashboard/design/properties/includeInSidebar"
	"github.com/resultra/resultra/webui/dashboard/design/properties/userRole"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/dashboard/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, dashboardName.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, includeInSidebar.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, userRole.TemplateFileList...)

}
