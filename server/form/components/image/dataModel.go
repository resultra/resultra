// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package image

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
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

func saveNewImage(trackerDBHandle *sql.DB, params NewImageParams) (*Image, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validImageFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewImage: %v", fieldErr)
	}

	properties := newDefaultImageProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newImage := Image{ParentFormID: params.ParentFormID,
		ImageID:    uniqueID.GenerateUniqueID(),
		Properties: properties}

	if err := saveImage(trackerDBHandle, newImage); err != nil {
		return nil, fmt.Errorf("saveNewImage: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newImage)

	return &newImage, nil

}

func getImage(trackerDBHandle *sql.DB, parentFormID string, imageID string) (*Image, error) {

	imageProps := newDefaultImageProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, imageEntityKind, parentFormID, imageID, &imageProps); getErr != nil {
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

func GetImages(trackerDBHandle *sql.DB, parentFormID string) ([]Image, error) {
	return getImagesFromSrc(trackerDBHandle, parentFormID)
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

func updateExistingImage(trackerDBHandle *sql.DB, imageID string, updatedImage *Image) (*Image, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, imageEntityKind, updatedImage.ParentFormID,
		updatedImage.ImageID, updatedImage.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingImage: error updating existing text box component: %v", updateErr)
	}

	return updatedImage, nil

}
