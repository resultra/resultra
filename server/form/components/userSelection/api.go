package userSelection

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	userSelectionRouter := mux.NewRouter()

	userSelectionRouter.HandleFunc("/api/frm/userSelection/new", newUserSelection)
	userSelectionRouter.HandleFunc("/api/frm/userSelection/resize", resizeUserSelection)
	userSelectionRouter.HandleFunc("/api/frm/userSelection/setLabelFormat", setLabelFormat)

	http.Handle("/api/frm/userSelection/", userSelectionRouter)
}

func newUserSelection(w http.ResponseWriter, r *http.Request) {

	userSelectionParams := NewUserSelectionParams{}
	if err := api.DecodeJSONRequest(r, &userSelectionParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if userSelectionRef, err := saveNewUserSelection(userSelectionParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *userSelectionRef)
	}

}

func processUserSelectionPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater UserSelectionPropUpdater) {
	if userSelectionRef, err := updateUserSelectionProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userSelectionRef)
	}
}

func resizeUserSelection(w http.ResponseWriter, r *http.Request) {
	var resizeParams UserSelectionResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}
