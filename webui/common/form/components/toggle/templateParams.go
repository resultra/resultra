package toggle

import (
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ToggleDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
}

type ToggleViewTemplateParams struct {
	ElemPrefix         string
	CommentPanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams ToggleDesignTemplateParams
var ViewTemplateParams ToggleViewTemplateParams

func init() {

	elemPrefix := "toggle_"

	DesignTemplateParams = ToggleDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix, HideNoLabelOption: true,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "toggleLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "toggleDelete", "Delete Toggle"),
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "toggleVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "togglePerms"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "toggleFormat"},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Validation", PanelID: "toggleValidation"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Toggle",
			FieldInfoPrompt: `Toggle values are stored in fields. Either a new field can be created to store the values, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this toggle's values.`}}

	ViewTemplateParams = ToggleViewTemplateParams{
		ElemPrefix:         elemPrefix,
		CommentPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Comments", PanelID: "toggleComments"}}

}
