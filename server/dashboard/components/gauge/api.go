// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package gauge

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {

	gaugeRouter := mux.NewRouter()

	gaugeRouter.HandleFunc("/api/dashboard/gauge/new", newGaugeAPI)

	gaugeRouter.HandleFunc("/api/dashboard/gauge/setTitle", setGaugeTitle)
	gaugeRouter.HandleFunc("/api/dashboard/gauge/setDimensions", setGaugeDimensions)

	gaugeRouter.HandleFunc("/api/dashboard/gauge/setValSummary", setValSummary)
	gaugeRouter.HandleFunc("/api/dashboard/gauge/setRange", setRange)
	gaugeRouter.HandleFunc("/api/dashboard/gauge/setThresholds", setThresholds)
	gaugeRouter.HandleFunc("/api/dashboard/gauge/setHelpPopupMsg", setHelpPopupMsg)

	gaugeRouter.HandleFunc("/api/dashboard/gauge/setDefaultFilterRules", setDefaultFilterRules)
	gaugeRouter.HandleFunc("/api/dashboard/gauge/setPreFilterRules", setPreFilterRules)

	http.Handle("/api/dashboard/gauge/", gaugeRouter)
}

func newGaugeAPI(w http.ResponseWriter, r *http.Request) {

	var params NewGaugeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if gaugeRef, err := newGauge(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, gaugeRef)
	}

}

func processGaugePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater GaugePropertyUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if gaugeRef, err := updateGaugeProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, gaugeRef)
	}
}

func setGaugeTitle(w http.ResponseWriter, r *http.Request) {
	var titleParams SetGaugeTitleParams
	if err := api.DecodeJSONRequest(r, &titleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, titleParams)
}

func setGaugeDimensions(w http.ResponseWriter, r *http.Request) {

	var params SetGaugeDimensionsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setDefaultFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetDefaultFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setPreFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetPreFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setValSummary(w http.ResponseWriter, r *http.Request) {
	var params SetValSummaryParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setRange(w http.ResponseWriter, r *http.Request) {
	var params SetRangeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setThresholds(w http.ResponseWriter, r *http.Request) {
	var params SetThresholdsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params SetHelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}
