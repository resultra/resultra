// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package defaultValues

import (
	"resultra/tracker/webui/common/field"
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
