package controller

import (
	"appengine"
	"errors"
	"fmt"
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
		writeJSONResponse(w, JSONParams{"layoutID": layoutID})
	}

}

func newLayoutContainer(w http.ResponseWriter, r *http.Request) {

	containerParams := datamodel.NewUninitializedLayoutContainerParams()
	if err := decodeJSONRequest(r, &containerParams); err != nil {
		writeErrorResponse(w, err)
		return
	}

	if len(containerParams.ContainerID) == 0 {
		writeErrorResponse(w, errors.New("ERROR: API: newLayoutContainer: Missing placeholder ID in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if containerID, err := datamodel.NewLayoutContainer(appEngCntxt, containerParams); err != nil {
		writeErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{
			"layoutContainerID": containerID,
			"placeholderID":     containerParams.ContainerID})
	}

}

func resizeLayoutContainer(w http.ResponseWriter, r *http.Request) {

	resizeParams := datamodel.NewUninitializedLayoutContainerParams()
	if err := decodeJSONRequest(r, &resizeParams); err != nil {
		writeErrorResponse(w, err)
		return
	}

	if len(resizeParams.ContainerID) == 0 {
		writeErrorResponse(w, errors.New("ERROR: API: newLayoutContainer: Missing placeholder ID in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)

	if err := datamodel.ResizeLayoutContainer(appEngCntxt, resizeParams); err != nil {
		writeErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{})
	}

}

func getLayoutContainers(w http.ResponseWriter, r *http.Request) {

	var jsonParams JSONParams
	if err := decodeJSONRequest(r, &jsonParams); err != nil {
		writeErrorResponse(w, err)
		return
	}

	layoutID, found := jsonParams["layoutID"]
	if found != true || len(layoutID) == 0 {
		writeErrorResponse(w, fmt.Errorf("Missing layoutID parameter in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)

	if layoutContainers, err := datamodel.GetLayoutContainers(appEngCntxt, layoutID); err != nil {
		writeErrorResponse(w, err)
	} else {
		writeJSONResponse(w, layoutContainers)
	}

}
