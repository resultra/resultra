// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"html/template"
	"net/http"
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
)

var contentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/common/settingsTOC.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}

	contentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TOCContentTemplParams struct {
	IsSingleUserWorkspace bool
}

func tocContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := TOCContentTemplParams{
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace()}

	if err := contentTemplates.ExecuteTemplate(w, "adminSettingsTOC", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
