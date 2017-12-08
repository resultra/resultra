package userRoleProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var userRoleTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/userRole/userRoleProps/editRoleProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		adminCommon.TemplateFileList,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	userRoleTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserRoleTemplParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	WorkspaceName   string
	RoleID          string
	RoleName        string
	CurrUserIsAdmin bool
}

func editRolePropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	roleID := vars["roleID"]

	log.Println("editRolePropsPage: viewing/editing admin settings for role ID = ", roleID)

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

	roleInfo, err := userRole.GetUserRole(trackerDBHandle, roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("editRolePropsPage: viewing/editing admin settings for role = %+v", roleInfo)

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, roleInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
		return
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	//	elemPrefix := "userRole_"
	templParams := UserRoleTemplParams{
		Title:           "User Role Settings",
		DatabaseID:      roleInfo.ParentDatabaseID,
		DatabaseName:    dbInfo.DatabaseName,
		WorkspaceName:   workspaceName,
		RoleID:          roleID,
		RoleName:        roleInfo.RoleName,
		CurrUserIsAdmin: isAdmin}

	if err := userRoleTemplates.ExecuteTemplate(w, "editUserRolePropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
