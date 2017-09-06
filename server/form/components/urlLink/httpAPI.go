package urlLink

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	urlLinkRouter := mux.NewRouter()

	urlLinkRouter.HandleFunc("/api/frm/urlLink/new", newUrlLink)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/resize", resizeUrlLink)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setLabelFormat", setLabelFormat)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setVisibility", setVisibility)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setPermissions", setPermissions)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setValidation", setValidation)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setClearValueSupported", setClearValueSupported)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setHelpPopupMsg", setHelpPopupMsg)

	urlLinkRouter.HandleFunc("/api/frm/urlLink/validateInput", validateInputAPI)

	http.Handle("/api/frm/urlLink/", urlLinkRouter)
}

func newUrlLink(w http.ResponseWriter, r *http.Request) {

	urlLinkParams := NewUrlLinkParams{}
	if err := api.DecodeJSONRequest(r, &urlLinkParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if urlLinkRef, err := saveNewUrlLink(urlLinkParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *urlLinkRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params UrlLinkValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(params)
	api.WriteJSONResponse(w, validationResp)
}

func processUrlLinkPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater UrlLinkPropUpdater) {
	if urlLinkRef, err := updateUrlLinkProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, urlLinkRef)
	}
}

func resizeUrlLink(w http.ResponseWriter, r *http.Request) {
	var resizeParams UrlLinkResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}
