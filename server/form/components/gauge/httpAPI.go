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

	gaugeRouter.HandleFunc("/api/frm/gauge/new", newGauge)
	gaugeRouter.HandleFunc("/api/frm/gauge/resize", resizeGauge)
	gaugeRouter.HandleFunc("/api/frm/gauge/setRange", setRange)
	gaugeRouter.HandleFunc("/api/frm/gauge/setThresholds", setThresholds)
	gaugeRouter.HandleFunc("/api/frm/gauge/setLabelFormat", setLabelFormat)
	gaugeRouter.HandleFunc("/api/frm/gauge/setVisibility", setVisibility)
	gaugeRouter.HandleFunc("/api/frm/gauge/setValueFormat", setValueFormat)
	gaugeRouter.HandleFunc("/api/frm/gauge/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/frm/gauge/", gaugeRouter)
}

func newGauge(w http.ResponseWriter, r *http.Request) {

	params := NewGaugeParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if gaugeRef, err := saveNewGauge(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *gaugeRef)
	}

}

func processGaugePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater GaugePropUpdater) {
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

func resizeGauge(w http.ResponseWriter, r *http.Request) {
	var resizeParams GaugeResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, resizeParams)
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

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params GaugeLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params GaugeVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setValueFormat(w http.ResponseWriter, r *http.Request) {
	var params GaugeValueFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processGaugePropUpdate(w, r, params)
}
