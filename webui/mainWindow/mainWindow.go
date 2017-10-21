package mainWindow

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var mainWindowTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/mainWindow/mainWindow.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		common.TemplateFileList}
	mainWindowTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type MainWindowTemplateParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	CurrUserIsAdmin bool
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/mainWindow/{databaseID}", viewMainWindow)
}

func viewMainWindow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	databaseID := vars["databaseID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := mainWindowTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {

		dbInfo, getErr := databaseController.GetDatabaseInfo(databaseID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

		templParams := MainWindowTemplateParams{Title: "Main Window",
			DatabaseID:      dbInfo.DatabaseID,
			DatabaseName:    dbInfo.DatabaseName,
			CurrUserIsAdmin: isAdmin}

		if err := mainWindowTemplates.ExecuteTemplate(w, "mainWindow", templParams); err != nil {
			api.WriteErrorResponse(w, err)
		}

	}
}
