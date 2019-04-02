// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formButton

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/defaultValues"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type PopupBehaviorPropParams struct {
	DefaultValuePanelParams defaultValues.DefaultValuesPanelTemplateParams
}

type ButtonTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ButtonLabelPanelParams   inputProperties.FormButtonLabelPropertyTemplateParams
	PopupBehaviorPanelParams propertiesSidebar.PanelTemplateParams
	PopupBehaviorPropParams  PopupBehaviorPropParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
}

var TemplateParams ButtonTemplateParams

func init() {

	elemPrefix := "button_"
	visibilityElemPrefix := "buttonVisibility_"

	popupBehaviorParams := PopupBehaviorPropParams{
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	TemplateParams = ButtonTemplateParams{
		ElemPrefix:        elemPrefix,
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "buttonFormat"},

		ButtonLabelPanelParams: inputProperties.FormButtonLabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Button Label", PanelID: "formButtonButtonLabel"}},

		DeletePanelParams:        delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "buttonDelete", "Delete Form Button"),
		VisibilityPanelParams:    visibility.NewComponentVisibilityTempalteParams(visibilityElemPrefix, "buttonVisibility"),
		PopupBehaviorPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Form Behavior", PanelID: "buttonPopupForm"},
		PopupBehaviorPropParams:  popupBehaviorParams}
}
