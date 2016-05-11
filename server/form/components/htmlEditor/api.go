package htmlEditor

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	htmlEditorRouter := mux.NewRouter()

	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/new", newHtmlEditor)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/resize", resizeHtmlEditor)
	htmlEditorRouter.HandleFunc("/api/frm/htmlEditor/reposition", repositionHtmlEditor)

	http.Handle("/api/frm/htmlEditor/", htmlEditorRouter)
}

func newHtmlEditor(w http.ResponseWriter, r *http.Request) {

	params := NewHtmlEditorParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if checkBoxRef, err := saveNewHtmlEditor(appEngCntxt, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *checkBoxRef)
	}

}

func processHtmlEditorPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater HtmlEditorPropUpdater) {
	appEngCntxt := appengine.NewContext(r)
	if checkBoxRef, err := updateHtmlEditorProps(appEngCntxt, propUpdater); err != nil {
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

func repositionHtmlEditor(w http.ResponseWriter, r *http.Request) {
	var reposParams HtmlEditorRepositionParams
	if err := api.DecodeJSONRequest(r, &reposParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHtmlEditorPropUpdate(w, r, reposParams)
}
