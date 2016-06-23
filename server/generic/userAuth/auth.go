package userAuth

import (
	"fmt"
	//	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength int = 8

func validateWellFormedPassword(password string) error {
	if len(password) < minPasswordLength {
		return fmt.Errorf("validateWellFormedPassword: password must be at least 8 characters")
	}
	return nil
}

func generatePasswordHash(password string) (string, error) {

	if validatePassErr := validateWellFormedPassword(password); validatePassErr != nil {
		return "", fmt.Errorf("generatePasswordHash: invalid password: %v", validatePassErr)
	}

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return "", fmt.Errorf("generatePasswordHash: couldn't save user: %v", hashErr)
	}
	hashStr := string(hash)

	return hashStr, nil

}

type LoginParams NewUserParams

func loginUser(params LoginParams) *AuthResponse {

	user, getResp := getUser(params.EmailAddr)
	if !getResp.Success {
		return getResp
	}

	pwVerify := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if pwVerify != nil {
		return newAuthResponse(false, "Incorrect password")
	}
	return newAuthResponse(true, "Login successful")
}
