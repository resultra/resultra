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
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

// Parse all the HTML templates at once. Individual templates can then
// be referenced throughout this package using htmlTemplates.ExecuteTemplate(...)
var designFormTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/form/designForm.html",
		"static/form/viewForm.html",
		"static/form/designFormProperties.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		components.TemplateFileList}
	designFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FormElemTemplateParams struct {
	ElemPrefix string
}

type FormTemplateParams struct {
	NamePanelParams   propertiesSidebar.PanelTemplateParams
	FilterPanelParams propertiesSidebar.PanelTemplateParams
}

type DesignFormTemplateParams struct {
	Title            string
	FormID           string
	FormName         string
	TableID          string
	CheckboxParams   FormElemTemplateParams
	DatePickerParams FormElemTemplateParams
	TextBoxParams    FormElemTemplateParams
	HtmlEditorParams FormElemTemplateParams
	ImageParams      FormElemTemplateParams
	FormParams       FormTemplateParams
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

	formParams := FormTemplateParams{
		NamePanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Form Name", PanelID: "formName"},
		FilterPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "formFilter"}}

	templParams := DesignFormTemplateParams{
		Title:            "Design Form",
		FormID:           formRef.FormID,
		TableID:          formRef.TableID,
		FormName:         formRef.Name,
		CheckboxParams:   FormElemTemplateParams{ElemPrefix: "checkbox_"},
		DatePickerParams: FormElemTemplateParams{ElemPrefix: "datePicker_"},
		TextBoxParams:    FormElemTemplateParams{ElemPrefix: "textBox_"},
		HtmlEditorParams: FormElemTemplateParams{ElemPrefix: "htmlEditor_"},
		ImageParams:      FormElemTemplateParams{ElemPrefix: "image_"},
		FormParams:       formParams}

	err := designFormTemplates.ExecuteTemplate(w, "designForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
