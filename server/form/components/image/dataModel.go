package image

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const imageEntityKind string = "image"

type ImageProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

type Image struct {
	ParentFormID string          `json:"parentFormID"`
	ImageID      string          `json:"imageID"`
	Properties   ImageProperties `json:"properties"`
}

type NewImageParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func validImageFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeFile {
		return true
	} else {
		return false
	}
}

func saveNewImage(params NewImageParams) (*Image, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validImageFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: %v", compLinkErr)
	}

	properties := ImageProperties{
		Geometry:      params.Geometry,
		ComponentLink: params.ComponentLink}

	newImage := Image{ParentFormID: params.ParentFormID,
		ImageID:    uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := common.SaveNewFormComponent(imageEntityKind,
		newImage.ParentFormID, newImage.ImageID, newImage.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewImage: Unable to save image form component with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: saveNewImage: Created new image component: %+v", newImage)

	return &newImage, nil

}

func getImage(parentFormID string, imageID string) (*Image, error) {

	imageProps := ImageProperties{}
	if getErr := common.GetFormComponent(imageEntityKind, parentFormID, imageID, &imageProps); getErr != nil {
		return nil, fmt.Errorf("getImage: Unable to retrieve image form component: %v", getErr)
	}

	image := Image{
		ParentFormID: parentFormID,
		ImageID:      imageID,
		Properties:   imageProps}

	return &image, nil
}

func GetImages(parentFormID string) ([]Image, error) {

	images := []Image{}
	addImage := func(imageID string, encodedProps string) error {

		var imageProps ImageProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &imageProps); decodeErr != nil {
			return fmt.Errorf("GetImages: can't decode properties: %v", encodedProps)
		}

		currImage := Image{
			ParentFormID: parentFormID,
			ImageID:      imageID,
			Properties:   imageProps}
		images = append(images, currImage)

		return nil
	}
	if getErr := common.GetFormComponents(imageEntityKind, parentFormID, addImage); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get image form components: %v")
	}

	return images, nil

}

func updateExistingImage(imageID string, updatedImage *Image) (*Image, error) {

	if updateErr := common.UpdateFormComponent(imageEntityKind, updatedImage.ParentFormID,
		updatedImage.ImageID, updatedImage.Properties); updateErr != nil {
	}

	return updatedImage, nil

}
