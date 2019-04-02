// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package htmlEditor

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type HTMLEditorDesignTemplateParams struct {
	ElemPrefix               string
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type HTMLEditorViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams HTMLEditorDesignTemplateParams
var ViewTemplateParams HTMLEditorViewTemplateParams

func init() {

	elemPrefix := "htmlEditor_"

	DesignTemplateParams = HTMLEditorDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "htmlEditorLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "editorDelete", "Delete Editor"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "htmlEditorHelp"),
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "htmlEditorValidation"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "editorVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "editorPerms"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Editor",
			FieldInfoPrompt: `Content provided in an editor is stored in fields. Either a new field can be created for this
					editor, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this editor's content.`}}

	ViewTemplateParams = HTMLEditorViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "htmlEditorTimeline"}}
}
