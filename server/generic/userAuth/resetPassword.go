package userAuth

import (
	"database/sql"
	"fmt"
	"log"
	"net/smtp"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

func sendPasswordResetEmail(trackerDBHandle *sql.DB, emailAddr string, userID string) error {

	resetID := uniqueID.GenerateSnowflakeID()
	resetTimestamp := time.Now().UTC()

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO password_reset_links (reset_id, reset_timestamp_utc, user_id) VALUES ($1,$2,$3)`,
		resetID,
		resetTimestamp,
		userID); insertErr != nil {
		return fmt.Errorf("sendPasswordResetEmail: database insert failed:  %v", insertErr)
	}

	from := "admin-email-test@resultra.com"
	pass := "here4test"
	mailSrv := "smtp.gmail.com"

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, pass, mailSrv)

	to := emailAddr

	resetLink := "https://www.resultra.com/resetPassword/" + resetID

	body := "\nA password reset request has been submitted to your Resultra account. \n" +
		"If you submitted this request, please visit the following link to\n" +
		"reset your password:\n\n" +
		resetLink + "\n\n" +
		"If you've received this email in error, please ignore."

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Resultra Password Reset Link\n\n" +
		body

	mailErr := smtp.SendMail("smtp.gmail.com:587", auth,
		from, []string{to}, []byte(msg))

	if mailErr != nil {
		log.Printf("smtp error: %s", mailErr)
		return fmt.Errorf("sendPasswordResetEmail: failure sending mail: %v", mailErr)
	}

	return nil

}

type PasswordResetParams struct {
	EmailAddr string `json:"emailAddr"`
}

func sendResetPasswordLink(trackerDBHandle *sql.DB, params PasswordResetParams) *AuthResponse {

	userInfo, err := GetUserInfoByEmail(trackerDBHandle, params.EmailAddr)
	if err != nil {
		return newAuthResponse(false, fmt.Sprintf("Unable to send password link: %v", err))
	}

	if sendErr := sendPasswordResetEmail(trackerDBHandle, params.EmailAddr, userInfo.UserID); sendErr != nil {
		log.Printf("sendResetPasswordLink: failure sending password reset email: %v", sendErr)
		return newAuthResponse(false,
			fmt.Sprintf("System failure sending password link. Please contact support try again later."))
	}

	log.Printf("Sending password reset link: email = %v, user ID = %v", params.EmailAddr, userInfo.UserID)
	return newAuthResponse(true, "Password reset link sent.")
}
