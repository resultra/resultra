// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRoleProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/databaseController"
	adminCommon "resultra/tracker/webui/admin/common"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var userRoleTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/userRole/userRoleProps/editRoleProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		adminCommon.TemplateFileList,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	userRoleTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserRoleTemplParams struct {
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	RoleID                string
	RoleName              string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func editRolePropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	roleID := vars["roleID"]

	log.Println("editRolePropsPage: viewing/editing admin settings for role ID = ", roleID)

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

	roleInfo, err := userRole.GetUserRole(trackerDBHandle, roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("editRolePropsPage: viewing/editing admin settings for role = %+v", roleInfo)

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, roleInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
		return
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	//	elemPrefix := "userRole_"
	templParams := UserRoleTemplParams{
		Title:                 "User Role Settings",
		DatabaseID:            roleInfo.ParentDatabaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		RoleID:                roleID,
		RoleName:              roleInfo.RoleName,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
		CurrUserIsAdmin:       isAdmin}

	if err := userRoleTemplates.ExecuteTemplate(w, "editUserRolePropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
