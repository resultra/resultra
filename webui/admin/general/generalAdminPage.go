package general

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/userRole"
	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var generalTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/general/generalAdminPage.html",
		"static/admin/general/generalProperties.html",
		"static/admin/general/saveTemplate.html",
		"static/admin/general/trackerDescription.html",
		"static/admin/general/activeTracker.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	generalTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	CurrUserIsAdmin bool
}

func generalAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	currUserIsAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := TemplParams{
		Title:           "General Settings",
		DatabaseID:      databaseID,
		DatabaseName:    dbInfo.DatabaseName,
		CurrUserIsAdmin: currUserIsAdmin}

	if err := generalTemplates.ExecuteTemplate(w, "generalAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
