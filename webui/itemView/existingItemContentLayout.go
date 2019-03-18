// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemView

import (
	"html/template"
	"net/http"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
)

var existingContentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/itemView/existingItemContentLayout.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		common.TemplateFileList}

	existingContentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ExistingContentTemplParams struct{}

func existingItemContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := ExistingContentTemplParams{}

	if err := existingContentTemplates.ExecuteTemplate(w, "existingItemContentLayout", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func existingItemOffPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := ExistingContentTemplParams{}

	if err := existingContentTemplates.ExecuteTemplate(w, "existingItemOffPageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
