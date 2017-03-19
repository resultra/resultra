package rating

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	ratingRouter := mux.NewRouter()

	ratingRouter.HandleFunc("/api/frm/rating/new", newRating)
	ratingRouter.HandleFunc("/api/frm/rating/resize", resizeRating)
	ratingRouter.HandleFunc("/api/frm/rating/setTooltips", setTooltips)
	ratingRouter.HandleFunc("/api/frm/rating/setIcon", setIcon)
	ratingRouter.HandleFunc("/api/frm/rating/setLabelFormat", setLabelFormat)
	ratingRouter.HandleFunc("/api/frm/rating/setVisibility", setVisibility)

	http.Handle("/api/frm/rating/", ratingRouter)
}

func newRating(w http.ResponseWriter, r *http.Request) {

	ratingParams := NewRatingParams{}
	if err := api.DecodeJSONRequest(r, &ratingParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if ratingRef, err := saveNewRating(ratingParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *ratingRef)
	}

}

func processRatingPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater RatingPropUpdater) {
	if ratingRef, err := updateRatingProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, ratingRef)
	}
}

func resizeRating(w http.ResponseWriter, r *http.Request) {
	var resizeParams RatingResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, resizeParams)
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

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params RatingVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processRatingPropUpdate(w, r, params)
}
