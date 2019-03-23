// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"resultra/tracker/server/common/databaseWrapper"
	"strings"
)

func promptUserInputString(prompt string, defaultInput string) string {

	readUserInput := func() string {

		if len(defaultInput) > 0 {
			fmt.Printf("%v [%v]:", prompt, defaultInput)
		} else {
			fmt.Printf("%v:", prompt)
		}

		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimRight(userInput, "\n")

		if len(userInput) == 0 && len(defaultInput) > 0 {
			return defaultInput
		} else if len(userInput) > 0 {
			return userInput
		} else {
			return ""
		}
	}

	userInput := ""
	for userInput = readUserInput(); userInput == ""; userInput = readUserInput() {
	}

	return userInput
}

func main() {

	databaseHost := promptUserInputString("Database host name", "localhost")
	databaseUserName := promptUserInputString("Database user name", "postgres")
	databasePassword := promptUserInputString("Enter database password", "")

	fmt.Printf("\n")
	fmt.Printf("Database configuration details:\n")
	fmt.Printf("Database host: %v\n", databaseHost)
	fmt.Printf("Database user name: %v\n", databaseUserName)
	fmt.Printf("Database password: %v\n", databasePassword)
	fmt.Printf("\n")

	createDatabaseAndWebUser := func() error {
		connectStr := fmt.Sprintf("host=%s user=%s password=%s sslmode=disable",
			databaseHost, databaseUserName, databasePassword)

		dbHandle, openErr := sql.Open("postgres", connectStr)
		if openErr != nil {
			return fmt.Errorf("can't establish connection to database: %v", openErr)
		}
		defer dbHandle.Close()

		_, createDBErr := dbHandle.Exec(`CREATE DATABASE resultra`)
		if createDBErr != nil {
			return fmt.Errorf("can't create main resultra database: %v", createDBErr)
		}

		// TODO - don't hard-code the password
		_, createUserErr := dbHandle.Exec(`CREATE USER resultra_web_user WITH NOSUPERUSER NOCREATEDB NOCREATEROLE PASSWORD 'resultrawebpw'`)
		if createUserErr != nil {
			return fmt.Errorf("can't create main resultra database: %v", createUserErr)
		}
		return nil
	}

	connectToResultraDB := func() (*sql.DB, error) {
		connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			databaseHost, databaseUserName, databasePassword, "resultra")

		dbHandle, openErr := sql.Open("postgres", connectStr)
		if openErr != nil {
			return nil, fmt.Errorf("can't establish connection to database: %v", openErr)
		}
		return dbHandle, nil
	}

	initDatabaseTables := func(dbHandle *sql.DB) error {

		initErr := databaseWrapper.InitNewTrackerDatabaseToDest(dbHandle)
		if initErr != nil {
			return fmt.Errorf("can't initialize tracker database: %v", initErr)
		}

		return nil
	}

	grantWebUserPerms := func(dbHandle *sql.DB) error {

		permsSchema := `REVOKE ALL ON DATABASE resultra FROM public;` +
			`GRANT CONNECT ON DATABASE resultra TO resultra_web_user;` +
			`REVOKE ALL ON SCHEMA public FROM public;` +
			`GRANT USAGE ON SCHEMA public TO resultra_web_user;` +
			`REVOKE ALL ON ALL TABLES IN SCHEMA public FROM PUBLIC;` +
			`GRANT SELECT,UPDATE,INSERT,DELETE ON ALL TABLES in SCHEMA public to resultra_web_user;` +
			`ALTER DEFAULT PRIVILEGES FOR ROLE resultra_web_user IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO resultra_web_user;`

		if _, permsErr := dbHandle.Exec(permsSchema); permsErr != nil {
			return fmt.Errorf("can't initialize tracker database permissions: %v", permsErr)
		}

		return nil
	}

	confirm := promptUserInputString("Ready to create account (enter y or n)", "")
	if confirm == "y" || confirm == "Y" {

		fmt.Printf("Setting up database ...")

		createErr := createDatabaseAndWebUser()
		if createErr != nil {
			log.Fatalf("Error setting up tracker database: %v", createErr)
		}

		resultraDBHandle, dbErr := connectToResultraDB()
		if dbErr != nil {
			log.Fatalf("Error setting up tracker database: %v", createErr)
		}
		defer resultraDBHandle.Close()

		initErr := initDatabaseTables(resultraDBHandle)
		if createErr != nil {
			log.Fatalf("Error initializing tracker database tables: %v", initErr)
		}

		grantErr := grantWebUserPerms(resultraDBHandle)
		if grantErr != nil {
			log.Fatalf("Error initializing tracker database tables: %v", grantErr)
		}

		fmt.Printf(" done.\n")

	} else {
		fmt.Println("Database setup aborted.")
	}

}
