package image

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type ImageIDInterface interface {
	getImageID() string
	getParentFormID() string
}

type ImageIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	ImageID      string `json:"imageID"`
}

func (idHeader ImageIDHeader) getImageID() string {
	return idHeader.ImageID
}

func (idHeader ImageIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type ImagePropUpdater interface {
	ImageIDInterface
	updateProps(image *Image) error
}

func updateImageProps(propUpdater ImagePropUpdater) (*Image, error) {

	// Retrieve the bar chart from the data store
	imageForUpdate, getErr := getImage(propUpdater.getParentFormID(), propUpdater.getImageID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to get existing image: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(imageForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to update existing image properties: %v", propUpdateErr)
	}

	image, updateErr := updateExistingImage(propUpdater.getImageID(), imageForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to update existing image properties: datastore update error =  %v", updateErr)
	}

	return image, nil
}

type ImageResizeParams struct {
	ImageIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams ImageResizeParams) updateProps(image *Image) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set image dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	image.Properties.Geometry = updateParams.Geometry

	return nil
}

type AttachmentLabelFormatParams struct {
	ImageIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams AttachmentLabelFormatParams) updateProps(image *Image) error {

	// TODO - Validate format is well-formed.

	image.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}
