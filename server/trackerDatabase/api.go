// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package trackerDatabase

import (
	"fmt"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
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
