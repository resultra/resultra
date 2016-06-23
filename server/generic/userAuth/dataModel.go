package userAuth

import (
	"fmt"
	"github.com/gocql/gocql"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

type User struct {
	UserID       string
	EmailAddr    string
	PasswordHash string
}

type UserInfo struct {
	EmailAddr string
}

type NewUserParams struct {
	EmailAddr string `json:"emailAddr"`
	Password  string `json:"password"`
}

func saveNewUser(params NewUserParams) *AuthResponse {

	pwHash, hashErr := generatePasswordHash(params.Password)
	if hashErr != nil {
		return newAuthResponse(false, "System error: failed to create login credentials")
	}

	_, verifyNotExistingUserResp := getUser(params.EmailAddr)
	if verifyNotExistingUserResp.Success {
		return newAuthResponse(false, "Registration failed: user with same email already exists")
	}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return newAuthResponse(false, "System error: failed to create database session")
	}
	defer dbSession.Close()

	userID := gocql.TimeUUID().String()

	if insertErr := dbSession.Query(`INSERT INTO users (user_id, email_addr, password_hash) VALUES (?,?,?)`,
		userID, params.EmailAddr, pwHash).Exec(); insertErr != nil {
		return newAuthResponse(false, "System error: failed to create login credentials")
	}

	return newAuthResponse(true, "Registration complete")
}

func getUser(emailAddr string) (*User, *AuthResponse) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, newAuthResponse(false, "System error: failed to create database session")
	}
	defer dbSession.Close()

	var user User
	user.EmailAddr = emailAddr
	getErr := dbSession.Query(
		`SELECT user_id, password_hash 
			FROM users 
			WHERE email_addr=? LIMIT 1`,
		emailAddr).Scan(&user.UserID, &user.PasswordHash)
	if getErr != nil {

		return nil, newAuthResponse(false, fmt.Sprintf("Can't find user with email: %v", emailAddr))
	}

	return &user, newAuthResponse(true, "Successfully retrieved user information")
}

func getUserInfoByID(userID string) (*UserInfo, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("System error: failed to create database session")
	}
	defer dbSession.Close()

	var userInfo UserInfo
	getErr := dbSession.Query(
		`SELECT email_addr 
			FROM users 
			WHERE user_id=? LIMIT 1`,
		userID).Scan(&userInfo.EmailAddr)
	if getErr != nil {

		return nil, fmt.Errorf("Can't find user with id: %v", userID)
	}

	return &userInfo, nil
}
