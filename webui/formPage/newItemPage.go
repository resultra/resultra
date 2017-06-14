package formPage

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/formLink"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var newItemFormTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/formPage/newItemPage.html", "static/formPage/common.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	newItemFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func newItemFormPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	formLinkID := vars["formLinkID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := newItemFormTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {

		formLink, getFormLinkErr := formLink.GetFormLink(formLinkID)
		if getFormLinkErr != nil {
			api.WriteErrorResponse(w, getFormLinkErr)
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

		if err := submitFormTemplates.ExecuteTemplate(w, "newItemFormPage", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
