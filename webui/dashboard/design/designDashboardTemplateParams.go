package design

import (
	"fmt"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/database"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type DashboardTemplateParams struct {
	Title           string
	DashboardID     string
	DatabaseID      string
	DatabaseName    string
	DashboardName   string
	NamePanelParams propertiesSidebar.PanelTemplateParams
	RolePanelParams propertiesSidebar.PanelTemplateParams
	ComponentParams components.ComponentDesignTemplateParams
}

func createDashboardTemplateParams(dashboardForDesign *dashboard.Dashboard) (*DashboardTemplateParams, error) {

	dashboardDB, err := database.GetDatabase(dashboardForDesign.ParentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("createDashboardTemplateParams: %v", err)
	}

	templParams := DashboardTemplateParams{
		Title:           "Design Dashboard",
		DashboardID:     dashboardForDesign.DashboardID,
		DatabaseID:      dashboardForDesign.ParentDatabaseID,
		DatabaseName:    dashboardDB.Name,
		DashboardName:   dashboardForDesign.Name,
		NamePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Dashboard Name", PanelID: "dashboardName"},
		RolePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles & Privileges", PanelID: "dashboardRoles"},
		ComponentParams: components.DesignTemplateParams}

	return &templParams, nil
}
