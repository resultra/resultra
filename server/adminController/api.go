package adminController

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
	adminRouter := mux.NewRouter()

	adminRouter.HandleFunc("/api/admin/getUserRoleInfo", getUserRoleInfoAPI)

	http.Handle("/api/admin/", adminRouter)
}

func getUserRoleInfoAPI(w http.ResponseWriter, r *http.Request) {

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
