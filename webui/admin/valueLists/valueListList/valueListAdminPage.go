package valueListList

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

var valueListTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/valueLists/valueListList/valueListAdminPage.html",
		"static/admin/valueLists/valueListList/valueListList.html",
		"static/admin/valueLists/valueListList/newValueListDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	valueListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FieldTemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
}

func valueListAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	templParams := FieldTemplParams{
		Title:        "Forms",
		DatabaseID:   databaseID,
		DatabaseName: dbInfo.DatabaseName}

	if err := valueListTemplates.ExecuteTemplate(w, "valueListAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
