package userRoleController

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/userRole"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	roleRouter := mux.NewRouter()

	roleRouter.HandleFunc("/api/userRole/validateRoleName", validateRoleNameAPI)
	roleRouter.HandleFunc("/api/userRole/setName", setNameAPI)
	roleRouter.HandleFunc("/api/userRole/newRole", newRoleAPI)
	roleRouter.HandleFunc("/api/userRole/get", getRoleAPI)

	roleRouter.HandleFunc("/api/userRole/getListRolePrivs", getListRolePrivsAPI)
	roleRouter.HandleFunc("/api/userRole/setListRolePrivs", setListRolePrivsAPI)

	roleRouter.HandleFunc("/api/userRole/getNewItemRolePrivs", getNewItemRolePrivsAPI)
	roleRouter.HandleFunc("/api/userRole/setNewItemRolePrivs", setNewItemRolePrivsAPI)

	roleRouter.HandleFunc("/api/userRole/getAlertRolePrivs", getAlertRolePrivsAPI)
	roleRouter.HandleFunc("/api/userRole/setAlertRolePrivs", setAlertRolePrivsAPI)

	roleRouter.HandleFunc("/api/userRole/getRoleListPrivs", getRoleListPrivsAPI)
	roleRouter.HandleFunc("/api/userRole/getRoleDashboardPrivs", getRoleDashboardPrivsAPI)

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

	var params userRole.NewDatabaseRoleParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if newRole, newErr := userRole.NewDatabaseRole(params); newErr != nil {
		api.WriteErrorResponse(w, newErr)
	} else {
		api.WriteJSONResponse(w, newRole)
	}

}

type GetRoleParams struct {
	RoleID string `json:"roleID"`
}

func getRoleAPI(w http.ResponseWriter, r *http.Request) {

	var params GetRoleParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	/*	if verifyErr := VerifyCurrUserIsDatabaseAdminForRole(r, params.RoleID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	} */

	if roleInfo, getErr := userRole.GetUserRole(params.RoleID); getErr != nil {
		api.WriteErrorResponse(w, getErr)
	} else {
		api.WriteJSONResponse(w, roleInfo)
	}

}

func setNameAPI(w http.ResponseWriter, r *http.Request) {
	var params userRole.SetRoleNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	userRole.ProcessRolePropUpdate(w, r, params)
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

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForItemList(r, params.ListID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if listRolePrivs, err := userRole.GetListRolePrivs(params.ListID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, listRolePrivs)

	}

}

func setListRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.SetListRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForItemList(r, params.ListID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.SetListRolePrivs(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResponse := true
		api.WriteJSONResponse(w, successResponse)
	}
}

func setNewItemRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.SetNewItemFormLinkRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForNewItemLink(r, params.LinkID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.SetNewItemFormLinkRolePrivs(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResponse := true
		api.WriteJSONResponse(w, successResponse)
	}
}

func getNewItemRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.GetNewItemPrivParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForUserRole(r, params.RoleID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if roleNewItemPrivs, err := userRole.GetNewItemPrivs(params.RoleID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, roleNewItemPrivs)

	}

}

func setAlertRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.SetAlertRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForAlert(r, params.AlertID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.SetAlertRolePrivs(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResponse := true
		api.WriteJSONResponse(w, successResponse)
	}
}

func getAlertRolePrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.GetAlertPrivParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForUserRole(r, params.RoleID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if roleAlertPrivs, err := userRole.GetAlertPrivs(params.RoleID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, roleAlertPrivs)

	}

}

func getRoleListPrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.GetRoleListPrivParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForUserRole(r, params.RoleID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if roleListPrivs, err := getRoleListPrivsWithDefaults(params.RoleID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, roleListPrivs)

	}

}

type DashboardRolePrivsParams struct {
	DashboardID string `json:"dashboardID"`
}

func setDashboardRolePrivsAPI(w http.ResponseWriter, r *http.Request) {
	var params userRole.SetDashboardRolePrivsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForDashboard(r, params.DashboardID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.SetDashboardRolePrivs(params); err != nil {
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

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForDashboard(r, params.DashboardID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if dashboardRolePrivs, err := userRole.GetDashboardRolePrivs(params.DashboardID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dashboardRolePrivs)

	}

}

func getRoleDashboardPrivsAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.GetRoleDashboardPrivParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForUserRole(r, params.RoleID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if roleDashboardPrivs, err := getRoleDashboardPrivsWithDefaults(params.RoleID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, roleDashboardPrivs)

	}

}

func getDatabaseRolesAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.DatabaseRolesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if roles, err := userRole.GetDatabaseRoles(params.DatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, roles)

	}

}
