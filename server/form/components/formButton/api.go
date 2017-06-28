package formButton

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	buttonRouter := mux.NewRouter()

	buttonRouter.HandleFunc("/api/frm/formButton/new", newButton)

	buttonRouter.HandleFunc("/api/frm/formButton/get", getButtonAPI)

	buttonRouter.HandleFunc("/api/frm/formButton/resize", resizeButton)
	buttonRouter.HandleFunc("/api/frm/formButton/setPopupBehavior", setPopupBehavior)
	buttonRouter.HandleFunc("/api/frm/formButton/setDefaultVals", setDefaultVals)
	buttonRouter.HandleFunc("/api/frm/formButton/setSize", setSize)
	buttonRouter.HandleFunc("/api/frm/formButton/setColorScheme", setColorScheme)
	buttonRouter.HandleFunc("/api/frm/formButton/setIcon", setIcon)
	buttonRouter.HandleFunc("/api/frm/formButton/setVisibility", setVisibility)

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

type GetButtonParams struct {
	ButtonID string `json:"buttonID"`
}

func getButtonAPI(w http.ResponseWriter, r *http.Request) {

	params := GetButtonParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if buttonRef, err := getButtonFromButtonID(params.ButtonID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *buttonRef)
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

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params ButtonVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}
