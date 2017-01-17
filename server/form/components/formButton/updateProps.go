package formButton

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
)

type ButtonIDInterface interface {
	getButtonID() string
	getParentFormID() string
}

type ButtonIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	ButtonID     string `json:"buttonID"`
}

func (idHeader ButtonIDHeader) getButtonID() string {
	return idHeader.ButtonID
}

func (idHeader ButtonIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type ButtonPropUpdater interface {
	ButtonIDInterface
	updateProps(button *FormButton) error
}

func updateButtonProps(propUpdater ButtonPropUpdater) (*FormButton, error) {

	// Retrieve the bar chart from the data store
	buttonForUpdate, getErr := getButton(propUpdater.getParentFormID(), propUpdater.getButtonID())
	if getErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to get existing button: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(buttonForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to update existing button properties: %v", propUpdateErr)
	}

	updatedButton, updateErr := updateExistingButton(buttonForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to update existing button properties: datastore update error =  %v", updateErr)
	}

	return updatedButton, nil
}

type ButtonResizeParams struct {
	ButtonIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams ButtonResizeParams) updateProps(buttonForUpdate *FormButton) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set comment box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	buttonForUpdate.Properties.Geometry = updateParams.Geometry

	return nil
}
