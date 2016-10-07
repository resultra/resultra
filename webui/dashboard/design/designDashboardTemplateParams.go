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
	RolePanelParams  propertiesSidebar.PanelTemplateParams
	ComponentParams  components.ComponentDesignTemplateParams
}

func createDashboardTemplateParams(dashboardForDesign *dashboard.Dashboard) DashboardTemplateParams {

	templParams := DashboardTemplateParams{
		Title:            "Design Dashboard",
		DashboardID:      dashboardForDesign.DashboardID,
		DatabaseID:       dashboardForDesign.ParentDatabaseID,
		DashboardName:    dashboardForDesign.Name,
		NamePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Dashboard Name", PanelID: "dashboardName"},
		StylePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Style", PanelID: "dashboardStyle"},
		RolePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles & Privileges", PanelID: "dashboardRoles"},
		ComponentParams:  components.DesignTemplateParams}

	return templParams
}
