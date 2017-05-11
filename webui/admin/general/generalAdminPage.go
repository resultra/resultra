package general

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"

	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var generalTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/general/generalAdminPage.html",
		"static/admin/general/generalProperties.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	generalTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
}

func generalAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	templParams := TemplParams{
		Title:        "General Settings",
		DatabaseID:   databaseID,
		DatabaseName: dbInfo.DatabaseName}

	if err := generalTemplates.ExecuteTemplate(w, "generalAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
