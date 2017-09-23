package image

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const imageEntityKind string = "image"

type Image struct {
	ParentTableID string          `json:"parentTableID"`
	ImageID       string          `json:"imageID"`
	ColType       string          `json:"colType"`
	ColumnID      string          `json:"columnID"`
	Properties    ImageProperties `json:"properties"`
}

type NewImageParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validImageFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeImage {
		return true
	} else {
		return false
	}
}

func saveImage(newImage Image) error {
	if saveErr := common.SaveNewTableColumn(imageEntityKind,
		newImage.ParentTableID, newImage.ImageID, newImage.Properties); saveErr != nil {
		return fmt.Errorf("saveImage: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewImage(params NewImageParams) (*Image, error) {

	if fieldErr := field.ValidateField(params.FieldID, validImageFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewImage: %v", fieldErr)
	}

	properties := newDefaultImageProperties()
	properties.FieldID = params.FieldID

	imageID := uniqueID.GenerateSnowflakeID()
	newImage := Image{ParentTableID: params.ParentTableID,
		ImageID:    imageID,
		ColumnID:   imageID,
		Properties: properties,
		ColType:    imageEntityKind}

	if err := saveImage(newImage); err != nil {
		return nil, fmt.Errorf("saveNewImage: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newImage)

	return &newImage, nil

}

func getImage(parentTableID string, imageID string) (*Image, error) {

	imageProps := newDefaultImageProperties()
	if getErr := common.GetTableColumn(imageEntityKind, parentTableID, imageID, &imageProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	image := Image{
		ParentTableID: parentTableID,
		ImageID:       imageID,
		ColumnID:      imageID,
		Properties:    imageProps,
		ColType:       imageEntityKind}

	return &image, nil
}

func GetImages(parentTableID string) ([]Image, error) {

	images := []Image{}
	addImage := func(imageID string, encodedProps string) error {

		imageProps := newDefaultImageProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &imageProps); decodeErr != nil {
			return fmt.Errorf("GetImages: can't decode properties: %v", encodedProps)
		}

		currImage := Image{
			ParentTableID: parentTableID,
			ImageID:       imageID,
			ColumnID:      imageID,
			Properties:    imageProps,
			ColType:       imageEntityKind}
		images = append(images, currImage)

		return nil
	}
	if getErr := common.GetTableColumns(imageEntityKind, parentTableID, addImage); getErr != nil {
		return nil, fmt.Errorf("GetImages: Can't get text boxes: %v")
	}

	return images, nil

}

func CloneImages(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcImagees, err := GetImages(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneImagees: %v", err)
	}

	for _, srcImage := range srcImagees {
		remappedImageID := remappedIDs.AllocNewOrGetExistingRemappedID(srcImage.ImageID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcImage.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
		destProperties, err := srcImage.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
		destImage := Image{
			ParentTableID: remappedFormID,
			ImageID:       remappedImageID,
			ColumnID:      remappedImageID,
			Properties:    *destProperties,
			ColType:       imageEntityKind}
		if err := saveImage(destImage); err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
	}

	return nil
}

func updateExistingImage(imageID string, updatedImage *Image) (*Image, error) {

	if updateErr := common.UpdateTableColumn(imageEntityKind, updatedImage.ParentTableID,
		updatedImage.ImageID, updatedImage.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingImage: error updating existing text box component: %v", updateErr)
	}

	return updatedImage, nil

}
