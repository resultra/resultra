package design

import (
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/webui/common/form/components/caption"
	"resultra/datasheet/webui/common/form/components/checkBox"
	"resultra/datasheet/webui/common/form/components/comment"
	"resultra/datasheet/webui/common/form/components/datePicker"
	"resultra/datasheet/webui/common/form/components/emailAddr"
	"resultra/datasheet/webui/common/form/components/formButton"
	"resultra/datasheet/webui/common/form/components/gauge"
	"resultra/datasheet/webui/common/form/components/header"
	"resultra/datasheet/webui/common/form/components/htmlEditor"
	"resultra/datasheet/webui/common/form/components/image"
	"resultra/datasheet/webui/common/form/components/label"
	"resultra/datasheet/webui/common/form/components/numberInput"
	"resultra/datasheet/webui/common/form/components/progress"
	"resultra/datasheet/webui/common/form/components/rating"
	"resultra/datasheet/webui/common/form/components/selection"
	"resultra/datasheet/webui/common/form/components/socialButton"
	"resultra/datasheet/webui/common/form/components/textBox"
	"resultra/datasheet/webui/common/form/components/toggle"
	"resultra/datasheet/webui/common/form/components/urlLink"
	"resultra/datasheet/webui/common/form/components/userSelection"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type FormPropertyTemplateParams struct {
	NamePanelParams       propertiesSidebar.PanelTemplateParams
	FilterPanelParams     propertiesSidebar.PanelTemplateParams
	RolesPanelParams      propertiesSidebar.PanelTemplateParams
	NewItemPanelParams    propertiesSidebar.PanelTemplateParams
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
	ToggleParams        toggle.ToggleDesignTemplateParams
	DatePickerParams    datePicker.DatePickerDesignTemplateParams
	TextBoxParams       textBox.TextboxDesignTemplateParams
	NumberInputParams   numberInput.NumberInputDesignTemplateParams
	SelectionParams     selection.SelectionDesignTemplateParams
	HtmlEditorParams    htmlEditor.HTMLEditorDesignTemplateParams
	ImageParams         image.ImageDesignTemplateParams
	CommentParams       comment.CommentDesignTemplateParams
	RatingParams        rating.RatingDesignTemplateParams
	UserSelectionParams userSelection.UserSelectionDesignTemplateParams
	ProgressParams      progress.ProgressDesignTemplateParams
	GaugeParams         gauge.GaugeDesignTemplateParams
	HeaderParams        header.HeaderTemplateParams
	CaptionParams       caption.CaptionTemplateParams
	ButtonParams        formButton.ButtonTemplateParams
	SocialButtonParams  socialButton.SocialButtonDesignTemplateParams
	LabelParams         label.LabelDesignTemplateParams
	EmailAddrParams     emailAddr.EmailAddrDesignTemplateParams
	UrlLinkParams       urlLink.UrlLinkDesignTemplateParams
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
		NewItemPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "New Items",
			PanelID: "newItems"},
		SortPanelParams:       propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Sorting", PanelID: "formSort"},
		FilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

	templParams := DesignFormTemplateParams{
		Title:        "Design Form",
		DatabaseID:   formInfo.DatabaseID,
		DatabaseName: formInfo.DatabaseName,
		FormID:       formInfo.FormID,
		FormName:     formInfo.FormName,

		CheckboxParams:      checkBox.DesignTemplateParams,
		ToggleParams:        toggle.DesignTemplateParams,
		DatePickerParams:    datePicker.DesignTemplateParams,
		TextBoxParams:       textBox.DesignTemplateParams,
		NumberInputParams:   numberInput.DesignTemplateParams,
		SelectionParams:     selection.DesignTemplateParams,
		UserSelectionParams: userSelection.DesignTemplateParams,
		ProgressParams:      progress.DesignTemplateParams,
		GaugeParams:         gauge.DesignTemplateParams,
		HtmlEditorParams:    htmlEditor.DesignTemplateParams,
		ImageParams:         image.DesignTemplateParams,
		CommentParams:       comment.DesignTemplateParams,
		RatingParams:        rating.DesignTemplateParams,
		SocialButtonParams:  socialButton.DesignTemplateParams,
		HeaderParams:        header.TemplateParams,
		CaptionParams:       caption.TemplateParams,
		ButtonParams:        formButton.TemplateParams,
		LabelParams:         label.DesignTemplateParams,
		EmailAddrParams:     emailAddr.DesignTemplateParams,
		UrlLinkParams:       urlLink.DesignTemplateParams,
		FormPropertyParams:  formPropParams}

	return templParams

}
