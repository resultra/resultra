// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package numberInput

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type NumberInputDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	SpinnerPanelParams       propertiesSidebar.PanelTemplateParams
	CondFormatPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	ClearValuePanelParams    propertiesSidebar.PanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type NumberInputViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams NumberInputDesignTemplateParams
var ViewTemplateParams NumberInputViewTemplateParams

func init() {

	elemPrefix := "numberInput_"

	DesignTemplateParams = NumberInputDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "numberInputLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "numberInputDelete", "Delete Text Box"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Number Format", PanelID: "numberInputFormat"},
		CondFormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Conditional Format", PanelID: "condNumberInputFormat"},
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "numberInputClearValue"},
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "numberInputHelp"),
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "numberInputValidation"},
		SpinnerPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Spinner Buttons", PanelID: "numberInputSpinner"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "numberInputVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "numberInputPerms"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Number Input",
			FieldInfoPrompt:    `Values from number inputs are stored in fields. Either a new field can be created, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this number input's values.`}}

	ViewTemplateParams = NumberInputViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "numberInputTimeline"}}

}
