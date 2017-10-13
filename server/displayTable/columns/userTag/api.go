package userTag

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	userTagRouter := mux.NewRouter()

	userTagRouter.HandleFunc("/api/tableView/userTag/new", newUserTag)

	userTagRouter.HandleFunc("/api/tableView/userTag/get", getUserTagAPI)

	userTagRouter.HandleFunc("/api/tableView/userTag/setLabelFormat", setLabelFormat)
	userTagRouter.HandleFunc("/api/tableView/userTag/setClearValueSupported", setClearValueSupported)
	userTagRouter.HandleFunc("/api/tableView/userTag/setHelpPopupMsg", setHelpPopupMsg)

	userTagRouter.HandleFunc("/api/tableView/userTag/setSelectableRoles", setSelectableRoles)
	userTagRouter.HandleFunc("/api/tableView/userTag/setCurrUserSelectable", setCurrUserSelectable)

	userTagRouter.HandleFunc("/api/tableView/userTag/setPermissions", setPermissions)
	userTagRouter.HandleFunc("/api/tableView/userTag/setValidation", setValidation)
	userTagRouter.HandleFunc("/api/tableView/userTag/validateInput", validateInputAPI)

	http.Handle("/api/tableView/userTag/", userTagRouter)
}

func newUserTag(w http.ResponseWriter, r *http.Request) {

	userTagParams := NewUserTagParams{}
	if err := api.DecodeJSONRequest(r, &userTagParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if userTagRef, err := saveNewUserTag(userTagParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *userTagRef)
	}

}

type GetUserTagParams struct {
	ParentTableID string `json:"parentTableID"`
	UserTagID     string `json:"userTagID"`
}

func getUserTagAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUserTagParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	textInput, err := getUserTag(params.ParentTableID, params.UserTagID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *textInput)
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

func processUserTagPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater UserTagPropUpdater) {
	if userTagRef, err := updateUserTagProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userTagRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UserTagLabelFormatParams
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