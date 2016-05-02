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

// Parse all the HTML templates at once. Individual templates can then
// be referenced throughout this package using htmlTemplates.ExecuteTemplate(...)
var designFormTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/form/designForm.html",
		"static/form/viewForm.html",
		"static/form/properties.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		common.TemplateFileList,
		components.TemplateFileList}
	designFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FormElemTemplateParams struct {
	ElemPrefix string
}

type DesignFormTemplateParams struct {
	Title          string
	FormID         string
	FormName       string
	CheckboxParams FormElemTemplateParams
	TextBoxParams  FormElemTemplateParams
}

func designForm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	formID := vars["formID"]
	log.Println("Design Form: editing for form with ID = ", formID)

	appEngContext := appengine.NewContext(r)
	formRef, getErr := form.GetFormRef(appEngContext, form.GetFormParams{formID})
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := DesignFormTemplateParams{
		Title:          "Edit Layout",
		FormID:         formRef.FormID,
		FormName:       formRef.Name,
		CheckboxParams: FormElemTemplateParams{ElemPrefix: "checkbox_"},
		TextBoxParams:  FormElemTemplateParams{ElemPrefix: "textBox_"}}

	err := designFormTemplates.ExecuteTemplate(w, "designForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
