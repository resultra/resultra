package userRole

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	roleRouter := mux.NewRouter()

	roleRouter.HandleFunc("/api/userRole/validateRoleName", validateRoleNameAPI)
	roleRouter.HandleFunc("/api/userRole/newRole", newRoleAPI)

	http.Handle("/api/userRole/", roleRouter)
}

type RoleNameValidationParams struct {
	Name string `json:"name"`
}

func validateRoleNameAPI(w http.ResponseWriter, r *http.Request) {

	roleName := r.FormValue("roleName")
	log.Printf("Role Name: %v", roleName)

	if generic.WellFormedItemName(roleName) {
		response := true
		api.WriteJSONResponse(w, response)
		return
	}

	//	response := true
	response := "Invalid role name"
	api.WriteJSONResponse(w, response)
}

func newRoleAPI(w http.ResponseWriter, r *http.Request) {

	var params NewDatabaseRoleWithPrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if newErr := newDatabaseRoleWithPrivs(params); newErr != nil {
		api.WriteErrorResponse(w, newErr)
	} else {
		successResponse := true
		api.WriteJSONResponse(w, successResponse)
	}

}
