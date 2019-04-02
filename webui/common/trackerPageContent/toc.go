// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package trackerPageContent

import (
	"html/template"
	"net/http"
	"github.com/resultra/resultra/webui/generic"
)

var tocTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/common/trackerPageContent/toc.html"}

	templateFileLists := [][]string{baseTemplateFiles}

	tocTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TOCTemplParams struct{}

func tocPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := TOCTemplParams{}

	if err := tocTemplates.ExecuteTemplate(w, "databaseTOC", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
