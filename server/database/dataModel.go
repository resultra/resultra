package database

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
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

func saveNewDatabase(params NewDatabaseParams) (*Database, error) {

	sanitizedDbName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	databaseID := uniqueID.GenerateSnowflakeID()

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO databases VALUES ($1,$2)`,
		databaseID, sanitizedDbName); insertErr != nil {
		return nil, fmt.Errorf("saveNewDatabase: insert failed: error = %v", insertErr)
	}

	newDatabase := Database{DatabaseID: databaseID, Name: sanitizedDbName}

	return &newDatabase, nil
}
