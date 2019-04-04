// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package templatePage

import (
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
			IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace()}
		err := homePageTemplates.ExecuteTemplate(respWriter, "templatePagePublic", templParams)
		if err != nil {
			http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		}
	} else {
		templParams := PageInfo{Title: "Template Page - Signed In",
			IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace()}
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
