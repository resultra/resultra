package recordFilter

import (
	"resultra/datasheet/webui/common/field"
)

type FilterPanelTemplateParams struct {
	ElemPrefix           string
	FieldSelectionParams field.FieldSelectionDropdownTemplateParams
}

func NewFilterPanelTemplateParams(elemPrefix string) FilterPanelTemplateParams {

	fieldSelectionParams := field.FieldSelectionDropdownTemplateParams{
		ElemPrefix:     elemPrefix,
		ButtonTitle:    "Add Filter Condition",
		ButtonIconName: "glyphicon-plus"}

	filterPanelParams := FilterPanelTemplateParams{
		ElemPrefix:           elemPrefix,
		FieldSelectionParams: fieldSelectionParams}

	return filterPanelParams
}
