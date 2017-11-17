package trackerDatabase

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

func processDatabasePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater DatabasePropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedDB, err := updateDatabaseProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedDB)
	}
}

func SetActiveAPI(w http.ResponseWriter, r *http.Request) {
	var params SetDatabaseActiveParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatabasePropUpdate(w, r, params)
}

func SetNameAPI(w http.ResponseWriter, r *http.Request) {
	var params SetDatabaseNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatabasePropUpdate(w, r, params)
}

func SetListOrderAPI(w http.ResponseWriter, r *http.Request) {
	var params SetListOrderParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatabasePropUpdate(w, r, params)
}

func SetDashboardOrderAPI(w http.ResponseWriter, r *http.Request) {
	var params SetDashboardOrderParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatabasePropUpdate(w, r, params)
}

func SetFormLinkOrderAPI(w http.ResponseWriter, r *http.Request) {
	var params SetFormLinkOrderParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatabasePropUpdate(w, r, params)
}

func SetDescriptionAPI(w http.ResponseWriter, r *http.Request) {
	var params SetDescriptionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatabasePropUpdate(w, r, params)
}

func ValidateDatabaseNameAPI(w http.ResponseWriter, r *http.Request) {

	databaseID := r.FormValue("databaseID")
	databaseName := r.FormValue("databaseName")

	if err := validateDatabaseName(databaseID, databaseName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func ValidateNewTrackerNameAPI(w http.ResponseWriter, r *http.Request) {

	trackerName := r.FormValue("trackerName")

	if err := validateNewTrackerName(trackerName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
