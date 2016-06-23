package userAuth

import (
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var authCookieStore *sessions.CookieStore

func init() {

	// TODO (Important) replace the session key with a key from a config file.
	// Both the authentication and encryption keys are 32 bytes for AES-256
	// http://www.gorillatoolkit.org/pkg/sessions#NewCookieStore
	authCookieStore = sessions.NewCookieStore([]byte("nRrHLlHcHH0u7fUxyzHje9m7uJ5SnJzP"),
		[]byte("CAp1KsJncuMzARfookqSFLqsBi5ag2bE"))
	authCookieStore.MaxAge(3600 * 8) // 8 hours
}

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

func loginUser(rw http.ResponseWriter, req *http.Request, params LoginParams) *AuthResponse {

	user, getResp := getUser(params.EmailAddr)
	if !getResp.Success {
		return getResp
	}

	pwVerify := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if pwVerify != nil {
		return newAuthResponse(false, "Incorrect password")
	}

	authSession, sessErr := authCookieStore.Get(req, "auth")
	if sessErr != nil {
		return newAuthResponse(false, "System error: couldn't create session")
	}
	authSession.Values["user_id"] = user.UserID
	if saveErr := authSession.Save(req, rw); saveErr != nil {
		return newAuthResponse(false, "System error: couldn't save login session")
	}

	return newAuthResponse(true, "Login successful")
}
