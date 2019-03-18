// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAdmin

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var registrationPageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/userAdmin/userRegistrationPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	registrationPageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type RegistrationPageInfo struct {
	Title            string `json:"title"`
	WorkspaceName    string `json:"workspaceName"`
	InviteID         string `json:"inviteID"`
	InviteeEmailAddr string `json:"InviteeEmailAddr"`
}

func registerNewUser(respWriter http.ResponseWriter, req *http.Request) {

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
	inviteID := vars["inviteID"]

	inviteInfo, inviteErr := userAuth.GetInviteInfo(trackerDBHandle, inviteID)
	if inviteErr != nil {
		http.Error(respWriter, inviteErr.Error(), http.StatusInternalServerError)
		return
	}

	templParams := RegistrationPageInfo{
		Title:            "Resultra Workspace - Register New Account",
		WorkspaceName:    workspaceName,
		InviteID:         inviteID,
		InviteeEmailAddr: inviteInfo.InviteeEmail}
	templErr := registrationPageTemplates.ExecuteTemplate(respWriter, "userRegistrationPage", templParams)
	if templErr != nil {
		http.Error(respWriter, templErr.Error(), http.StatusInternalServerError)
		return
	}

}
