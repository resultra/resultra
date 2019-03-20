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
	"resultra/tracker/server/common/runtimeConfig"
)

type TransactionEmailParams struct {
	ToAddress string
	Subject   string
	Body      string
}

func SendTransactionEmail(params TransactionEmailParams) error {

	config := runtimeConfig.CurrRuntimeConfig.TransactionalEmailConfig

	if config == nil {
		log.Println("WARNING: Transactional email not configured: Email not sent: to = %v, subject = %v, body = %v",
			params.ToAddress, params.Subject, params.Body)
		return nil
	}

	from := config.FromEmailAddr
	pass := config.SMTPPassword
	mailSrv := config.SMTPServerAddress
	sendMailSrv := fmt.Sprintf("%s:%d", config.SMTPServerAddress, *config.SMTPPort)
	sendMailUser := config.SMTPUserName

	// Set up authentication information.
	auth := smtp.PlainAuth("", sendMailUser, pass, mailSrv)

	msg := "From: " + from + "\n" +
		"To: " + params.ToAddress + "\n" +
		"Subject:" + params.Subject + "\n\n" +
		params.Body

	mailErr := smtp.SendMail(sendMailSrv, auth,
		from, []string{params.ToAddress}, []byte(msg))

	if mailErr != nil {
		log.Printf("smtp error: %s", mailErr)
		return fmt.Errorf("SendTransactionEmail: failure sending mail: %v", mailErr)
	}

	return nil

}
