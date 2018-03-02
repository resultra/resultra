package mainAdminPage

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/common/runtimeConfig"
	"resultra/datasheet/server/databaseController"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/server/workspace"
	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"

	"resultra/datasheet/webui/admin/alerts/alertList"
	"resultra/datasheet/webui/admin/collaborators/collaboratorList"
	"resultra/datasheet/webui/admin/dashboards"
	"resultra/datasheet/webui/admin/fields/fieldList"
	"resultra/datasheet/webui/admin/formLink/formLinkList"
	"resultra/datasheet/webui/admin/forms/formList"
	"resultra/datasheet/webui/admin/general"
	"resultra/datasheet/webui/admin/itemList/itemListList"
	"resultra/datasheet/webui/admin/tables/tableList"
	"resultra/datasheet/webui/admin/userRole/userRoleList"
	"resultra/datasheet/webui/admin/valueLists/valueListList"
)

var mainAdminPageTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/mainAdminPage/mainAdminPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList,
		general.TemplateFileList,
		formList.TemplateFileList,
		formLinkList.TemplateFileList,
		tableList.TemplateFileList,
		itemListList.TemplateFileList,
		fieldList.TemplateFileList,
		valueListList.TemplateFileList,
		dashboards.TemplateFileList,
		alertList.TemplateFileList,
		userRoleList.TemplateFileList,
		collaboratorList.TemplateFileList}

	mainAdminPageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func mainAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

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

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	currUserIsAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := TemplParams{
		Title:                 "Settings",
		DatabaseID:            databaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		CurrUserIsAdmin:       currUserIsAdmin,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace}

	if err := mainAdminPageTemplates.ExecuteTemplate(w, "mainAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
