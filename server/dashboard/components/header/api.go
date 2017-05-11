package header

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	headerRouter := mux.NewRouter()

	headerRouter.HandleFunc("/api/dashboard/header/new", newHeaderAPI)

	headerRouter.HandleFunc("/api/dashboard/header/setTitle", setHeaderTitle)
	headerRouter.HandleFunc("/api/dashboard/header/setDimensions", setHeaderDimensions)
	headerRouter.HandleFunc("/api/dashboard/header/setSize", setSize)
	headerRouter.HandleFunc("/api/dashboard/header/setUnderlined", setUnderline)

	http.Handle("/api/dashboard/header/", headerRouter)
}

func newHeaderAPI(w http.ResponseWriter, r *http.Request) {

	var params NewHeaderParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if headerRef, err := newHeader(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}

}

func processHeaderPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater HeaderPropertyUpdater) {

	if headerRef, err := updateHeaderProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func setHeaderTitle(w http.ResponseWriter, r *http.Request) {
	var titleParams SetHeaderTitleParams
	if err := api.DecodeJSONRequest(r, &titleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, titleParams)
}

func setHeaderDimensions(w http.ResponseWriter, r *http.Request) {

	var params SetHeaderDimensionsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setSize(w http.ResponseWriter, r *http.Request) {

	var params SetHeaderSizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setUnderline(w http.ResponseWriter, r *http.Request) {

	var params SetHeaderUnderlineParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}