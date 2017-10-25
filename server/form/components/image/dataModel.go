package image

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

const imageEntityKind string = "image"

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
	if fieldType == field.FieldTypeImage {
		return true
	} else {
		return false
	}
}

func saveImage(destDBHandle *sql.DB, newImage Image) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, imageEntityKind,
		newImage.ParentFormID, newImage.ImageID, newImage.Properties); saveErr != nil {
		return fmt.Errorf("saveImage: Unable to save image: %v", saveErr)
	}
	return nil

}

func saveNewImage(params NewImageParams) (*Image, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validImageFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewImage: %v", fieldErr)
	}

	properties := newDefaultImageProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newImage := Image{ParentFormID: params.ParentFormID,
		ImageID:    uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveImage(databaseWrapper.DBHandle(), newImage); err != nil {
		return nil, fmt.Errorf("saveNewImage: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newImage)

	return &newImage, nil

}

func getImage(parentFormID string, imageID string) (*Image, error) {

	imageProps := newDefaultImageProperties()
	if getErr := common.GetFormComponent(imageEntityKind, parentFormID, imageID, &imageProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	image := Image{
		ParentFormID: parentFormID,
		ImageID:      imageID,
		Properties:   imageProps}

	return &image, nil
}

func getImagesFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Image, error) {

	imagees := []Image{}
	addImage := func(imageID string, encodedProps string) error {

		imageProps := newDefaultImageProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &imageProps); decodeErr != nil {
			return fmt.Errorf("GetImage: can't decode properties: %v", encodedProps)
		}

		currImage := Image{
			ParentFormID: parentFormID,
			ImageID:      imageID,
			Properties:   imageProps}
		imagees = append(imagees, currImage)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, imageEntityKind, parentFormID, addImage); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return imagees, nil

}

func GetImages(parentFormID string) ([]Image, error) {
	return getImagesFromSrc(databaseWrapper.DBHandle(), parentFormID)
}

func CloneImages(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcImage, err := getImagesFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneImage: %v", err)
	}

	for _, srcImage := range srcImage {
		remappedImageID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcImage.ImageID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcImage.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneImage: %v", err)
		}
		destProperties, err := srcImage.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneImage: %v", err)
		}
		destImage := Image{
			ParentFormID: remappedFormID,
			ImageID:      remappedImageID,
			Properties:   *destProperties}
		if err := saveImage(cloneParams.DestDBHandle, destImage); err != nil {
			return fmt.Errorf("CloneImage: %v", err)
		}
	}

	return nil
}

func updateExistingImage(imageID string, updatedImage *Image) (*Image, error) {

	if updateErr := common.UpdateFormComponent(imageEntityKind, updatedImage.ParentFormID,
		updatedImage.ImageID, updatedImage.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingImage: error updating existing text box component: %v", updateErr)
	}

	return updatedImage, nil

}
