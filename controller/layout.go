package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newLayout(w http.ResponseWriter, r *http.Request) {

	var layoutParam map[string]string
	if err := decodeJSONRequest(r, layoutParam); err != nil {
		writeErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if layoutID, err := datamodel.NewLayout(appEngCntxt, layoutParam["name"]); err != nil {
		writeErrorResponse(w, err)
	} else {
		writeJSONResponse(w, map[string]string{"layoutID": layoutID})
	}

}

func newLayoutContainer(w http.ResponseWriter, r *http.Request) {

	containerParams := datamodel.NewUninitializedLayoutContainerParams()
	if err := decodeJSONRequest(r, &containerParams); err != nil {
		writeErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if containerID, err := datamodel.NewLayoutContainer(appEngCntxt, containerParams); err != nil {
		writeErrorResponse(w, err)
	} else {
		writeJSONResponse(w, map[string]string{"layoutContainerID": containerID})
	}

}
