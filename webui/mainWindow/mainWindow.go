package mainWindow

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"resultra/datasheet/server/common/runtimeConfig"
	"resultra/datasheet/server/common/userAuth"
	//	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	//	"resultra/datasheet/server/userRole"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/alertListView"
	"resultra/datasheet/webui/common"
	dashboardComponents "resultra/datasheet/webui/dashboard/components"
	dashboardView "resultra/datasheet/webui/dashboard/view"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/itemList"
	"resultra/datasheet/webui/itemView"
	"resultra/datasheet/webui/thirdParty"
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

	isSingleUser := runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace

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
			IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
			/*ItemListParams:        itemList.ViewListTemplParams,*/
			WorkspaceName: workspaceName,
			/*DashboardParams:       dashboardView.ViewTemplateParams*/
		}

		if err := mainWindowTemplates.ExecuteTemplate(w, "mainWindow", templParams); err != nil {
			api.WriteErrorResponse(w, err)
		}

	}
}
