package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"resultra/datasheet/server/common/databaseWrapper"
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

	fmt.Printf("\n")
	fmt.Printf("Account details:\n")
	fmt.Printf("Email: [%v]\n", owner_email)
	fmt.Printf("First: [%v]\n", owner_first)
	fmt.Printf("Last: [%v]\n", owner_last)

	fmt.Printf("Subdomain: [%v]\n", account_subdomain)

	fmt.Printf("\n")

	confirm := promptUserInputString("Ready to create account (enter y or n): ")
	if confirm == "y" {
		fmt.Printf("Creating account ...")

		if err := databaseWrapper.InitAccountInfoConnection(); err != nil {
			log.Fatal(err)
		}

		hostName := account_subdomain + ".resultra.com"
		newAcctInfo := databaseWrapper.NewAccountInfo{
			HostName:  hostName,
			FirstName: owner_first,
			LastName:  owner_last,
			Email:     owner_email}

		if err := databaseWrapper.CreateNewAccount(newAcctInfo); err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" done.\n")

	} else {
		fmt.Println("Account creation aborted.")
	}

}
