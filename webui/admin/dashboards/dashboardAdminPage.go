package dashboards

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"

	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var formsTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/dashboards/dashboardAdminPage.html",
		"static/admin/dashboards/dashboardList.html",
		"static/admin/dashboards/newDashboardDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	formsTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
}

func dashboardAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	templParams := TemplParams{
		Title:        "Dashboards",
		DatabaseID:   databaseID,
		DatabaseName: dbInfo.DatabaseName}

	if err := formsTemplates.ExecuteTemplate(w, "dashboardAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
