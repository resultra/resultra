// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"strings"
)

func init() {

}

func promptUserInputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt)
	user_input, _ := reader.ReadString('\n')
	user_input = strings.TrimRight(user_input, "\n")
	return user_input
}

func main() {

	owner_first := promptUserInputString("Enter account owner first name: ")
	owner_last := promptUserInputString("Enter account owner last name: ")
	owner_email := promptUserInputString("Enter account owner email address: ")
	account_subdomain := promptUserInputString("Subdomain name: ")
	database_hostname := promptUserInputString("Tracker database host name: ")

	fmt.Printf("\n")
	fmt.Printf("Account details:\n")
	fmt.Printf("Email: [%v]\n", owner_email)
	fmt.Printf("First: [%v]\n", owner_first)
	fmt.Printf("Last: [%v]\n", owner_last)

	fmt.Printf("Subdomain: [%v]\n", account_subdomain)
	fmt.Printf("Database hostname: [%v]\n", database_hostname)

	fmt.Printf("\n")

	confirm := promptUserInputString("Ready to create account (enter y or n): ")
	if confirm == "y" {
		fmt.Printf("Creating account ...")

		if err := databaseWrapper.InitAccountInfoConnection(); err != nil {
			log.Fatal(err)
		}

		hostName := account_subdomain + ".resultra.com"
		newAcctInfo := databaseWrapper.NewAccountInfo{
			HostName:   hostName,
			FirstName:  owner_first,
			LastName:   owner_last,
			Email:      owner_email,
			DBHostName: database_hostname}

		if err := databaseWrapper.CreateNewAccount(newAcctInfo); err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" done.\n")

	} else {
		fmt.Println("Account creation aborted.")
	}

}
