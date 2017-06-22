package formButton

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	buttonRouter := mux.NewRouter()

	buttonRouter.HandleFunc("/api/tableView/formButton/new", newButton)

	buttonRouter.HandleFunc("/api/tableView/formButton/setPopupBehavior", setPopupBehavior)

	buttonRouter.HandleFunc("/api/tableView/formButton/setDefaultVals", setDefaultVals)
	buttonRouter.HandleFunc("/api/tableView/formButton/setSize", setSize)
	buttonRouter.HandleFunc("/api/tableView/formButton/setColorScheme", setColorScheme)
	buttonRouter.HandleFunc("/api/tableView/formButton/setIcon", setIcon)

	http.Handle("/api/tableView/formButton/", buttonRouter)
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

type GetButtonParams struct {
	ParentTableID string `json:"parentTableID"`
	ButtonID      string `json:"buttonID"`
}

func getNoteAPI(w http.ResponseWriter, r *http.Request) {

	var params GetButtonParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	formButton, err := getButton(params.ParentTableID, params.ButtonID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *formButton)
}

func processButtonPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ButtonPropUpdater) {
	if headerRef, err := updateButtonProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func setPopupBehavior(w http.ResponseWriter, r *http.Request) {
	var params ButtonBehaviorParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setDefaultVals(w http.ResponseWriter, r *http.Request) {
	var params ButtonDefaultValParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setSize(w http.ResponseWriter, r *http.Request) {
	var params ButtonSizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setColorScheme(w http.ResponseWriter, r *http.Request) {
	var params ButtonColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setIcon(w http.ResponseWriter, r *http.Request) {
	var params ButtonIconParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}
