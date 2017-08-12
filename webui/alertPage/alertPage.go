package alertPage

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

var alertPageTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/alertPage/alertPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	alertPageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type AlertPageTemplateParams struct {
	Title           string
	CurrUserIsAdmin bool
	DatabaseID      string
	DatabaseName    string
}

func alertPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	databaseID := vars["databaseID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := alertPageTemplates.ExecuteTemplate(w, "userSignInPage", nil)
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

		templParams := AlertPageTemplateParams{Title: "Alerts",
			DatabaseID:      dbInfo.DatabaseID,
			CurrUserIsAdmin: isAdmin,
			DatabaseName:    dbInfo.DatabaseName}

		if err := alertPageTemplates.ExecuteTemplate(w, "alertPage", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
