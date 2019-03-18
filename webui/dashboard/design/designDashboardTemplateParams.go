package design

import (
	"fmt"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/dashboard"
	"resultra/tracker/server/trackerDatabase"
	"resultra/tracker/server/userRole"
	"resultra/tracker/webui/dashboard/components"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type DashboardTemplateParams struct {
	Title                       string
	DashboardID                 string
	DatabaseID                  string
	DatabaseName                string
	DashboardName               string
	WorkspaceName               string
	CurrUserIsAdmin             bool
	IsSingleUserWorkspace       bool
	NamePanelParams             propertiesSidebar.PanelTemplateParams
	RolePanelParams             propertiesSidebar.PanelTemplateParams
	IncludeInSidebarPanelParams propertiesSidebar.PanelTemplateParams
	ComponentParams             components.ComponentDesignTemplateParams
}

func createDashboardTemplateParams(r *http.Request, dashboardForDesign *dashboard.Dashboard,
	workspaceName string) (*DashboardTemplateParams, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		return nil, dbErr
	}

	dashboardDB, err := trackerDatabase.GetDatabase(trackerDBHandle, dashboardForDesign.ParentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("createDashboardTemplateParams: %v", err)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dashboardForDesign.ParentDatabaseID)

	templParams := DashboardTemplateParams{
		Title:                       "Design Dashboard",
		DashboardID:                 dashboardForDesign.DashboardID,
		DatabaseID:                  dashboardForDesign.ParentDatabaseID,
		DatabaseName:                dashboardDB.Name,
		WorkspaceName:               workspaceName,
		DashboardName:               dashboardForDesign.Name,
		CurrUserIsAdmin:             isAdmin,
		IsSingleUserWorkspace:       runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
		NamePanelParams:             propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Dashboard Name", PanelID: "dashboardName"},
		IncludeInSidebarPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Include in Sidebar", PanelID: "includeInSidebar"},
		RolePanelParams:             propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles & Privileges", PanelID: "dashboardRoles"},
		ComponentParams:             components.DesignTemplateParams}

	return &templParams, nil
}
