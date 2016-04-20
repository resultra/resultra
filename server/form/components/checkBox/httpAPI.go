package checkBox

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	checkBoxRouter := mux.NewRouter()

	checkBoxRouter.HandleFunc("/api/frm/checkBox/new", newCheckBox)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/resize", resizeCheckBox)
	checkBoxRouter.HandleFunc("/api/frm/checkBox/reposition", repositionCheckBox)

	http.Handle("/api/frm/checkBox/", checkBoxRouter)
}

func newCheckBox(w http.ResponseWriter, r *http.Request) {

	checkBoxParams := NewCheckBoxParams{}
	if err := api.DecodeJSONRequest(r, &checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if checkBoxRef, err := saveNewCheckBox(appEngCntxt, checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *checkBoxRef)
	}

}

func processCheckBoxPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater CheckBoxPropUpdater) {
	appEngCntxt := appengine.NewContext(r)
	if checkBoxRef, err := updateCheckBoxProps(appEngCntxt, propUpdater); err != nil {
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

func repositionCheckBox(w http.ResponseWriter, r *http.Request) {
	var reposParams CheckBoxRepositionParams
	if err := api.DecodeJSONRequest(r, &reposParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, reposParams)
}
