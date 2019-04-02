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
	"github.com/resultra/resultra/server/displayTable/columns/common"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
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

func saveImage(destDBHandle *sql.DB, newImage Image) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, imageEntityKind,
		newImage.ParentTableID, newImage.ImageID, newImage.Properties); saveErr != nil {
		return fmt.Errorf("saveImage: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewImage(trackerDBHandle *sql.DB, params NewImageParams) (*Image, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validImageFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewImage: %v", fieldErr)
	}

	properties := newDefaultImageProperties()
	properties.FieldID = params.FieldID

	imageID := uniqueID.GenerateUniqueID()
	newImage := Image{ParentTableID: params.ParentTableID,
		ImageID:    imageID,
		ColumnID:   imageID,
		Properties: properties,
		ColType:    imageEntityKind}

	if err := saveImage(trackerDBHandle, newImage); err != nil {
		return nil, fmt.Errorf("saveNewImage: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newImage)

	return &newImage, nil

}

func getImage(trackerDBHandle *sql.DB, parentTableID string, imageID string) (*Image, error) {

	imageProps := newDefaultImageProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, imageEntityKind, parentTableID, imageID, &imageProps); getErr != nil {
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

func getImagesFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]Image, error) {

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
	if getErr := common.GetTableColumns(srcDBHandle, imageEntityKind, parentTableID, addImage); getErr != nil {
		return nil, fmt.Errorf("GetImages: Can't get text boxes: %v")
	}

	return images, nil

}

func GetImages(trackerDBHandle *sql.DB, parentTableID string) ([]Image, error) {
	return getImagesFromSrc(trackerDBHandle, parentTableID)
}

func CloneImages(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcImagees, err := getImagesFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneImagees: %v", err)
	}

	for _, srcImage := range srcImagees {
		remappedImageID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcImage.ImageID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcImage.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
		destProperties, err := srcImage.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
		destImage := Image{
			ParentTableID: remappedFormID,
			ImageID:       remappedImageID,
			ColumnID:      remappedImageID,
			Properties:    *destProperties,
			ColType:       imageEntityKind}
		if err := saveImage(cloneParams.DestDBHandle, destImage); err != nil {
			return fmt.Errorf("CloneImages: %v", err)
		}
	}

	return nil
}

func updateExistingImage(trackerDBHandle *sql.DB, imageID string, updatedImage *Image) (*Image, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, imageEntityKind, updatedImage.ParentTableID,
		updatedImage.ImageID, updatedImage.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingImage: error updating existing text box component: %v", updateErr)
	}

	return updatedImage, nil

}
