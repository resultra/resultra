package toggle

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	toggleRouter := mux.NewRouter()

	toggleRouter.HandleFunc("/api/frm/toggle/new", newToggle)
	toggleRouter.HandleFunc("/api/frm/toggle/resize", resizeToggle)

	toggleRouter.HandleFunc("/api/frm/toggle/setOffColorScheme", setOffColorScheme)
	toggleRouter.HandleFunc("/api/frm/toggle/setOnColorScheme", setOnColorScheme)
	toggleRouter.HandleFunc("/api/frm/toggle/setOffLabel", setOffLabel)
	toggleRouter.HandleFunc("/api/frm/toggle/setOnLabel", setOnLabel)
	toggleRouter.HandleFunc("/api/frm/toggle/setClearValueSupported", setClearValueSupported)
	toggleRouter.HandleFunc("/api/frm/toggle/setHelpPopupMsg", setHelpPopupMsg)

	toggleRouter.HandleFunc("/api/frm/toggle/setLabelFormat", setLabelFormat)
	toggleRouter.HandleFunc("/api/frm/toggle/setVisibility", setVisibility)
	toggleRouter.HandleFunc("/api/frm/toggle/setPermissions", setPermissions)

	toggleRouter.HandleFunc("/api/frm/toggle/setValidation", setValidation)
	toggleRouter.HandleFunc("/api/frm/toggle/validateInput", validateInputAPI)

	http.Handle("/api/frm/toggle/", toggleRouter)
}

func newToggle(w http.ResponseWriter, r *http.Request) {

	toggleParams := NewToggleParams{}
	if err := api.DecodeJSONRequest(r, &toggleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if toggleRef, err := saveNewToggle(trackerDBHandle, toggleParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *toggleRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params ToggleValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	validationResp := validateInput(trackerDBHandle, params)
	api.WriteJSONResponse(w, validationResp)
}

func processTogglePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TogglePropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if toggleRef, err := updateToggleProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, toggleRef)
	}
}

func resizeToggle(w http.ResponseWriter, r *http.Request) {
	var resizeParams ToggleResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, resizeParams)
}

func setOffColorScheme(w http.ResponseWriter, r *http.Request) {
	var params ToggleOffColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setOnColorScheme(w http.ResponseWriter, r *http.Request) {
	var params ToggleOnColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setOffLabel(w http.ResponseWriter, r *http.Request) {
	var params ToggleOffLabelParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setOnLabel(w http.ResponseWriter, r *http.Request) {
	var params ToggleOnLabelParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params ToggleLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params ToggleVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TogglePermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params ToggleValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params ToggleClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}
