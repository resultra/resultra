package collaboratorProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var userRoleTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/collaborators/collaboratorProps/collaboratorProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	userRoleTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserRoleTemplParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	UserID          string
	CollaboratorID  string
	UserName        string
	CurrUserIsAdmin bool
}

func editCollabPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	collaboratorID := vars["collaboratorID"]
	databaseID := vars["databaseID"]

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	collabInfo, collabErr := userRole.GetCollaboratorByID(trackerDBHandle, collaboratorID)
	if collabErr != nil {
		http.Error(w, collabErr.Error(), http.StatusInternalServerError)
		return

	}

	userInfo, err := userAuth.GetUserInfoByID(trackerDBHandle, collabInfo.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
		return
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	//	elemPrefix := "userRole_"
	templParams := UserRoleTemplParams{
		Title:           "Collaborator Settings",
		DatabaseID:      dbInfo.DatabaseID,
		DatabaseName:    dbInfo.DatabaseName,
		UserID:          collabInfo.UserID,
		CollaboratorID:  collaboratorID,
		UserName:        userInfo.UserName,
		CurrUserIsAdmin: isAdmin}

	if err := userRoleTemplates.ExecuteTemplate(w, "collabPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
