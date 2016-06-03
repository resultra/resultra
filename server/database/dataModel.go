package database

import (
	"appengine"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const DatabaseEntityKind string = "Database"

type Database struct {
	DatabaseID string `json:"databaseID"`
	Name       string `json:"name"`
}

type NewDatabaseParams struct {
	Name string `json:"name"`
}

func saveNewDatabase(appEngContext appengine.Context, params NewDatabaseParams) (*Database, error) {

	sanitizedDbName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	// TODO - Validate name is unique

	newDatabase := Database{DatabaseID: uniqueID.GenerateUniqueID(), Name: sanitizedDbName}

	insertErr := datastoreWrapper.InsertNewRootEntity(
		appEngContext, DatabaseEntityKind, &newDatabase)
	if insertErr != nil {
		return nil, insertErr
	}

	return &newDatabase, nil
}
