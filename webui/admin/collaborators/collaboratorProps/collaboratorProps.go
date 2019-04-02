// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package collaboratorProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/userRole"
	adminCommon "github.com/resultra/resultra/webui/admin/common"

	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/workspace"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
)

var userRoleTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/collaborators/collaboratorProps/collaboratorProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	userRoleTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserRoleTemplParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	WorkspaceName   string
	UserID          string
	CollaboratorID  string
	UserName        string
	CurrUserIsAdmin bool
}

func editCollabPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	collaboratorID := vars["collaboratorID"]
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

	collabInfo, collabErr := userRole.GetCollaboratorByID(trackerDBHandle, collaboratorID)
	if collabErr != nil {
		http.Error(w, collabErr.Error(), http.StatusInternalServerError)
		return

	}

	userInfo, err := userAuth.GetUserInfoByID(trackerDBHandle, collabInfo.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
		return
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	//	elemPrefix := "userRole_"
	templParams := UserRoleTemplParams{
		Title:           "Collaborator Settings",
		DatabaseID:      dbInfo.DatabaseID,
		DatabaseName:    dbInfo.DatabaseName,
		WorkspaceName:   workspaceName,
		UserID:          collabInfo.UserID,
		CollaboratorID:  collaboratorID,
		UserName:        userInfo.UserName,
		CurrUserIsAdmin: isAdmin}

	if err := userRoleTemplates.ExecuteTemplate(w, "collabPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
