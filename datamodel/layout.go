package datamodel

import (
	"appengine"
	"appengine/datastore"
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
	// ContainerID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	ContainerID    string `json:"containerID" datastore:"-"`    // don't save to datastore
	ParentLayoutID string `json:"parentLayoutID" datastore:"-"` // don't save to datastore
	PositionTop    int    `json:"positionTop"`
	PositionLeft   int    `json:"positionLeft"`
	SizeWidth      int    `json:"sizeWidth"`
	SizeHeight     int    `json:"sizeHeight"`
}

func NewUninitializedLayoutContainerParams() LayoutContainerParams {
	// Use -1 for top and left, so a failure of a client to initialize
	// can be detected.
	return LayoutContainerParams{"", "", -1, -1, -1, -1}
}

func NewLayoutContainer(appEngContext appengine.Context, containerParams LayoutContainerParams) (string, error) {

	if containerParams.PositionTop < 0 || containerParams.PositionLeft < 0 ||
		containerParams.SizeWidth <= 0 || containerParams.SizeHeight <= 0 {
		return "", fmt.Errorf("Invalid layout container parameters: %+v", containerParams)
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

	log.Printf("INFO: API: NewLayout: Created new Layout container: id=%v params=%+v",
		containerID, containerParams)

	return containerID, nil

}

func ResizeLayoutContainer(appEngContext appengine.Context, resizeParams LayoutContainerParams) error {

	if resizeParams.PositionTop < 0 || resizeParams.PositionLeft < 0 ||
		resizeParams.SizeWidth <= 0 || resizeParams.SizeHeight <= 0 {
		return fmt.Errorf("Invalid layout container resize parameters: %+v", resizeParams)
	}

	parentLayoutKey, err := getExistingRootEntityKey(appEngContext, layoutEntityKind,
		resizeParams.ParentLayoutID)
	if err != nil {
		return err
	}

	if updateErr := updateExistingEntity(appEngContext,
		resizeParams.ContainerID, layoutContainerEntityKind,
		parentLayoutKey, &resizeParams); updateErr != nil {
		return updateErr
	}

	return nil

}

func GetLayoutContainers(appEngContext appengine.Context, parentLayoutID string) ([]LayoutContainerParams, error) {

	parentLayoutKey, err := getExistingRootEntityKey(appEngContext, layoutEntityKind,
		parentLayoutID)
	if err != nil {
		return nil, err
	}

	containerQuery := datastore.NewQuery(layoutContainerEntityKind).Ancestor(parentLayoutKey)
	var layoutContainers []LayoutContainerParams
	keys, err := containerQuery.GetAll(appEngContext, &layoutContainers)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve layout containers: layout id=%v key=%+v", parentLayoutID, parentLayoutKey)
	} else {
		layoutContainersWithIDs := make([]LayoutContainerParams, len(layoutContainers))
		for i, c := range layoutContainers {
			containerKey := keys[i]
			containerID, encodeErr := encodeUniqueEntityIDToStr(containerKey)
			if encodeErr != nil {
				return nil, fmt.Errorf("Failed to encode unique ID for layout container: key=%+v, encode err=%v", containerKey, encodeErr)
			}
			c.ParentLayoutID = parentLayoutID
			c.ContainerID = containerID
			layoutContainersWithIDs[i] = c
		}
		return layoutContainersWithIDs, nil
	}

}
