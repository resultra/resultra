package templatePage

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var homePageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/templatePage/templatePageSignedIn.html",
		"static/templatePage/templatePropertiesDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	homePageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type PageInfo struct {
	Title                 string `json:"title"`
	IsSingleUserWorkspace bool
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/templatePage/mainContent", templatePageMainContent)
	mainRouter.HandleFunc("/templatePage/offPageContent", templatePageOffPageContent)
}

func templatePageMainContent(respWriter http.ResponseWriter, req *http.Request) {

	_, authErr := userAuth.GetCurrentUserInfo(req)
	if authErr != nil {
		log.Printf("user not authorized: %v", authErr)
		templParams := PageInfo{Title: "Template Page - Signed out",
			IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace}
		err := homePageTemplates.ExecuteTemplate(respWriter, "templatePagePublic", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		}
	} else {
		templParams := PageInfo{Title: "Template Page - Signed In",
			IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace}
		err := homePageTemplates.ExecuteTemplate(respWriter, "templatePageSignedIn", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		}

	}

}

type OffPageContentTemplParams struct{}

func templatePageOffPageContent(respWriter http.ResponseWriter, req *http.Request) {

	templParams := OffPageContentTemplParams{}
	err := homePageTemplates.ExecuteTemplate(respWriter, "templatePageOffPageContent", templParams)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
