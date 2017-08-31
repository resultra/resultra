package socialButton

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type SocialButtonDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	TooltipPanelParams       propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type SocialButtonViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams SocialButtonDesignTemplateParams
var ViewTemplateParams SocialButtonViewTemplateParams

func init() {

	elemPrefix := "socialButton_"

	DesignTemplateParams = SocialButtonDesignTemplateParams{
		ElemPrefix:            elemPrefix,
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "socialButtonVisibility"),
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "socialButtonLabel"}},
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "socialButtonHelp"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "socialButtonDelete", "Delete Button"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "socialButtonFormat"},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "socialButtonPerms"),
		TooltipPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Rating Descriptions", PanelID: "socialButtonTooltip"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Social Button",
			FieldInfoPrompt: `Social button values are stored in user fields. Either a new field can be created to store the values, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the values.`}}

	ViewTemplateParams = SocialButtonViewTemplateParams{
		ElemPrefix: elemPrefix}

}
