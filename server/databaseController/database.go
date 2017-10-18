package databaseController

import (
	"net/http"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/trackerDatabase"
	"resultra/datasheet/server/userRole"
)

func createNewDatabase(req *http.Request, dbParams trackerDatabase.NewDatabaseParams) (*UserTrackingDatabaseInfo, error) {

	userID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, userErr
	}

	var newDB *trackerDatabase.Database
	var newDBErr error

	if (dbParams.TemplateDatabaseID != nil) && (len(*dbParams.TemplateDatabaseID) > 0) {

		newDBFromTemplateParams := trackerDatabase.CloneDatabaseParams{
			SourceDatabaseID: *dbParams.TemplateDatabaseID,
			NewName:          dbParams.Name,
			IsTemplate:       false,
			CreatedByUserID:  userID}
		newDB, newDBErr = cloneIntoNewTrackerDatabase(newDBFromTemplateParams)
		if newDBErr != nil {
			return nil, newDBErr
		}
	} else {
		dbParams.CreatedByUserID = userID
		dbParams.IsTemplate = false
		newDB, newDBErr = trackerDatabase.SaveNewEmptyDatabase(dbParams)
		if newDBErr != nil {
			return nil, newDBErr
		}
	}

	if adminErr := userRole.AddDatabaseAdmin(newDB.DatabaseID, userID); adminErr != nil {
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
