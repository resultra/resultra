package userRole

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/stringValidation"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	roleRouter := mux.NewRouter()

	roleRouter.HandleFunc("/api/userRole/validateRoleName", validateRoleNameAPI)
	roleRouter.HandleFunc("/api/userRole/newRole", newRoleAPI)

	roleRouter.HandleFunc("/api/userRole/getListRolePrivs", getListRolePrivsAPI)
	roleRouter.HandleFunc("/api/userRole/setListRolePrivs", setListRolePrivsAPI)

	roleRouter.HandleFunc("/api/userRole/getDatabaseRoles", getDatabaseRolesAPI)

	roleRouter.HandleFunc("/api/userRole/getDashboardRolePrivs", getDashboardRolePrivsAPI)
	roleRouter.HandleFunc("/api/userRole/setDashboardRolePrivs", setDashboardRolePrivsAPI)

	http.Handle("/api/userRole/", roleRouter)
}

type RoleNameValidationParams struct {
	Name string `json:"name"`
}

func validateRoleNameAPI(w http.ResponseWriter, r *http.Request) {

	roleName := r.FormValue("roleName")
	log.Printf("Role Name: %v", roleName)

	if stringValidation.WellFormedItemName(roleName) {
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

type ListRolePrivsParams struct {
	ListID string `json:"listID"`
}

func getListRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params ListRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := VerifyCurrUserIsDatabaseAdminForItemList(r, params.ListID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if listRolePrivs, err := getListRolePrivs(params.ListID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, listRolePrivs)

	}

}

func setListRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params SetListRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := VerifyCurrUserIsDatabaseAdminForItemList(r, params.ListID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := setListRolePrivs(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResponse := true
		api.WriteJSONResponse(w, successResponse)
	}
}

type DashboardRolePrivsParams struct {
	DashboardID string `json:"dashboardID"`
}

func setDashboardRolePrivsAPI(w http.ResponseWriter, r *http.Request) {
	var params SetDashboardRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := VerifyCurrUserIsDatabaseAdminForDashboard(r, params.DashboardID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := setDashboardRolePrivs(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResponse := true
		api.WriteJSONResponse(w, successResponse)
	}

}

func getDashboardRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params DashboardRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := VerifyCurrUserIsDatabaseAdminForDashboard(r, params.DashboardID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if dashboardRolePrivs, err := getDashboardRolePrivs(params.DashboardID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dashboardRolePrivs)

	}

}

func getDatabaseRolesAPI(w http.ResponseWriter, r *http.Request) {

	var params DatabaseRolesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if roles, err := getDatabaseRoles(params.DatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, roles)

	}

}
