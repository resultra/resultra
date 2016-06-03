package dashboard

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const dashboardEntityKind string = "Dashboard"

type Dashboard struct {
	DashboardID      string `json:"dashboardID"`
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
}

const dashboardIDFieldName string = "DashboardID"

type NewDashboardParams struct {
	DatabaseID string `json:"databaseID"`
	Name       string `json:"name"`
}

func NewDashboard(appEngContext appengine.Context, params NewDashboardParams) (*Dashboard, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	var newDashboard = Dashboard{
		DashboardID:      uniqueID.GenerateUniqueID(),
		ParentDatabaseID: params.DatabaseID,
		Name:             sanitizedName}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, dashboardEntityKind, &newDashboard)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new dashboard: error inserting into datastore: %v", insertErr)
	}

	return &newDashboard, nil

}

func GetDashboard(appEngContext appengine.Context, dashboardID string) (*Dashboard, error) {

	var dashboard Dashboard

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, dashboardEntityKind,
		dashboardIDFieldName, dashboardID, &dashboard); getErr != nil {
		return nil, fmt.Errorf("GetDashboard: Unable to get dashboard from datastore: error = %v", getErr)
	}

	return &dashboard, nil

}

type GetDashboardDataParams struct {
	DashboardID string `json:"dashboardID"`
}

type DashboardDataRef struct {
	BarChartsData []barChart.BarChartData `json:"barChartsData"`
}

func GetDashboardData(appEngContext appengine.Context, params GetDashboardDataParams) (*DashboardDataRef, error) {

	barChartData, getBarChartsErr := barChart.GetDashboardBarChartsData(appEngContext, params.DashboardID)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard barchart data: error = %v", getBarChartsErr)
	}

	dashboardDataRef := DashboardDataRef{
		BarChartsData: barChartData}

	return &dashboardDataRef, nil
}
