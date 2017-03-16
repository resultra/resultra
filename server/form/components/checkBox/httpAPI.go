package checkBox

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	checkBoxRouter := mux.NewRouter()

	checkBoxRouter.HandleFunc("/api/frm/checkBox/new", newCheckBox)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/resize", resizeCheckBox)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setColorScheme", setColorScheme)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setStrikethrough", setStrikethrough)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/setLabelFormat", setLabelFormat)

	http.Handle("/api/frm/checkBox/", checkBoxRouter)
}

func newCheckBox(w http.ResponseWriter, r *http.Request) {

	checkBoxParams := NewCheckBoxParams{}
	if err := api.DecodeJSONRequest(r, &checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if checkBoxRef, err := saveNewCheckBox(checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *checkBoxRef)
	}

}

func processCheckBoxPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater CheckBoxPropUpdater) {
	if checkBoxRef, err := updateCheckBoxProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, checkBoxRef)
	}
}

func resizeCheckBox(w http.ResponseWriter, r *http.Request) {
	var resizeParams CheckBoxResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, resizeParams)
}

func setColorScheme(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setStrikethrough(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxStrikethroughParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}
