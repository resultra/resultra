package tableList

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/userRole"

	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var tableTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/tables/tableList/tableAdminPage.html",
		"static/admin/tables/tableList/tableList.html",
		"static/admin/tables/tableList/newTableDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	tableTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	CurrUserIsAdmin bool
}

func tableAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := TemplParams{
		Title:           "Tables",
		DatabaseID:      databaseID,
		DatabaseName:    dbInfo.DatabaseName,
		CurrUserIsAdmin: isAdmin}

	if err := tableTemplates.ExecuteTemplate(w, "tableAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
