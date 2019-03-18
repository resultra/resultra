// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package general

import (
	"html/template"
	"net/http"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
	adminCommon "resultra/tracker/webui/workspaceAdmin/common"
)

var generalTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/workspaceAdmin/general/workspaceAdminPage.html",
		"static/workspaceAdmin/general/workspaceNameProperty.html",
		"static/workspaceAdmin/general/userAccountProperties.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList,
		adminCommon.TemplateFileList}

	generalTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title         string
	WorkspaceName string
}

func workspaceAdminPage(w http.ResponseWriter, r *http.Request) {

	userInfo, authErr := userAuth.GetCurrentUserInfo(r)
	if (authErr != nil) || (!userInfo.IsWorkspaceAdmin) {
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

	templParams := TemplParams{
		Title:         "General Settings",
		WorkspaceName: workspaceName}

	if err := generalTemplates.ExecuteTemplate(w, "workspaceAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
