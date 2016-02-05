package datamodel

import (
	"appengine"
	"fmt"
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
	layoutID, insertErr := insertNewEntity(appEngContext, layoutEntityKind, nil, &newLayout)
	if insertErr != nil {
		return "", insertErr
	}

	// TODO - verify IntID != 0
	log.Printf("NewLayout: Created new Layout: id= %v, name='%v'", layoutID, sanitizedLayoutName)

	return layoutID, nil

}

type LayoutContainerParams struct {
	// PlaceholderID is a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	PlaceholderID  string `json:"placeholderID" datastore:"-"`  // don't save to datastore
	ParentLayoutID string `json:"parentLayoutID" datastore:"-"` // don't save to datastore
	PositionTop    int    `json:"positionTop"`
	PositionLeft   int    `json:"positionLeft"`
}

func NewUninitializedLayoutContainerParams() LayoutContainerParams {
	// Use -1 for top and left, so a failure of a client to initialize
	// can be detected.
	return LayoutContainerParams{"", "", -1, -1}
}

func NewLayoutContainer(appEngContext appengine.Context, containerParams LayoutContainerParams) (string, error) {

	if containerParams.PositionTop < 0 || containerParams.PositionLeft < 0 {
		return "", fmt.Errorf("Invalid layout container position: top=%v, left=%v",
			containerParams.PositionTop, containerParams.PositionLeft)
	}

	parentLayoutKey, err := getExistingRootEntityKey(appEngContext, layoutEntityKind,
		containerParams.ParentLayoutID)
	if err != nil {
		return "", err
	}

	containerID, insertErr := insertNewEntity(appEngContext, layoutContainerEntityKind,
		parentLayoutKey, &containerParams)
	if insertErr != nil {
		return "", insertErr
	}

	log.Printf("NewLayout: Created new Layout container: id=%v", containerID)

	return containerID, nil

}
