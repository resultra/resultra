// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package homePage

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var homePageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/homePage/homePageSignedIn.html",
		"static/homePage/trackerList.html",
		"static/homePage/newTrackerDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	homePageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/homePage", home)
	mainRouter.HandleFunc("/homePageSignedOut", home)
	mainRouter.HandleFunc("/homePage/offPageContent", homePageOffPageContent)
}

type PageInfo struct {
	Title                    string `json:"title"`
	WorkspaceName            string `json:"workspaceName"`
	CurrUserIsWorkspaceAdmin bool   `json:"currUserIsWorkspaceAdmin"`
	IsSingleUserWorkspace    bool   `json:"isSingleUserWorkspace"`
}

func home(respWriter http.ResponseWriter, req *http.Request) {

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

	isSingleUser := runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace

	if isSingleUser {
		authResp := userAuth.LoginSingleUser(respWriter, req)

		if !authResp.Success {
			log.Printf("Ping: single user not logged int: %v", authResp.Msg)
		}

	}

	log.Printf("workspace name: %v", workspaceName)

	userInfo, authErr := userAuth.GetCurrentUserInfo(req)
	if authErr != nil {
		http.Error(respWriter, authErr.Error(), http.StatusInternalServerError)
		return
	} else {

		templParams := PageInfo{
			Title:                    "Resultra Workspace - Signed In",
			WorkspaceName:            workspaceName,
			CurrUserIsWorkspaceAdmin: userInfo.IsWorkspaceAdmin,
			IsSingleUserWorkspace:    isSingleUser}
		err := homePageTemplates.ExecuteTemplate(respWriter, "homePageSignedIn", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}

type OffPageContentTemplParams struct{}

func homePageOffPageContent(respWriter http.ResponseWriter, req *http.Request) {

	templParams := OffPageContentTemplParams{}
	err := homePageTemplates.ExecuteTemplate(respWriter, "homePageOffPageContent", templParams)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
