package image

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	imageRouter := mux.NewRouter()

	imageRouter.HandleFunc("/api/frm/image/new", newImage)
	imageRouter.HandleFunc("/api/frm/image/resize", resizeImage)
	imageRouter.HandleFunc("/api/frm/image/setLabelFormat", setLabelFormat)
	imageRouter.HandleFunc("/api/frm/image/setPermissions", setPermissions)

	http.Handle("/api/frm/image/", imageRouter)
}

func newImage(w http.ResponseWriter, r *http.Request) {

	imageParams := NewImageParams{}
	if err := api.DecodeJSONRequest(r, &imageParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if imageRef, err := saveNewImage(imageParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *imageRef)
	}

}

func processImagePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ImagePropUpdater) {
	if imageRef, err := updateImageProps(propUpdater); err != nil {
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
