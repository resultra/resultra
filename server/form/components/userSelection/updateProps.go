package userSelection

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type UserSelectionIDInterface interface {
	getUserSelectionID() string
	getParentFormID() string
}

type UserSelectionIDHeader struct {
	ParentFormID    string `json:"parentFormID"`
	UserSelectionID string `json:"userSelectionID"`
}

func (idHeader UserSelectionIDHeader) getUserSelectionID() string {
	return idHeader.UserSelectionID
}

func (idHeader UserSelectionIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type UserSelectionPropUpdater interface {
	UserSelectionIDInterface
	updateProps(userSelection *UserSelection) error
}

func updateUserSelectionProps(propUpdater UserSelectionPropUpdater) (*UserSelection, error) {

	// Retrieve the bar chart from the data store
	userSelectionForUpdate, getErr := getUserSelection(propUpdater.getParentFormID(), propUpdater.getUserSelectionID())
	if getErr != nil {
		return nil, fmt.Errorf("updateUserSelectionProps: Unable to get existing userSelection: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(userSelectionForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateUserSelectionProps: Unable to update existing userSelection properties: %v", propUpdateErr)
	}

	updatedUserSelection, updateErr := updateExistingUserSelection(userSelectionForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateUserSelectionProps: Unable to update existing userSelection properties: datastore update error =  %v", updateErr)
	}

	return updatedUserSelection, nil
}

type UserSelectionResizeParams struct {
	UserSelectionIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams UserSelectionResizeParams) updateProps(userSelection *UserSelection) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	userSelection.Properties.Geometry = updateParams.Geometry

	return nil
}

type UserSelectionLabelFormatParams struct {
	UserSelectionIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams UserSelectionLabelFormatParams) updateProps(userSelection *UserSelection) error {

	// TODO - Validate format is well-formed.

	userSelection.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}
