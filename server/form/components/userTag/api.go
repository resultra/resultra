package userTag

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

func init() {
	userTagRouter := mux.NewRouter()

	userTagRouter.HandleFunc("/api/frm/userTag/new", newUserTag)
	userTagRouter.HandleFunc("/api/frm/userTag/resize", resizeUserTag)
	userTagRouter.HandleFunc("/api/frm/userTag/setLabelFormat", setLabelFormat)
	userTagRouter.HandleFunc("/api/frm/userTag/setVisibility", setVisibility)
	userTagRouter.HandleFunc("/api/frm/userTag/setPermissions", setPermissions)
	userTagRouter.HandleFunc("/api/frm/userTag/setClearValueSupported", setClearValueSupported)
	userTagRouter.HandleFunc("/api/frm/userTag/setHelpPopupMsg", setHelpPopupMsg)
	userTagRouter.HandleFunc("/api/frm/userTag/setSelectableRoles", setSelectableRoles)
	userTagRouter.HandleFunc("/api/frm/userTag/setCurrUserSelectable", setCurrUserSelectable)

	userTagRouter.HandleFunc("/api/frm/userTag/setValidation", setValidation)
	userTagRouter.HandleFunc("/api/frm/userTag/validateInput", validateInputAPI)

	http.Handle("/api/frm/userTag/", userTagRouter)
}

func newUserTag(w http.ResponseWriter, r *http.Request) {

	userTagParams := NewUserTagParams{}
	if err := api.DecodeJSONRequest(r, &userTagParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if userTagRef, err := saveNewUserTag(trackerDBHandle, userTagParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *userTagRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params ValidateInputParams
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

func processUserTagPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater UserTagPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if userTagRef, err := updateUserTagProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userTagRef)
	}
}

func resizeUserTag(w http.ResponseWriter, r *http.Request) {
	var resizeParams UserTagResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UserTagLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params UserTagVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params UserTagPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params UserTagValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params UserTagClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setSelectableRoles(w http.ResponseWriter, r *http.Request) {
	var params SelectableRoleParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setCurrUserSelectable(w http.ResponseWriter, r *http.Request) {
	var params CurrUserSelectableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}
