package itemList

import (
	"resultra/datasheet/webui/common/form/components"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ViewListTemplateParams struct {
	DisplayPanelParams      propertiesSidebar.PanelTemplateParams
	FilteringPanelParams    propertiesSidebar.PanelTemplateParams
	SortPanelParams         propertiesSidebar.PanelTemplateParams
	ComponentParams         components.ComponentViewTemplateParams
	FilterConfigPanelParams recordFilter.FilterPanelTemplateParams
}

var ViewListTemplParams ViewListTemplateParams

func init() {

	elemPrefix := "form_"

	ViewListTemplParams = ViewListTemplateParams{
		DisplayPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "View",
			PanelID: "viewListDisplay"},
		FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering",
			PanelID: "viewFormFiltering"},
		FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		SortPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Sorting",
			PanelID: "viewFormSorting"},
		ComponentParams: components.ViewTemplateParams}

}
