package textInput

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	textInputRouter := mux.NewRouter()

	textInputRouter.HandleFunc("/api/tableView/textInput/new", newTextInput)

	textInputRouter.HandleFunc("/api/tableView/textInput/get", getTextInputAPI)

	textInputRouter.HandleFunc("/api/tableView/textInput/setLabelFormat", setLabelFormat)
	textInputRouter.HandleFunc("/api/tableView/textInput/setPermissions", setPermissions)
	textInputRouter.HandleFunc("/api/tableView/textInput/setValueList", setValueList)
	textInputRouter.HandleFunc("/api/tableView/textInput/setValidation", setValidation)

	textInputRouter.HandleFunc("/api/tableView/textInput/validateInput", validateInputAPI)

	http.Handle("/api/tableView/textInput/", textInputRouter)
}

func newTextInput(w http.ResponseWriter, r *http.Request) {

	textInputParams := NewTextInputParams{}
	if err := api.DecodeJSONRequest(r, &textInputParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if textInputRef, err := saveNewTextInput(textInputParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *textInputRef)
	}

}

type GetTextInputParams struct {
	ParentTableID string `json:"parentTableID"`
	TextInputID   string `json:"textInputID"`
}

func getTextInputAPI(w http.ResponseWriter, r *http.Request) {

	var params GetTextInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	textInput, err := getTextInput(params.ParentTableID, params.TextInputID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *textInput)
}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params TextInputValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(params)
	api.WriteJSONResponse(w, validationResp)
}

func processTextInputPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TextInputPropUpdater) {
	if textInputRef, err := updateTextInputProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, textInputRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params TextInputLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TextInputPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setValueList(w http.ResponseWriter, r *http.Request) {
	var params TextInputValueListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params TextInputValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}
