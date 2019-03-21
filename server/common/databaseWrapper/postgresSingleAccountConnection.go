// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package databaseWrapper

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type PostgresSingleAccountDatabaseConfig struct {
	TrackerDBHostName string `json:"databaseHostName"`
	TrackerUserName   string `json:"databaseUserName"`
	TrackerPassword   string `json:"databasePassword"`
	TrackerDBName     string `json:"trackerDatabaseName"`
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
			"PostgresSingleAccountDatabaseConfig.connectToAccountTrackerDatabase: can't establish connection to account info database (ping failed): %v", err)
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
