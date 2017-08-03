package formPage

import (
	"fmt"

	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/formLink"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var newItemFormTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/formPage/newItemPage.html",
		"static/formPage/common.html"}

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

		hasPrivs, privsErr := userRole.CurrentUserHasNewItemLinkPrivs(r, formDBInfo.DatabaseID, formLinkID)
		if privsErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}
		if !hasPrivs {
			api.WriteErrorResponse(w, fmt.Errorf("ERROR: No permissions to add new items with this page"))
			return
		}

		isAdmin := userRole.CurrUserIsDatabaseAdmin(r, formDBInfo.DatabaseID)

		templParams := SubmitFormPageTemplateParams{Title: "New Item - " + formDBInfo.FormName,
			FormID:          formLink.FormID,
			FormName:        formDBInfo.FormName,
			DatabaseID:      formDBInfo.DatabaseID,
			CurrUserIsAdmin: isAdmin,
			FormLinkID:      formLink.LinkID,
			LinkName:        formLink.Name,
			DatabaseName:    formDBInfo.DatabaseName}

		if err := newItemFormTemplates.ExecuteTemplate(w, "newItemFormPage", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
