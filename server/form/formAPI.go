package form

import (
	"appengine"
	"net/http"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/textBox"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

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

type FormInfo struct {
	TextBoxes  []textBox.TextBoxRef   `json:"textBoxes"`
	CheckBoxes []checkBox.CheckBoxRef `json:"checkBoxes"`
}

func getFormInfo(w http.ResponseWriter, r *http.Request) {

	var parentFormID datastoreWrapper.UniqueRootIDHeader
	if err := api.DecodeJSONRequest(r, &parentFormID); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)

	textBoxRefs, err := textBox.GetTextBoxes(appEngCntxt, parentFormID.UniqueID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	checkBoxRefs, err := checkBox.GetCheckBoxes(appEngCntxt, parentFormID.UniqueID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	formInfoInfo := FormInfo{
		TextBoxes:  textBoxRefs,
		CheckBoxes: checkBoxRefs}

	api.WriteJSONResponse(w, formInfoInfo)

}
