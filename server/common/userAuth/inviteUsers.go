// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/generic/timestamp"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"log"
)

func sendUserInviteEmail(trackerDBHandle *sql.DB, fromUserID string, inviteeEmailAddr string) error {

	inviteID := uniqueID.GenerateUniqueID()
	inviteTimestamp := timestamp.CurrentTimestampUTC()
	inviteMsg := "" // placeholder for future custom invite message

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO user_invites 
			(invite_id, invite_timestamp_utc, from_user_id,invitee_email_addr,invite_msg) 
			VALUES ($1,$2,$3,$4,$5)`,
		inviteID,
		inviteTimestamp,
		fromUserID,
		inviteeEmailAddr,
		inviteMsg); insertErr != nil {
		return fmt.Errorf("sendUserInviteEmail: database insert failed:  %v", insertErr)
	}

	inviteLink := "https://www.resultra.com/register/" + inviteID

	body := "\nYou been invited to register a Resultra account. \n" +
		"Please visit the following link to register your account:\n\n" +
		inviteLink + "\n\n" +
		"If you've received this email in error, please ignore."

	emailParams := TransactionEmailParams{
		ToAddress: inviteeEmailAddr,
		Subject:   "Resultra Membership Invitation",
		Body:      body}

	emailErr := SendTransactionEmail(emailParams)
	if emailErr != nil {
		return fmt.Errorf("sendPasswordResetEmail: mail send failed:  %v", emailErr)
	}

	return nil

}

type UserInviteParams struct {
	EmailAddrs []string `json:"emailAddrs"`
}

func sendUserInvites(trackerDBHandle *sql.DB, fromUserID string, params UserInviteParams) *AuthResponse {

	for _, currInviteEmailAddr := range params.EmailAddrs {
		if sendInviteErr := sendUserInviteEmail(trackerDBHandle, fromUserID, currInviteEmailAddr); sendInviteErr != nil {
			log.Printf("setUserInvites: failure sending invitation: %v", sendInviteErr)
			return newAuthResponse(false,
				fmt.Sprintf("System failure sending invitation. Please contact support or try again later."))
		}
	}

	return newAuthResponse(true, "Invitations sent.")

}

type UserInviteInfo struct {
	InviteID     string `json:"inviteID"`
	InviteeEmail string `json:"inviteeEmail"`
}

func GetInviteInfo(trackerDBHandle *sql.DB, inviteID string) (*UserInviteInfo, error) {
	inviteInfo := UserInviteInfo{}

	getErr := trackerDBHandle.QueryRow(`SELECT invite_id, invitee_email_addr
		 FROM user_invites
		 WHERE invite_id=$1 LIMIT 1`, inviteID).Scan(
		&inviteInfo.InviteID,
		&inviteInfo.InviteeEmail)
	if getErr != nil {
		return nil, fmt.Errorf("GetInviteInfo: Unabled to get user invitation info info: invite ID = %v: datastore err=%v",
			inviteID, getErr)
	}
	return &inviteInfo, nil

}
