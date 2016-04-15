package textBox

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/common/datastoreWrapper"
	"resultra/datasheet/server/dataModel"
	"resultra/datasheet/server/field"
)

const layoutContainerEntityKind string = "LayoutContainer"
const textBoxEntityKind string = "TextBox"

// A LayoutContainer represents what is actually stored in the datastore
// for each layout container.
type LayoutContainer struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

// NEW CODE for TEXT BOX
type UniqueID struct {
	ParentID string `json:"uniqueID"`
	ObjectID string `json:"objectID"`
}

type UniqueIDHeader struct {
	UniqueID UniqueID `json:"uniqueID"`
}

type TextBox struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

type TextBoxRef struct {
	UniqueIDHeader
	FieldRef field.FieldRef
	Geometry common.LayoutGeometry
}

// LayoutContainerParams is a parameter block used to create new and complete
// LayoutContainers, or to retrieve them from the datastore. The ID's
// are the encoded/stringified datastore keys, suitable for passing back
// and forth to clients of this package.
type LayoutContainerParams struct {
	// ContainerID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	ParentLayoutID string                `json:"parentLayoutID" datastore:"-"` // don't save to datastore
	ContainerID    string                `json:"containerID" datastore:"-"`    // don't save to datastore
	FieldID        string                `json:"fieldID" datastore:"-"`
	Geometry       common.LayoutGeometry `json:"geometry"`
}

func NewUninitializedLayoutContainerParams() LayoutContainerParams {
	// Use -1 for top and left, so a failure of a client to initialize
	// can be detected.
	return LayoutContainerParams{"", "", "", common.NewUnitializedLayoutGeometry()}
}

// ResizeContainerParams has a subset of LayoutContainer properties
// which are modified when resizing the container (fieldID is absent)
type ResizeContainerParams struct {
	ParentLayoutID string `json:"parentLayoutID" datastore:"-"`
	ContainerID    string `json:"containerID" datastore:"-"`
	Geometry       common.LayoutGeometry
}

func NewUninitializedResizeLayoutContainerParams() ResizeContainerParams {
	// Use -1 for top and left, so a failure of a client to initialize
	// can be detected.
	return ResizeContainerParams{"", "", common.NewUnitializedLayoutGeometry()}
}

func NewLayoutContainer(appEngContext appengine.Context, containerParams LayoutContainerParams) (string, error) {

	if !common.ValidGeometry(containerParams.Geometry) {
		return "", fmt.Errorf("Invalid layout container parameters: %+v", containerParams)
	}

	parentLayoutKey, err := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.LayoutEntityKind,
		containerParams.ParentLayoutID)
	if err != nil {
		return "", err
	}

	fieldKey, fieldErr := field.GetExistingFieldKey(appEngContext, containerParams.FieldID)
	if fieldErr != nil {
		return "", fmt.Errorf("NewLayoutContainer: Can't create layout container with field ID = '%v': datastore error=%v",
			containerParams.FieldID, fieldErr)
	}

	newLayoutContainer := LayoutContainer{fieldKey, containerParams.Geometry}

	containerID, insertErr := datastoreWrapper.InsertNewEntity(appEngContext, layoutContainerEntityKind,
		parentLayoutKey, &newLayoutContainer)
	if insertErr != nil {
		return "", insertErr
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: id=%v params=%+v",
		containerID, containerParams)

	return containerID, nil

}

// NEW CODE for TEXT BOX

type NewTextBoxParams struct {
	// ContainerID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	ParentID string                `json:"parentID"`
	FieldID  string                `json:"fieldID"`
	Geometry common.LayoutGeometry `json:"geometry"`
}

func NewTextBox(appEngContext appengine.Context, params NewTextBoxParams) (*TextBoxRef, error) {
	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	parentKey, err := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.LayoutEntityKind,
		params.ParentID)
	if err != nil {
		return nil, err
	}

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, field.GetFieldParams{params.FieldID})
	if fieldErr != nil {
		return nil, fmt.Errorf("NewTextBox: Can't text box with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	/*		if (fieldRef.FieldInfo.type == field.FieldTypeText) || (fieldRef.FieldInfo.type == field.FieldTypeNumber) {
				return "", fmt.Errorf("NewTextBox: Can't create text box - incompatible field type")
			}
	*/
	newTextBox := TextBox{Field: fieldKey, Geometry: params.Geometry}

	textBoxID, insertErr := datastoreWrapper.InsertNewEntity(appEngContext, textBoxEntityKind,
		parentKey, &newTextBox)
	if insertErr != nil {
		return nil, insertErr
	}

	textBoxUniqueID := UniqueID{params.ParentID, textBoxID}
	textBoxRef := TextBoxRef{UniqueIDHeader{UniqueID: textBoxUniqueID}, *fieldRef, params.Geometry}

	log.Printf("INFO: API: NewLayout: Created new Layout container: id=%v params=%+v", textBoxID, params)

	return &textBoxRef, nil

}

func ResizeLayoutContainer(appEngContext appengine.Context, resizeParams ResizeContainerParams) error {

	if !common.ValidGeometry(resizeParams.Geometry) {
		return fmt.Errorf("Invalid layout container resize parameters: %+v", resizeParams)
	}

	parentLayoutKey, err := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.LayoutEntityKind,
		resizeParams.ParentLayoutID)
	if err != nil {
		return err
	}

	// Retrieve the entire LayoutContainer, but overwrite just the Geometry property.
	var layoutContainerForUpdate LayoutContainer
	getErr := datastoreWrapper.GetChildEntityByID(resizeParams.ContainerID, appEngContext,
		layoutContainerEntityKind,
		parentLayoutKey, &layoutContainerForUpdate)
	if getErr != nil {
		return fmt.Errorf("Can't resize container: Error retrieving existing container for update: %v", getErr)
	}

	// Update the geometry properties
	layoutContainerForUpdate.Geometry = resizeParams.Geometry

	if updateErr := datastoreWrapper.UpdateExistingEntity(appEngContext,
		resizeParams.ContainerID, layoutContainerEntityKind,
		parentLayoutKey, &layoutContainerForUpdate); updateErr != nil {
		return updateErr
	}

	return nil

}

func GetLayoutContainers(appEngContext appengine.Context, parentLayoutID string) ([]LayoutContainerParams, error) {

	parentLayoutKey, err := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.LayoutEntityKind,
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
			containerID, encodeErr := datastoreWrapper.EncodeUniqueEntityIDToStr(containerKey)
			if encodeErr != nil {
				return nil, fmt.Errorf("Failed to encode unique ID for layout container: key=%+v, encode err=%v",
					containerKey, encodeErr)
			}

			fieldID, fieldIDEncodeErr := datastoreWrapper.EncodeUniqueEntityIDToStr(c.Field)
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
