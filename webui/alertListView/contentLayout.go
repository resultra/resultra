// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package alertListView

import (
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/dashboard/components"
	"github.com/resultra/resultra/webui/generic"
	"html/template"
	"net/http"
)

var contentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/alertListView/contentLayout.html",
		"static/alertListView/notificationList.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		components.TemplateFileList}

	contentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ContentTemplParams struct{}

func alertListContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := ContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "alertListContentLayout", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
