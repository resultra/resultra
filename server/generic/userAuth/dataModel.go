package userAuth

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
	"strings"
)

type User struct {
	UserID       string
	UserName     string
	FirstName    string
	LastName     string
	EmailAddr    string
	PasswordHash string
}

type UserInfo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
}

type NewUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	EmailAddr string `json:"emailAddr"`
	Password  string `json:"password"`
}

// Strip leading and trailing whitespace from most registration parameters.
func (rawParams NewUserParams) sanitize() NewUserParams {
	sanitizedParams := NewUserParams{
		FirstName: strings.TrimSpace(rawParams.FirstName),
		LastName:  strings.TrimSpace(rawParams.LastName),
		UserName:  strings.TrimSpace(rawParams.UserName),
		EmailAddr: strings.TrimSpace(rawParams.EmailAddr),
		Password:  rawParams.Password}
	return sanitizedParams
}

func saveNewUser(rawParams NewUserParams) *AuthResponse {

	params := rawParams.sanitize()

	firstNameResp := validateWellFormedRealName(params.FirstName)
	if !firstNameResp.Success {
		return newAuthResponse(false, "First name is required")
	}

	lastNameResp := validateWellFormedRealName(params.LastName)
	if !lastNameResp.Success {
		return newAuthResponse(false, "Last name is required")
	}

	userNameResp := validateNewUserName(params.UserName)
	if !userNameResp.Success {
		return userNameResp
	}

	sanitizedEmail := strings.TrimSpace(params.EmailAddr)
	emailResp := validateWellFormedEmailAddr(sanitizedEmail)
	if !emailResp.Success {
		return emailResp
	}

	passwordValResp := validatePasswordStrength(params.Password)
	if !passwordValResp.ValidPassword {
		return newAuthResponse(false, passwordValResp.Msg)
	}

	pwHash, hashErr := generatePasswordHash(params.Password)
	if hashErr != nil {
		return newAuthResponse(false, "System error: failed to create login credentials")
	}

	_, verifyNotExistingUserResp := getUser(params.EmailAddr)
	if verifyNotExistingUserResp.Success {
		return newAuthResponse(false, "Registration failed: user with same email already exists")
	}

	userID := uniqueID.GenerateSnowflakeID()

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO users (user_id, email_addr, user_name, first_name,last_name, password_hash) 
				VALUES ($1,$2,$3,$4,$5,$6)`,
		userID, params.EmailAddr, params.UserName,
		params.FirstName, params.LastName, pwHash); insertErr != nil {
		return newAuthResponse(false, "System error: failed to create login credentials")
	}

	return newAuthResponse(true, "Registration complete")
}

func getUser(emailAddr string) (*User, *AuthResponse) {

	var user User
	user.EmailAddr = emailAddr
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT user_id, password_hash 
			FROM users 
			WHERE email_addr=$1 LIMIT 1`,
		emailAddr).Scan(&user.UserID, &user.PasswordHash)
	if getErr != nil {
		return nil, newAuthResponse(false, fmt.Sprintf("Can't find user with email: %v", emailAddr))
	}

	return &user, newAuthResponse(true, "Successfully retrieved user information")
}

func getUserInfoByID(userID string) (*UserInfo, error) {

	var userInfo UserInfo
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT first_name,last_name,user_name 
			FROM users 
			WHERE user_id=$1 LIMIT 1`,
		userID).Scan(&userInfo.FirstName, &userInfo.LastName, &userInfo.UserName)
	if getErr != nil {
		return nil, fmt.Errorf("Can't find user with id: %v", userID)
	}

	return &userInfo, nil
}
