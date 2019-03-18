package recordFilter

import (
	"resultra/tracker/webui/common/field"
)

type FilterPanelTemplateParams struct {
	ElemPrefix           string
	FieldSelectionParams field.FieldSelectionDropdownTemplateParams
}

func NewFilterPanelTemplateParams(elemPrefix string) FilterPanelTemplateParams {

	fieldSelectionParams := field.FieldSelectionDropdownTemplateParams{
		ElemPrefix:     elemPrefix,
		ButtonTitle:    "Add Condition",
		ButtonIconName: "glyphicon-plus-sign"}

	filterPanelParams := FilterPanelTemplateParams{
		ElemPrefix:           elemPrefix,
		FieldSelectionParams: fieldSelectionParams}

	return filterPanelParams
}
