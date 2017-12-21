package formButton

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/defaultValues"
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
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
