// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemView

import (
	"html/template"
	"net/http"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
)

var contentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/itemView/newItemContentLayout.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		common.TemplateFileList}

	contentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ContentTemplParams struct{}

func newItemContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := ContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "newItemContentLayout", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func newItemOffPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := ContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "newItemOffPageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
