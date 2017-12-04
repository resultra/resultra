package homePage

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var homePageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/homePage/homePagePublic.html",
		"static/homePage/homePageSignedIn.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	homePageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/", home)
}

type PageInfo struct {
	Title         string `json:"title"`
	WorkspaceName string `json:"workspaceName"`
}

func home(respWriter http.ResponseWriter, req *http.Request) {

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

	log.Printf("workspace name: %v", workspaceName)

	_, authErr := userAuth.GetCurrentUserInfo(req)
	if authErr != nil {
		log.Printf("user not authorized: %v", authErr)
		templParams := PageInfo{
			Title:         "Resultra Workspace - Signed out",
			WorkspaceName: workspaceName}
		err := homePageTemplates.ExecuteTemplate(respWriter, "homePagePublic", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {

		templParams := PageInfo{
			Title:         "Resultra Workspace - Signed In",
			WorkspaceName: workspaceName}
		err := homePageTemplates.ExecuteTemplate(respWriter, "homePageSignedIn", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}
