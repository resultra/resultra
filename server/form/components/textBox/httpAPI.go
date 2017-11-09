package textBox

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

func init() {
	textBoxRouter := mux.NewRouter()

	textBoxRouter.HandleFunc("/api/frm/textBox/new", newTextBox)
	textBoxRouter.HandleFunc("/api/frm/textBox/resize", resizeTextBox)
	textBoxRouter.HandleFunc("/api/frm/textBox/setLabelFormat", setLabelFormat)
	textBoxRouter.HandleFunc("/api/frm/textBox/setVisibility", setVisibility)
	textBoxRouter.HandleFunc("/api/frm/textBox/setPermissions", setPermissions)
	textBoxRouter.HandleFunc("/api/frm/textBox/setValueList", setValueList)
	textBoxRouter.HandleFunc("/api/frm/textBox/setValidation", setValidation)
	textBoxRouter.HandleFunc("/api/frm/textBox/setClearValueSupported", setClearValueSupported)
	textBoxRouter.HandleFunc("/api/frm/textBox/setHelpPopupMsg", setHelpPopupMsg)

	textBoxRouter.HandleFunc("/api/frm/textBox/validateInput", validateInputAPI)

	http.Handle("/api/frm/textBox/", textBoxRouter)
}

func newTextBox(w http.ResponseWriter, r *http.Request) {

	textBoxParams := NewTextBoxParams{}
	if err := api.DecodeJSONRequest(r, &textBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if textBoxRef, err := saveNewTextBox(trackerDBHandle, textBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *textBoxRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params TextBoxValidateInputParams
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

func processTextBoxPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TextBoxPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if textBoxRef, err := updateTextBoxProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, textBoxRef)
	}
}

func resizeTextBox(w http.ResponseWriter, r *http.Request) {
	var resizeParams TextBoxResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params TextBoxLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params TextBoxVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TextBoxPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setValueList(w http.ResponseWriter, r *http.Request) {
	var params TextBoxValueListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params TextBoxValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params TextBoxClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}
