package database

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
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

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("saveNewDatabase: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	databaseID := gocql.TimeUUID().String()
	if insertErr := dbSession.Query(`INSERT INTO database (databaseID, name) VALUES (?, ?)`,
		databaseID, sanitizedDbName).Exec(); insertErr != nil {
		fmt.Errorf("saveNewDatabase: Can't create database: unable to create database: error = %v", insertErr)
	}

	newDatabase := Database{DatabaseID: databaseID, Name: sanitizedDbName}

	return &newDatabase, nil
}
