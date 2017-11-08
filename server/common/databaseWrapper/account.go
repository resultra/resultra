package databaseWrapper

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"resultra/datasheet/server/generic/uniqueID"
	"strings"
)

var accountInfoDBHandle *sql.DB

func InitAccountInfoConnection() error {

	var err error

	accountInfoDBHandle, err = sql.Open("postgres",
		"user=devuser dbname=resultra_accounts password=here4dev sslmode=disable")
	if err != nil {
		return fmt.Errorf("can't establish connection to account info database: %v", err)
	}

	// Configure the maximum number of open connections to be less than the limit supported by Postgres
	// If postgress supports 100, then 75 allows for a value which is safely below the maximum and also
	// allow connections from other clients (e.g., for administration)
	// Configure the maximum number of open connections.
	accountInfoDBHandle.SetMaxOpenConns(75)

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if err := accountInfoDBHandle.Ping(); err != nil {
		return fmt.Errorf("can't establish connection to account info database: %v", err)
	}

	return nil

}

type AccountInfo struct {
	AccountID string
	FirstName string
	LastName  string
	Email     string
}

func AddAccount(newAcctInfo AccountInfo) error {

	dbHandle := AccountInfoDBHandle()

	if dbHandle == nil {
		return fmt.Errorf("AddAccount: can't add account - connection to account info database not established")
	}

	if _, insertErr := dbHandle.Exec(`INSERT INTO account_info (account_id,owner_first,owner_last,owner_email) 
			VALUES ($1,$2,$3,$4)`,
		newAcctInfo.AccountID,
		newAcctInfo.FirstName,
		newAcctInfo.LastName,
		newAcctInfo.Email); insertErr != nil {
		return fmt.Errorf("AddAccount: can't add account: error = %v", insertErr)
	}
	return nil

}

func AddAccountHostMapping(hostName string, accountID string) error {
	dbHandle := AccountInfoDBHandle()

	if dbHandle == nil {
		return fmt.Errorf("AddAccountHostMapping: can't add host mapping - connection to account info database not established")
	}
	if _, insertErr := dbHandle.Exec(`INSERT INTO host_mappings (account_id,host_name) VALUES ($1,$2)`,
		accountID, hostName); insertErr != nil {
		return fmt.Errorf("AddAccountHostMapping: can't add host mapping: error = %v", insertErr)
	}
	return nil

}

func CreateAccountTrackerDatabase(accountID string) error {

	dbHandle := AccountInfoDBHandle()

	if dbHandle == nil {
		return fmt.Errorf("CreateAccountTrackerDatabase: can't add host mapping - connection to account info database not established")
	}

	trackerDBName := "trackers_" + accountID
	if _, createtErr := dbHandle.Exec(`CREATE DATABASE ` + trackerDBName); createtErr != nil {
		return fmt.Errorf("CreateAccountTrackerDatabase: can't create account's tracker database: error = %v", createtErr)
	}

	trackerDBHandle, err := sql.Open("postgres",
		"user=devuser dbname="+trackerDBName+" password=here4dev sslmode=disable")
	if err != nil {
		return fmt.Errorf("CreateAccountTrackerDatabase: can't establish connection to newly created tracker database: %v", err)
	}

	if initErr := initNewTrackerDatabaseToDest(trackerDBHandle); initErr != nil {
		return fmt.Errorf("CreateAccountTrackerDatabase: can't initialize newly created tracker database: %v", initErr)
	}

	return nil
}

type NewAccountInfo struct {
	HostName  string
	FirstName string
	LastName  string
	Email     string
}

func CreateNewAccount(newAccountInfo NewAccountInfo) error {

	accountID := uniqueID.GenerateV4UUIDNoDashes()

	accountInfo := AccountInfo{
		AccountID: accountID,
		FirstName: newAccountInfo.FirstName,
		LastName:  newAccountInfo.LastName,
		Email:     newAccountInfo.Email}
	if err := AddAccount(accountInfo); err != nil {
		return fmt.Errorf("CreateNewAccount: %v", err)
	}

	if err := AddAccountHostMapping(newAccountInfo.HostName, accountID); err != nil {
		return fmt.Errorf("CreateNewAccount: %v", err)
	}

	if err := CreateAccountTrackerDatabase(accountID); err != nil {
		return fmt.Errorf("CreateNewAccount: %v", err)
	}

	return nil
}

// Return the hostname without any trailing information for the port number
// the service may be running on.
func AccountHostNameFromReq(req *http.Request) string {

	fullyQualifiedHostName := req.Host

	hostParts := strings.Split(fullyQualifiedHostName, ":")
	return hostParts[0]

}

func AccountInfoDBHandle() *sql.DB {
	return accountInfoDBHandle
}
