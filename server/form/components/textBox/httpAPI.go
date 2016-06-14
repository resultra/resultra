package textBox

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	textBoxRouter := mux.NewRouter()

	textBoxRouter.HandleFunc("/api/frm/textBox/new", newTextBox)
	textBoxRouter.HandleFunc("/api/frm/textBox/resize", resizeTextBox)
	textBoxRouter.HandleFunc("/api/frm/textBox/reposition", repositionTextBox)

	http.Handle("/api/frm/textBox/", textBoxRouter)
}

func newTextBox(w http.ResponseWriter, r *http.Request) {

	textBoxParams := NewTextBoxParams{}
	if err := api.DecodeJSONRequest(r, &textBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if textBoxRef, err := saveNewTextBox(textBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *textBoxRef)
	}

}

func processTextBoxPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TextBoxPropUpdater) {
	if textBoxRef, err := updateTextBoxProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, textBoxRef)
	}
}

func resizeTextBox(w http.ResponseWriter, r *http.Request) {
	var resizeParams TextBoxResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, resizeParams)
}

func repositionTextBox(w http.ResponseWriter, r *http.Request) {
	var reposParams TextBoxRepositionParams
	if err := api.DecodeJSONRequest(r, &reposParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, reposParams)
}
