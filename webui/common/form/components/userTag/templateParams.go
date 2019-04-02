// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userTag

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type UserTagDesignTemplateParams struct {
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

type UserTagViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams UserTagDesignTemplateParams
var ViewTemplateParams UserTagViewTemplateParams

func init() {

	elemPrefix := "userTag_"

	DesignTemplateParams = UserTagDesignTemplateParams{
		ElemPrefix:            elemPrefix,
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "userTagVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "userTagPerms"),
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "userClearValue"},
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "htmlSelectionHelp"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "userTagDelete", "Delete User Selection"),
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "userTagLabel"}},
		ValidationPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "userTagValidation"},
		SelectableUserPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Selectable Users", PanelID: "userTagSelectableUsers"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New User Selection",
			FieldInfoPrompt: `User selections are stored in fields. Either a new field can be created to store the values, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the user selections.`}}

	ViewTemplateParams = UserTagViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "userTagTimeline"}}

}
