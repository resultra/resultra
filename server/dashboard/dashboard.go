package dashboard

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/trackerDatabase"
	"resultra/datasheet/server/userRole"
)

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

func saveDashboard(destDBHandle *sql.DB, newDashboard Dashboard) error {

	encodedDashboardProps, err := generic.EncodeJSONString(newDashboard.Properties)
	if err != nil {
		return fmt.Errorf("NewDashboard: failure encoding properties: error = %v", err)
	}

	if _, err := destDBHandle.Exec(`INSERT INTO dashboards (database_id, dashboard_id, name,properties) 
			VALUES ($1,$2,$3,$4)`,
		newDashboard.ParentDatabaseID, newDashboard.DashboardID, newDashboard.Name, encodedDashboardProps); err != nil {
		return fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", err)
	}

	return nil
}

func NewDashboard(trackerDBHandle *sql.DB, params NewDashboardParams) (*Dashboard, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	dashboardProps := DashboardProperties{
		Layout: componentLayout.ComponentLayout{}}

	var newDashboard = Dashboard{
		DashboardID:      uniqueID.GenerateSnowflakeID(),
		ParentDatabaseID: params.DatabaseID,
		Name:             sanitizedName,
		Properties:       dashboardProps}

	if err := saveDashboard(trackerDBHandle, newDashboard); err != nil {
		fmt.Errorf("NewDashboard: Can't create dashboard: unable to create dashboard: error = %v", err)
	}

	return &newDashboard, nil

}

type GetDashboardParams struct {
	DashboardID string `json:"dashboardID"`
}

func GetDashboard(trackerDBHandle *sql.DB, dashboardID string) (*Dashboard, error) {

	dashboardName := ""
	databaseID := ""
	encodedProps := ""
	getErr := trackerDBHandle.QueryRow(`SELECT database_id,name,properties
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

func getAllDashboardsFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]Dashboard, error) {

	rows, err := srcDBHandle.Query(
		`SELECT database_id,dashboard_id,name,properties
		 FROM dashboards
		 WHERE database_id = $1`,
		parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetAllDashboards: Failure querying database: %v", err)
	}
	defer rows.Close()

	dashboards := []Dashboard{}
	for rows.Next() {
		currDashboard := Dashboard{}
		encodedProps := ""

		if err := rows.Scan(&currDashboard.ParentDatabaseID,
			&currDashboard.DashboardID,
			&currDashboard.Name, &encodedProps); err != nil {
			return nil, fmt.Errorf("GetAllDashboards: Failure querying database: %v", err)
		}

		var dashboardProps DashboardProperties
		if err := generic.DecodeJSONString(encodedProps, &dashboardProps); err != nil {
			return nil, fmt.Errorf("GetAllDashboards: can't decode properties: %v,error=%v", encodedProps, err)
		}
		currDashboard.Properties = dashboardProps

		dashboards = append(dashboards, currDashboard)
	}

	return dashboards, nil

}

func GetAllDashboards(trackerDBHandle *sql.DB, parentDatabaseID string) ([]Dashboard, error) {
	return getAllDashboardsFromSrc(trackerDBHandle, parentDatabaseID)
}

type GetUserDashboardListParams struct {
	DatabaseID string `json:"databaseID"`
}

func getUserDashboards(req *http.Request, databaseID string) ([]Dashboard, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	allDashboards, err := GetAllDashboards(trackerDBHandle, databaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserDashboards: %v", err)
	}

	if userRole.CurrUserIsDatabaseAdmin(req, databaseID) {
		return allDashboards, nil
	}

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("getUserDashboards: can't verify user: %v", userErr)
	}

	viewableDashboards, privsErr := userRole.GetDashboardsWithUserViewPrivs(trackerDBHandle, databaseID, currUserID)
	if privsErr != nil {
		return nil, fmt.Errorf("getUserDashboards: can't verify user: %v", privsErr)
	}

	userDashboards := []Dashboard{}
	for _, currDashboard := range allDashboards {
		_, foundPriv := viewableDashboards[currDashboard.DashboardID]
		if foundPriv {
			userDashboards = append(userDashboards, currDashboard)
		}
	}

	return userDashboards, nil
}

func orderDashboardsByManualOrder(unorderedDashboards []Dashboard, manualOrder []string) []Dashboard {
	// Map the listID -> ListInfo.
	infoByID := map[string]Dashboard{}
	for _, currInfo := range unorderedDashboards {
		infoByID[currInfo.DashboardID] = currInfo
	}
	// Iterate throught the manually ordered list of ListIDs, pull items from listInfoByID in
	// the order they are encountered in the ordered list, then re-append the ListInfo's into a
	// new ordered list in the same order they are found.
	orderedInfo := []Dashboard{}
	for _, currID := range manualOrder {
		dashboardInfo, foundInfo := infoByID[currID]
		if foundInfo {
			orderedInfo = append(orderedInfo, dashboardInfo)
			delete(infoByID, currID)
		}
	}
	for _, currInfo := range infoByID {
		orderedInfo = append(orderedInfo, currInfo)
	}
	return orderedInfo

}

func GetAllSortedDashboard(trackerDBHandle *sql.DB, parentDatabaseID string) ([]Dashboard, error) {

	unorderedDashboards, err := GetAllDashboards(trackerDBHandle, parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetAllSortedItemLists: %v")
	}

	db, getErr := trackerDatabase.GetDatabase(trackerDBHandle, parentDatabaseID)
	if getErr != nil {
		return nil, fmt.Errorf("getDatabaseInfo: Unable to get existing database: %v", getErr)
	}

	orderedDashboards := orderDashboardsByManualOrder(unorderedDashboards, db.Properties.DashboardOrder)

	return orderedDashboards, nil

}

func CloneDashboards(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	remappedDatabaseID, err := cloneParams.IDRemapper.GetExistingRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneDashboards: Error getting remapped database ID: %v", err)
	}

	dashboards, err := getAllDashboardsFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneDashboards: Error getting dashboards for parent database ID = %v: %v",
			cloneParams.SourceDatabaseID, err)
	}

	for _, currDashboard := range dashboards {

		destDashboard := currDashboard
		destDashboard.ParentDatabaseID = remappedDatabaseID

		destDashboardID, err := cloneParams.IDRemapper.AllocNewRemappedID(currDashboard.DashboardID)
		if err != nil {
			return fmt.Errorf("CloneDashboards: %v", err)
		}
		destDashboard.DashboardID = destDashboardID

		destProps, err := currDashboard.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneDashboards: %v", err)
		}
		destDashboard.Properties = *destProps

		if err := saveDashboard(cloneParams.DestDBHandle, destDashboard); err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}

		if err := cloneDashboardComponents(cloneParams, currDashboard.DashboardID); err != nil {
			return fmt.Errorf("CloneDashboards: %v", err)
		}

	}

	return nil

}

func updateExistingDashboard(trackerDBHandle *sql.DB, dashboardID string, updatedDB *Dashboard) (*Dashboard, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedDB.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingDatabase: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE dashboards 
				SET name=$1,properties=$2
				WHERE dashboard_id=$3`,
		updatedDB.Name, encodedProps, dashboardID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingDashboard: Can't update dashboard properties %v: error = %v",
			dashboardID, updateErr)
	}

	return updatedDB, nil

}

