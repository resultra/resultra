package trackerMainPage

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var pageTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/trackerMainPage/trackerMainPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}
	pageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	DatabaseID   string
	DatabaseName string
	Title        string
}

func trackerMainPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]
	log.Println("Admin page: viewing/editing admin settings for database ID = ", databaseID)

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	templParams := TemplParams{
		Title:        "Tracker Main Page",
		DatabaseName: dbInfo.DatabaseName,
		DatabaseID:   databaseID}

	err := pageTemplates.ExecuteTemplate(w, "trackerMainPage", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
