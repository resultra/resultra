package admin

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/webui/admin/dashboards"
	"resultra/datasheet/webui/admin/formLink/formLinkList"
	"resultra/datasheet/webui/admin/general"
	"resultra/datasheet/webui/admin/globals"
	"resultra/datasheet/webui/admin/userRole/userRoleList"
	"resultra/datasheet/webui/admin/users"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var adminTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/admin.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		userRole.TemplateFileList,
		users.TemplateFileList,
		formLinkList.TemplateFileList,
		dashboards.TemplateFileList,
		general.TemplateFileList,
		globals.TemplateFileList}
	adminTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type AdminTemplParams struct {
	DatabaseID   string
	DatabaseName string
	Title        string
}

func adminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]
	log.Println("Admin page: viewing/editing admin settings for database ID = ", databaseID)

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	templParams := AdminTemplParams{
		Title:        "Tracker Settings",
		DatabaseName: dbInfo.DatabaseName,
		DatabaseID:   databaseID}

	err := adminTemplates.ExecuteTemplate(w, "adminPage", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
