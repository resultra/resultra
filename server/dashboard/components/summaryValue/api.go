// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package summaryValue

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {

	summaryValRouter := mux.NewRouter()

	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/new", newSummaryValAPI)

	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/setTitle", setSummaryValTitle)
	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/setDimensions", setSummaryValDimensions)

	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/setValSummary", setValSummary)
	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/setThresholds", setThresholds)

	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/setDefaultFilterRules", setDefaultFilterRules)
	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/setPreFilterRules", setPreFilterRules)
	summaryValRouter.HandleFunc("/api/dashboard/summaryVal/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/dashboard/summaryVal/", summaryValRouter)
}

func newSummaryValAPI(w http.ResponseWriter, r *http.Request) {

	var params NewSummaryValParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if summaryValRef, err := newSummaryVal(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryValRef)
	}

}

func processSummaryValPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SummaryValPropertyUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if summaryValRef, err := updateSummaryValProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryValRef)
	}
}

func setSummaryValTitle(w http.ResponseWriter, r *http.Request) {
	var titleParams SetSummaryValTitleParams
	if err := api.DecodeJSONRequest(r, &titleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryValPropUpdate(w, r, titleParams)
}

func setSummaryValDimensions(w http.ResponseWriter, r *http.Request) {

	var params SetSummaryValDimensionsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryValPropUpdate(w, r, params)
}

func setDefaultFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetDefaultFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryValPropUpdate(w, r, params)
}

func setPreFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetPreFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryValPropUpdate(w, r, params)
}

func setValSummary(w http.ResponseWriter, r *http.Request) {
	var params SetValSummaryParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryValPropUpdate(w, r, params)
}

func setThresholds(w http.ResponseWriter, r *http.Request) {
	var params SetThresholdsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryValPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params SetHelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryValPropUpdate(w, r, params)
}
