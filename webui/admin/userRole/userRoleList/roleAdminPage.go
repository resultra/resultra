// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRoleList

import (
	"html/template"
	"net/http"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/databaseController"

	"github.com/gorilla/mux"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	"resultra/tracker/server/workspace"
	adminCommon "resultra/tracker/webui/admin/common"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var roleTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/userRole/userRoleList/dashboardPrivs.html",
		"static/admin/userRole/userRoleList/formPrivs.html",
		"static/admin/userRole/userRoleList/newUserRoleDialog.html",
		"static/admin/userRole/userRoleList/roleAdminPage.html",
		"static/admin/userRole/userRoleList/userRole.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	roleTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type RoleTemplParams struct {
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	IsSingleUserWorkspace bool
	CurrUserIsAdmin       bool
}

func roleAdminPage(w http.ResponseWriter, r *http.Request) {

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

	templParams := RoleTemplParams{
		Title:                 "Role",
		DatabaseID:            databaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),
		CurrUserIsAdmin:       isAdmin}

	if err := roleTemplates.ExecuteTemplate(w, "roleAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
