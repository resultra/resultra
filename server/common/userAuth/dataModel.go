package userAuth

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
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
	UserID           string         `json:"userID"`
	FirstName        string         `json:"firstName"`
	LastName         string         `json:"lastName"`
	UserName         string         `json:"userName"`
	IsWorkspaceAdmin bool           `json:"isWorkspaceAdmin"`
	Properties       UserProperties `json:"properties"`
}

type AdminUserInfo struct {
	UserID           string         `json:"userID"`
	FirstName        string         `json:"firstName"`
	LastName         string         `json:"lastName"`
	UserName         string         `json:"userName"`
	EmailAddress     string         `json:"emailAddress"`
	IsWorkspaceAdmin bool           `json:"isWorkspaceAdmin"`
	IsActive         bool           `json:"isActive"`
	Properties       UserProperties `json:"properties"`
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

func saveNewUser(trackerDBHandle *sql.DB, rawParams NewUserParams) *AuthResponse {

	params := rawParams.sanitize()

	firstNameResp := validateWellFormedRealName(params.FirstName)
	if !firstNameResp.Success {
		return newAuthResponse(false, "First name is required")
	}

	lastNameResp := validateWellFormedRealName(params.LastName)
	if !lastNameResp.Success {
		return newAuthResponse(false, "Last name is required")
	}

	userNameResp := validateNewUserName(trackerDBHandle, params.UserName)
	if !userNameResp.Success {
		return userNameResp
	}

	sanitizedEmail := strings.TrimSpace(params.EmailAddr)
	emailResp := ValidateWellFormedEmailAddr(sanitizedEmail)
	if !emailResp.Success {
		return emailResp
	}

	passwordValResp := validatePasswordStrength(params.Password)
	if !passwordValResp.ValidPassword {
		return newAuthResponse(false, passwordValResp.Msg)
	}

	props := newDefaultUserProperties()
	encodedProps, encodeErr := generic.EncodeJSONString(props)
	if encodeErr != nil {
		return newAuthResponse(false, "System error: failed to create login credentials")
	}

	pwHash, hashErr := generatePasswordHash(params.Password)
	if hashErr != nil {
		log.Printf("saveNewUser: system failure registering user: %v", hashErr)
		return newAuthResponse(false, "System error: failed to create login credentials")
	}

	_, verifyNotExistingUserResp := getUser(trackerDBHandle, params.EmailAddr)
	if verifyNotExistingUserResp.Success {
		return newAuthResponse(false, "Registration failed: user with same email already exists")
	}

	userID := uniqueID.GenerateSnowflakeID()

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO users (user_id, email_addr, user_name, first_name,last_name, password_hash,properties) 
				VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		userID, params.EmailAddr, params.UserName,
		params.FirstName, params.LastName, pwHash, encodedProps); insertErr != nil {
		log.Printf("saveNewUser: system failure registering user: %v", insertErr)

		return newAuthResponse(false, "System error: failed to create login credentials")
	}

	return newAuthResponse(true, "Registration complete")
}

type RegisterSingleUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
}

const singleUserDummyEmailAddr string = `single.user.email.not.applicable@example.com`
const singleUserDummyPassword string = `singleUserpwNotApplicable$!$`

func saveNewSingleUser(trackerDBHandle *sql.DB, params RegisterSingleUserParams) *AuthResponse {
	saveUserParams := NewUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		UserName:  params.UserName,
		EmailAddr: singleUserDummyEmailAddr,
		Password:  singleUserDummyPassword}
	return saveNewUser(trackerDBHandle, saveUserParams)
}

func getUser(trackerDBHandle *sql.DB, emailAddr string) (*User, *AuthResponse) {

	var user User
	user.EmailAddr = emailAddr
	getErr := trackerDBHandle.QueryRow(
		`SELECT user_id, password_hash 
			FROM users 
			WHERE email_addr=$1 LIMIT 1`,
		emailAddr).Scan(&user.UserID, &user.PasswordHash)
	if getErr != nil {
		return nil, newAuthResponse(false, fmt.Sprintf("Can't find user with email: %v", emailAddr))
	}

	return &user, newAuthResponse(true, "Successfully retrieved user information")
}

func validateUniqueUserName(trackerDBHandle *sql.DB, userName string) (bool, error) {

	upperUserName := strings.ToUpper(userName)

	rows, queryErr := trackerDBHandle.Query(
		`SELECT user_id FROM users WHERE UPPER(user_name)=$1`, upperUserName)
	if queryErr != nil {
		return false, fmt.Errorf("validateUniqueUserName: Can't query database for user name: %v", queryErr)
	}
	defer rows.Close()

	existingUserNameAlreadyUsed := rows.Next()
	if existingUserNameAlreadyUsed {
		return false, nil
	}

	return true, nil

}

func validateUniqueEmail(trackerDBHandle *sql.DB, emailAddr string) (bool, error) {

	upperEmail := strings.ToUpper(emailAddr)

	rows, queryErr := trackerDBHandle.Query(
		`SELECT user_id FROM users WHERE UPPER(email_addr)=$1`, upperEmail)
	if queryErr != nil {
		return false, fmt.Errorf("validateUniqueEmail: Can't query database for email address: %v", queryErr)
	}
	defer rows.Close()

	existingEmailAlreadyUsed := rows.Next()
	if existingEmailAlreadyUsed {
		return false, nil
	}

	return true, nil

}

