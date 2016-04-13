package form

import (
	"appengine"
	"net/http"
	"resultra/datasheet/server/common/api"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/textBox"
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

type LayoutEditInfo struct {
	LayoutContainers []textBox.LayoutContainerParams `json:"layoutContainers"`
	FieldsByType     field.FieldsByType              `json:"fieldsByType"`
}

func getLayoutEditInfo(w http.ResponseWriter, r *http.Request) {

	layoutContainers, err := textBox.GetLayoutContainersFromRequest(r)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	fieldsByType, err := field.GetFieldsByType(appEngCntxt)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	layoutEditInfo := LayoutEditInfo{layoutContainers, *fieldsByType}

	api.WriteJSONResponse(w, layoutEditInfo)

}
