// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package design

import (
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/userRole"
	"github.com/resultra/resultra/webui/common/form/components/attachment"
	"github.com/resultra/resultra/webui/common/form/components/caption"
	"github.com/resultra/resultra/webui/common/form/components/checkBox"
	"github.com/resultra/resultra/webui/common/form/components/comment"
	"github.com/resultra/resultra/webui/common/form/components/datePicker"
	"github.com/resultra/resultra/webui/common/form/components/emailAddr"
	"github.com/resultra/resultra/webui/common/form/components/file"
	"github.com/resultra/resultra/webui/common/form/components/formButton"
	"github.com/resultra/resultra/webui/common/form/components/gauge"
	"github.com/resultra/resultra/webui/common/form/components/header"
	"github.com/resultra/resultra/webui/common/form/components/htmlEditor"
	"github.com/resultra/resultra/webui/common/form/components/image"
	"github.com/resultra/resultra/webui/common/form/components/label"
	"github.com/resultra/resultra/webui/common/form/components/numberInput"
	"github.com/resultra/resultra/webui/common/form/components/progress"
	"github.com/resultra/resultra/webui/common/form/components/rating"
	"github.com/resultra/resultra/webui/common/form/components/selection"
	"github.com/resultra/resultra/webui/common/form/components/socialButton"
	"github.com/resultra/resultra/webui/common/form/components/textBox"
	"github.com/resultra/resultra/webui/common/form/components/toggle"
	"github.com/resultra/resultra/webui/common/form/components/urlLink"
	"github.com/resultra/resultra/webui/common/form/components/userSelection"
	"github.com/resultra/resultra/webui/common/form/components/userTag"
	"github.com/resultra/resultra/webui/common/recordFilter"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
	"net/http"
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
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	FormID                string
	FormName              string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
	CheckboxParams        checkBox.CheckboxDesignTemplateParams
	ToggleParams          toggle.ToggleDesignTemplateParams
	DatePickerParams      datePicker.DatePickerDesignTemplateParams
	TextBoxParams         textBox.TextboxDesignTemplateParams
	NumberInputParams     numberInput.NumberInputDesignTemplateParams
	SelectionParams       selection.SelectionDesignTemplateParams
	HtmlEditorParams      htmlEditor.HTMLEditorDesignTemplateParams
	AttachmentParams      attachment.ImageDesignTemplateParams
	ImageParams           image.ImageDesignTemplateParams
	CommentParams         comment.CommentDesignTemplateParams
	RatingParams          rating.RatingDesignTemplateParams
	UserSelectionParams   userSelection.UserSelectionDesignTemplateParams
	UserTagParams         userTag.UserTagDesignTemplateParams
	ProgressParams        progress.ProgressDesignTemplateParams
	GaugeParams           gauge.GaugeDesignTemplateParams
	HeaderParams          header.HeaderTemplateParams
	CaptionParams         caption.CaptionTemplateParams
	ButtonParams          formButton.ButtonTemplateParams
	SocialButtonParams    socialButton.SocialButtonDesignTemplateParams
	LabelParams           label.LabelDesignTemplateParams
	EmailAddrParams       emailAddr.EmailAddrDesignTemplateParams
	UrlLinkParams         urlLink.UrlLinkDesignTemplateParams
	FileParams            file.FileDesignTemplateParams
	FormPropertyParams    FormPropertyTemplateParams
}

var designFormTemplateParams DesignFormTemplateParams

func createDesignFormTemplateParams(r *http.Request,
	formInfo *databaseController.FormDatabaseInfo, workspaceName string) DesignFormTemplateParams {

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

	currUserIsAdmin := userRole.CurrUserIsDatabaseAdmin(r, formInfo.DatabaseID)

	templParams := DesignFormTemplateParams{
		Title:                 "Design Form",
		DatabaseID:            formInfo.DatabaseID,
		DatabaseName:          formInfo.DatabaseName,
		FormID:                formInfo.FormID,
		FormName:              formInfo.FormName,
		WorkspaceName:         workspaceName,
		CurrUserIsAdmin:       currUserIsAdmin,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),

		CheckboxParams:      checkBox.DesignTemplateParams,
		ToggleParams:        toggle.DesignTemplateParams,
		DatePickerParams:    datePicker.DesignTemplateParams,
		TextBoxParams:       textBox.DesignTemplateParams,
		NumberInputParams:   numberInput.DesignTemplateParams,
		SelectionParams:     selection.DesignTemplateParams,
		UserSelectionParams: userSelection.DesignTemplateParams,
		UserTagParams:       userTag.DesignTemplateParams,
		ProgressParams:      progress.DesignTemplateParams,
		GaugeParams:         gauge.DesignTemplateParams,
		HtmlEditorParams:    htmlEditor.DesignTemplateParams,
		AttachmentParams:    attachment.DesignTemplateParams,
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
		FileParams:          file.DesignTemplateParams,
		FormPropertyParams:  formPropParams}

	return templParams

}
