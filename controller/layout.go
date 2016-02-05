package controller

import (
	"appengine"
	"errors"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newLayout(w http.ResponseWriter, r *http.Request) {

	var layoutParam map[string]string
	if err := decodeJSONRequest(r, &layoutParam); err != nil {
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

	if len(containerParams.PlaceholderID) == 0 {
		writeErrorResponse(w, errors.New("ERROR: API: newLayoutContainer: Missing placeholder ID in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if containerID, err := datamodel.NewLayoutContainer(appEngCntxt, containerParams); err != nil {
		writeErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{
			"layoutContainerID": containerID,
			"placeholderID":     containerParams.PlaceholderID})
	}

}
