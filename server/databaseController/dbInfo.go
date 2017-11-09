package databaseController

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/itemList"
	"resultra/datasheet/server/trackerDatabase"
)

type DatabaseInfoParams struct {
	DatabaseID string `json:"databaseID"`
}

func getDatabaseDashboardsInfo(trackerDBHandle *sql.DB, params DatabaseInfoParams) ([]dashboard.Dashboard, error) {

	dashboards, err := dashboard.GetAllSortedDashboard(trackerDBHandle, params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getDatabaseDashboardsInfo: %v", err)
	}

	return dashboards, nil

}

func getDatabaseFormsInfo(trackerDBHandle *sql.DB, params DatabaseInfoParams) ([]form.Form, error) {

	formsInfo, getFormsErr := form.GetAllForms(trackerDBHandle, params.DatabaseID)
	if getFormsErr != nil {
		return nil, fmt.Errorf("getDatabaseFormsInfo: Failure querying database: %v", getFormsErr)
	}

	return formsInfo, nil
}

func getDatabaseItemListInfo(trackerDBHandle *sql.DB, params DatabaseInfoParams) ([]itemList.ItemList, error) {

	listInfo, getsListsErr := itemList.GetAllSortedItemLists(trackerDBHandle, params.DatabaseID)
	if getsListsErr != nil {
		return nil, fmt.Errorf("getDatabaseItemListInfo: %v", getsListsErr)
	}
	return listInfo, nil
}

type DatabaseContentsInfo struct {
	DatabaseInfo   trackerDatabase.Database `json:"databaseInfo"`
	FormsInfo      []form.Form              `json:"formsInfo"`
	ListsInfo      []itemList.ItemList      `json:"listsInfo"`
	DashboardsInfo []dashboard.Dashboard    `json:"dashboardsInfo"`
}

func getDatabaseInfo(trackerDBHandle *sql.DB, params DatabaseInfoParams) (*DatabaseContentsInfo, error) {

	db, getErr := trackerDatabase.GetDatabase(trackerDBHandle, params.DatabaseID)
	if getErr != nil {
		return nil, fmt.Errorf("getDatabaseInfo: Unable to get existing database: %v", getErr)
	}

	formsInfo, formsErr := getDatabaseFormsInfo(trackerDBHandle, params)
	if formsErr != nil {
		return nil, formsErr
	}

	dashboardsInfo, dashboardsErr := getDatabaseDashboardsInfo(trackerDBHandle, params)
	if dashboardsErr != nil {
		return nil, dashboardsErr
	}

	listsInfo, err := getDatabaseItemListInfo(trackerDBHandle, params)
	if err != nil {
		return nil, err
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

func GetFormDatabaseInfo(trackerDBHandle *sql.DB, formID string) (*FormDatabaseInfo, error) {

	var formDBInfo FormDatabaseInfo
	getErr := trackerDBHandle.QueryRow(`
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

func GetDashboardDatabaseInfo(trackerDBHandle *sql.DB, dashboardID string) (*DashboardDatabaseInfo, error) {

	var dashDBInfo DashboardDatabaseInfo
	getErr := trackerDBHandle.QueryRow(`
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

func GetDatabaseInfo(trackerDBHandle *sql.DB, databaseID string) (*DatabaseInfo, error) {

	var dbInfo DatabaseInfo
	getErr := trackerDBHandle.QueryRow(`
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