func validateUniqueDashboardName(trackerDBHandle *sql.DB, databaseID string, dashboardID string, dashboardName string) error {
	// Query to validate the name is unique:
	// 1. Select all the dashboards in the same database
	// 2. Include dashboards with the same name.
	// 3. Exclude dashboards with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT dashboards.dashboard_id 
			FROM dashboards,databases
			WHERE databases.database_id=$1 AND
				dashboards.database_id=databases.database_id AND 
				dashboards.name=$2 AND dashboards.dashboard_id<>$3`,
		databaseID, dashboardName, dashboardID)
	if queryErr != nil {
		return fmt.Errorf("System error validating dashboard name (%v)", queryErr)
	}
	defer rows.Close()

	existingDashboardNameUsedByAnotherDashboard := rows.Next()
	if existingDashboardNameUsedByAnotherDashboard {
		return fmt.Errorf("Invalid dashboard name - names must be unique")
	}

	return nil

}

func getDashboardDatabaseID(trackerDBHandle *sql.DB, dashboardID string) (string, error) {

	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
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

func validateNewDashboardName(trackerDBHandle *sql.DB, databaseID string, dashboardName string) error {

	if !stringValidation.WellFormedItemName(dashboardName) {
		return fmt.Errorf("Invalid dashboard name")
	}

	// No dashboard will have an empty dashboardID, so this will cause test for unique
	// dashboard names to return true if any dashboard already has the given dashboardName.
	dashboardID := ""
	if uniqueErr := validateUniqueDashboardName(trackerDBHandle, databaseID, dashboardID, dashboardName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateDashboardName(trackerDBHandle *sql.DB, dashboardID string, dashboardName string) error {

	if !stringValidation.WellFormedItemName(dashboardName) {
		return fmt.Errorf("Invalid dashboard name")
	}

	databaseID, err := getDashboardDatabaseID(trackerDBHandle, dashboardID)
	if err != nil {
		return fmt.Errorf("System error validating name")
	}

	if uniqueErr := validateUniqueDashboardName(trackerDBHandle, databaseID, dashboardID, dashboardName); uniqueErr != nil {
		return uniqueErr
	}

	return nil

}

func validateComponentTitle(title string) error {

	if !stringValidation.WellFormedItemName(title) {
		return fmt.Errorf("Invalid title")
	}

	return nil
}
