// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

import (
	"fmt"
	"log"
	"net/smtp"
)

type TransactionEmailParams struct {
	ToAddress string
	Subject   string
	Body      string
}

func SendTransactionEmail(params TransactionEmailParams) error {
	from := "admin-email-test@resultra.com"
	pass := "here4test"
	mailSrv := "smtp.gmail.com"

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, pass, mailSrv)

	msg := "From: " + from + "\n" +
		"To: " + params.ToAddress + "\n" +
		"Subject:" + params.Subject + "\n\n" +
		params.Body

	mailErr := smtp.SendMail("smtp.gmail.com:587", auth,
		from, []string{params.ToAddress}, []byte(msg))

	if mailErr != nil {
		log.Printf("smtp error: %s", mailErr)
		return fmt.Errorf("SendTransactionEmail: failure sending mail: %v", mailErr)
	}

	return nil

}
