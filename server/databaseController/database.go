// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package databaseController

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
	"github.com/resultra/resultra/server/userRole"
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
		newDB, newDBErr = CloneIntoNewTrackerDatabase(&newDBFromTemplateParams)
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
		IsActive:     true,
		IsAdmin:      true}

	return &newDBInfo, nil
}
