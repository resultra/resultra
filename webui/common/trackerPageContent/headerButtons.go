// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package trackerPageContent

import (
	"github.com/gorilla/mux"

	"html/template"
	"net/http"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	"resultra/tracker/webui/common/alert"
	"resultra/tracker/webui/generic"
)

var headerTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/common/trackerPageContent/headerButtons.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		alert.TemplateFileList}

	headerTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type HeaderTemplParams struct {
	CurrUserIsAdmin bool
	DatabaseID      string
}

func headerButtonsContent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, databaseID)

	templParams := HeaderTemplParams{
		CurrUserIsAdmin: isAdmin,
		DatabaseID:      databaseID}

	if err := headerTemplates.ExecuteTemplate(w, "trackerHeaderButtons", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
