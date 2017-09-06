package urlLink

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	urlLinkRouter := mux.NewRouter()

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/new", newUrlLink)

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/get", getUrlLinkAPI)

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setLabelFormat", setLabelFormat)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setPermissions", setPermissions)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setValidation", setValidation)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setClearValueSupported", setClearValueSupported)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setHelpPopupMsg", setHelpPopupMsg)

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/validateInput", validateInputAPI)

	http.Handle("/api/tableView/urlLink/", urlLinkRouter)
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

type GetUrlLinkParams struct {
	ParentTableID string `json:"parentTableID"`
	UrlLinkID   string `json:"urlLinkID"`
}

func getUrlLinkAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUrlLinkParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	urlLink, err := getUrlLink(params.ParentTableID, params.UrlLinkID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *urlLink)
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

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkLabelFormatParams
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
