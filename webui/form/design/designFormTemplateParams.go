package design

import (
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/webui/common/form/components/checkBox"
	"resultra/datasheet/webui/common/form/components/comment"
	"resultra/datasheet/webui/common/form/components/datePicker"
	"resultra/datasheet/webui/common/form/components/header"
	"resultra/datasheet/webui/common/form/components/htmlEditor"
	"resultra/datasheet/webui/common/form/components/image"
	"resultra/datasheet/webui/common/form/components/rating"
	"resultra/datasheet/webui/common/form/components/selection"
	"resultra/datasheet/webui/common/form/components/textBox"
	"resultra/datasheet/webui/common/form/components/userSelection"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type FormPropertyTemplateParams struct {
	NamePanelParams       propertiesSidebar.PanelTemplateParams
	FilterPanelParams     propertiesSidebar.PanelTemplateParams
	RolesPanelParams      propertiesSidebar.PanelTemplateParams
	SortPanelParams       propertiesSidebar.PanelTemplateParams
	FilterPropPanelParams recordFilter.FilterPanelTemplateParams
}

// Aggregate the template parameters from all the form components, then
// combine them with the paramers for the form itself.
type DesignFormTemplateParams struct {
	Title               string
	DatabaseID          string
	DatabaseName        string
	FormID              string
	FormName            string
	CheckboxParams      checkBox.CheckboxDesignTemplateParams
	DatePickerParams    datePicker.DatePickerDesignTemplateParams
	TextBoxParams       textBox.TextboxDesignTemplateParams
	SelectionParams     selection.SelectionDesignTemplateParams
	HtmlEditorParams    htmlEditor.HTMLEditorDesignTemplateParams
	ImageParams         image.ImageDesignTemplateParams
	CommentParams       comment.CommentDesignTemplateParams
	RatingParams        rating.RatingDesignTemplateParams
	UserSelectionParams userSelection.UserSelectionDesignTemplateParams
	HeaderParams        header.HeaderTemplateParams
	FormPropertyParams  FormPropertyTemplateParams
}

var designFormTemplateParams DesignFormTemplateParams

func createDesignFormTemplateParams(formInfo *databaseController.FormDatabaseInfo) DesignFormTemplateParams {

	elemPrefix := "form_"

	formPropParams := FormPropertyTemplateParams{
		NamePanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Form Name", PanelID: "formName"},
		FilterPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "formFilter"},
		RolesPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles and Privileges",
			PanelID: "formRoles"},
		SortPanelParams:       propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Sorting", PanelID: "formSort"},
		FilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

	templParams := DesignFormTemplateParams{
		Title:        "Design Form",
		DatabaseID:   formInfo.DatabaseID,
		DatabaseName: formInfo.DatabaseName,
		FormID:       formInfo.FormID,
		FormName:     formInfo.FormName,

		CheckboxParams:      checkBox.DesignTemplateParams,
		DatePickerParams:    datePicker.DesignTemplateParams,
		TextBoxParams:       textBox.DesignTemplateParams,
		SelectionParams:     selection.DesignTemplateParams,
		UserSelectionParams: userSelection.DesignTemplateParams,
		HtmlEditorParams:    htmlEditor.DesignTemplateParams,
		ImageParams:         image.DesignTemplateParams,
		CommentParams:       comment.DesignTemplateParams,
		RatingParams:        rating.DesignTemplateParams,
		HeaderParams:        header.TemplateParams,
		FormPropertyParams:  formPropParams}

	return templParams

}
