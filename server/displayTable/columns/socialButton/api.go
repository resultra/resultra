package socialButton

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	socialButtonRouter := mux.NewRouter()

	socialButtonRouter.HandleFunc("/api/tableView/socialButton/new", newSocialButton)

	socialButtonRouter.HandleFunc("/api/tableView/socialButton/get", getSocialButtonAPI)

	socialButtonRouter.HandleFunc("/api/tableView/socialButton/setIcon", setIcon)
	socialButtonRouter.HandleFunc("/api/tableView/socialButton/setPermissions", setPermissions)
	socialButtonRouter.HandleFunc("/api/tableView/socialButton/setLabelFormat", setLabelFormat)
	socialButtonRouter.HandleFunc("/api/tableView/socialButton/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/tableView/socialButton/", socialButtonRouter)
}

func newSocialButton(w http.ResponseWriter, r *http.Request) {

	params := NewSocialButtonParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if socialButtonRef, err := saveNewSocialButton(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *socialButtonRef)
	}

}

type GetSocialButtonParams struct {
	ParentTableID  string `json:"parentTableID"`
	SocialButtonID string `json:"socialButtonID"`
}

func getSocialButtonAPI(w http.ResponseWriter, r *http.Request) {

	var params GetSocialButtonParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	rating, err := getSocialButton(trackerDBHandle, params.ParentTableID, params.SocialButtonID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *rating)
}

func processSocialButtonPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SocialButtonPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if socialButtonRef, err := updateSocialButtonProps(trackerDBHandle, propUpdater); err != nil {
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
