package image

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const imageEntityKind string = "Image"

type Image struct {
	ParentFormID string                `json:"parentFormID"`
	ImageID      string                `json:"imageID"`
	FieldID      string                `json:"fieldID"`
	Geometry     common.LayoutGeometry `json:"geometry"`
}

const imageParentFormIDFieldName string = "ParentFormID"
const imageIDFieldName string = "ImageID"

type NewImageParams struct {
	ParentID           string                `json:"parentID"`
	FieldParentTableID string                `json:"fieldParentTableID"`
	FieldID            string                `json:"fieldID"`
	Geometry           common.LayoutGeometry `json:"geometry"`
}

func validImageFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeFile {
		return true
	} else {
		return false
	}
}

func saveNewImage(appEngContext appengine.Context, params NewImageParams) (*Image, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	field, fieldErr := field.GetField(appEngContext, params.FieldParentTableID, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validImageFieldType(field.Type) {
		return nil, fmt.Errorf("NewImage: Invalid field type: expecting file field, got %v", field.Type)
	}

	newImage := Image{ParentFormID: params.ParentID,
		FieldID:  params.FieldID,
		ImageID:  uniqueID.GenerateUniqueID(),
		Geometry: params.Geometry}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, imageEntityKind, &newImage)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new image component: error inserting into datastore: %v", insertErr)
	}

	log.Printf("INFO: API: saveNewImage: Created new image component: %+v", newImage)

	return &newImage, nil

}

func getImage(appEngContext appengine.Context, imageID string) (*Image, error) {

	var image Image
	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, imageEntityKind,
		imageIDFieldName, imageID, &image); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to image container from datastore: error = %v", getErr)
	}
	return &image, nil
}

func GetImages(appEngContext appengine.Context, parentFormID string) ([]Image, error) {

	var images []Image

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentFormID,
		imageEntityKind, imageParentFormIDFieldName, &images)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve layout containers: form id=%v", parentFormID)
	}

	return images, nil

}

func updateExistingImage(appEngContext appengine.Context, imageID string, updatedImage *Image) (*Image, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		imageID, imageEntityKind, imageIDFieldName, updatedImage); updateErr != nil {
		return nil, fmt.Errorf("updateExistingImage: Error updating existing image component: error = %v", updateErr)
	}

	return updatedImage, nil

}
