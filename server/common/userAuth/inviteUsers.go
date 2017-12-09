package userAuth

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

func sendUserInviteEmail(trackerDBHandle *sql.DB, fromUserID string, inviteeEmailAddr string) error {

	inviteID := uniqueID.GenerateSnowflakeID()
	inviteTimestamp := time.Now().UTC()
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
