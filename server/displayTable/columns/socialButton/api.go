package socialButton

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	socialButtonRouter := mux.NewRouter()

	socialButtonRouter.HandleFunc("api/tableView/socialButton/new", newSocialButton)
	socialButtonRouter.HandleFunc("api/tableView/socialButton/setIcon", setIcon)
	socialButtonRouter.HandleFunc("api/tableView/socialButton/setLabelFormat", setLabelFormat)
	socialButtonRouter.HandleFunc("api/tableView/socialButton/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("api/tableView/socialButton/", socialButtonRouter)
}

func newSocialButton(w http.ResponseWriter, r *http.Request) {

	params := NewSocialButtonParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if socialButtonRef, err := saveNewSocialButton(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *socialButtonRef)
	}

}

func processSocialButtonPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SocialButtonPropUpdater) {
	if socialButtonRef, err := updateSocialButtonProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, socialButtonRef)
	}
}

func setIcon(w http.ResponseWriter, r *http.Request) {
	var params SocialButtonIconParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params SocialButtonLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params SocialButtonPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}
