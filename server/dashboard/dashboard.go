package dashboard

import (
	"fmt"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
)

type Dashboard struct {
	DashboardID      string `json:"dashboardID"`
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
}

type NewDashboardParams struct {
	DatabaseID string `json:"databaseID"`
	Name       string `json:"name"`
}

func NewDashboard(params NewDashboardParams) (*Dashboard, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	var newDashboard = Dashboard{
		DashboardID:      databaseWrapper.GlobalUniqueID(),
		ParentDatabaseID: params.DatabaseID,
		Name:             sanitizedName}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO dashboards (database_id, dashboard_id, name) VALUES ($1,$2,$3)`,
		newDashboard.ParentDatabaseID, newDashboard.DashboardID, newDashboard.Name); insertErr != nil {
		fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", insertErr)
	}

	return &newDashboard, nil

}

func GetDashboard(databaseID string, dashboardID string) (*Dashboard, error) {

	dashboardName := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT name FROM dashboards
		 WHERE database_id=$1 AND dashboard_id=$2 LIMIT 1`,
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

func GetDashboardData(params GetDashboardDataParams) (*DashboardDataRef, error) {

	barChartData, getBarChartsErr := barChart.GetDashboardBarChartsData(params.DashboardID)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard barchart data: error = %v", getBarChartsErr)
	}

	dashboardDataRef := DashboardDataRef{
		BarChartsData: barChartData}

	return &dashboardDataRef, nil
}
