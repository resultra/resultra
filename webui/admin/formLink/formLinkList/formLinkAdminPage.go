package formLinkList

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"

	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var formLinkTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/formLink/formLinkList/formLinkAdminPage.html",
		"static/admin/formLink/formLinkList/formLinkList.html",
		"static/admin/formLink/formLinkList/newFormLinkDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	formLinkTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
}

func formLinkAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	templParams := TemplParams{
		Title:        "Forms",
		DatabaseID:   databaseID,
		DatabaseName: dbInfo.DatabaseName}

	if err := formLinkTemplates.ExecuteTemplate(w, "formLinkAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
