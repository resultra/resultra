package userSelection

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type UserSelectionIDInterface interface {
	getUserSelectionID() string
	getParentTableID() string
}

type UserSelectionIDHeader struct {
	ParentTableID   string `json:"parentTableID"`
	UserSelectionID string `json:"userSelectionID"`
}

func (idHeader UserSelectionIDHeader) getUserSelectionID() string {
	return idHeader.UserSelectionID
}

func (idHeader UserSelectionIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type UserSelectionPropUpdater interface {
	UserSelectionIDInterface
	updateProps(userSelection *UserSelection) error
}

func updateUserSelectionProps(propUpdater UserSelectionPropUpdater) (*UserSelection, error) {

	// Retrieve the bar chart from the data store
	userSelectionForUpdate, getErr := getUserSelection(propUpdater.getParentTableID(), propUpdater.getUserSelectionID())
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

type UserSelectionLabelFormatParams struct {
	UserSelectionIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams UserSelectionLabelFormatParams) updateProps(userSelection *UserSelection) error {

	// TODO - Validate format is well-formed.

	userSelection.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type UserSelectionPermissionParams struct {
	UserSelectionIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams UserSelectionPermissionParams) updateProps(userSelection *UserSelection) error {

	userSelection.Properties.Permissions = updateParams.Permissions

	return nil
}

type UserSelectionValidationParams struct {
	UserSelectionIDHeader
	Validation ValidationProperties `json:"validation"`
}

func (updateParams UserSelectionValidationParams) updateProps(userSelection *UserSelection) error {

	userSelection.Properties.Validation = updateParams.Validation

	return nil
}