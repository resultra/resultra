package formPage

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"resultra/tracker/server/databaseController"
	"resultra/tracker/server/formLink"
	"resultra/tracker/server/generic/api"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
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
