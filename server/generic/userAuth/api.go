package userAuth

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForPkgImport struct{ DummyVal int64 }

func init() {
	authRouter := mux.NewRouter()

	authRouter.HandleFunc("/auth/register", registerNewUserAPI)
	authRouter.HandleFunc("/auth/login", loginUserAPI)

	http.Handle("/auth/", authRouter)
}

func registerNewUserAPI(w http.ResponseWriter, r *http.Request) {
	var params NewUserParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newUserResp := saveNewUser(params)
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
