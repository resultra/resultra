package image

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/form/components/common"
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

func updateImageProps(trackerDBHandle *sql.DB, propUpdater ImagePropUpdater) (*Image, error) {

	// Retrieve the bar chart from the data store
	imageForUpdate, getErr := getImage(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getImageID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(imageForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	image, updateErr := updateExistingImage(trackerDBHandle, propUpdater.getImageID(), imageForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateImageProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return image, nil
}

type ImageResizeParams struct {
	ImageIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams ImageResizeParams) updateProps(image *Image) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set text box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	image.Properties.Geometry = updateParams.Geometry

	return nil
}

type ImageLabelFormatParams struct {
	ImageIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams ImageLabelFormatParams) updateProps(image *Image) error {

	// TODO - Validate format is well-formed.

	image.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type ImageVisibilityParams struct {
	ImageIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams ImageVisibilityParams) updateProps(image *Image) error {

	// TODO - Validate conditions

	image.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type ImagePermissionParams struct {
	ImageIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams ImagePermissionParams) updateProps(image *Image) error {

	image.Properties.Permissions = updateParams.Permissions

	return nil
}

type ImageValidationParams struct {
	ImageIDHeader
	Validation ImageValidationProperties `json:"validation"`
}

func (updateParams ImageValidationParams) updateProps(image *Image) error {

	image.Properties.Validation = updateParams.Validation

	return nil
}

type ImageClearValueSupportedParams struct {
	ImageIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams ImageClearValueSupportedParams) updateProps(image *Image) error {

	image.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	ImageIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(image *Image) error {

	image.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
