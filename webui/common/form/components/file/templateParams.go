// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package file

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type FileDesignTemplateParams struct {
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

type FileViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams FileDesignTemplateParams
var ViewTemplateParams FileViewTemplateParams

func init() {

	elemPrefix := "file_"

	DesignTemplateParams = FileDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "fileLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "fileDelete", "Delete File Attachment"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "fileHelp"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "fileFormat"},
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "fileClearValue"},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "fileValidation"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "fileVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "filePerms"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New File Attachment",
			FieldInfoPrompt:    `File attachments are stored in fields. Either a new field can be created, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the file.'`}}

	ViewTemplateParams = FileViewTemplateParams{
		ElemPrefix: elemPrefix}

}
