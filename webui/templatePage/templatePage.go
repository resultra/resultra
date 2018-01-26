package templatePage

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/runtimeConfig"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
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

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/templates", home)
}

type PageInfo struct {
	Title                 string `json:"title"`
	IsSingleUserWorkspace bool
}

func home(respWriter http.ResponseWriter, req *http.Request) {

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
