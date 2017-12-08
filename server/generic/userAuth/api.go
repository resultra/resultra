package userAuth

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForPkgImport struct{ DummyVal int64 }

func init() {
	authRouter := mux.NewRouter()

	authRouter.HandleFunc("/auth/register", registerNewUserAPI)
	authRouter.HandleFunc("/auth/login", loginUserAPI)
	authRouter.HandleFunc("/auth/signout", signoutUserAPI)

	authRouter.HandleFunc("/auth/getCurrentUserInfo", getCurrentUserInfoAPI)
	authRouter.HandleFunc("/auth/getUserInfo", getUserInfoAPI)
	authRouter.HandleFunc("/auth/getUsersInfo", getUsersInfoAPI)

	authRouter.HandleFunc("/auth/getAllUsersInfo", getAllUsersInfoAPI)
	authRouter.HandleFunc("/auth/getAdminUserInfo", getAdminUserInfoAPI)

	authRouter.HandleFunc("/auth/validateName", validateNameAPI)
	authRouter.HandleFunc("/auth/validateNewUserName", validateNewUserNameAPI)
	authRouter.HandleFunc("/auth/validateNewUserEmail", validateNewUserEmailAPI)

	authRouter.HandleFunc("/auth/validateExistingUserEmail", validateExistingUserEmailAPI)
	authRouter.HandleFunc("/auth/sendResetPasswordLink", sendResetPasswordLinkAPI)
	authRouter.HandleFunc("/auth/sendResetPasswordLinkByUserID", sendResetPasswordLinkByUserIDAPI)

	authRouter.HandleFunc("/auth/resetPassword", resetPasswordAPI)
	authRouter.HandleFunc("/auth/setUserActive", setUserActiveAPI)

	authRouter.HandleFunc("/auth/validatePasswordStrength", validatePasswordStrengthAPI)

	authRouter.HandleFunc("/auth/searchUsers", searchUsersAPI)

	http.Handle("/auth/", authRouter)
}

func registerNewUserAPI(w http.ResponseWriter, r *http.Request) {
	var params NewUserParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	newUserResp := saveNewUser(trackerDBHandle, params)
	api.WriteJSONResponse(w, newUserResp)

}

func sendResetPasswordLinkAPI(w http.ResponseWriter, r *http.Request) {
	var params PasswordResetParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	resetResp := sendResetPasswordLink(trackerDBHandle, params)
	api.WriteJSONResponse(w, resetResp)

}

func sendResetPasswordLinkByUserIDAPI(w http.ResponseWriter, r *http.Request) {
	var params PasswordResetByUserIdParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	resetResp := sendResetPasswordLinkByUserID(trackerDBHandle, params)
	api.WriteJSONResponse(w, resetResp)

}
func loginUserAPI(w http.ResponseWriter, r *http.Request) {
	var params LoginParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	authResp := loginUser(w, r, params)
	api.WriteJSONResponse(w, authResp)

}

func resetPasswordAPI(w http.ResponseWriter, r *http.Request) {

	var params PasswordResetEntryParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	authResp := resetPassword(trackerDBHandle, params)
	api.WriteJSONResponse(w, authResp)

}

func setUserActiveAPI(w http.ResponseWriter, r *http.Request) {

	var params SetUserActiveParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	setErr := setUserActive(trackerDBHandle, params)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func signoutUserAPI(w http.ResponseWriter, r *http.Request) {
	authResp := signOutUser(w, r)
	api.WriteJSONResponse(w, authResp)

}

func getCurrentUserInfoAPI(w http.ResponseWriter, r *http.Request) {
	userInfo, err := GetCurrentUserInfo(r)

	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, userInfo)
}

type GetUserInfoParams struct {
	UserID string `json:"userID"`
}

func getUserInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUserInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	userInfo, err := GetUserInfoByID(trackerDBHandle, params.UserID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, userInfo)
}

func getAdminUserInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUserInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	userInfo, err := getAdminUserInfoByID(trackerDBHandle, params.UserID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, userInfo)
}

type GetUsersInfoParams struct {
	UserIDs []string `json:"userIDs"`
}

func getUsersInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUsersInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	usersInfo := []UserInfo{}

	for _, currUserID := range params.UserIDs {
		userInfo, err := GetUserInfoByID(trackerDBHandle, currUserID)
		if err != nil {
			api.WriteErrorResponse(w, err)
			return
		}
		usersInfo = append(usersInfo, *userInfo)
	}

	api.WriteJSONResponse(w, usersInfo)
}

type GetAllUsersInfoParams struct {
	IncludeInactiveUsers bool `json:"includeInactiveUsers"`
}

func getAllUsersInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetAllUsersInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	allUsersInfo, userErr := getAllUsersInfo(trackerDBHandle)

	if userErr != nil {
		api.WriteErrorResponse(w, userErr)
		return
	}

	api.WriteJSONResponse(w, allUsersInfo)

}

func searchUsersAPI(w http.ResponseWriter, r *http.Request) {

	searchTerm := r.FormValue("searchTerm")
	page := r.FormValue("page")

	log.Printf("Search users: term = %v, page = %v", searchTerm, page)

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	results, err := searchUsers(trackerDBHandle, searchTerm)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	api.WriteJSONResponse(w, results)

}

func validateNewUserNameAPI(w http.ResponseWriter, r *http.Request) {

	userName := r.FormValue("userName")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	isValid, err := validateUniqueUserName(trackerDBHandle, userName)

	if err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := isValid
	api.WriteJSONResponse(w, response)

}

func validateNewUserEmailAPI(w http.ResponseWriter, r *http.Request) {

	emailAddr := r.FormValue("emailAddr")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	isValid, err := validateUniqueEmail(trackerDBHandle, emailAddr)

	if err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := isValid
	api.WriteJSONResponse(w, response)

}

func validateExistingUserEmailAPI(w http.ResponseWriter, r *http.Request) {

	emailAddr := r.FormValue("emailAddr")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	isValid, err := validateExistingEmail(trackerDBHandle, emailAddr)

	if err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := isValid
	api.WriteJSONResponse(w, response)

}

func validateNameAPI(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	nameResp := validateWellFormedRealName(name)
	isValid := nameResp.Success

	response := isValid
	api.WriteJSONResponse(w, response)

}

func validatePasswordStrengthAPI(w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("password")

	pwResp := validatePasswordStrength(password)

	response := pwResp.ValidPassword
	api.WriteJSONResponse(w, response)

}
