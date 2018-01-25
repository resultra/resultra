package adminController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	adminRouter := mux.NewRouter()

	adminRouter.HandleFunc("/api/admin/getUserRoleInfo", getUserRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/getRoleInfo", getRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/getRoleCollaborators", getRoleCollaboratorAPI)

	adminRouter.HandleFunc("/api/admin/addCollaborator", addCollaboratorAPI)
	adminRouter.HandleFunc("/api/admin/getAllCollaboratorInfo", getAllCollaboratorInfoAPI)

	adminRouter.HandleFunc("/api/admin/getSingleUserRoleInfo", getSingleUserRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/setUserRoleInfo", setUserRoleInfoAPI)
	adminRouter.HandleFunc("/api/admin/deleteCollaborator", deleteCollaboratorAPI)

	adminRouter.HandleFunc("/api/admin/ping", pingAPI)

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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if userRoleInfo, err := getUserRolesInfo(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userRoleInfo)
	}

}

func getRoleCollaboratorAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.GetRoleCollaboratorsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if roleCollabInfo, err := userRole.GetRoleCollaborators(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, roleCollabInfo)
	}

}

func getAllCollaboratorInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params userRole.GetAllCollaborUserInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if collabInfo, err := userRole.GetAllCollaboratorUserInfo(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, collabInfo)
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
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if roleInfo, err := getRoleInfo(trackerDBHandle, params); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if collaboratorUserRoleInfo, err := addCollaborator(trackerDBHandle, params); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if userRolesInfo, err := userRole.GetCollaboratorRoleInfoAPI(trackerDBHandle, params); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.SetCollaboratorRoleInfo(trackerDBHandle, params); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, r, params.DatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := userRole.DeleteCollaborator(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		successResult := true
		api.WriteJSONResponse(w, successResult)

	}

}

func pingAPI(w http.ResponseWriter, r *http.Request) {

	api.WriteJSONResponse(w, true)
}
