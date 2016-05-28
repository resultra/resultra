package dashboard

import (
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type DashboardTemplateParams struct {
	Title            string
	DashboardID      string
	DatabaseID       string
	DashboardName    string
	NamePanelParams  propertiesSidebar.PanelTemplateParams
	StylePanelParams propertiesSidebar.PanelTemplateParams
}

func createDashboardTemplateParams(dashboardRef *dashboard.DashboardRef) DashboardTemplateParams {

	templParams := DashboardTemplateParams{
		Title:            "Design Dashboard",
		DashboardID:      dashboardRef.DashboardID,
		DatabaseID:       dashboardRef.DatabaseID,
		DashboardName:    dashboardRef.Name,
		NamePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Dashboard Name", PanelID: "dashboardName"},
		StylePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Style", PanelID: "dashboardStyle"}}

	return templParams
}
