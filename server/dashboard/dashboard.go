package dashboard

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
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
		DashboardID:      gocql.TimeUUID().String(),
		ParentDatabaseID: params.DatabaseID,
		Name:             sanitizedName}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if insertErr := dbSession.Query(`INSERT INTO dashboard (databaseID, dashboardID, name) VALUES (?,?,?)`,
		newDashboard.ParentDatabaseID, newDashboard.DashboardID, newDashboard.Name).Exec(); insertErr != nil {
		fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", insertErr)
	}

	return &newDashboard, nil

}

func GetDashboard(appEngContext appengine.Context, databaseID string, dashboardID string) (*Dashboard, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("GetDashboard: Unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	dashboardName := ""
	getErr := dbSession.Query(`SELECT name FROM dashboard
		 WHERE databaseID=? AND dashboardID=? LIMIT 1`,
		databaseID, dashboardID).Scan(&dashboardName)
	if getErr != nil {
		return nil, fmt.Errorf("GetForm: Unabled to get dashboard: datastore err=%v", getErr)
	}

	getDashboard := Dashboard{
		ParentDatabaseID: databaseID,
		DashboardID:      dashboardID,
		Name:             dashboardName}

	return &getDashboard, nil

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
