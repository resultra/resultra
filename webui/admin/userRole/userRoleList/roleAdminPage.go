package userRoleList

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

var roleTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/userRole/userRoleList/dashboardPrivs.html",
		"static/admin/userRole/userRoleList/formPrivs.html",
		"static/admin/userRole/userRoleList/newRoleDashboardPrivsDialogPanel.html",
		"static/admin/userRole/userRoleList/newRoleFormPrivsDialogPanel.html",
		"static/admin/userRole/userRoleList/newRoleRoleNameDialogPanel.html",
		"static/admin/userRole/userRoleList/newUserRoleDialog.html",
		"static/admin/userRole/userRoleList/roleAdminPage.html",
		"static/admin/userRole/userRoleList/userRole.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	roleTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FieldTemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
}

func roleAdminPage(w http.ResponseWriter, r *http.Request) {

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

	if err := roleTemplates.ExecuteTemplate(w, "roleAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
