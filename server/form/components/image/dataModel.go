package image

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const imageEntityKind string = "Image"

type Image struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

type ImageRef struct {
	ImageID  string                `json:"imageID"`
	FieldRef field.FieldRef        `json:"fieldRef"`
	Geometry common.LayoutGeometry `json:"geometry"`
}

type NewImageParams struct {
	ParentID string                `json:"parentID"`
	FieldID  string                `json:"fieldID"`
	Geometry common.LayoutGeometry `json:"geometry"`
}

func validImageFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeFile {
		return true
	} else {
		return false
	}
}

func saveNewImage(appEngContext appengine.Context, params NewImageParams) (*ImageRef, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validImageFieldType(fieldRef.FieldInfo.Type) {
		return nil, fmt.Errorf("NewImage: Invalid field type: expecting file field, got %v", fieldRef.FieldInfo.Type)
	}

	newImage := Image{Field: fieldKey, Geometry: params.Geometry}

	imageID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext, params.ParentID, imageEntityKind, &newImage)
	if insertErr != nil {
		return nil, insertErr
	}

	imageRef := ImageRef{
		ImageID:  imageID,
		FieldRef: *fieldRef,
		Geometry: params.Geometry}

	log.Printf("INFO: API: saveNewImage: Created new image container: id=%v params=%+v", imageID, params)

	return &imageRef, nil

}

func getImage(appEngContext appengine.Context, imageID string) (*Image, error) {

	var image Image
	if getErr := datastoreWrapper.GetEntity(appEngContext, imageID, &image); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to image container from datastore: error = %v", getErr)
	}
	return &image, nil
}

func GetImages(appEngContext appengine.Context, parentFormID string) ([]ImageRef, error) {

	var imagees []Image
	imageIDs, getErr := datastoreWrapper.GetAllChildEntities(appEngContext, parentFormID, imageEntityKind, &imagees)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve layout containers: form id=%v", parentFormID)
	}

	imageRefs := make([]ImageRef, len(imagees))
	for imageIter, currImage := range imagees {

		imageID := imageIDs[imageIter]

		fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, currImage.Field)
		if fieldErr != nil {
			return nil, fmt.Errorf("GetImagees: Error retrieving field for image: error = %v", fieldErr)
		}

		imageRefs[imageIter] = ImageRef{
			ImageID:  imageID,
			FieldRef: *fieldRef,
			Geometry: currImage.Geometry}

	} // for each image
	return imageRefs, nil

}

func updateExistingImage(appEngContext appengine.Context, imageID string, updatedImage *Image) (*ImageRef, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntity(appEngContext, imageID, updatedImage); updateErr != nil {
		return nil, fmt.Errorf("updateExistingImage: Error updating image: error = %v", updateErr)
	}

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, updatedImage.Field)
	if fieldErr != nil {
		return nil, fmt.Errorf("updateExistingImage: Error retrieving field for image: error = %v", fieldErr)
	}

	imageRef := ImageRef{
		ImageID:  imageID,
		FieldRef: *fieldRef,
		Geometry: updatedImage.Geometry}

	return &imageRef, nil

}
