// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package users

import (
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

var formsTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/workspaceAdmin/users/userManagement.html",
		"static/workspaceAdmin/users/userRegistrationProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	formsTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserTemplParams struct {
	Title         string
	WorkspaceName string
}

func userAdminPage(w http.ResponseWriter, r *http.Request) {

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

	templParams := UserTemplParams{
		Title:         "User Management",
		WorkspaceName: workspaceName}

	if err := formsTemplates.ExecuteTemplate(w, "userMgmtAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
