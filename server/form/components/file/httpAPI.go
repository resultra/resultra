package file

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	fileRouter := mux.NewRouter()

	fileRouter.HandleFunc("/api/frm/file/new", newFile)
	fileRouter.HandleFunc("/api/frm/file/resize", resizeFile)
	fileRouter.HandleFunc("/api/frm/file/setLabelFormat", setLabelFormat)
	fileRouter.HandleFunc("/api/frm/file/setVisibility", setVisibility)
	fileRouter.HandleFunc("/api/frm/file/setPermissions", setPermissions)
	fileRouter.HandleFunc("/api/frm/file/setValidation", setValidation)
	fileRouter.HandleFunc("/api/frm/file/setClearValueSupported", setClearValueSupported)
	fileRouter.HandleFunc("/api/frm/file/setHelpPopupMsg", setHelpPopupMsg)

	fileRouter.HandleFunc("/api/frm/file/validateInput", validateInputAPI)

	http.Handle("/api/frm/file/", fileRouter)
}

func newFile(w http.ResponseWriter, r *http.Request) {

	fileParams := NewFileParams{}
	if err := api.DecodeJSONRequest(r, &fileParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if fileRef, err := saveNewFile(fileParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *fileRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params FileValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(params)
	api.WriteJSONResponse(w, validationResp)
}

func processFilePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater FilePropUpdater) {
	if fileRef, err := updateFileProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fileRef)
	}
}

func resizeFile(w http.ResponseWriter, r *http.Request) {
	var resizeParams FileResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params FileLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params FileVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params FilePermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params FileValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params FileClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}
