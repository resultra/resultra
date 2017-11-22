package databaseWrapper

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Interface for handling connections to different types of databases
type TrackerDatabaseConnection interface {
	InitConnection() error
	GetTrackerDBHandle(r *http.Request) (*sql.DB, error)
}

var dbConnection TrackerDatabaseConnection
var factoryTemplateDBConnection TrackerDatabaseConnection

func GetTrackerDatabaseHandle(r *http.Request) (*sql.DB, error) {

	if dbConnection == nil {
		return nil, fmt.Errorf("GetTrackerDatabaseHandle: uninitialized database connection")
	}
	return dbConnection.GetTrackerDBHandle(r)

}

func InitConnectionConfiguration(conn TrackerDatabaseConnection) error {

	if err := conn.InitConnection(); err != nil {
		return fmt.Errorf("InitConnectionConfiguration: %v", err)
	}

	dbConnection = conn

	return nil

}

func InitFactoryTemplateConnectionConfiguration(conn TrackerDatabaseConnection) error {

	if err := conn.InitConnection(); err != nil {
		return fmt.Errorf("InitFactoryTemplateConnectionConfiguration: %v", err)
	}

	factoryTemplateDBConnection = conn

	return nil

}

func FactoryTemplateDatabaseIsConfigured() bool {
	if factoryTemplateDBConnection != nil {
		return true
	}
	return false
}

func GetFactoryTemplateTrackerDatabaseHandle(r *http.Request) (*sql.DB, error) {

	if !FactoryTemplateDatabaseIsConfigured() {
		return nil, fmt.Errorf("GetFactoryTemplateTrackerDatabaseHandle: uninitialized database connection")
	}
	return factoryTemplateDBConnection.GetTrackerDBHandle(r)

}
