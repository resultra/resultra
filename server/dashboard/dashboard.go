package dashboard

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/common/datastoreWrapper"
	"resultra/datasheet/server/dashboard/barChart"
	"resultra/datasheet/server/dataModel"
)

type Dashboard struct {
	Name string `json:"name"`
}

type DashboardRef struct {
	DashboardID string `json:"dashboardID"`
	Name        string `json:"name"`
}

func NewDashboard(appEngContext appengine.Context, dashboardName string) (*DashboardRef, error) {

	sanitizedName, sanitizeErr := common.SanitizeName(dashboardName)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	var newDashboard = Dashboard{sanitizedName}
	dashboardID, insertErr := datastoreWrapper.InsertNewEntity(appEngContext, dataModel.DashboardEntityKind, nil, &newDashboard)
	if insertErr != nil {
		return nil, insertErr
	}

	log.Printf("NewDashboard: Created new dashboard: id= %v, name='%v'", dashboardID, sanitizedName)
	dashboardRef := DashboardRef{dashboardID, sanitizedName}

	return &dashboardRef, nil

}

func GetDashboardRef(appEngContext appengine.Context, dashboardID string) (*DashboardRef, error) {

	var dashboard DashboardRef
	getErr := datastoreWrapper.GetRootEntityByID(appEngContext, dataModel.DashboardEntityKind, dashboardID, &dashboard)
	if getErr != nil {
		return nil, fmt.Errorf("GetDashboardRef: Can't get dashboard: Error retrieving existing dashboard: dashboard ID=%v, err = %v", dashboardID, getErr)
	}

	return &DashboardRef{dashboardID, dashboard.Name}, nil

}

func getDashboardKey(appEngContext appengine.Context, dashboardID string) (*datastore.Key, error) {
	dashboardKey, getDashboardErr := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.DashboardEntityKind,
		dashboardID)
	if getDashboardErr != nil {
		return nil, fmt.Errorf("getDashboardKey: Invalid dashboard: %v", getDashboardErr)
	}
	return dashboardKey, nil
}

type GetDashboardDataParams struct {
	DashboardID string `json:"dashboardID"`
}

type DashboardDataRef struct {
	BarChartsData []barChart.BarChartData `json:"barChartsData"`
}

func GetDashboardData(appEngContext appengine.Context, params GetDashboardDataParams) (*DashboardDataRef, error) {

	parentDashboardKey, err := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.DashboardEntityKind,
		params.DashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard: error = %v", err)
	}

	barChartData, getBarChartsErr := barChart.GetDashboardBarChartsData(appEngContext, params.DashboardID, parentDashboardKey)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard barchart data: error = %v", getBarChartsErr)
	}

	dashboardDataRef := DashboardDataRef{
		BarChartsData: barChartData}

	return &dashboardDataRef, nil
}
