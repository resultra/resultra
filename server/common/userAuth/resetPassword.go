package userAuth

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/generic/timestamp"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

func sendPasswordResetEmail(trackerDBHandle *sql.DB, emailAddr string, userID string) error {

	resetID := uniqueID.GenerateSnowflakeID()
	resetTimestamp := timestamp.CurrentTimestampUTC()

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO password_reset_links (reset_id, reset_timestamp_utc, user_id) VALUES ($1,$2,$3)`,
		resetID,
		resetTimestamp,
		userID); insertErr != nil {
		return fmt.Errorf("sendPasswordResetEmail: database insert failed:  %v", insertErr)
	}

	resetLink := "https://www.resultra.com/resetPassword/" + resetID

	body := "\nA password reset request has been submitted to your Resultra account. \n" +
		"If you submitted this request, please visit the following link to\n" +
		"reset your password:\n\n" +
		resetLink + "\n\n" +
		"If you've received this email in error, please ignore."

	emailParams := TransactionEmailParams{
		ToAddress: emailAddr,
		Subject:   "Resultra Passsword Reset Link",
		Body:      body}

	emailErr := SendTransactionEmail(emailParams)
	if emailErr != nil {
		return fmt.Errorf("sendPasswordResetEmail: mail send failed:  %v", emailErr)
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

type PasswordResetByUserIdParams struct {
	UserID string `json:"userID"`
}

func sendResetPasswordLinkByUserID(trackerDBHandle *sql.DB, params PasswordResetByUserIdParams) *AuthResponse {
	userInfo, err := getAdminUserInfoByID(trackerDBHandle, params.UserID)
	if err != nil {
		log.Printf("sendResetPasswordLinkByUserID: failure sending password reset email: %v", err)
		return newAuthResponse(false, fmt.Sprintf("Unable to send password link: can't retrieve user information"))
	}
	sendParams := PasswordResetParams{EmailAddr: userInfo.EmailAddress}

	return sendResetPasswordLink(trackerDBHandle, sendParams)
}

type PasswordResetInfo struct {
	ResetID   string
	UserID    string
	ResetTime time.Time
}

func GetPasswordResetInfo(trackerDBHandle *sql.DB, resetID string) (*PasswordResetInfo, error) {
	var resetInfo PasswordResetInfo

	getErr := trackerDBHandle.QueryRow(
		`SELECT reset_id, user_id,  reset_timestamp_utc
			FROM password_reset_links 
			WHERE reset_id=$1 LIMIT 1`, resetID).Scan(&resetInfo.ResetID, &resetInfo.UserID, &resetInfo.ResetTime)
	if getErr != nil {
		return nil, fmt.Errorf("Invalid password reset information")
	}

	return &resetInfo, nil

}

func setUserPassword(trackerDBHandle *sql.DB, userID string, password string) error {

	passwordValResp := validatePasswordStrength(password)
	if !passwordValResp.ValidPassword {
		return fmt.Errorf("setUserPassword: system failure setting password: %v", passwordValResp.Msg)
	}

	pwHash, hashErr := generatePasswordHash(password)
	if hashErr != nil {
		return fmt.Errorf("setUserPassword: system failure setting password: %v", hashErr)
	}

	if _, updateErr := trackerDBHandle.Exec(
		`UPDATE users set password_hash=$1 where user_id=$2`, pwHash, userID); updateErr != nil {
		return fmt.Errorf("setUserPassword: set password failed:  %v", updateErr)
	}

	trackerDBHandle.Exec(`DELETE from password_reset_links where user_id=$2`, userID)

	return nil

}

type PasswordResetEntryParams struct {
	ResetID     string `json:"resetID"`
	NewPassword string `json:"newPassword"`
}

func resetPassword(trackerDBHandle *sql.DB, params PasswordResetEntryParams) *AuthResponse {

	resetInfo, infoErr := GetPasswordResetInfo(trackerDBHandle, params.ResetID)
	if infoErr != nil {
		log.Printf("resetPassword: error: %v", infoErr)
		return newAuthResponse(false, fmt.Sprintf("Failure setting password. Invalid reset link."))
	}

	if setErr := setUserPassword(trackerDBHandle, resetInfo.UserID, params.NewPassword); setErr != nil {
		log.Printf("resetPassword: error: %v", setErr)
		return newAuthResponse(false, fmt.Sprintf("Failure setting password. system error."))
	}

	return newAuthResponse(true, fmt.Sprintf("Password changed."))

}
