package design

import (
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/webui/form/components/checkBox"
	"resultra/datasheet/webui/form/components/datePicker"
	"resultra/datasheet/webui/form/components/header"
	"resultra/datasheet/webui/form/components/htmlEditor"
	"resultra/datasheet/webui/form/components/image"
	"resultra/datasheet/webui/form/components/textBox"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type FormTemplateParams struct {
	NamePanelParams   propertiesSidebar.PanelTemplateParams
	FilterPanelParams propertiesSidebar.PanelTemplateParams
	RolesPanelParams  propertiesSidebar.PanelTemplateParams
}

// Aggregate the template parameters from all the form components, then
// combine them with the paramers for the form itself.
type DesignFormTemplateParams struct {
	Title            string
	DatabaseID       string
	FormID           string
	FormName         string
	TableID          string
	CheckboxParams   checkBox.CheckboxTemplateParams
	DatePickerParams datePicker.DatePickerTemplateParams
	TextBoxParams    textBox.TextboxTemplateParams
	HtmlEditorParams htmlEditor.HTMLEditorTemplateParams
	ImageParams      image.ImageTemplateParams
	HeaderParams     header.HeaderTemplateParams
	FormParams       FormTemplateParams
}

var designFormTemplateParams DesignFormTemplateParams

func createDesignFormTemplateParams(formInfo *databaseController.FormDatabaseInfo) DesignFormTemplateParams {

	formParams := FormTemplateParams{
		NamePanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Form Name", PanelID: "formName"},
		FilterPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "formFilter"},
		RolesPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles and Privileges",
			PanelID: "formRoles"},
	}

	templParams := DesignFormTemplateParams{
		Title:            "Design Form",
		DatabaseID:       formInfo.DatabaseID,
		FormID:           formInfo.FormID,
		TableID:          formInfo.TableID,
		FormName:         formInfo.FormName,
		CheckboxParams:   checkBox.TemplateParams,
		DatePickerParams: datePicker.TemplateParams,
		TextBoxParams:    textBox.TemplateParams,
		HtmlEditorParams: htmlEditor.TemplateParams,
		ImageParams:      image.TemplateParams,
		HeaderParams:     header.TemplateParams,
		FormParams:       formParams}

	return templParams

}
