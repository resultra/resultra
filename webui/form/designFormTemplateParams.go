package form

import (
	"resultra/datasheet/server/form"
	"resultra/datasheet/webui/form/components/checkBox"
	"resultra/datasheet/webui/form/components/datePicker"
	"resultra/datasheet/webui/form/components/htmlEditor"
	"resultra/datasheet/webui/form/components/image"
	"resultra/datasheet/webui/form/components/textBox"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type FormTemplateParams struct {
	NamePanelParams   propertiesSidebar.PanelTemplateParams
	FilterPanelParams propertiesSidebar.PanelTemplateParams
}

// Aggregate the template parameters from all the form components, then
// combine them with the paramers for the form itself.
type DesignFormTemplateParams struct {
	Title            string
	FormID           string
	FormName         string
	TableID          string
	CheckboxParams   checkBox.CheckboxTemplateParams
	DatePickerParams datePicker.DatePickerTemplateParams
	TextBoxParams    textBox.TextboxTemplateParams
	HtmlEditorParams htmlEditor.HTMLEditorTemplateParams
	ImageParams      image.ImageTemplateParams
	FormParams       FormTemplateParams
}

var designFormTemplateParams DesignFormTemplateParams

func createDesignFormTemplateParams(formToDesign *form.Form) DesignFormTemplateParams {

	formParams := FormTemplateParams{
		NamePanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Form Name", PanelID: "formName"},
		FilterPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "formFilter"}}

	templParams := DesignFormTemplateParams{
		Title:            "Design Form",
		FormID:           formToDesign.FormID,
		TableID:          formToDesign.ParentTableID,
		FormName:         formToDesign.Name,
		CheckboxParams:   checkBox.TemplateParams,
		DatePickerParams: datePicker.TemplateParams,
		TextBoxParams:    textBox.TemplateParams,
		HtmlEditorParams: htmlEditor.TemplateParams,
		ImageParams:      image.TemplateParams,
		FormParams:       formParams}

	return templParams

}
