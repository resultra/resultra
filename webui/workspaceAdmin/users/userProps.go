// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package users

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/workspace"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
	adminCommon "github.com/resultra/resultra/webui/workspaceAdmin/common"
)

var userTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/workspaceAdmin/users/userProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	userTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserPropsTemplParams struct {
	Title         string
	WorkspaceName string
	UserID        string
}

func userPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID := vars["userID"]

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

	templParams := UserPropsTemplParams{
		Title:         "User Account Settings",
		WorkspaceName: workspaceName,
		UserID:        userID}

	if err := userTemplates.ExecuteTemplate(w, "userPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
