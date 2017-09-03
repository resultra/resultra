package label

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type LabelDesignTemplateParams struct {
	ElemPrefix               string
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type LabelViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams LabelDesignTemplateParams
var ViewTemplateParams LabelViewTemplateParams

func init() {

	elemPrefix := "label_"

	DesignTemplateParams = LabelDesignTemplateParams{
		ElemPrefix:            elemPrefix,
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "labelVisibility"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "labelPerms"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "htmlSelectionHelp"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "labelDelete", "Delete label"),
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "labelLabel"}},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "labelValidation"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Label",
			FieldInfoPrompt: `Labels are stored in fields. Either a new field can be created to store the labels, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the labels.`}}

	ViewTemplateParams = LabelViewTemplateParams{
		ElemPrefix: elemPrefix}

}
