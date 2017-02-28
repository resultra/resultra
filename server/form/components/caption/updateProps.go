package caption

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/stringValidation"
)

type CaptionIDInterface interface {
	getCaptionID() string
	getParentFormID() string
}

type CaptionIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	CaptionID    string `json:"captionID"`
}

func (idCaption CaptionIDHeader) getCaptionID() string {
	return idCaption.CaptionID
}

func (idCaption CaptionIDHeader) getParentFormID() string {
	return idCaption.ParentFormID
}

type CaptionPropUpdater interface {
	CaptionIDInterface
	updateProps(caption *Caption) error
}

func updateCaptionProps(propUpdater CaptionPropUpdater) (*Caption, error) {

	// Retrieve the bar chart from the data store
	captionForUpdate, getErr := getCaption(propUpdater.getParentFormID(), propUpdater.getCaptionID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCaptionProps: Unable to get existing caption: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(captionForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCaptionProps: Unable to update existing caption properties: %v", propUpdateErr)
	}

	updatedCaption, updateErr := updateExistingCaption(captionForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCaptionProps: Unable to update existing caption properties: datastore update error =  %v", updateErr)
	}

	return updatedCaption, nil
}

type CaptionLabelParams struct {
	CaptionIDHeader
	Label string `json:"label"`
}

func (updateParams CaptionLabelParams) updateProps(captionForUpdate *Caption) error {

	if !stringValidation.WellFormedItemLabel(updateParams.Label) {
		return fmt.Errorf("Update caption label: invalid label: %v", updateParams.Label)
	}

	captionForUpdate.Properties.Label = updateParams.Label

	return nil
}

type CaptionResizeParams struct {
	CaptionIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams CaptionResizeParams) updateProps(captionForUpdate *Caption) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set comment box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	captionForUpdate.Properties.Geometry = updateParams.Geometry

	return nil
}
