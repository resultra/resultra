package numberInput

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/form/components/common/delete"
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/newFormElemDialog"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/common/form/components/common/visibility"
	"resultra/tracker/webui/generic/propertiesSidebar"
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
