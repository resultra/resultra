package progress

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	progressRouter := mux.NewRouter()

	progressRouter.HandleFunc("/api/frm/progress/new", newProgress)
	progressRouter.HandleFunc("/api/frm/progress/resize", resizeProgress)
	progressRouter.HandleFunc("/api/frm/progress/setRange", setRange)
	progressRouter.HandleFunc("/api/frm/progress/setThresholds", setThresholds)
	progressRouter.HandleFunc("/api/frm/progress/setLabelFormat", setLabelFormat)
	progressRouter.HandleFunc("/api/frm/progress/setVisibility", setVisibility)
	progressRouter.HandleFunc("/api/frm/progress/setValueFormat", setValueFormat)
	progressRouter.HandleFunc("/api/frm/progress/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/frm/progress/", progressRouter)
}

func newProgress(w http.ResponseWriter, r *http.Request) {

	params := NewProgressParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if progressRef, err := saveNewProgress(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *progressRef)
	}

}

func processProgressPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ProgressPropUpdater) {
	if progressRef, err := updateProgressProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, progressRef)
	}
}

func resizeProgress(w http.ResponseWriter, r *http.Request) {
	var resizeParams ProgressResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processProgressPropUpdate(w, r, resizeParams)
}

func setRange(w http.ResponseWriter, r *http.Request) {
	var params SetRangeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processProgressPropUpdate(w, r, params)
}

func setThresholds(w http.ResponseWriter, r *http.Request) {
	var params SetThresholdsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processProgressPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params ProgressLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processProgressPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params ProgressVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processProgressPropUpdate(w, r, params)
}

func setValueFormat(w http.ResponseWriter, r *http.Request) {
	var params ProgressValueFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processProgressPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processProgressPropUpdate(w, r, params)
}
