package databaseWrapper

import (
	"database/sql"
	"fmt"
	"github.com/hashicorp/golang-lru"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type PostgresMultipleAccountDatabaseConfig struct {
	TrackerUserName         string `json:"trackerUserName"`
	TrackerPassword         string `json:"trackerPassword"`
	AccountHostName         string `json:"accountHostName"`
	AccountUserName         string `json:"accountUserName"`
	AccountPassword         string `json:"accountPassword"`
	AccountConnectionCache  *lru.Cache
	AccountInfoDBConnection *sql.DB
}

const maxOpenAccountConnections int = 20
const maxOpenConnectionPerAccount int = 4

func (config *PostgresMultipleAccountDatabaseConfig) connectToAccountTrackerDatabase(accountDBHostName string, accountDBName string) (*sql.DB, error) {

	connectionStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		accountDBHostName, config.TrackerUserName, accountDBName, config.TrackerPassword)

	trackerDBHandle, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("PostgresMultipleAccountDatabaseConfig.connectToAccountTrackerDatabase: can't establish connection to tracker database: %v", err)
	}

	// Only a few open connections are needed to the account database.
	trackerDBHandle.SetMaxOpenConns(maxOpenConnectionPerAccount)

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if err := trackerDBHandle.Ping(); err != nil {
		return nil, fmt.Errorf(
			"PostgresMultipleAccountDatabaseConfig.connectToAccountTrackerDatabase:: can't establish connection to account info database (ping failed): %v", err)
	}

	return trackerDBHandle, nil

}

type AccountConnectionInfo struct {
	trackerDBConnection *sql.DB
	accountInfo         *AccountTrackerDBInfo
}

func (config *PostgresMultipleAccountDatabaseConfig) InitConnection() error {

	closeConnectionWhenConnectionEvictedCallback := func(accountName interface{}, connInfo interface{}) {
		acctConnInfo := connInfo.(*AccountConnectionInfo)
		closeErr := acctConnInfo.trackerDBConnection.Close()
		if closeErr != nil {
			accountNameStr := accountName.(string)
			log.Printf("Failure closing tracker database connection for account=%v", accountNameStr)
		}
	}

	connectionCache, err := lru.NewWithEvict(maxOpenAccountConnections, closeConnectionWhenConnectionEvictedCallback)
	if err != nil {
		return fmt.Errorf("PostgresMultipleAccountDatabaseConfig: can't create connection cache: %v", err)
	}
	config.AccountConnectionCache = connectionCache

	accountDB, err := connectToAccountInfoDB(config.AccountHostName, config.AccountUserName, config.AccountPassword)
	if err != nil {
		return fmt.Errorf("PostgresMultipleAccountDatabaseConfig: failure connecting to accounts database: %v", err)
	}
	config.AccountInfoDBConnection = accountDB

	return nil
}

func (config PostgresMultipleAccountDatabaseConfig) getTrackerDB(r *http.Request) (*sql.DB, error) {
	return nil, fmt.Errorf("not implemented")
}

func (config PostgresMultipleAccountDatabaseConfig) GetTrackerDBHandle(req *http.Request) (*sql.DB, error) {

	accountHostName := AccountHostNameFromReq(req)

	cachedConnection, cachedConnectionFound := config.AccountConnectionCache.Get(accountHostName)
	if cachedConnectionFound {
		connectionInfo := cachedConnection.(*AccountConnectionInfo)
		return connectionInfo.trackerDBConnection, nil
	} else {

		accountTrackerDBInfo, err := getHostAccountTrackerDBInfo(config.AccountInfoDBConnection, accountHostName)
		if err != nil {
			return nil, fmt.Errorf("PostgresMultipleAccountDatabaseConfig: unable to get tracker database info for account host = %v", accountHostName)
		}

		dbConnection, err := config.connectToAccountTrackerDatabase(
			accountTrackerDBInfo.DBHostName, accountTrackerDBInfo.DBName)
		if err != nil {
			return nil, fmt.Errorf("PostgresMultipleAccountDatabaseConfig: can't connect to tracker database info for account host = %v: error = %v",
				accountHostName, err)

		}

		acctConnectionInfo := AccountConnectionInfo{
			trackerDBConnection: dbConnection,
			accountInfo:         accountTrackerDBInfo}

		config.AccountConnectionCache.Add(accountHostName, &acctConnectionInfo)

		return dbConnection, nil
	}

}
