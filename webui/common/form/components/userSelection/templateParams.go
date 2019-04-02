// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userSelection

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type UserSelectionDesignTemplateParams struct {
	ElemPrefix                string
	ValidationPanelParams     propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams  newFormElemDialog.TemplateParams
	LabelPanelParams          label.LabelPropertyTemplateParams
	VisibilityPanelParams     visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams     permissions.PermissionsPropertyTemplateParams
	SelectableUserPanelParams propertiesSidebar.PanelTemplateParams
	DeletePanelParams         delete.DeletePropertyPanelTemplateParams
	ClearValuePanelParams     propertiesSidebar.PanelTemplateParams
	HelpPopupParams           inputProperties.HelpPopupPropertyTemplateParams
}

type UserSelectionViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams UserSelectionDesignTemplateParams
var ViewTemplateParams UserSelectionViewTemplateParams

func init() {

	elemPrefix := "userSelection_"

	DesignTemplateParams = UserSelectionDesignTemplateParams{
		ElemPrefix:            elemPrefix,
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "userSelectionVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "userSelectionPerms"),
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "userClearValue"},
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "htmlSelectionHelp"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "userSelectionDelete", "Delete User Selection"),
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "userSelectionLabel"}},
		ValidationPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "userSelectionValidation"},
		SelectableUserPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Selectable Users", PanelID: "userSelectionSelectableUsers"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New User Selection",
			FieldInfoPrompt: `User selections are stored in fields. Either a new field can be created to store the values, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the user selections.`}}

	ViewTemplateParams = UserSelectionViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "userSelectionTimeline"}}

}
