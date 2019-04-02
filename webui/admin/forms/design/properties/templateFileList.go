// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package properties

import (
	"github.com/resultra/resultra/webui/admin/forms/design/properties/formName"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/admin/forms/design/properties/properties.html"}

	TemplateFileList = append(TemplateFileList, formName.TemplateFileList...)

}