func validateExistingEmail(trackerDBHandle *sql.DB, emailAddr string) (bool, error) {

	upperEmail := strings.ToUpper(emailAddr)

	rows, queryErr := trackerDBHandle.Query(
		`SELECT user_id FROM users WHERE UPPER(email_addr)=$1`, upperEmail)
	if queryErr != nil {
		return false, fmt.Errorf("validateUniqueEmail: Can't query database for email address: %v", queryErr)
	}
	defer rows.Close()

	existingEmailAlreadyUsed := rows.Next()
	if existingEmailAlreadyUsed {
		return true, nil
	}

	return false, nil

}

func GetUserInfoByID(trackerDBHandle *sql.DB, userID string) (*UserInfo, error) {

	var userInfo UserInfo
	userInfo.UserID = userID

	encodedProps := ""

	getErr := trackerDBHandle.QueryRow(
		`SELECT first_name,last_name,user_name,is_workspace_admin,properties
			FROM users 
			WHERE user_id=$1 LIMIT 1`,
		userID).Scan(&userInfo.FirstName,
		&userInfo.LastName, &userInfo.UserName,
		&userInfo.IsWorkspaceAdmin,
		&encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("Can't find user with id: %v: error = $v", userID, getErr)
	}

	userProps := newDefaultUserProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &userProps); decodeErr != nil {
		return nil, fmt.Errorf("getAlert: can't decode properties: %v", encodedProps)
	}
	userInfo.Properties = userProps

	return &userInfo, nil
}

func getAdminUserInfoByID(trackerDBHandle *sql.DB, userID string) (*AdminUserInfo, error) {

	var userInfo AdminUserInfo
	userInfo.UserID = userID

	encodedProps := ""

	getErr := trackerDBHandle.QueryRow(
		`SELECT first_name,last_name,user_name,is_workspace_admin,properties,is_active,email_addr
			FROM users 
			WHERE user_id=$1 LIMIT 1`,
		userID).Scan(&userInfo.FirstName,
		&userInfo.LastName, &userInfo.UserName,
		&userInfo.IsWorkspaceAdmin,
		&encodedProps,
		&userInfo.IsActive,
		&userInfo.EmailAddress)
	if getErr != nil {
		return nil, fmt.Errorf("Can't find user with id: %v: error = $v", userID, getErr)
	}

	userProps := newDefaultUserProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &userProps); decodeErr != nil {
		return nil, fmt.Errorf("getAlert: can't decode properties: %v", encodedProps)
	}
	userInfo.Properties = userProps

	return &userInfo, nil
}

func GetUserInfoByEmail(trackerDBHandle *sql.DB, emailAddr string) (*UserInfo, error) {

	var userInfo UserInfo

	upperEmail := strings.ToUpper(emailAddr)

	encodedProps := ""

	getErr := trackerDBHandle.QueryRow(
		`SELECT user_id,first_name,last_name,user_name,is_workspace_admin,properties
			FROM users 
			WHERE UPPER(email_addr)=$1 LIMIT 1`,
		upperEmail).Scan(
		&userInfo.UserID,
		&userInfo.FirstName,
		&userInfo.LastName,
		&userInfo.UserName,
		&userInfo.IsWorkspaceAdmin,
		&encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("Can't find user with given email address: %v", emailAddr)
	}

	userProps := newDefaultUserProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &userProps); decodeErr != nil {
		return nil, fmt.Errorf("getAlert: can't decode properties: %v", encodedProps)
	}
	userInfo.Properties = userProps

	return &userInfo, nil
}

func getAllUsersInfo(trackerDBHandle *sql.DB) ([]AdminUserInfo, error) {
	rows, queryErr := trackerDBHandle.Query(
		`SELECT user_id,first_name,last_name,user_name,email_addr,is_workspace_admin,properties,is_active FROM users`)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllUsersInfo: Can't query database for users: %v", queryErr)
	}
	defer rows.Close()

	allUserInfo := []AdminUserInfo{}

	for rows.Next() {

		var userInfo AdminUserInfo
		encodedProps := ""

		if scanErr := rows.Scan(
			&userInfo.UserID,
			&userInfo.FirstName,
			&userInfo.LastName,
			&userInfo.UserName,
			&userInfo.EmailAddress,
			&userInfo.IsWorkspaceAdmin,
			&encodedProps,
			&userInfo.IsActive); scanErr != nil {
			return nil, fmt.Errorf("getAllUsersInfo: Failure querying database: %v", scanErr)
		}

		userProps := newDefaultUserProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &userProps); decodeErr != nil {
			return nil, fmt.Errorf("getAlert: can't decode properties: %v", encodedProps)
		}
		userInfo.Properties = userProps

		allUserInfo = append(allUserInfo, userInfo)

	}

	return allUserInfo, nil

}

func updateUserProperties(trackerDBHandle *sql.DB, userID string, props UserProperties) error {

	encodedProps, encodeErr := generic.EncodeJSONString(props)
	if encodeErr != nil {
		return fmt.Errorf("updateUserProperties: can't update user properties: %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(
		`UPDATE users set properties=$1 where user_id=$2`,
		encodedProps, userID); updateErr != nil {

		return fmt.Errorf("updateUserProperties: system failure updating user properties: %v", updateErr)
	}
	return nil

}

type SetUserActiveParams struct {
	UserID   string `json:"userID"`
	IsActive bool   `json:"isActive"`
}

func setUserActive(trackerDBHandle *sql.DB, params SetUserActiveParams) error {
	if _, updateErr := trackerDBHandle.Exec(
		`UPDATE users set is_active=$1 where user_id=$2`,
		params.IsActive, params.UserID); updateErr != nil {

		return fmt.Errorf("updateUserProperties: system failure updating user properties: %v", updateErr)
	}
	return nil

}
