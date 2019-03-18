package tableProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/databaseController"
	"resultra/tracker/server/displayTable"
	"resultra/tracker/server/userRole"
	adminCommon "resultra/tracker/webui/admin/common"

	"resultra/tracker/server/workspace"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
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
	ElemPrefix            string
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	TableID               string
	TableName             string
	SiteBaseURL           string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/table/{tableID}", editPropsPage)
	mainRouter.HandleFunc("/admin/table/tableProps/offPageContent", tablePropsOffPageContent)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tableID := vars["tableID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	tableInfo, err := displayTable.GetTable(trackerDBHandle, tableID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbInfo, err := databaseController.GetDatabaseInfo(trackerDBHandle, tableInfo.ParentDatabaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	elemPrefix := "tableProps_"
	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := TemplParams{
		ElemPrefix:            elemPrefix,
		Title:                 "Table view properties",
		DatabaseID:            dbInfo.DatabaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
		TableID:               tableID,
		TableName:             tableInfo.Name,
		SiteBaseURL:           runtimeConfig.GetSiteBaseURL(),
		CurrUserIsAdmin:       isAdmin}

	if err := tablePropTemplates.ExecuteTemplate(w, "tablePropsAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

type DialogTemplParams struct {
	ElemPrefix string
}

func tablePropsOffPageContent(w http.ResponseWriter, r *http.Request) {

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	elemPrefix := "tableProps_"

	templParams := DialogTemplParams{
		ElemPrefix: elemPrefix}

	if err := tablePropTemplates.ExecuteTemplate(w, "tablePropsOffPageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
