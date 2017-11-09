package databaseWrapper

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"resultra/datasheet/server/common/runtimeConfig"
)

var dbHandle *sql.DB

func TrackerDatabaseFileExists(databaseFileName string) bool {
	if _, err := os.Stat(databaseFileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func InitDatabaseConnection() error {

	var err error

	databaseFileName := runtimeConfig.CurrRuntimeConfig.TrackerDatabaseFileName()

	dbFileAlreadyExists := TrackerDatabaseFileExists(databaseFileName)

	dbHandle, err = sql.Open("sqlite3", databaseFileName)
	if err != nil {
		return fmt.Errorf("can't establish connection to database: %v", err)
	}

	// Configure the maximum number of open connections.
	dbHandle.SetMaxOpenConns(75)

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if err := dbHandle.Ping(); err != nil {
		return fmt.Errorf("can't establish connection to database: %v", err)
	}

	log.Printf("Database connection established: %v", databaseFileName)

	if !dbFileAlreadyExists {
		log.Printf("New database found, initializing: %v", databaseFileName)
		if initErr := initNewTrackerDatabase(); initErr != nil {
			return fmt.Errorf("failure initializing tracker database: %v", initErr)
		} else {
			log.Printf("New database initialization complete: %v", databaseFileName)
		}
	} else {
		log.Printf("Existing tracker database found.")
	}

	return nil
}

func DBHandle() *sql.DB {
	return dbHandle
}

func GetTrackerDatabaseHandle(r *http.Request) (*sql.DB, error) {
	return DBHandle(), nil
}
