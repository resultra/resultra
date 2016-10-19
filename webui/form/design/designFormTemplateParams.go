package design

import (
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/webui/form/components/checkBox"
	"resultra/datasheet/webui/form/components/datePicker"
	"resultra/datasheet/webui/form/components/header"
	"resultra/datasheet/webui/form/components/htmlEditor"
	"resultra/datasheet/webui/form/components/image"
	"resultra/datasheet/webui/form/components/rating"
	"resultra/datasheet/webui/form/components/textBox"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type FormTemplateParams struct {
	NamePanelParams   propertiesSidebar.PanelTemplateParams
	FilterPanelParams propertiesSidebar.PanelTemplateParams
	RolesPanelParams  propertiesSidebar.PanelTemplateParams
	SortPanelParams   propertiesSidebar.PanelTemplateParams
}

// Aggregate the template parameters from all the form components, then
// combine them with the paramers for the form itself.
type DesignFormTemplateParams struct {
	Title            string
	DatabaseID       string
	FormID           string
	FormName         string
	TableID          string
	CheckboxParams   checkBox.CheckboxDesignTemplateParams
	DatePickerParams datePicker.DatePickerDesignTemplateParams
	TextBoxParams    textBox.TextboxDesignTemplateParams
	HtmlEditorParams htmlEditor.HTMLEditorDesignTemplateParams
	ImageParams      image.ImageDesignTemplateParams
	RatingParams     rating.RatingDesignTemplateParams
	HeaderParams     header.HeaderTemplateParams
	FormParams       FormTemplateParams
}

var designFormTemplateParams DesignFormTemplateParams

func createDesignFormTemplateParams(formInfo *databaseController.FormDatabaseInfo) DesignFormTemplateParams {

	formParams := FormTemplateParams{
		NamePanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Form Name", PanelID: "formName"},
		FilterPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "formFilter"},
		RolesPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles and Privileges",
			PanelID: "formRoles"},
		SortPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Sorting", PanelID: "formSort"},
	}

	templParams := DesignFormTemplateParams{
		Title:            "Design Form",
		DatabaseID:       formInfo.DatabaseID,
		FormID:           formInfo.FormID,
		TableID:          formInfo.TableID,
		FormName:         formInfo.FormName,
		CheckboxParams:   checkBox.DesignTemplateParams,
		DatePickerParams: datePicker.DesignTemplateParams,
		TextBoxParams:    textBox.DesignTemplateParams,
		HtmlEditorParams: htmlEditor.DesignTemplateParams,
		ImageParams:      image.DesignTemplateParams,
		RatingParams:     rating.DesignTemplateParams,
		HeaderParams:     header.TemplateParams,
		FormParams:       formParams}

	return templParams

}
