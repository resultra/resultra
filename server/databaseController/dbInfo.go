package databaseController

import (
	"fmt"
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
		`SELECT forms.form_id, forms.name FROM forms,data_tables,databases WHERE 
			databases.database_id=$1 AND 
			data_tables.database_id = databases.database_id AND
			forms.table_id = data_tables.table_id`,
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

type DatabaseContentsInfo struct {
	FormsInfo      []FormInfo      `json:"formsInfo"`
	DashboardsInfo []DashboardInfo `json:"dashboardsInfo"`
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

	dbInfo := DatabaseContentsInfo{
		FormsInfo:      formsInfo,
		DashboardsInfo: dashboardsInfo}

	return &dbInfo, nil
}

type FormDatabaseInfo struct {
	FormID       string
	TableID      string
	DatabaseID   string
	DatabaseName string
	FormName     string
}

func GetFormDatabaseInfo(formID string) (*FormDatabaseInfo, error) {

	var formDBInfo FormDatabaseInfo
	getErr := databaseWrapper.DBHandle().QueryRow(`
			SELECT 
				databases.database_id, databases.name AS database_name, data_tables.table_id, forms.form_id, forms.name 
			FROM 
				forms,data_tables,databases 
			WHERE 
				forms.form_id = $1 AND 
				data_tables.database_id = databases.database_id AND
				forms.table_id = data_tables.table_id`, formID).Scan(
		&formDBInfo.DatabaseID, &formDBInfo.DatabaseName, &formDBInfo.TableID, &formDBInfo.FormID, &formDBInfo.FormName)
	if getErr != nil {
		return nil, fmt.Errorf("GetFormDatabaseInfo: Unabled to get form info: form id = %v: datastore err=%v",
			formID, getErr)
	}

	return &formDBInfo, nil

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
