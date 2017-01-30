package defaultValues

import (
	"resultra/datasheet/webui/common/field"
)

type DefaultValuesPanelTemplateParams struct {
	ElemPrefix           string
	FieldSelectionParams field.FieldSelectionDropdownTemplateParams
}

func NewDefaultValuesTemplateParams(elemPrefix string) DefaultValuesPanelTemplateParams {

	fieldSelectionParams := field.FieldSelectionDropdownTemplateParams{
		ElemPrefix:     elemPrefix,
		ButtonTitle:    "Add Default Value",
		ButtonIconName: "glyphicon-plus"}

	defaultValuePanelParams := DefaultValuesPanelTemplateParams{
		ElemPrefix:           elemPrefix,
		FieldSelectionParams: fieldSelectionParams}

	return defaultValuePanelParams
}
