package databaseWrapper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"resultra/datasheet/server/common/runtimeConfig"
)

var dbHandle *sql.DB

func InitDatabaseConnection() error {

	var err error

	databaseFileName := runtimeConfig.CurrRuntimeConfig.TrackerDatabaseFileName()
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

	return nil
}

func DBHandle() *sql.DB {
	return dbHandle
}
