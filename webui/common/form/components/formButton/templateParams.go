package formButton

import (
	"resultra/datasheet/webui/common/defaultValues"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type PopupBehaviorPropParams struct {
	DefaultValuePanelParams defaultValues.DefaultValuesPanelTemplateParams
}

type ButtonTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	LinkedFormPanelParams    propertiesSidebar.PanelTemplateParams
	PopupBehaviorPanelParams propertiesSidebar.PanelTemplateParams
	PopupBehaviorPropParams  PopupBehaviorPropParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
}

var TemplateParams ButtonTemplateParams

func init() {

	elemPrefix := "button_"
	visibilityElemPrefix := "buttonVisibility_"

	popupBehaviorParams := PopupBehaviorPropParams{
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	TemplateParams = ButtonTemplateParams{
		ElemPrefix:               elemPrefix,
		FormatPanelParams:        propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "buttonFormat"},
		LinkedFormPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Linked Form", PanelID: "buttonForm"},
		VisibilityPanelParams:    visibility.NewComponentVisibilityTempalteParams(visibilityElemPrefix, "buttonVisibility"),
		PopupBehaviorPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Popup Form Behavior", PanelID: "buttonPopupForm"},
		PopupBehaviorPropParams:  popupBehaviorParams}
}
