package dashboard

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/dashboard/components/summaryTable"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

type DashboardProperties struct {
	Layout componentLayout.ComponentLayout `json:"layout"`
}

type Dashboard struct {
	DashboardID      string              `json:"dashboardID"`
	ParentDatabaseID string              `json:"parentDatabaseID"`
	Name             string              `json:"name"`
	Properties       DashboardProperties `json:"properties"`
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

	dashboardProps := DashboardProperties{
		Layout: componentLayout.ComponentLayout{}}
	encodedDashboardProps, encodeErr := generic.EncodeJSONString(dashboardProps)
	if encodeErr != nil {
		return nil, fmt.Errorf("NewDashboard: failure encoding properties: error = %v", encodeErr)
	}

	var newDashboard = Dashboard{
		DashboardID:      uniqueID.GenerateSnowflakeID(),
		ParentDatabaseID: params.DatabaseID,
		Name:             sanitizedName,
		Properties:       dashboardProps}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO dashboards (database_id, dashboard_id, name,properties) 
			VALUES ($1,$2,$3,$4)`,
		newDashboard.ParentDatabaseID, newDashboard.DashboardID, newDashboard.Name, encodedDashboardProps); insertErr != nil {
		fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", insertErr)
	}

	return &newDashboard, nil

}

type GetDashboardParams struct {
	DashboardID string `json:"dashboardID"`
}

func GetDashboard(dashboardID string) (*Dashboard, error) {

	dashboardName := ""
	databaseID := ""
	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT database_id,name,properties
		 FROM dashboards
		 WHERE dashboard_id=$1 LIMIT 1`, dashboardID).Scan(&databaseID, &dashboardName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetDashboard: Unabled to get dashboard with ID = %v: datastore err=%v", dashboardID, getErr)
	}

	var dashboardProps DashboardProperties
	if decodeErr := generic.DecodeJSONString(encodedProps, &dashboardProps); decodeErr != nil {
		return nil, fmt.Errorf("GetDashboard: can't decode properties: %v", encodedProps)
	}

	getDashboard := Dashboard{
		ParentDatabaseID: databaseID,
		DashboardID:      dashboardID,
		Name:             dashboardName,
		Properties:       dashboardProps}

	return &getDashboard, nil

}

func updateExistingDashboard(dashboardID string, updatedDB *Dashboard) (*Dashboard, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedDB.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingDatabase: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE dashboards 
				SET name=$1,properties=$2
				WHERE dashboard_id=$3`,
		updatedDB.Name, encodedProps, dashboardID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingDashboard: Can't update dashboard properties %v: error = %v",
			dashboardID, updateErr)
	}

	return updatedDB, nil

}

type GetDashboardDataParams struct {
	DashboardID string `json:"dashboardID"`
}

type DashboardDataRef struct {
	Dashboard         Dashboard                       `json:"dashboard"`
	BarChartsData     []barChart.BarChartData         `json:"barChartsData"`
	SummaryTablesData []summaryTable.SummaryTableData `json:"summaryTablesData"`
}

func GetDashboardData(params GetDashboardDataParams) (*DashboardDataRef, error) {

	dashboard, err := GetDashboard(params.DashboardID)
	if err != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard: error = %v", err)
	}

	barChartData, getBarChartsErr := barChart.GetDashboardBarChartsData(params.DashboardID)
	if getBarChartsErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard barchart data: error = %v", getBarChartsErr)
	}

	summaryTablesData, getTableErr := summaryTable.GetDashboardSummaryTablesData(params.DashboardID)
	if getTableErr != nil {
		return nil, fmt.Errorf("GetDashboardData: Can't retrieve dashboard summary tables data: error = %v", getTableErr)
	}

	dashboardDataRef := DashboardDataRef{
		Dashboard:         *dashboard,
		BarChartsData:     barChartData,
		SummaryTablesData: summaryTablesData}

	return &dashboardDataRef, nil
}

func validateUniqueDashboardName(databaseID string, dashboardID string, dashboardName string) error {
	// Query to validate the name is unique:
	// 1. Select all the dashboards in the same database
	// 2. Include dashboards with the same name.
	// 3. Exclude dashboards with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT dashboards.dashboard_id 
			FROM dashboards,databases
			WHERE databases.database_id=$1 AND
				dashboards.database_id=databases.database_id AND 
				dashboards.name=$2 AND dashboards.dashboard_id<>$3`,
		databaseID, dashboardName, dashboardID)
	if queryErr != nil {
		return fmt.Errorf("System error validating dashboard name (%v)", queryErr)
	}

	existingDashboardNameUsedByAnotherDashboard := rows.Next()
	if existingDashboardNameUsedByAnotherDashboard {
		return fmt.Errorf("Invalid dashboard name - names must be unique")
	}

	return nil

}

func getDashboardDatabaseID(dashboardID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT databases.database_id 
			FROM databases,dashboards 
			WHERE dashboards.dashboard_id=$1 
				AND dashboards.database_id=databases.database_id LIMIT 1`,
		dashboardID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getDashboardDatabaseID: can't get database for dashboard = %v: err=%v",
			dashboardID, getErr)
	}

	return databaseID, nil

}

func validateNewDashboardName(databaseID string, dashboardName string) error {

	if !generic.WellFormedItemName(dashboardName) {
		return fmt.Errorf("Invalid dashboard name")
	}

	// No dashboard will have an empty dashboardID, so this will cause test for unique
	// dashboard names to return true if any dashboard already has the given dashboardName.
	dashboardID := ""
	if uniqueErr := validateUniqueDashboardName(databaseID, dashboardID, dashboardName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateDashboardName(dashboardID string, dashboardName string) error {

	if !generic.WellFormedItemName(dashboardName) {
		return fmt.Errorf("Invalid dashboard name")
	}

	databaseID, err := getDashboardDatabaseID(dashboardID)
	if err != nil {
		return fmt.Errorf("System error validating name")
	}

	if uniqueErr := validateUniqueDashboardName(databaseID, dashboardID, dashboardName); uniqueErr != nil {
		return uniqueErr
	}

	return nil

}

func validateComponentTitle(title string) error {

	if !generic.WellFormedItemName(title) {
		return fmt.Errorf("Invalid title")
	}

	return nil
}
