// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package setupPage

import (
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/workspace"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
	"html/template"
	"net/http"
)

var setupPageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/setupPage/adminSetupPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	setupPageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type SetupPageInfo struct {
	Title         string `json:"title"`
	WorkspaceName string `json:"workspaceName"`
}

func setupAdminUser(respWriter http.ResponseWriter, req *http.Request) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		http.Error(respWriter, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(respWriter, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	templParams := SetupPageInfo{
		Title:         "Setup Resultra",
		WorkspaceName: workspaceName}
	err := setupPageTemplates.ExecuteTemplate(respWriter, "adminSetupPage", templParams)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
