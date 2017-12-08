package userAdmin

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	//	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var homePageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/userAdmin/passwordResetPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	homePageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type PageInfo struct {
	Title         string `json:"title"`
	WorkspaceName string `json:"workspaceName"`
	ResetID       string `json:"resetID"`
}

func resetPassword(respWriter http.ResponseWriter, req *http.Request) {

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

	resetID := vars["resetID"]

	log.Printf("password reset: reset id = %v", resetID)

	templParams := PageInfo{
		Title:         "Resultra Workspace - Reset Password",
		WorkspaceName: workspaceName,
		ResetID:       resetID}
	err := homePageTemplates.ExecuteTemplate(respWriter, "resetPasswordPage", templParams)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
