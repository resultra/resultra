// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
