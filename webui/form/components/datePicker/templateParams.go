package datePicker

import (
	"resultra/datasheet/webui/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type DatePickerDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

type DatePickerViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams DatePickerDesignTemplateParams
var ViewTemplateParams DatePickerViewTemplateParams

func init() {
	DesignTemplateParams = DatePickerDesignTemplateParams{
		ElemPrefix:        "datePicker_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "datePickerFormat"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  "datePicker_",
			DialogTitle: "New Date Picker",
			FieldInfoPrompt: `Values from date pickers are stored in fields. Either a new field can be created for this
					date picker, or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this date picker's values.`}}

	ViewTemplateParams = DatePickerViewTemplateParams{
		ElemPrefix:          "datePicker_",
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "datePickerTimeline"}}

}
