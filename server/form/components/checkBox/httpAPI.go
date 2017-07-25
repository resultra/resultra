package checkBox

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	checkBoxRouter := mux.NewRouter()

	checkBoxRouter.HandleFunc("/api/frm/checkBox/new", newCheckBox)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/resize", resizeCheckBox)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setColorScheme", setColorScheme)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setStrikethrough", setStrikethrough)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setLabelFormat", setLabelFormat)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setVisibility", setVisibility)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setPermissions", setPermissions)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setClearValueSupported", setClearValueSupported)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setHelpPopupMsg", setHelpPopupMsg)

	checkBoxRouter.HandleFunc("/api/frm/checkBox/setValidation", setValidation)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/validateInput", validateInputAPI)

	http.Handle("/api/frm/checkBox/", checkBoxRouter)
}

func newCheckBox(w http.ResponseWriter, r *http.Request) {

	checkBoxParams := NewCheckBoxParams{}
	if err := api.DecodeJSONRequest(r, &checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if checkBoxRef, err := saveNewCheckBox(checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *checkBoxRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params CheckBoxValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(params)
	api.WriteJSONResponse(w, validationResp)
}

func processCheckBoxPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater CheckBoxPropUpdater) {
	if checkBoxRef, err := updateCheckBoxProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, checkBoxRef)
	}
}

func resizeCheckBox(w http.ResponseWriter, r *http.Request) {
	var resizeParams CheckBoxResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, resizeParams)
}

func setColorScheme(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setStrikethrough(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxStrikethroughParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}
