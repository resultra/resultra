// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package valueListList

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/databaseController"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	"resultra/tracker/server/workspace"
	adminCommon "resultra/tracker/webui/admin/common"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var valueListTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/valueLists/valueListList/valueListAdminPage.html",
		"static/admin/valueLists/valueListList/valueListList.html",
		"static/admin/valueLists/valueListList/newValueListDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	valueListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FieldTemplParams struct {
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func valueListAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := FieldTemplParams{
		Title:                 "Value Lists",
		DatabaseID:            databaseID,
		DatabaseName:          dbInfo.DatabaseName,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
		WorkspaceName:         workspaceName,
		CurrUserIsAdmin:       isAdmin}

	if err := valueListTemplates.ExecuteTemplate(w, "valueListAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
