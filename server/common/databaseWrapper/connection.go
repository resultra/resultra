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
