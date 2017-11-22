package databaseController

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/trackerDatabase"
	"resultra/datasheet/server/userRole"
)

func createNewDatabase(trackerDBHandle *sql.DB, req *http.Request, dbParams trackerDatabase.NewDatabaseParams) (*UserTrackingDatabaseInfo, error) {

	userID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, userErr
	}

	var newDB *trackerDatabase.Database
	var newDBErr error

	if (dbParams.TemplateDatabaseID != nil) && (len(*dbParams.TemplateDatabaseID) > 0) {

		if dbParams.TemplateSource == nil {
			return nil, fmt.Errorf("createNewDatabase: missing templates source for new database")
		}
		templSource := *dbParams.TemplateSource
		var srcTrackerDBHandle *sql.DB
		if templSource == trackerDatabase.NewDatabaseTemplateSourceFactory {
			factoryDBHandle, err := databaseWrapper.GetFactoryTemplateTrackerDatabaseHandle(req)
			if err != nil {
				return nil, fmt.Errorf("createNewDatabase: can't get factory template database handle: %v", err)
			}
			srcTrackerDBHandle = factoryDBHandle
		} else if templSource == trackerDatabase.NewDatabaseTemplateSourceAccount {
			srcTrackerDBHandle = trackerDBHandle
		} else {
			return nil, fmt.Errorf("createNewDatabase: can't get factory template database handle: unrecognized template source %v", templSource)
		}

		newDBFromTemplateParams := trackerDatabase.CloneDatabaseParams{
			SourceDatabaseID: *dbParams.TemplateDatabaseID,
			NewName:          dbParams.Name,
			IsTemplate:       false,
			CreatedByUserID:  userID,
			SrcDBHandle:      srcTrackerDBHandle,
			DestDBHandle:     trackerDBHandle,
			IDRemapper:       uniqueID.UniqueIDRemapper{}}
		newDB, newDBErr = cloneIntoNewTrackerDatabase(&newDBFromTemplateParams)
		if newDBErr != nil {
			return nil, newDBErr
		}
	} else {
		dbParams.CreatedByUserID = userID
		dbParams.IsTemplate = false
		newDB, newDBErr = trackerDatabase.SaveNewEmptyDatabase(trackerDBHandle, dbParams)
		if newDBErr != nil {
			return nil, newDBErr
		}
	}

	if adminErr := userRole.AddDatabaseAdmin(trackerDBHandle, newDB.DatabaseID, userID); adminErr != nil {
		return nil, adminErr
	}

	// Return a structure which includes not only the name and ID, but also the information
	// about the current user's permissions. Rather than returning the raw database struct,
	// this UserTrackingDatabaseInfo struct allows the tracking database to be displayed in a
	// list with options to change settings, etc.
	newDBInfo := UserTrackingDatabaseInfo{
		DatabaseID:   newDB.DatabaseID,
		DatabaseName: newDB.Name,
		IsAdmin:      true}

	return &newDBInfo, nil
}
