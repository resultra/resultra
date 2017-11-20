package databaseWrapper

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type PostgresSingleAccountDatabaseConfig struct {
	TrackerDBHostName string `json:"trackerDBHostName"`
	TrackerUserName   string `json:"trackerUserName"`
	TrackerPassword   string `json:"trackerPassword"`
	TrackerDBName     string `json:"trackerDBName"`
	TrackerDBHandle   *sql.DB
}

const maxSingleAccountConnections int = 10

func (config *PostgresSingleAccountDatabaseConfig) connectToTrackerDatabase() (*sql.DB, error) {

	connectionStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.TrackerDBHostName, config.TrackerUserName, config.TrackerDBName, config.TrackerPassword)

	trackerDBHandle, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("PostgresSingleAccountDatabaseConfig.connectToTrackerDatabase: can't establish connection to tracker database: %v", err)
	}

	// Only a few open connections are needed to the account database.
	trackerDBHandle.SetMaxOpenConns(maxSingleAccountConnections)

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if err := trackerDBHandle.Ping(); err != nil {
		return nil, fmt.Errorf(
			"PostgresMultipleAccountDatabaseConfig.connectToAccountTrackerDatabase:: can't establish connection to account info database (ping failed): %v", err)
	}

	return trackerDBHandle, nil

}

func (config *PostgresSingleAccountDatabaseConfig) InitConnection() error {

	trackerDB, err := config.connectToTrackerDatabase()
	if err != nil {
		return fmt.Errorf("PostgresSingleAccountDatabaseConfig: failure connecting to tracker database: %v", err)
	}
	log.Printf("PostgresSingleAccountDatabaseConfig.InitConnection: connected to tracker database: host=%v, user=%v\n",
		config.TrackerDBHostName, config.TrackerUserName)
	config.TrackerDBHandle = trackerDB

	return nil
}

func (config PostgresSingleAccountDatabaseConfig) GetTrackerDBHandle(req *http.Request) (*sql.DB, error) {

	if config.TrackerDBHandle == nil {
		return nil, fmt.Errorf("PostgresSingleAccountDatabaseConfig: uninitialized database connection")
	}

	return config.TrackerDBHandle, nil
}
