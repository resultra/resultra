package databaseController

import (
	"net/http"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"
)

func createNewDatabase(req *http.Request, dbParams database.NewDatabaseParams) (*database.Database, error) {

	userID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, userErr
	}

	newDB, newDBErr := database.SaveNewDatabase(dbParams)
	if newDBErr != nil {
		return nil, newDBErr
	}

	if adminErr := userRole.AddDatabaseAdmin(newDB.DatabaseID, userID); adminErr != nil {
		return nil, adminErr
	}

	return newDB, nil
}
