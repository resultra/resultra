package userSelection

import (
	"resultra/datasheet/webui/common/form/components/common/newFormElemDialog"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type UserSelectionDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
}

type UserSelectionViewTemplateParams struct {
	ElemPrefix          string
	TimelinePanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams UserSelectionDesignTemplateParams
var ViewTemplateParams UserSelectionViewTemplateParams

func init() {

	elemPrefix := "userSelection_"

	DesignTemplateParams = UserSelectionDesignTemplateParams{
		ElemPrefix:        elemPrefix,
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "userSelectionFormat"},
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:  elemPrefix,
			DialogTitle: "New User Selection",
			FieldInfoPrompt: `User selections are stored in fields. Either a new field can be created to store the values, 
					or an existing field can be used.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store the user selections.`}}

	ViewTemplateParams = UserSelectionViewTemplateParams{
		ElemPrefix:          elemPrefix,
		TimelinePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Timeline", PanelID: "userSelectionTimeline"}}

}
