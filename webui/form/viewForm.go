package form

import (
	"appengine"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/webui/common"
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
		components.TemplateFileList}
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

	log.Println("view form: : form ID = %v", formID)
	appEngContext := appengine.NewContext(r)
	formToView, getErr := form.GetForm(appEngContext, form.GetFormParams{tableID, formID})
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
