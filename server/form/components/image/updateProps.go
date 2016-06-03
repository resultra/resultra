package image

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/common"
)

type ImageIDInterface interface {
	getImageID() string
}

type ImageIDHeader struct {
	ImageID string `json:"imageID"`
}

func (idHeader ImageIDHeader) getImageID() string {
	return idHeader.ImageID
}

type ImagePropUpdater interface {
	ImageIDInterface
	updateProps(image *Image) error
}

func updateImageProps(appEngContext appengine.Context, propUpdater ImagePropUpdater) (*Image, error) {

	// Retrieve the bar chart from the data store
	imageForUpdate, getErr := getImage(appEngContext, propUpdater.getImageID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to get existing image: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(imageForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to update existing image properties: %v", propUpdateErr)
	}

	image, updateErr := updateExistingImage(appEngContext, propUpdater.getImageID(), imageForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to update existing image properties: datastore update error =  %v", updateErr)
	}

	return image, nil
}

type ImageResizeParams struct {
	ImageIDHeader
	Geometry common.LayoutGeometry `json:"geometry"`
}

func (updateParams ImageResizeParams) updateProps(image *Image) error {

	if !common.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set image dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	image.Geometry = updateParams.Geometry

	return nil
}

type ImageRepositionParams struct {
	ImageIDHeader
	Position common.LayoutPosition `json:"position"`
}

func (updateParams ImageRepositionParams) updateProps(image *Image) error {

	if err := image.Geometry.SetPosition(updateParams.Position); err != nil {
		return fmt.Errorf("Error setting position for image: Invalid geometry: %v", err)
	}

	return nil
}
