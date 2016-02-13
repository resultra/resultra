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
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if layoutID, err := datamodel.NewLayout(appEngCntxt, layoutParam["name"]); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{"layoutID": layoutID})
	}

}

func newLayoutContainer(w http.ResponseWriter, r *http.Request) {

	containerParams := datamodel.NewUninitializedLayoutContainerParams()
	if err := decodeJSONRequest(r, &containerParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	if len(containerParams.ContainerID) == 0 {
		WriteErrorResponse(w, errors.New("ERROR: API: newLayoutContainer: Missing placeholder ID in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if containerID, err := datamodel.NewLayoutContainer(appEngCntxt, containerParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{
			"layoutContainerID": containerID,
			"placeholderID":     containerParams.ContainerID})
	}

}

func resizeLayoutContainer(w http.ResponseWriter, r *http.Request) {

	resizeParams := datamodel.NewUninitializedResizeLayoutContainerParams()
	if err := decodeJSONRequest(r, &resizeParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	if len(resizeParams.ContainerID) == 0 {
		WriteErrorResponse(w, errors.New("ERROR: API: newLayoutContainer: Missing container ID in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)

	if err := datamodel.ResizeLayoutContainer(appEngCntxt, resizeParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{})
	}

}

func getLayoutIDFromRequestParams(r *http.Request) (string, error) {
	var jsonParams JSONParams
	if err := decodeJSONRequest(r, &jsonParams); err != nil {
		return "", err
	}

	layoutID, found := jsonParams["layoutID"]
	if found != true || len(layoutID) == 0 {
		return "", fmt.Errorf("Missing layoutID parameter in request")
	}

	return layoutID, nil
}

func getLayoutContainersFromRequest(r *http.Request) ([]datamodel.LayoutContainerParams, error) {

	layoutID, err := getLayoutIDFromRequestParams(r)
	if err != nil {
		return nil, err
	}

	appEngCntxt := appengine.NewContext(r)

	if layoutContainers, err := datamodel.GetLayoutContainers(appEngCntxt, layoutID); err != nil {
		return nil, err
	} else {
		return layoutContainers, nil
	}

}

func getLayoutContainers(w http.ResponseWriter, r *http.Request) {

	if layoutContainers, err := getLayoutContainersFromRequest(r); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, layoutContainers)
	}

}

type LayoutEditInfo struct {
	LayoutContainers []datamodel.LayoutContainerParams `json:"layoutContainers"`
	FieldsByType     datamodel.FieldsByType            `json:"fieldsByType"`
}

func getLayoutEditInfo(w http.ResponseWriter, r *http.Request) {

	layoutContainers, err := getLayoutContainersFromRequest(r)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	fieldsByType, err := datamodel.GetFieldsByType(appEngCntxt)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	layoutEditInfo := LayoutEditInfo{layoutContainers, *fieldsByType}

	writeJSONResponse(w, layoutEditInfo)

}
