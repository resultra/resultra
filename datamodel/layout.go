package datamodel

import (
	"appengine"
	"log"
)

const layoutEntityKind string = "Layout"
const layoutContainerEntityKind string = "LayoutContainer"

type Layout struct {
	Name string `json:"name"`
}

func NewLayout(appEngContext appengine.Context, layoutName string) (string, error) {

	sanitizedLayoutName, sanitizeErr := sanitizeName(layoutName)
	if sanitizeErr != nil {
		return "", sanitizeErr
	}

	var newLayout = Layout{sanitizedLayoutName}
	layoutID, insertErr := insertNewEntity(appEngContext, layoutEntityKind, &newLayout)
	if insertErr != nil {
		return "", insertErr
	}

	// TODO - verify IntID != 0
	log.Printf("NewLayout: Created new Layout: id= %v, name='%v'", layoutID, sanitizedLayoutName)

	return layoutID, nil

}

type LayoutContainerParams struct {
	ParentLayoutID string `json:"parentLayoutID" datastore:"-"` // don't save to datastore
	PositionTop    int    `json:"positionTop"`
	PositionLeft   int    `json:"positionLeft"`
}

func NewUninitializedLayoutContainerParams() LayoutContainerParams {
	// Use -1 for top and left, so a failure of a client to initialize
	// can be detected.
	return LayoutContainerParams{"", -1, -1}
}

func NewLayoutContainer(appEngContext appengine.Context, containerParams LayoutContainerParams) (string, error) {

	var parentLayout = Layout{}
	if err := getEntityByID(containerParams.ParentLayoutID, appEngContext,
		layoutEntityKind, &parentLayout); err != nil {
		return "", err
	}

	containerID, insertErr := insertNewEntity(appEngContext, layoutContainerEntityKind, &containerParams)
	if insertErr != nil {
		return "", insertErr
	}

	log.Printf("NewLayout: Created new Layout container: id=%v", containerID)

	return containerID, nil

}
