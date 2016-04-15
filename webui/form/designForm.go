package form

import (
	"appengine"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/api"
	"resultra/datasheet/server/form"
)

// Parse all the HTML templates at once. Individual templates can then
// be referenced throughout this package using htmlTemplates.ExecuteTemplate(...)
var designFormTemplates = template.Must(template.ParseFiles(
	"static/common/common.html",
	"static/field/calcField.html",
	"static/form/designForm.html",
	"static/form/checkBox/newCheckBoxDialog.html",
	"static/form/textBox/newTextBoxDialog.html",
	"static/form/viewForm.html",
	"static/form/common/newFormElemDialog.html",
	"static/form/checkBox/checkboxProp.html"))

type FormElemTemplateParams struct {
	ElemPrefix string
}

type DesignFormTemplateParams struct {
	Title          string
	LayoutID       string
	LayoutName     string
	CheckboxParams FormElemTemplateParams
}

func designForm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	layoutID := vars["layoutID"]
	log.Println("Design Form: editing for layout with ID = ", layoutID)

	appEngContext := appengine.NewContext(r)
	layoutRef, getErr := form.GetLayoutRef(appEngContext, form.GetLayoutParams{layoutID})
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := DesignFormTemplateParams{
		Title:          "Edit Layout",
		LayoutID:       layoutID,
		LayoutName:     layoutRef.Layout.Name,
		CheckboxParams: FormElemTemplateParams{ElemPrefix: "checkbox_"}}

	err := designFormTemplates.ExecuteTemplate(w, "designForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
