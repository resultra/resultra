package tableProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/displayTable"
	"resultra/datasheet/server/userRole"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/server/common/runtimeConfig"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var tablePropTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/tables/tableProps/tablePropsPage.html",
		"static/admin/tables/tableProps/tableCols.html",
		"static/admin/tables/tableProps/newColDialog.html",
		"static/admin/tables/tableProps/newColNewOrExistingFieldDialogPanel.html",
		"static/admin/tables/tableProps/newColNewFieldDialogPanel.html",
		"static/admin/tables/tableProps/newColColTypeDialogPanel.html",
		"static/admin/tables/tableProps/newFormButtonColDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	tablePropTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	ElemPrefix      string
	Title           string
	DatabaseID      string
	DatabaseName    string
	TableID         string
	TableName       string
	SiteBaseURL     string
	CurrUserIsAdmin bool
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/table/{tableID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tableID := vars["tableID"]

	tableInfo, err := displayTable.GetTable(tableID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbInfo, err := databaseController.GetDatabaseInfo(tableInfo.ParentDatabaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	elemPrefix := "tableProps_"
	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := TemplParams{
		ElemPrefix:      elemPrefix,
		Title:           "Table view properties",
		DatabaseID:      dbInfo.DatabaseID,
		DatabaseName:    dbInfo.DatabaseName,
		TableID:         tableID,
		TableName:       tableInfo.Name,
		SiteBaseURL:     runtimeConfig.GetSiteBaseURL(),
		CurrUserIsAdmin: isAdmin}

	if err := tablePropTemplates.ExecuteTemplate(w, "tablePropsAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
