// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package design

import (
	"fmt"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/server/dashboard"
	"github.com/resultra/resultra/server/trackerDatabase"
	"github.com/resultra/resultra/server/userRole"
	"github.com/resultra/resultra/webui/dashboard/components"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
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
		IsSingleUserWorkspace:       runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),
		NamePanelParams:             propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Dashboard Name", PanelID: "dashboardName"},
		IncludeInSidebarPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Include in Sidebar", PanelID: "includeInSidebar"},
		RolePanelParams:             propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Roles & Privileges", PanelID: "dashboardRoles"},
		ComponentParams:             components.DesignTemplateParams}

	return &templParams, nil
}
