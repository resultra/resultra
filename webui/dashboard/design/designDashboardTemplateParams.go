package design

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/trackerDatabase"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type DashboardTemplateParams struct {
	Title           string
	DashboardID     string
	DatabaseID      string
	DatabaseName    string
	DashboardName   string
	CurrUserIsAdmin bool
	NamePanelParams propertiesSidebar.PanelTemplateParams
	RolePanelParams propertiesSidebar.PanelTemplateParams
	ComponentParams components.ComponentDesignTemplateParams
}

func createDashboardTemplateParams(r *http.Request, dashboardForDesign *dashboard.Dashboard) (*DashboardTemplateParams, error) {

	dashboardDB, err := trackerDatabase.GetDatabase(dashboardForDesign.ParentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("createDashboardTemplateParams: %v", err)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dashboardForDesign.ParentDatabaseID)

	templParams := DashboardTemplateParams{
		Title:           "Design Dashboard",
		DashboardID:     dashboardForDesign.DashboardID,
		DatabaseID:      dashboardForDesign.ParentDatabaseID,
		DatabaseName:    dashboardDB.Name,
		DashboardName:   dashboardForDesign.Name,
		CurrUserIsAdmin: isAdmin,
		NamePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Dashboard Name", PanelID: "dashboardName"},
		RolePanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles & Privileges", PanelID: "dashboardRoles"},
		ComponentParams: components.DesignTemplateParams}

	return &templParams, nil
}
