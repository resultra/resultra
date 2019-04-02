// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package colProps

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/defaultValues"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type FormButtonColPropsTemplateParams struct {
	ElemPrefix              string
	LabelPanelParams        label.LabelPropertyTemplateParams
	ButtonLabelPanelParams  inputProperties.FormButtonLabelPropertyTemplateParams
	PermissionPanelParams   permissions.PermissionsPropertyTemplateParams
	DefaultValuePanelParams defaultValues.DefaultValuesPanelTemplateParams
}

func newFormButtonTemplateParams() FormButtonColPropsTemplateParams {

	elemPrefix := "button_"

	templParams := FormButtonColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "formButtonLabel"}},

		ButtonLabelPanelParams: inputProperties.FormButtonLabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Button Label", PanelID: "formButtonButtonLabel"}},

		PermissionPanelParams:   permissions.NewPermissionTemplateParams(elemPrefix, "formButtonPerms"),
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	return templParams

}
