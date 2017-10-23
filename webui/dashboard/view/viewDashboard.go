package view

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/webui/dashboard/components"
)

type ViewDashboardInfo struct {
	DatabaseID      string
	DatabaseName    string
	DashboardID     string
	DashboardName   string
	CurrUserIsAdmin bool
	Title           string
	ComponentParams components.ComponentViewTemplateParams
}

func getViewDashboardInfo(r *http.Request, dashboardID string) (*ViewDashboardInfo, error) {

	dashboardDbInfo, getErr := databaseController.GetDashboardDatabaseInfo(dashboardID)
	if getErr != nil {
		return nil, getErr
	}

	hasViewPrivs, privsErr := userRole.CurrentUserHasDashboardViewPrivs(r, dashboardDbInfo.DatabaseID, dashboardID)
	if privsErr != nil {
		return nil, privsErr
	}
	if !hasViewPrivs {
		return nil, fmt.Errorf("ERROR: No permissions to view this dashboard")
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dashboardDbInfo.DatabaseID)

	dashboardInfo := ViewDashboardInfo{
		DatabaseID:      dashboardDbInfo.DatabaseID,
		DatabaseName:    dashboardDbInfo.DatabaseName,
		DashboardID:     dashboardDbInfo.DashboardID,
		DashboardName:   dashboardDbInfo.DashboardName,
		CurrUserIsAdmin: isAdmin}

	return &dashboardInfo, nil

}
