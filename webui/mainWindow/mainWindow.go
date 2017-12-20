package mainWindow

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/alertListView"
	"resultra/datasheet/webui/common"
	dashboardCommon "resultra/datasheet/webui/dashboard/common"
	dashboardComponents "resultra/datasheet/webui/dashboard/components"
	dashboardView "resultra/datasheet/webui/dashboard/view"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/itemList"
	"resultra/datasheet/webui/itemView"
	"resultra/datasheet/webui/thirdParty"
)

var mainWindowTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/mainWindow/mainWindow.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		common.TemplateFileList,
		itemList.TemplateFileList,
		dashboardComponents.TemplateFileList,
		dashboardCommon.TemplateFileList,
		dashboardView.TemplateFileList,
		alertListView.TemplateFileList,
		itemView.TemplateFileList}

	mainWindowTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type MainWindowTemplateParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	WorkspaceName   string
	CurrUserIsAdmin bool
	ItemListParams  itemList.ViewListTemplateParams
	DashboardParams dashboardView.ViewDashboardTemplateParams
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/main/{databaseID}", viewMainWindow)
}

func viewMainWindow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	databaseID := vars["databaseID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {

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

		dbInfo, getErr := databaseController.GetDatabaseInfo(trackerDBHandle, databaseID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

		templParams := MainWindowTemplateParams{
			Title:           dbInfo.DatabaseName,
			DatabaseID:      dbInfo.DatabaseID,
			DatabaseName:    dbInfo.DatabaseName,
			CurrUserIsAdmin: isAdmin,
			ItemListParams:  itemList.ViewListTemplParams,
			WorkspaceName:   workspaceName,
			DashboardParams: dashboardView.ViewTemplateParams}

		if err := mainWindowTemplates.ExecuteTemplate(w, "mainWindow", templParams); err != nil {
			api.WriteErrorResponse(w, err)
		}

	}
}
