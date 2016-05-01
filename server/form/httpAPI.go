package form

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/textBox"
	"resultra/datasheet/server/generic/api"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	apiRouter.HandleFunc("/api/newLayout", newLayout)
}

func init() {

	formRouter := mux.NewRouter()

	formRouter.HandleFunc("/api/frm/getFormInfo", getFormInfo)
	formRouter.HandleFunc("/api/frm/new", newFormAPI)
	formRouter.HandleFunc("/api/frm/get", getFormAPI)

	http.Handle("/api/frm/", formRouter)
}

func newLayout(w http.ResponseWriter, r *http.Request) {

	var layoutParam map[string]string
	if err := api.DecodeJSONRequest(r, &layoutParam); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if layoutID, err := NewLayout(appEngCntxt, layoutParam["name"]); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, api.JSONParams{"layoutID": layoutID})
	}

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
	if formRef, err := getForm(appEngCntxt, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formRef)
	}

}

type FormInfo struct {
	TextBoxes  []textBox.TextBoxRef   `json:"textBoxes"`
	CheckBoxes []checkBox.CheckBoxRef `json:"checkBoxes"`
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

	formInfoInfo := FormInfo{
		TextBoxes:  textBoxRefs,
		CheckBoxes: checkBoxRefs}

	api.WriteJSONResponse(w, formInfoInfo)

}
