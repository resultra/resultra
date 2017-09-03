package label

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	labelRouter := mux.NewRouter()

	labelRouter.HandleFunc("/api/frm/label/new", newLabel)
	labelRouter.HandleFunc("/api/frm/label/resize", resizeLabel)
	labelRouter.HandleFunc("/api/frm/label/setLabelFormat", setLabelFormat)
	labelRouter.HandleFunc("/api/frm/label/setVisibility", setVisibility)
	labelRouter.HandleFunc("/api/frm/label/setPermissions", setPermissions)
	labelRouter.HandleFunc("/api/frm/label/setHelpPopupMsg", setHelpPopupMsg)

	labelRouter.HandleFunc("/api/frm/label/setValidation", setValidation)
	labelRouter.HandleFunc("/api/frm/label/validateInput", validateInputAPI)

	http.Handle("/api/frm/label/", labelRouter)
}

func newLabel(w http.ResponseWriter, r *http.Request) {

	labelParams := NewLabelParams{}
	if err := api.DecodeJSONRequest(r, &labelParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if labelRef, err := saveNewLabel(labelParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *labelRef)
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

func processLabelPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater LabelPropUpdater) {
	if labelRef, err := updateLabelProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, labelRef)
	}
}

func resizeLabel(w http.ResponseWriter, r *http.Request) {
	var resizeParams LabelResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params LabelLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params LabelVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params LabelPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params LabelValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}
