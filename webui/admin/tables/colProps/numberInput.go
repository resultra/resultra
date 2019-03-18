// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package colProps

import (
	"resultra/tracker/webui/common/form/components/common/label"
	"resultra/tracker/webui/common/form/components/common/permissions"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type NumberInputColPropsTemplateParams struct {
	ElemPrefix            string
	FormatPanelParams     propertiesSidebar.PanelTemplateParams
	ValidationPanelParams propertiesSidebar.PanelTemplateParams
	SpinnerPanelParams    propertiesSidebar.PanelTemplateParams
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newNumberInputTemplateParams() NumberInputColPropsTemplateParams {

	elemPrefix := "numberInput_"

	templParams := NumberInputColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "numberInputLabel"}},
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Number Format", PanelID: "numberInputFormat"},
		ValidationPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Input Validation", PanelID: "numberInputValidation"},
		SpinnerPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Spinner Buttons", PanelID: "numberInputSpinner"},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "numberInputPerms")}

	return templParams

}
