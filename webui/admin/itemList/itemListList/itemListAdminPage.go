package itemListList

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

var formsTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/itemList/itemListList/itemListAdminPage.html",
		"static/admin/itemList/itemListList/listList.html",
		"static/admin/itemList/itemListList/newListDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	formsTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FieldTemplParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	CurrUserIsAdmin bool
}

func itemListAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	currUserIsAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := FieldTemplParams{
		Title:           "Forms",
		DatabaseID:      databaseID,
		DatabaseName:    dbInfo.DatabaseName,
		CurrUserIsAdmin: currUserIsAdmin}

	if err := formsTemplates.ExecuteTemplate(w, "itemListAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
