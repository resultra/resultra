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
	"github.com/resultra/resultra/server/generic/uniqueID"
	"net/http"
	"strings"
)

var accountInfoDBHandle *sql.DB

const trackerAccountDBPrefix string = "trackers_"

func connectToAccountInfoDB(hostName string, userName string, password string) (*sql.DB, error) {

	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		hostName, userName, password, "resultra_accounts")

	accountDBHandle, err := sql.Open("postgres", connectStr)
	if err != nil {
		return nil, fmt.Errorf("connectToAccountInfoDB: can't establish connection to account info database: %v", err)
	}

	// Only a few open connections are needed to the account database.
	accountDBHandle.SetMaxOpenConns(5)

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if err := accountDBHandle.Ping(); err != nil {
		return nil, fmt.Errorf("connectToAccountInfoDB: can't establish connection to account info database (ping failed): %v", err)
	}

	return accountDBHandle, nil

}

func InitAccountInfoConnection() error {

	var err error

	accountInfoDBHandle, err = connectToAccountInfoDB("localhost", "devuser", "here4dev")
	if err != nil {
		return fmt.Errorf("can't establish connection to account info database: %v", err)
	}

	return nil

}

type AccountInfo struct {
	AccountID  string
	FirstName  string
	LastName   string
	Email      string
	DBHostName string
}

func AddAccount(newAcctInfo AccountInfo) error {

	dbHandle := AccountInfoDBHandle()

	if dbHandle == nil {
		return fmt.Errorf("AddAccount: can't add account - connection to account info database not established")
	}

	if _, insertErr := dbHandle.Exec(`INSERT INTO account_info (account_id,owner_first,owner_last,owner_email,db_host_name) 
			VALUES ($1,$2,$3,$4,$5)`,
		newAcctInfo.AccountID,
		newAcctInfo.FirstName,
		newAcctInfo.LastName,
		newAcctInfo.Email,
		newAcctInfo.DBHostName); insertErr != nil {
		return fmt.Errorf("AddAccount: can't add account: error = %v", insertErr)
	}
	return nil

}

type AccountTrackerDBInfo struct {
	DBHostName string
	DBName     string
}

func trackerDBNameFromAccountID(accountID string) string {
	trackerDBName := trackerAccountDBPrefix + accountID
	return trackerDBName
}

func getHostAccountTrackerDBInfo(accountDB *sql.DB, accountHostName string) (*AccountTrackerDBInfo, error) {

	accountID := ""
	dbHostName := ""
	getErr := accountDB.QueryRow(`SELECT account_info.account_id, account_info.db_host_name
		FROM account_info, host_mappings
		WHERE host_mappings.host_name=$1 AND host_mappings.account_id=account_info.account_id LIMIT 1`, accountHostName).Scan(&accountID, &dbHostName)
	if getErr != nil {
		return nil, fmt.Errorf("GetHostAccountTrackerDBInfo: Unabled to get account information for account host name = %s: error = %v", accountHostName, getErr)
	}

	dbInfo := AccountTrackerDBInfo{
		DBHostName: dbHostName,
		DBName:     trackerDBNameFromAccountID(accountID)}

	return &dbInfo, nil

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

	trackerDBName := trackerDBNameFromAccountID(accountID)
	if _, createtErr := dbHandle.Exec(`CREATE DATABASE ` + trackerDBName); createtErr != nil {
		return fmt.Errorf("CreateAccountTrackerDatabase: can't create account's tracker database: error = %v", createtErr)
	}

	trackerDBHandle, err := sql.Open("postgres",
		"user=devuser dbname="+trackerDBName+" password=here4dev sslmode=disable")
	if err != nil {
		return fmt.Errorf("CreateAccountTrackerDatabase: can't establish connection to newly created tracker database: %v", err)
	}

	if initErr := InitNewTrackerDatabaseToDest(trackerDBHandle); initErr != nil {
		return fmt.Errorf("CreateAccountTrackerDatabase: can't initialize newly created tracker database: %v", initErr)
	}

	return nil
}

type NewAccountInfo struct {
	HostName   string
	FirstName  string
	LastName   string
	Email      string
	DBHostName string
}

func CreateNewAccount(newAccountInfo NewAccountInfo) error {

	accountID := uniqueID.GenerateV4UUIDNoDashes()

	accountInfo := AccountInfo{
		AccountID:  accountID,
		FirstName:  newAccountInfo.FirstName,
		LastName:   newAccountInfo.LastName,
		Email:      newAccountInfo.Email,
		DBHostName: newAccountInfo.DBHostName}
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
