package textBox

import (
	"appengine"
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/api"
)

func getLayoutIDFromRequestParams(r *http.Request) (string, error) {
	var jsonParams api.JSONParams
	if err := api.DecodeJSONRequest(r, &jsonParams); err != nil {
		return "", err
	}

	layoutID, found := jsonParams["layoutID"]
	if found != true || len(layoutID) == 0 {
		return "", fmt.Errorf("Missing layoutID parameter in request")
	}

	return layoutID, nil
}

func newTextBox(w http.ResponseWriter, r *http.Request) {

	textBoxParams := NewTextBoxParams{}
	if err := api.DecodeJSONRequest(r, &textBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if textBoxRef, err := NewTextBox(appEngCntxt, textBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *textBoxRef)
	}

}

func processTextBoxPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TextBoxPropUpdater) {
	appEngCntxt := appengine.NewContext(r)
	if textBoxRef, err := updateTextBoxProps(appEngCntxt, propUpdater); err != nil {
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
