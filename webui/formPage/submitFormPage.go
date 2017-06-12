package formPage

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/formLink"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var submitFormTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/formPage/submitFormPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	submitFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type SubmitFormPageTemplateParams struct {
	Title        string
	FormID       string
	FormLinkID   string
	FormName     string
	DatabaseID   string
	DatabaseName string
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

		formLink, getFormLinkErr := formLink.GetFormLinkFromSharedLinkID(sharedLinkID)
		if getFormLinkErr != nil {
			api.WriteErrorResponse(w, getFormLinkErr)
			return
		}

		if !formLink.SharedLinkEnabled {
			api.WriteErrorResponse(w, fmt.Errorf("Shared link disabled for form link"))
			return
		}

		formDBInfo, getErr := databaseController.GetFormDatabaseInfo(formLink.FormID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		templParams := SubmitFormPageTemplateParams{Title: "Submit Form",
			FormID:       formLink.FormID,
			FormName:     formDBInfo.FormName,
			DatabaseID:   formDBInfo.DatabaseID,
			FormLinkID:   formLink.LinkID,
			DatabaseName: formDBInfo.DatabaseName}

		if err := submitFormTemplates.ExecuteTemplate(w, "submitFormPage", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
