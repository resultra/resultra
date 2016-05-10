package form

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/datePicker"
	"resultra/datasheet/server/form/components/textBox"
	"resultra/datasheet/server/generic/api"
)

func init() {

	formRouter := mux.NewRouter()

	formRouter.HandleFunc("/api/frm/getFormInfo", getFormInfo)
	formRouter.HandleFunc("/api/frm/new", newFormAPI)
	formRouter.HandleFunc("/api/frm/get", getFormAPI)

	http.Handle("/api/frm/", formRouter)
}

func newFormAPI(w http.ResponseWriter, r *http.Request) {

	var params NewFormParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if formRef, err := newForm(appEngCntxt, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formRef)
	}

}

func getFormAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFormParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if formRef, err := GetFormRef(appEngCntxt, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formRef)
	}

}

type FormInfo struct {
	TextBoxes   []textBox.TextBoxRef       `json:"textBoxes"`
	CheckBoxes  []checkBox.CheckBoxRef     `json:"checkBoxes"`
	DatePickers []datePicker.DatePickerRef `json:"datePickers"`
}

type GetFormInfoParams struct {
	FormID string `json:"formID"`
}

func getFormInfo(w http.ResponseWriter, r *http.Request) {

	var params GetFormInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)

	textBoxRefs, err := textBox.GetTextBoxes(appEngCntxt, params.FormID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	checkBoxRefs, err := checkBox.GetCheckBoxes(appEngCntxt, params.FormID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	datePickerRefs, err := datePicker.GetDatePickers(appEngCntxt, params.FormID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	formInfoInfo := FormInfo{
		TextBoxes:   textBoxRefs,
		CheckBoxes:  checkBoxRefs,
		DatePickers: datePickerRefs}

	api.WriteJSONResponse(w, formInfoInfo)

}
