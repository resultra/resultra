package collaboratorProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/userAuth"
	adminCommon "resultra/datasheet/webui/admin/common"

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
	Title        string
	DatabaseID   string
	DatabaseName string
	UserID       string
	UserName     string
}

func editCollabPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID := vars["userID"]
	databaseID := vars["databaseID"]

	userInfo, err := userAuth.GetUserInfoByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
		return
	}

	//	elemPrefix := "userRole_"
	templParams := UserRoleTemplParams{
		Title:        "Collaborator Settings",
		DatabaseID:   dbInfo.DatabaseID,
		DatabaseName: dbInfo.DatabaseName,
		UserID:       userID,
		UserName:     userInfo.UserName}

	if err := userRoleTemplates.ExecuteTemplate(w, "collabPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
