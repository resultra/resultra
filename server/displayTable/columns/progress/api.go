package progress

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	progressRouter := mux.NewRouter()

	progressRouter.HandleFunc("/api/tableView/progress/new", newProgress)

	progressRouter.HandleFunc("/api/tableView/progress/get", getProgressAPI)

	progressRouter.HandleFunc("/api/tableView/progress/setRange", setRange)
	progressRouter.HandleFunc("/api/tableView/progress/setThresholds", setThresholds)
	progressRouter.HandleFunc("/api/tableView/progress/setLabelFormat", setLabelFormat)
	progressRouter.HandleFunc("/api/tableView/progress/setValueFormat", setValueFormat)
	progressRouter.HandleFunc("/api/tableView/progress/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/tableView/progress/", progressRouter)
}

func newProgress(w http.ResponseWriter, r *http.Request) {

	params := NewProgressParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if progressRef, err := saveNewProgress(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *progressRef)
	}

}

type GetProgressParams struct {
	ParentTableID string `json:"parentTableID"`
	ProgressID    string `json:"progressID"`
}

func getProgressAPI(w http.ResponseWriter, r *http.Request) {

	var params GetProgressParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	progress, err := getProgress(trackerDBHandle, params.ParentTableID, params.ProgressID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *progress)
}

func processProgressPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ProgressPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if progressRef, err := updateProgressProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, progressRef)
	}
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
