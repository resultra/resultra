package selection

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	selectionRouter := mux.NewRouter()

	selectionRouter.HandleFunc("/api/frm/selection/new", newSelection)
	selectionRouter.HandleFunc("/api/frm/selection/resize", resizeSelection)
	selectionRouter.HandleFunc("/api/frm/selection/setSelectableVals", setSelectionSelectableVals)
	selectionRouter.HandleFunc("/api/frm/selection/setLabelFormat", setLabelFormat)
	selectionRouter.HandleFunc("/api/frm/selection/setVisibility", setVisibility)

	http.Handle("/api/frm/selection/", selectionRouter)
}

func newSelection(w http.ResponseWriter, r *http.Request) {

	selectionParams := NewSelectionParams{}
	if err := api.DecodeJSONRequest(r, &selectionParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if selectionRef, err := saveNewSelection(selectionParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *selectionRef)
	}

}

func processSelectionPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SelectionPropUpdater) {
	if selectionRef, err := updateSelectionProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, selectionRef)
	}
}

func resizeSelection(w http.ResponseWriter, r *http.Request) {
	var resizeParams SelectionResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, resizeParams)
}

func setSelectionSelectableVals(w http.ResponseWriter, r *http.Request) {
	var params SelectionSelectableValsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)

}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params SelectionLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)

}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params SelectionVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)

}
