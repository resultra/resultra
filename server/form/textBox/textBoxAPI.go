package textBox

import (
	"appengine"
	"errors"
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/api"
)

func newLayoutContainer(w http.ResponseWriter, r *http.Request) {

	containerParams := NewUninitializedLayoutContainerParams()
	if err := api.DecodeJSONRequest(r, &containerParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if len(containerParams.ContainerID) == 0 {
		api.WriteErrorResponse(w, errors.New("ERROR: API: newLayoutContainer: Missing placeholder ID in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if containerID, err := NewLayoutContainer(appEngCntxt, containerParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, api.JSONParams{
			"layoutContainerID": containerID,
			"placeholderID":     containerParams.ContainerID})
	}

}

func resizeLayoutContainer(w http.ResponseWriter, r *http.Request) {

	resizeParams := NewUninitializedResizeLayoutContainerParams()
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if len(resizeParams.ContainerID) == 0 {
		api.WriteErrorResponse(w, errors.New("ERROR: API: newLayoutContainer: Missing container ID in request"))
		return
	}

	appEngCntxt := appengine.NewContext(r)

	if err := ResizeLayoutContainer(appEngCntxt, resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, api.JSONParams{})
	}

}

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

func GetLayoutContainersFromRequest(r *http.Request) ([]LayoutContainerParams, error) {

	layoutID, err := getLayoutIDFromRequestParams(r)
	if err != nil {
		return nil, err
	}

	appEngCntxt := appengine.NewContext(r)

	if layoutContainers, err := GetLayoutContainers(appEngCntxt, layoutID); err != nil {
		return nil, err
	} else {
		return layoutContainers, nil
	}

}

func getLayoutContainers(w http.ResponseWriter, r *http.Request) {

	if layoutContainers, err := GetLayoutContainersFromRequest(r); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, layoutContainers)
	}

}
