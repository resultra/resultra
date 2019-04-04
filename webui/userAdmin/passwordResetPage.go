// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAdmin

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"html/template"
	"log"
	"net/http"
	//	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/workspace"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
)

var homePageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/userAdmin/passwordResetPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	homePageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type PageInfo struct {
	Title         string `json:"title"`
	WorkspaceName string `json:"workspaceName"`
	ResetID       string `json:"resetID"`
}

func resetPassword(respWriter http.ResponseWriter, req *http.Request) {

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

	vars := mux.Vars(req)

	resetID := vars["resetID"]

	log.Printf("password reset: reset id = %v", resetID)

	templParams := PageInfo{
		Title:         "Resultra Workspace - Reset Password",
		WorkspaceName: workspaceName,
		ResetID:       resetID}
	err := homePageTemplates.ExecuteTemplate(respWriter, "resetPasswordPage", templParams)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
