// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemList

import (
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"html/template"
	"net/http"
)

var propertyTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/itemList/properties.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}

	propertyTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func propertySidebarContent(w http.ResponseWriter, r *http.Request) {

	if err := propertyTemplates.ExecuteTemplate(w, "listViewProps", ViewListTemplParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
