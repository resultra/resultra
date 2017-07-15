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
	userSelectionRouter.HandleFunc("/api/frm/userSelection/setVisibility", setVisibility)
	userSelectionRouter.HandleFunc("/api/frm/userSelection/setPermissions", setPermissions)
	userSelectionRouter.HandleFunc("/api/frm/userSelection/setClearValueSupported", setClearValueSupported)

	userSelectionRouter.HandleFunc("/api/frm/userSelection/setValidation", setValidation)
	userSelectionRouter.HandleFunc("/api/frm/userSelection/validateInput", validateInputAPI)

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

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params ValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(params)
	api.WriteJSONResponse(w, validationResp)
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

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}
