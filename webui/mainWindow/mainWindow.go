// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package mainWindow

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/generic/api"

	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/workspace"
	"github.com/resultra/resultra/webui/alertListView"
	"github.com/resultra/resultra/webui/common"
	dashboardComponents "github.com/resultra/resultra/webui/dashboard/components"
	dashboardView "github.com/resultra/resultra/webui/dashboard/view"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/itemList"
	"github.com/resultra/resultra/webui/itemView"
	"github.com/resultra/resultra/webui/thirdParty"
)

var mainWindowTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/mainWindow/mainWindow.html",
		"static/mainWindow/mainWindowSignedOut.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		common.TemplateFileList,
		itemList.TemplateFileList,
		dashboardComponents.TemplateFileList,
		dashboardView.TemplateFileList,
		alertListView.TemplateFileList,
		itemView.TemplateFileList}

	mainWindowTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type MainWindowTemplateParams struct {
	Title string
	//	DatabaseID            string
	//	DatabaseName          string
	WorkspaceName string
	//CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
	//	ItemListParams        itemList.ViewListTemplateParams
	//	DashboardParams       dashboardView.ViewDashboardTemplateParams
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/", viewMainWindow)
}

func viewMainWindow(w http.ResponseWriter, r *http.Request) {

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

	isSingleUser := runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace()

	if isSingleUser {
		authResp := userAuth.LoginSingleUser(w, r)

		if !authResp.Success {
			log.Printf("Ping: single user not logged int: %v", authResp.Msg)
		}

	}

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		templParams := MainWindowTemplateParams{
			Title:                 "Resultra Workspace - Signed out",
			WorkspaceName:         workspaceName,
			IsSingleUserWorkspace: isSingleUser}
		err := mainWindowTemplates.ExecuteTemplate(w, "mainWindowSignedOut", templParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {

		templParams := MainWindowTemplateParams{
			Title: workspaceName,
			//			DatabaseID:            dbInfo.DatabaseID,
			//			DatabaseName:          dbInfo.DatabaseName,
			//CurrUserIsAdmin:       isAdmin,
			IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),
			/*ItemListParams:        itemList.ViewListTemplParams,*/
			WorkspaceName: workspaceName,
			/*DashboardParams:       dashboardView.ViewTemplateParams*/
		}

		if err := mainWindowTemplates.ExecuteTemplate(w, "mainWindow", templParams); err != nil {
			api.WriteErrorResponse(w, err)
		}

	}
}
