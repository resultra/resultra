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
	textBoxRouter.HandleFunc("/api/frm/textBox/setValueFormat", setValueFormat)
	textBoxRouter.HandleFunc("/api/frm/textBox/setLabelFormat", setLabelFormat)
	textBoxRouter.HandleFunc("/api/frm/textBox/setVisibility", setVisibility)
	textBoxRouter.HandleFunc("/api/frm/textBox/setPermissions", setPermissions)

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

func setValueFormat(w http.ResponseWriter, r *http.Request) {
	var params TextBoxValueFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params TextBoxLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params TextBoxVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TextBoxPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextBoxPropUpdate(w, r, params)
}
