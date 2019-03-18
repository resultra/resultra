package userSelection

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form/components/common"
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

func updateUserSelectionProps(trackerDBHandle *sql.DB, propUpdater UserSelectionPropUpdater) (*UserSelection, error) {

	// Retrieve the bar chart from the data store
	userSelectionForUpdate, getErr := getUserSelection(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getUserSelectionID())
	if getErr != nil {
		return nil, fmt.Errorf("updateUserSelectionProps: Unable to get existing userSelection: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(userSelectionForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateUserSelectionProps: Unable to update existing userSelection properties: %v", propUpdateErr)
	}

	updatedUserSelection, updateErr := updateExistingUserSelection(trackerDBHandle, userSelectionForUpdate)
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

type UserSelectionClearValueSupportedParams struct {
	UserSelectionIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams UserSelectionClearValueSupportedParams) updateProps(userSelection *UserSelection) error {

	userSelection.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	UserSelectionIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(userSelection *UserSelection) error {

	userSelection.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}

type SelectableRoleParams struct {
	UserSelectionIDHeader
	SelectableRoles []string `json:"selectableRoles"`
}

func (updateParams SelectableRoleParams) updateProps(userSelection *UserSelection) error {

	// TODO - Validate the role list

	userSelection.Properties.SelectableRoles = updateParams.SelectableRoles

	return nil
}

type CurrUserSelectableParams struct {
	UserSelectionIDHeader
	CurrUserSelectable bool `json:"currUserSelectable"`
}

func (updateParams CurrUserSelectableParams) updateProps(userSelection *UserSelection) error {

	userSelection.Properties.CurrUserSelectable = updateParams.CurrUserSelectable

	return nil
}
