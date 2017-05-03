package gauge

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	gaugeRouter := mux.NewRouter()

	gaugeRouter.HandleFunc("/api/dashboard/gauge/new", newGaugeAPI)

	gaugeRouter.HandleFunc("/api/dashboard/gauge/setTitle", setGaugeTitle)
	gaugeRouter.HandleFunc("/api/dashboard/gauge/setDimensions", setGaugeDimensions)

	http.Handle("/api/dashboard/gauge/", gaugeRouter)
}

func newGaugeAPI(w http.ResponseWriter, r *http.Request) {

	var params NewGaugeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if gaugeRef, err := newGauge(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, gaugeRef)
	}

}

func processGaugePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater GaugePropertyUpdater) {

	if gaugeRef, err := updateGaugeProps(propUpdater); err != nil {
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
