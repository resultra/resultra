package attachment

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

// TODO - Finish refactoring
const imageEntityKind string = "attachment"

type Image struct {
	ParentFormID string          `json:"parentFormID"`
	ImageID      string          `json:"imageID"`
	Properties   ImageProperties `json:"properties"`
}

type NewImageParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validImageFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeAttachment {
		return true
	} else {
		return false
	}
}

func saveImage(destDBHandle *sql.DB, newImage Image) error {

	if saveErr := common.SaveNewFormComponent(destDBHandle, imageEntityKind,
		newImage.ParentFormID, newImage.ImageID, newImage.Properties); saveErr != nil {
		return fmt.Errorf("saveNewImage: Unable to save image form component: error = %v", saveErr)
	}
	return nil

}

func saveNewImage(params NewImageParams) (*Image, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validImageFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: %v", fieldErr)
	}

	properties := newDefaultAttachmentProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newImage := Image{ParentFormID: params.ParentFormID,
		ImageID:    uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := saveImage(databaseWrapper.DBHandle(), newImage); saveErr != nil {
		return nil, fmt.Errorf("saveNewImage: Unable to save image form component with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: saveNewImage: Created new image component: %+v", newImage)

	return &newImage, nil

}

func getImage(parentFormID string, imageID string) (*Image, error) {

	imageProps := newDefaultAttachmentProperties()
	if getErr := common.GetFormComponent(imageEntityKind, parentFormID, imageID, &imageProps); getErr != nil {
		return nil, fmt.Errorf("getImage: Unable to retrieve image form component: %v", getErr)
	}

	image := Image{
		ParentFormID: parentFormID,
		ImageID:      imageID,
		Properties:   imageProps}

	return &image, nil
}

func GetImagesFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Image, error) {

	images := []Image{}
	addImage := func(imageID string, encodedProps string) error {

		imageProps := newDefaultAttachmentProperties()
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
	if getErr := common.GetFormComponents(srcDBHandle, imageEntityKind, parentFormID, addImage); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get image form components: %v")
	}

	return images, nil

}

func GetImages(parentFormID string) ([]Image, error) {
	return GetImagesFromSrc(databaseWrapper.DBHandle(), parentFormID)
}

func CloneImages(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcImages, err := GetImagesFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneImages: %v", err)
	}

	for _, srcImage := range srcImages {
		remappedImageID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcImage.ImageID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcImage.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
		destProperties, err := srcImage.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
		destImage := Image{
			ParentFormID: remappedFormID,
			ImageID:      remappedImageID,
			Properties:   *destProperties}
		if err := saveImage(cloneParams.DestDBHandle, destImage); err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
	}

	return nil
}

func updateExistingImage(imageID string, updatedImage *Image) (*Image, error) {

	if updateErr := common.UpdateFormComponent(imageEntityKind, updatedImage.ParentFormID,
		updatedImage.ImageID, updatedImage.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingImage: error updating existing image component: %v", updateErr)
	}

	return updatedImage, nil

}
