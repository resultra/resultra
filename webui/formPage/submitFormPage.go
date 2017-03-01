package formPage

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var submitFormTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/formPage/submitFormPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}
	submitFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type SubmitFormPageTemplateParams struct {
	Title        string
	FormID       string
	FormName     string
	DatabaseID   string
	DatabaseName string
}

func submitFormPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	formID := vars["formID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := submitFormTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("Submit form: : formID = %v", formID)

		formDBInfo, getErr := databaseController.GetFormDatabaseInfo(formID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		templParams := SubmitFormPageTemplateParams{Title: "Submit Form",
			FormID:       formID,
			FormName:     formDBInfo.FormName,
			DatabaseID:   formDBInfo.DatabaseID,
			DatabaseName: formDBInfo.DatabaseName}

		if err := submitFormTemplates.ExecuteTemplate(w, "submitFormPage", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
