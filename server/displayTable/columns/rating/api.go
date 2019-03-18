// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package rating

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	ratingRouter := mux.NewRouter()

	ratingRouter.HandleFunc("/api/tableView/rating/new", newRating)

	ratingRouter.HandleFunc("/api/tableView/rating/get", getRatingAPI)

	ratingRouter.HandleFunc("/api/tableView/rating/setTooltips", setTooltips)
	ratingRouter.HandleFunc("/api/tableView/rating/setIcon", setIcon)
	ratingRouter.HandleFunc("/api/tableView/rating/setLabelFormat", setLabelFormat)
	ratingRouter.HandleFunc("/api/tableView/rating/setClearValueSupported", setClearValueSupported)
	ratingRouter.HandleFunc("/api/tableView/rating/setHelpPopupMsg", setHelpPopupMsg)
	ratingRouter.HandleFunc("/api/tableView/rating/setRange", setRange)

	ratingRouter.HandleFunc("/api/tableView/rating/setPermissions", setPermissions)
	ratingRouter.HandleFunc("/api/tableView/rating/setValidation", setValidation)
	ratingRouter.HandleFunc("/api/tableView/rating/validateInput", validateInputAPI)

	http.Handle("/api/tableView/rating/", ratingRouter)
}

func newRating(w http.ResponseWriter, r *http.Request) {

	ratingParams := NewRatingParams{}
	if err := api.DecodeJSONRequest(r, &ratingParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if ratingRef, err := saveNewRating(trackerDBHandle, ratingParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *ratingRef)
	}

}

type GetRatingParams struct {
	ParentTableID string `json:"parentTableID"`
	RatingID      string `json:"ratingID"`
}

func getRatingAPI(w http.ResponseWriter, r *http.Request) {

	var params GetRatingParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	rating, err := getRating(trackerDBHandle, params.ParentTableID, params.RatingID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *rating)
}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params RatingValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	validationResp := validateInput(trackerDBHandle, params)
	api.WriteJSONResponse(w, validationResp)
}

func processRatingPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater RatingPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if ratingRef, err := updateRatingProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, ratingRef)
	}
}

func setTooltips(w http.ResponseWriter, r *http.Request) {
	var tooltipParams RatingTooltipParams
	if err := api.DecodeJSONRequest(r, &tooltipParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, tooltipParams)
}

func setIcon(w http.ResponseWriter, r *http.Request) {
	var params RatingIconParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params RatingLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params RatingPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params RatingValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params RatingClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}

func setRange(w http.ResponseWriter, r *http.Request) {
	var params RangeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}
