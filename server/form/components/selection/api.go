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
