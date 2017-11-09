package userAuth

import (
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

func loginUserAPI(w http.ResponseWriter, r *http.Request) {
	var params LoginParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	authResp := loginUser(w, r, params)
	api.WriteJSONResponse(w, authResp)

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
