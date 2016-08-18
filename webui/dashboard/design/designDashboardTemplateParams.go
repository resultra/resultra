package design

import (
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type DashboardTemplateParams struct {
	Title            string
	DashboardID      string
	DatabaseID       string
	DashboardName    string
	NamePanelParams  propertiesSidebar.PanelTemplateParams
	StylePanelParams propertiesSidebar.PanelTemplateParams
	ComponentParams  components.ComponentTemplateParams
}

func createDashboardTemplateParams(dashboardForDesign *dashboard.Dashboard) DashboardTemplateParams {

	templParams := DashboardTemplateParams{
		Title:            "Design Dashboard",
		DashboardID:      dashboardForDesign.DashboardID,
		DatabaseID:       dashboardForDesign.ParentDatabaseID,
		DashboardName:    dashboardForDesign.Name,
		NamePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Dashboard Name", PanelID: "dashboardName"},
		StylePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Style", PanelID: "dashboardStyle"},
		ComponentParams:  components.TemplateParams}

	return templParams
}
