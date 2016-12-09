package databaseController

import (
	"fmt"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/generic/databaseWrapper"
)

type DatabaseInfoParams struct {
	DatabaseID string `json:"databaseID"`
}

type FormInfo struct {
	FormID string `json:"formID"`
	Name   string `json:"name"`
}

type DashboardInfo struct {
	DashboardID string `json:"dashboardID"`
	Name        string `json:"name"`
}

type ItemListInfo struct {
	ListID string `json:"listID"`
	Name   string `json:"name"`
}

func getDatabaseDashboardsInfo(params DatabaseInfoParams) ([]DashboardInfo, error) {
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT dashboards.dashboard_id, dashboards.name FROM dashboards,databases WHERE 
			databases.database_id=$1 AND 
			dashboards.database_id = databases.database_id`,
		params.DatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getDatabaseDashboardsInfo: Failure querying database: %v", queryErr)
	}

	dashboardsInfo := []DashboardInfo{}
	for rows.Next() {
		var currDashboardInfo DashboardInfo
		if scanErr := rows.Scan(&currDashboardInfo.DashboardID, &currDashboardInfo.Name); scanErr != nil {
			return nil, fmt.Errorf("getDatabaseDashboardsInfo: Failure querying database: %v", scanErr)
		}
		dashboardsInfo = append(dashboardsInfo, currDashboardInfo)
	}

	return dashboardsInfo, nil

}

func getDatabaseFormsInfo(params DatabaseInfoParams) ([]FormInfo, error) {
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT form_id, name FROM forms WHERE  database_id=$1`,
		params.DatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getDatabaseInfo: Failure querying database: %v", queryErr)
	}

	formsInfo := []FormInfo{}
	for rows.Next() {
		var currFormInfo FormInfo
		if scanErr := rows.Scan(&currFormInfo.FormID, &currFormInfo.Name); scanErr != nil {
			return nil, fmt.Errorf("getDatabaseInfo: Failure querying database: %v", scanErr)
		}
		formsInfo = append(formsInfo, currFormInfo)
	}

	return formsInfo, nil
}

func getDatabaseItemListInfo(params DatabaseInfoParams) ([]ItemListInfo, error) {
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT list_id, name FROM item_lists WHERE database_id=$1`,
		params.DatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getDatabaseInfo: Failure querying database: %v", queryErr)
	}

	listsInfo := []ItemListInfo{}
	for rows.Next() {
		var currListInfo ItemListInfo
		if scanErr := rows.Scan(&currListInfo.ListID, &currListInfo.Name); scanErr != nil {
			return nil, fmt.Errorf("getDatabaseInfo: Failure querying database: %v", scanErr)
		}
		listsInfo = append(listsInfo, currListInfo)
	}

	return listsInfo, nil
}

type DatabaseContentsInfo struct {
	DatabaseInfo   database.Database `json:"databaseInfo"`
	FormsInfo      []FormInfo        `json:"formsInfo"`
	ListsInfo      []ItemListInfo    `json:"listsInfo"`
	DashboardsInfo []DashboardInfo   `json:"dashboardsInfo"`
}

func getDatabaseInfo(params DatabaseInfoParams) (*DatabaseContentsInfo, error) {

	formsInfo, formsErr := getDatabaseFormsInfo(params)
	if formsErr != nil {
		return nil, formsErr
	}

	dashboardsInfo, dashboardsErr := getDatabaseDashboardsInfo(params)
	if dashboardsErr != nil {
		return nil, dashboardsErr
	}

	listsInfo, err := getDatabaseItemListInfo(params)
	if err != nil {
		return nil, err
	}

	db, getErr := database.GetDatabase(params.DatabaseID)
	if getErr != nil {
		return nil, fmt.Errorf("getDatabaseInfo: Unable to get existing database: %v", getErr)
	}

	dbInfo := DatabaseContentsInfo{
		DatabaseInfo:   *db,
		FormsInfo:      formsInfo,
		ListsInfo:      listsInfo,
		DashboardsInfo: dashboardsInfo}

	return &dbInfo, nil
}

type FormDatabaseInfo struct {
	FormID       string
	DatabaseID   string
	DatabaseName string
	FormName     string
}

func GetFormDatabaseInfo(formID string) (*FormDatabaseInfo, error) {

	var formDBInfo FormDatabaseInfo
	getErr := databaseWrapper.DBHandle().QueryRow(`
			SELECT 
				databases.database_id, databases.name AS database_name, forms.form_id, forms.name 
			FROM 
				forms,databases 
			WHERE 
				forms.form_id = $1 AND 
				forms.database_id = databases.database_id`, formID).Scan(
		&formDBInfo.DatabaseID, &formDBInfo.DatabaseName, &formDBInfo.FormID, &formDBInfo.FormName)
	if getErr != nil {
		return nil, fmt.Errorf("GetFormDatabaseInfo: Unabled to get form info: form id = %v: datastore err=%v",
			formID, getErr)
	}

	return &formDBInfo, nil

}

type DashboardDatabaseInfo struct {
	DashboardID   string
	DatabaseID    string
	DatabaseName  string
	DashboardName string
}

func GetDashboardDatabaseInfo(dashboardID string) (*DashboardDatabaseInfo, error) {

	var dashDBInfo DashboardDatabaseInfo
	getErr := databaseWrapper.DBHandle().QueryRow(`
			SELECT 
				databases.database_id, databases.name, dashboards.dashboard_id,dashboards.name
			FROM 
				dashboards,databases 
			WHERE 
				dashboards.dashboard_id = $1 AND 
				dashboards.database_id = databases.database_id`, dashboardID).Scan(
		&dashDBInfo.DatabaseID, &dashDBInfo.DatabaseName,
		&dashDBInfo.DashboardID, &dashDBInfo.DashboardName)
	if getErr != nil {
		return nil, fmt.Errorf("GetDashboardDatabaseInfo: Unabled to get dashboard info: dashboard id = %v: datastore err=%v",
			dashboardID, getErr)
	}

	return &dashDBInfo, nil

}

type DatabaseInfo struct {
	DatabaseID   string
	DatabaseName string
}

func GetDatabaseInfo(databaseID string) (*DatabaseInfo, error) {

	var dbInfo DatabaseInfo
	getErr := databaseWrapper.DBHandle().QueryRow(`
			SELECT 
				database_id, name 
			FROM 
				databases 
			WHERE 
				database_id = $1
			LIMIT 1`, databaseID).Scan(
		&dbInfo.DatabaseID, &dbInfo.DatabaseName)
	if getErr != nil {
		return nil, fmt.Errorf("GetDatabaseInfo: Unabled to get database info: database id = %v: datastore err=%v",
			databaseID, getErr)
	}

	return &dbInfo, nil

}
