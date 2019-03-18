package userAuth

import (
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/runtimeConfig"
)

var authCookieStore *sessions.CookieStore

func init() {

	// TODO (Important) replace the session key with a key from a config file.
	// Both the authentication and encryption keys are 32 bytes for AES-256
	// http://www.gorillatoolkit.org/pkg/sessions#NewCookieStore
	authCookieStore = sessions.NewCookieStore([]byte("nRrHLlHcHH0u7fUxyzHje9m7uJ5SnJzP"),
		[]byte("CAp1KsJncuMzARfookqSFLqsBi5ag2bE"))

	authCookieStore.MaxAge(3600 * 8) //  default to 8 hours
}

func updateCookieStoreAge() {
	numHoursSession := 8
	// If the login is for a single-user workspace (on the desktop), extend the login
	// time considerably. Since the user is logged in automatically, there is no need
	// to log them out automatically.
	if runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace {
		numHoursSession = 24 * 365 * 10 // 10 years!
	}

	authCookieStore.MaxAge(3600 * numHoursSession)

}

func generatePasswordHash(password string) (string, error) {

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return "", fmt.Errorf("generatePasswordHash: couldn't save user: %v", hashErr)
	}
	hashStr := string(hash)

	return hashStr, nil

}

type LoginParams NewUserParams

func loginUser(rw http.ResponseWriter, req *http.Request, params LoginParams) *AuthResponse {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		errMsg := fmt.Sprintf("System error: couldn't get tracker database: %v", dbErr)
		return newAuthResponse(false, errMsg)
	}

	user, getResp := getUser(trackerDBHandle, params.EmailAddr)
	if !getResp.Success {
		return getResp
	}

	pwVerify := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if pwVerify != nil {
		return newAuthResponse(false, "Incorrect password")
	}

	updateCookieStoreAge()

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

func LoginSingleUser(rw http.ResponseWriter, req *http.Request) *AuthResponse {
	loginParams := LoginParams{
		EmailAddr: singleUserDummyEmailAddr,
		Password:  singleUserDummyPassword}
	return loginUser(rw, req, loginParams)
}

func GetCurrentUserID(req *http.Request) (string, error) {

	authSession, sessErr := authCookieStore.Get(req, "auth")
	if sessErr != nil {
		return "", fmt.Errorf("GetCurrentUserID: Couldn't get session to authenticate user")
	}

	userID, userIDFound := authSession.Values["user_id"].(string)
	if !userIDFound {
		return "", fmt.Errorf("GetCurrentUserID: Can't get session value for user")
	}

	return userID, nil

}

func GetCurrentUserInfo(req *http.Request) (*UserInfo, error) {

	userID, getErr := GetCurrentUserID(req)
	if getErr != nil {
		return nil, fmt.Errorf("GetCurrentUserInfo: Can't get session value for user (user not signed in)")
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, fmt.Errorf("GetCurrentUserInfo: failure accessing tracker database for user: error %v", dbErr)
	}

	return GetUserInfoByID(trackerDBHandle, userID)

}

func signOutUser(rw http.ResponseWriter, req *http.Request) *AuthResponse {

	authSession, sessErr := authCookieStore.Get(req, "auth")
	if sessErr != nil {
		return newAuthResponse(false, "Couldn't get session information to sign out")
	}

	authSession.Options.MaxAge = -1 // kill the cookie
	if saveErr := authSession.Save(req, rw); saveErr != nil {
		return newAuthResponse(false, "Couldn't save session information while signing out")
	}

	return newAuthResponse(true, "Successfully signed out")

}
