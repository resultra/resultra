package userController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	//	"resultra/datasheet/server/generic/userAuth"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	userRouter := mux.NewRouter()

	userRouter.HandleFunc("/api/user/getRoleInfo", getRoleInfoAPI)

	http.Handle("/api/user/", userRouter)
}

func getRoleInfoAPI(w http.ResponseWriter, r *http.Request) {

	// TODO - Once filtering is implemented on a per form/dashboard basis,
	// pass in the parent filter.
	var params GetUserRolesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if userRoleInfo, err := getUserRolesInfo(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userRoleInfo)
	}

}
