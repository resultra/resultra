// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package selection

import (
	"resultra/tracker/webui/common/form/components/common/delete"
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/newFormElemDialog"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/common/form/components/common/visibility"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type SelectionDesignTemplateParams struct {
	ElemPrefix               string
	ValuesPanelParams        propertiesSidebar.PanelTemplateParams
	ValueListPanelParams     propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	ClearValuePanelParams    propertiesSidebar.PanelTemplateParams
}

type SelectionViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams SelectionDesignTemplateParams
var ViewTemplateParams SelectionViewTemplateParams

func init() {

	elemPrefix := "selection_"

	DesignTemplateParams = SelectionDesignTemplateParams{
		ElemPrefix: elemPrefix,

		ValuesPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Selectable Values", PanelID: "selectionValues"},
		ValueListPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Selectable Values", PanelID: "valueList"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "selectionVisibility"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "selectionDelete", "Delete Selection"),
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "selectionClearValue"},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "selectionPerms"),
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "selectionLabel"}},

		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Selection",
			FieldInfoPrompt: `Values from selections are stored in fields. Either a new field can be created for this
					selection, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this selection's values.`}}

	ViewTemplateParams = SelectionViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "selectionTimeline"}}

}
