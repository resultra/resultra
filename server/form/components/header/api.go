package header

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	headerRouter := mux.NewRouter()

	headerRouter.HandleFunc("/api/frm/header/new", newHeader)
	headerRouter.HandleFunc("/api/frm/header/resize", resizeHeader)
	headerRouter.HandleFunc("/api/frm/header/setLabel", setHeaderLabel)
	headerRouter.HandleFunc("/api/frm/header/setUnderlined", setUnderlined)
	headerRouter.HandleFunc("/api/frm/header/setSize", setSize)
	headerRouter.HandleFunc("/api/frm/header/setVisibility", setVisibility)

	http.Handle("/api/frm/header/", headerRouter)
}

func newHeader(w http.ResponseWriter, r *http.Request) {

	params := NewHeaderParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if headerRef, err := saveNewHeader(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *headerRef)
	}

}

func processHeaderPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater HeaderPropUpdater) {
	if headerRef, err := updateHeaderProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func resizeHeader(w http.ResponseWriter, r *http.Request) {
	var params HeaderResizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setHeaderLabel(w http.ResponseWriter, r *http.Request) {
	var params HeaderLabelParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setUnderlined(w http.ResponseWriter, r *http.Request) {
	var params HeaderUnderlinedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setSize(w http.ResponseWriter, r *http.Request) {
	var params HeaderSizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params HeaderVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}
