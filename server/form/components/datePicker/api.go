package datePicker

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	datePickerRouter := mux.NewRouter()

	datePickerRouter.HandleFunc("/api/frm/datePicker/new", newDatePicker)
	datePickerRouter.HandleFunc("/api/frm/datePicker/resize", resizeDatePicker)

	http.Handle("/api/frm/datePicker/", datePickerRouter)
}

func newDatePicker(w http.ResponseWriter, r *http.Request) {

	params := NewDatePickerParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if checkBoxRef, err := saveNewDatePicker(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *checkBoxRef)
	}

}

func processDatePickerPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater DatePickerPropUpdater) {
	if checkBoxRef, err := updateDatePickerProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, checkBoxRef)
	}
}

func resizeDatePicker(w http.ResponseWriter, r *http.Request) {
	var resizeParams DatePickerResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, resizeParams)
}
