package formButton

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	buttonRouter := mux.NewRouter()

	buttonRouter.HandleFunc("/api/frm/header/new", newButton)
	buttonRouter.HandleFunc("/api/frm/header/resize", resizeButton)

	http.Handle("/api/frm/formButton/", buttonRouter)
}

func newButton(w http.ResponseWriter, r *http.Request) {

	params := NewButtonParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if headerRef, err := saveNewButton(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *headerRef)
	}

}

func processButtonPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ButtonPropUpdater) {
	if headerRef, err := updateButtonProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func resizeButton(w http.ResponseWriter, r *http.Request) {
	var params ButtonResizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}
