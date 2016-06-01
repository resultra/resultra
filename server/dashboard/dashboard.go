package dashboard

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const dashboardEntityKind string = "Dashboard"

type Dashboard struct {
	Name string `json:"name"`
}

type DashboardRef struct {
	DatabaseID  string `json:"databaseID"`
	DashboardID string `json:"dashboardID"`
	Name        string `json:"name"`
}

type NewDashboardParams struct {
	DatabaseID string `json:"databaseID"`
	Name       string `json:"name"`
}

func NewDashboard(appEngContext appengine.Context, params NewDashboardParams) (*DashboardRef, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	if err := datastoreWrapper.ValidateEntityKind(params.DatabaseID, database.DatabaseEntityKind); err != nil {
		return nil, err
	}

	var newDashboard = Dashboard{sanitizedName}
	dashboardID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext, params.DatabaseID, dashboardEntityKind, &newDashboard)
	if insertErr != nil {
		return nil, insertErr
	}

	log.Printf("NewDashboard: Created new dashboard: id= %v, name='%v'", dashboardID, sanitizedName)
	dashboardRef := DashboardRef{DatabaseID: params.DatabaseID, DashboardID: dashboardID, Name: sanitizedName}

	return &dashboardRef, nil

}

func GetDashboardRef(appEngContext appengine.Context, dashboardID string) (*DashboardRef, error) {

	var dashboard DashboardRef
	getErr := datastoreWrapper.GetEntity(appEngContext, dashboardID, &dashboard)
	if getErr != nil {
		return nil, fmt.Errorf("GetDashboardRef: Can't get dashboard: Error retrieving existing dashboard: dashboard ID=%v, err = %v", dashboardID, getErr)
	}

	databaseID, getDatabaseIDErr := datastoreWrapper.GetParentID(dashboardID, database.DatabaseEntityKind)
	if getDatabaseIDErr != nil {
		return nil, fmt.Errorf("GetDashboardRef: Unable to get parent database ID: error = %v", getDatabaseIDErr)
	}

	return &DashboardRef{DatabaseID: databaseID, DashboardID: dashboardID, Name: dashboard.Name}, nil

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
