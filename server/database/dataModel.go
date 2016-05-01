package database

import (
	"appengine"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const databaseEntityKind string = "Database"

type Database struct {
	Name string
}

type NewDatabaseParams struct {
	Name string `json:"name"`
}

type DatabaseRef struct {
	DatabaseID string `json:"databaseID"`
	Name       string `json:"name"`
}

func saveNewDatabase(appEngContext appengine.Context, params NewDatabaseParams) (*DatabaseRef, error) {

	sanitizedDbName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	// TODO - Validate name is unique

	newDatabase := Database{Name: sanitizedDbName}

	databaseID, insertErr := datastoreWrapper.InsertNewRootEntity(
		appEngContext, databaseEntityKind, &newDatabase)
	if insertErr != nil {
		return nil, insertErr
	}

	dbRef := DatabaseRef{
		DatabaseID: databaseID,
		Name:       sanitizedDbName}

	return &dbRef, nil
}
