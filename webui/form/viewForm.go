package form

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/form/common/sort"
	"resultra/datasheet/webui/form/components"
	"resultra/datasheet/webui/generic"
)

var viewFormTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/form/viewForm.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		components.TemplateFileList,
		sort.TemplateFileList}
	viewFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ViewFormTemplateParams struct {
	Title    string
	FormID   string
	TableID  string
	FormName string
}

func viewForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	formID := vars["formID"]
	tableID := vars["tableID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := viewFormTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("view form: : form ID = %v", formID)
		formToView, getErr := form.GetForm(form.GetFormParams{tableID, formID})
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		templParams := ViewFormTemplateParams{Title: "View Form",
			FormID:   formID,
			TableID:  formToView.ParentTableID,
			FormName: formToView.Name}

		err := viewFormTemplates.ExecuteTemplate(w, "viewForm", templParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
