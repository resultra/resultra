package caption

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	captionRouter := mux.NewRouter()

	captionRouter.HandleFunc("/api/frm/caption/new", newCaption)
	captionRouter.HandleFunc("/api/frm/caption/resize", resizeCaption)
	captionRouter.HandleFunc("/api/frm/caption/setLabel", setCaptionLabel)
	captionRouter.HandleFunc("/api/frm/caption/setCaption", setCaptionCaption)
	captionRouter.HandleFunc("/api/frm/caption/setColorScheme", setColorScheme)

	http.Handle("/api/frm/caption/", captionRouter)
}

func newCaption(w http.ResponseWriter, r *http.Request) {

	params := NewCaptionParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if captionRef, err := saveNewCaption(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *captionRef)
	}

}

func processCaptionPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater CaptionPropUpdater) {
	if captionRef, err := updateCaptionProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, captionRef)
	}
}

func setCaptionLabel(w http.ResponseWriter, r *http.Request) {
	var labelParams CaptionLabelParams
	if err := api.DecodeJSONRequest(r, &labelParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCaptionPropUpdate(w, r, labelParams)
}

func setCaptionCaption(w http.ResponseWriter, r *http.Request) {
	var params CaptionCaptionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCaptionPropUpdate(w, r, params)
}

func resizeCaption(w http.ResponseWriter, r *http.Request) {
	var params CaptionResizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCaptionPropUpdate(w, r, params)
}

func setColorScheme(w http.ResponseWriter, r *http.Request) {
	var params CaptionColorParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCaptionPropUpdate(w, r, params)
}
