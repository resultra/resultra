package attachment

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

func init() {
	attachmentRouter := mux.NewRouter()

	attachmentRouter.HandleFunc("/api/frm/attachment/new", newImage)
	attachmentRouter.HandleFunc("/api/frm/attachment/resize", resizeImage)
	attachmentRouter.HandleFunc("/api/frm/attachment/setLabelFormat", setLabelFormat)
	attachmentRouter.HandleFunc("/api/frm/attachment/setPermissions", setPermissions)
	attachmentRouter.HandleFunc("/api/frm/attachment/setValidation", setValidation)
	attachmentRouter.HandleFunc("/api/frm/attachment/validateInput", validateInputAPI)
	attachmentRouter.HandleFunc("/api/frm/attachment/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/frm/attachment/", attachmentRouter)
}

func newImage(w http.ResponseWriter, r *http.Request) {

	imageParams := NewImageParams{}
	if err := api.DecodeJSONRequest(r, &imageParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if imageRef, err := saveNewImage(trackerDBHandle, imageParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *imageRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params ValidateInputParams
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

func processImagePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ImagePropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if imageRef, err := updateImageProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, imageRef)
	}
}

func resizeImage(w http.ResponseWriter, r *http.Request) {
	var resizeParams ImageResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var resizeParams AttachmentLabelFormatParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, resizeParams)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var resizeParams AttachmentPermissionParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, resizeParams)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var resizeParams AttachmentValidationParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, resizeParams)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, params)
}
