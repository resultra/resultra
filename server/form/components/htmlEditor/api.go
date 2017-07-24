package htmlEditor

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	htmlEditorRouter := mux.NewRouter()

	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/new", newHtmlEditor)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/resize", resizeHtmlEditor)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/setLabelFormat", setLabelFormat)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/setVisibility", setVisibility)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/setPermissions", setPermissions)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/setValidation", setValidation)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/validateInput", validateInputAPI)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/frm/htmlEditor/", htmlEditorRouter)
}

func newHtmlEditor(w http.ResponseWriter, r *http.Request) {

	params := NewHtmlEditorParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if editorRef, err := saveNewHtmlEditor(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *editorRef)
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

func processHtmlEditorPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater HtmlEditorPropUpdater) {
	if checkBoxRef, err := updateHtmlEditorProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, checkBoxRef)
	}
}

func resizeHtmlEditor(w http.ResponseWriter, r *http.Request) {
	var resizeParams HtmlEditorResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHtmlEditorPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params EditorLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHtmlEditorPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params EditorVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHtmlEditorPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params EditorPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHtmlEditorPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params EditorValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHtmlEditorPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHtmlEditorPropUpdate(w, r, params)
}
