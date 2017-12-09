package userAdmin

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var registrationPageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/userAdmin/userRegistrationPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	registrationPageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type RegistrationPageInfo struct {
	Title          string `json:"title"`
	WorkspaceName  string `json:"workspaceName"`
	RegistrationID string `json:"registrationID"`
}

func registerNewUser(respWriter http.ResponseWriter, req *http.Request) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		http.Error(respWriter, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(respWriter, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(req)
	registrationID := vars["registration"]

	templParams := RegistrationPageInfo{
		Title:          "Resultra Workspace - Register New Account",
		WorkspaceName:  workspaceName,
		RegistrationID: registrationID}
	err := registrationPageTemplates.ExecuteTemplate(respWriter, "userRegistrationPage", templParams)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
