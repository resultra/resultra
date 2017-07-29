package adminController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
	//	"resultra/datasheet/server/generic/userAuth"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	adminRouter := mux.NewRouter()

	adminRouter.HandleFunc("/api/admin/getUserRoleInfo", getUserRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/getRoleInfo", getRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/addCollaborator", addCollaboratorAPI)
	adminRouter.HandleFunc("/api/admin/getSingleUserRoleInfo", getSingleUserRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/setUserRoleInfo", setUserRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/deleteCollaborator", deleteCollaboratorAPI)

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

func getRoleInfoAPI(w http.ResponseWriter, r *http.Request) {

	// TODO - Once filtering is implemented on a per form/dashboard basis,
	// pass in the parent filter.
	var params GetRoleInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if roleInfo, err := getRoleInfo(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *roleInfo)
	}

}

func addCollaboratorAPI(w http.ResponseWriter, r *http.Request) {
	var params AddCollaboratorParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if collaboratorUserRoleInfo, err := addCollaborator(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, collaboratorUserRoleInfo)

	}

}

func getSingleUserRoleInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.GetCollaboratorRoleInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if userRolesInfo, err := userRole.GetCollaboratorRoleInfoAPI(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userRolesInfo)

	}

}

func setUserRoleInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.SetCollaboratorRoleInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.SetCollaboratorRoleInfo(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResult := true
		api.WriteJSONResponse(w, successResult)

	}

}

func deleteCollaboratorAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.DeleteCollaboratorParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.DeleteCollaborator(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResult := true
		api.WriteJSONResponse(w, successResult)

	}

}
