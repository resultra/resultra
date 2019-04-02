// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package emailAddr

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type EmailAddrDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	ClearValuePanelParams    propertiesSidebar.PanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type EmailAddrViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams EmailAddrDesignTemplateParams
var ViewTemplateParams EmailAddrViewTemplateParams

func init() {

	elemPrefix := "emailAddr_"

	DesignTemplateParams = EmailAddrDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "emailAddrLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "emailAddrDelete", "Delete Email Address Input"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "emailAddrHelp"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "emailAddrFormat"},
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "emailAddrClearValue"},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "emailAddrValidation"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "emailAddrVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "emailAddrPerms"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Email Address Input",
			FieldInfoPrompt:    `Email addresses are stored in fields. Either a new field can be created, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the email addresses.'`}}

	ViewTemplateParams = EmailAddrViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "emailAddrTimeline"}}

}
