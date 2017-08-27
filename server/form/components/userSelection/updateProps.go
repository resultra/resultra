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

type UserSelectionVisibilityParams struct {
	UserSelectionIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams UserSelectionVisibilityParams) updateProps(userSelection *UserSelection) error {

	// TODO - Validate conditions

	userSelection.Properties.VisibilityConditions = updateParams.VisibilityConditions

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
