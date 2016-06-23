package homePage

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var homePageTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/homePage/homePagePublic.html",
		"static/homePage/homePageSignedIn.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}
	homePageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/", home)
}

type PageInfo struct {
	Title string `json:"title"`
}

func home(respWriter http.ResponseWriter, req *http.Request) {

	_, authErr := userAuth.GetCurrentUserInfo(req)
	if authErr != nil {
		templParams := PageInfo{"Home Page - Signed out"}
		err := homePageTemplates.ExecuteTemplate(respWriter, "homePagePublic", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		}
	} else {
		templParams := PageInfo{"Home Page - Signed In"}
		err := homePageTemplates.ExecuteTemplate(respWriter, "homePageSignedIn", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		}

	}

}
