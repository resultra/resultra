package userRoleProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"

	"resultra/datasheet/server/userRole"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var userRoleTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/userRole/userRoleProps/editRoleProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}

	userRoleTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserRoleTemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
	RoleID       string
	RoleName     string
}

func editRolePropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	roleID := vars["roleID"]

	log.Println("editRolePropsPage: viewing/editing admin settings for role ID = ", roleID)

	roleInfo, err := userRole.GetUserRole(roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("editRolePropsPage: viewing/editing admin settings for role = %+v", roleInfo)

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(roleInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	//	elemPrefix := "userRole_"
	templParams := UserRoleTemplParams{
		Title:        "User Role Settings",
		DatabaseID:   roleInfo.ParentDatabaseID,
		DatabaseName: dbInfo.DatabaseName,
		RoleID:       roleID,
		RoleName:     roleInfo.RoleName}

	if err := userRoleTemplates.ExecuteTemplate(w, "editUserRolePropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
