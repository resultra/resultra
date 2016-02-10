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

type LayoutGeometry struct {
	PositionTop  int `json:"positionTop"`
	PositionLeft int `json:"positionLeft"`
	SizeWidth    int `json:"sizeWidth"`
	SizeHeight   int `json:"sizeHeight"`
}

func NewUnitializedLayoutGeometry() LayoutGeometry {
	return LayoutGeometry{-1, -1, -1, -1}
}

func validGeometry(geom LayoutGeometry) bool {
	if (geom.PositionTop >= 0) && (geom.PositionLeft >= 0) &&
		(geom.SizeWidth > 0) && (geom.SizeHeight > 0) {
		return true
	} else {
		return false
	}
}

// A LayoutContainer represents what is actually stored in the datastore
// for each layout container.
type LayoutContainer struct {
	Field    *datastore.Key
	Geometry LayoutGeometry
}

// LayoutContainerParams is a parameter block used to create new and complete
// LayoutContainers, or to retrieve them from the datastore. The ID's
// are the encoded/stringified datastore keys, suitable for passing back
// and forth to clients of this package.
type LayoutContainerParams struct {
	// ContainerID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	ParentLayoutID string         `json:"parentLayoutID" datastore:"-"` // don't save to datastore
	ContainerID    string         `json:"containerID" datastore:"-"`    // don't save to datastore
	FieldID        string         `json:"fieldID" datastore:"-"`
	Geometry       LayoutGeometry `json:"geometry"`
}

func NewUninitializedLayoutContainerParams() LayoutContainerParams {
	// Use -1 for top and left, so a failure of a client to initialize
	// can be detected.
	return LayoutContainerParams{"", "", "", NewUnitializedLayoutGeometry()}
}

// ResizeContainerParams has a subset of LayoutContainer properties
// which are modified when resizing the container (fieldID is absent)
type ResizeContainerParams struct {
	ParentLayoutID string `json:"parentLayoutID" datastore:"-"`
	ContainerID    string `json:"containerID" datastore:"-"`
	Geometry       LayoutGeometry
}

func NewUninitializedResizeLayoutContainerParams() ResizeContainerParams {
	// Use -1 for top and left, so a failure of a client to initialize
	// can be detected.
	return ResizeContainerParams{"", "", NewUnitializedLayoutGeometry()}
}

func NewLayoutContainer(appEngContext appengine.Context, containerParams LayoutContainerParams) (string, error) {

	if !validGeometry(containerParams.Geometry) {
		return "", fmt.Errorf("Invalid layout container parameters: %+v", containerParams)
	}

	parentLayoutKey, err := getExistingRootEntityKey(appEngContext, layoutEntityKind,
		containerParams.ParentLayoutID)
	if err != nil {
		return "", err
	}

	fieldKey, fieldErr := getExistingRootEntityKey(appEngContext, fieldEntityKind, containerParams.FieldID)
	if err != nil {
		return "", fmt.Errorf("Can't create layout container: invalid field ID '%v': datastore error=%v",
			containerParams.FieldID, fieldErr)
	}

	newLayoutContainer := LayoutContainer{fieldKey, containerParams.Geometry}

	containerID, insertErr := insertNewEntity(appEngContext, layoutContainerEntityKind,
		parentLayoutKey, &newLayoutContainer)
	if insertErr != nil {
		return "", insertErr
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: id=%v params=%+v",
		containerID, containerParams)

	return containerID, nil

}

func ResizeLayoutContainer(appEngContext appengine.Context, resizeParams ResizeContainerParams) error {

	if !validGeometry(resizeParams.Geometry) {
		return fmt.Errorf("Invalid layout container resize parameters: %+v", resizeParams)
	}

	parentLayoutKey, err := getExistingRootEntityKey(appEngContext, layoutEntityKind,
		resizeParams.ParentLayoutID)
	if err != nil {
		return err
	}

	// Retrieve the entire LayoutContainer, but overwrite just the Geometry property.
	var layoutContainerForUpdate LayoutContainer
	getErr := getChildEntityByID(resizeParams.ContainerID, appEngContext,
		layoutContainerEntityKind,
		parentLayoutKey, &layoutContainerForUpdate)
	if getErr != nil {
		return fmt.Errorf("Can't resize container: Error retrieving existing container for update: %v", getErr)
	}

	// Update the geometry properties
	layoutContainerForUpdate.Geometry = resizeParams.Geometry

	if updateErr := updateExistingEntity(appEngContext,
		resizeParams.ContainerID, layoutContainerEntityKind,
		parentLayoutKey, &layoutContainerForUpdate); updateErr != nil {
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

	// Retrieve the raw/datastore representation into LayoutContainer's then
	// build up corresponding LayoutContainerParams including the encoded/stringified
	// IDs (from the datastore keys), so clients of this package can make reference to
	// the layout container and it's references without exposing the datastore internals.
	containerQuery := datastore.NewQuery(layoutContainerEntityKind).Ancestor(parentLayoutKey)
	var layoutContainers []LayoutContainer
	keys, err := containerQuery.GetAll(appEngContext, &layoutContainers)

	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve layout containers: layout id=%v key=%+v", parentLayoutID, parentLayoutKey)
	} else {
		layoutContainersWithIDs := make([]LayoutContainerParams, len(layoutContainers))
		for i, c := range layoutContainers {

			containerKey := keys[i]
			containerID, encodeErr := encodeUniqueEntityIDToStr(containerKey)
			if encodeErr != nil {
				return nil, fmt.Errorf("Failed to encode unique ID for layout container: key=%+v, encode err=%v",
					containerKey, encodeErr)
			}

			fieldID, fieldIDEncodeErr := encodeUniqueEntityIDToStr(c.Field)
			if fieldIDEncodeErr != nil {
				return nil, fmt.Errorf("Failed to encode unique ID for layout container's field:  key=%+v, encode err=%v",
					c.Field, fieldIDEncodeErr)
			}

			containerParams := LayoutContainerParams{
				ParentLayoutID: parentLayoutID,
				ContainerID:    containerID,
				FieldID:        fieldID,
				Geometry:       c.Geometry}

			layoutContainersWithIDs[i] = containerParams
		}
		return layoutContainersWithIDs, nil
	}

}
