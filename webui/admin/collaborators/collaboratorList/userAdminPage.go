package collaboratorList

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"

	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var userTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/collaborators/collaboratorList/userAdminPage.html",
		"static/admin/collaborators/collaboratorList/users.html",
		"static/admin/collaborators/collaboratorList/addUserDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	userTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
}

func userAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	templParams := TemplParams{
		Title:        "Collaborators",
		DatabaseID:   databaseID,
		DatabaseName: dbInfo.DatabaseName}

	if err := userTemplates.ExecuteTemplate(w, "userAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
