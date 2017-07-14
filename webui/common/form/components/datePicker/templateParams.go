package datePicker

import (
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type DatePickerDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ValidationPanelParams    propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	PermissionPanelParams    permissions.PermissionsPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	ClearValuePanelParams    propertiesSidebar.PanelTemplateParams
}

type DatePickerViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams DatePickerDesignTemplateParams
var ViewTemplateParams DatePickerViewTemplateParams

func init() {

	elemPrefix := "datePicker_"

	DesignTemplateParams = DatePickerDesignTemplateParams{
		ElemPrefix:            elemPrefix,
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "datePickerFormat"},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "datePickerValidation"},
		ClearValuePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Clear Values", PanelID: "datePickerClearValue"},
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "datePickerVisibility"),
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "datePickerLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "datePickerDelete", "Delete Date Picker"),
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "datePickerPerms"),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New Date Picker",
			FieldInfoPrompt: `Values from date pickers are stored in fields. Either a new field can be created for this
					date picker, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this date picker's values.`}}

	ViewTemplateParams = DatePickerViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "datePickerTimeline"}}

}
