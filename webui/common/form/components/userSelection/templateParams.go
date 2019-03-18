package userSelection

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/form/components/common/delete"
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/newFormElemDialog"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/common/form/components/common/visibility"
	"resultra/tracker/webui/generic/propertiesSidebar"
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
