// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formPage

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/formLink"
	"github.com/resultra/resultra/server/generic/api"
	"github.com/resultra/resultra/server/userRole"

	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
)

var submitFormTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/formPage/submitFormPage.html", "static/formPage/common.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	submitFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func submitFormPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sharedLinkID := vars["sharedLinkID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := submitFormTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("Submit form: : shared link ID = %v", sharedLinkID)

		trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
		if dbErr != nil {
			http.Error(w, dbErr.Error(), http.StatusInternalServerError)
			return
		}

		formLink, getFormLinkErr := formLink.GetFormLinkFromSharedLinkID(trackerDBHandle, sharedLinkID)
		if getFormLinkErr != nil {
			api.WriteErrorResponse(w, getFormLinkErr)
			return
		}

		if !formLink.SharedLinkEnabled {
			api.WriteErrorResponse(w, fmt.Errorf("Shared link disabled for form link"))
			return
		}

		formDBInfo, getErr := databaseController.GetFormDatabaseInfo(trackerDBHandle, formLink.FormID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		isAdmin := userRole.CurrUserIsDatabaseAdmin(r, formDBInfo.DatabaseID)

		templParams := SubmitFormPageTemplateParams{Title: "Submit Form",
			FormID:          formLink.FormID,
			FormName:        formDBInfo.FormName,
			LinkName:        formLink.Name,
			DatabaseID:      formDBInfo.DatabaseID,
			CurrUserIsAdmin: isAdmin,
			FormLinkID:      formLink.LinkID,
			DatabaseName:    formDBInfo.DatabaseName}

		if err := submitFormTemplates.ExecuteTemplate(w, "submitFormPage", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
