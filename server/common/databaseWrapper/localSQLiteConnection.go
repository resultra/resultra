package databaseWrapper

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type LocalSQLiteTrackerDatabaseConnectionConfig struct {
	DatabaseBasePath string `json:"databaseBasePath"`
	DBHandle         *sql.DB
}

const permsOwnerReadWriteOnly os.FileMode = 0700

func TrackerDatabaseFileExists(databaseFileName string) bool {
	if _, err := os.Stat(databaseFileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (config LocalSQLiteTrackerDatabaseConnectionConfig) validateWellFormedDatabaseBasePath() error {

	if len(config.DatabaseBasePath) == 0 {
		return fmt.Errorf("configuration file missing database path configuration")
	}
	return nil

}

func (config LocalSQLiteTrackerDatabaseConnectionConfig) TrackerDatabaseFileName() string {
	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		panic(fmt.Sprintf("runtime config: tried to database path from invalid config: %v", err))
	}
	return (config.DatabaseBasePath) + `/trackers.db`
}

func (config *LocalSQLiteTrackerDatabaseConnectionConfig) InitConnection() error {

	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		return fmt.Errorf("runtime config: invalid/missing database base path for local database connection: %v", err)
	}

	log.Printf("Initializing local sqlite connection: %v \n", config.DatabaseBasePath)

	err := os.MkdirAll(config.DatabaseBasePath, permsOwnerReadWriteOnly)
	if err != nil {
		return fmt.Errorf("Error initializing tracker directory %v: %v",
			config.DatabaseBasePath, err)
	}

	databaseFileName := config.TrackerDatabaseFileName()

	dbFileAlreadyExists := TrackerDatabaseFileExists(databaseFileName)

	dbHandle, err := sql.Open("sqlite3", databaseFileName)
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
		if initErr := initNewTrackerDatabaseToDest(dbHandle); initErr != nil {
			return fmt.Errorf("failure initializing tracker database: %v", initErr)
		} else {
			log.Printf("New database initialization complete: %v", databaseFileName)
		}
	} else {
		log.Printf("Existing tracker database found.")
	}

	log.Printf("Done initializing local sqlite connection: %v \n", config.DatabaseBasePath)

	config.DBHandle = dbHandle

	return nil

}

func (config LocalSQLiteTrackerDatabaseConnectionConfig) GetAttachmentBasePath(r *http.Request) (string, error) {
	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		panic(fmt.Sprintf("runtime config: tried to retrieve attachment path from invalid config: %v", err))
	}
	return (config.DatabaseBasePath) + `/attachments`, nil
}

func (config LocalSQLiteTrackerDatabaseConnectionConfig) GetTrackerDBHandle(r *http.Request) (*sql.DB, error) {

	if config.DBHandle == nil {
		return nil, fmt.Errorf("LocalSQLiteTrackerDatabaseConnectionConfig: uninitialized database connection")
	}

	return config.DBHandle, nil

}

func (config LocalSQLiteTrackerDatabaseConnectionConfig) SaveAttachment(saveParams SaveAttachmentParams) error {

	attachmentBasePath, err := config.GetAttachmentBasePath(saveParams.HTTPReq)
	if err != nil {
		return fmt.Errorf("LocalSQLiteTrackerDatabaseConnectionConfig: can't get base path: %v", err)
	}

	if err := saveLocalAttachmentFile(attachmentBasePath, saveParams.ParentDatabaseID,
		saveParams.CloudFileName, saveParams.FileData); err != nil {
		return fmt.Errorf("LocalSQLiteTrackerDatabaseConnectionConfig: can't save file: %v", err)
	}

	return nil

}

func (config LocalSQLiteTrackerDatabaseConnectionConfig) ServeAttachment(serveParams ServeAttachmentParams) {

	attachmentBasePath, err := config.GetAttachmentBasePath(serveParams.HTTPReq)
	if err != nil {
		http.Error(serveParams.RespWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	serveLocalFileAttachment(attachmentBasePath, serveParams)

}
