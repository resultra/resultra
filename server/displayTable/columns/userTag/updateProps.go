package userTag

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type UserTagIDInterface interface {
	getUserTagID() string
	getParentTableID() string
}

type UserTagIDHeader struct {
	ParentTableID   string `json:"parentTableID"`
	UserTagID string `json:"userTagID"`
}

func (idHeader UserTagIDHeader) getUserTagID() string {
	return idHeader.UserTagID
}

func (idHeader UserTagIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type UserTagPropUpdater interface {
	UserTagIDInterface
	updateProps(userTag *UserTag) error
}

func updateUserTagProps(propUpdater UserTagPropUpdater) (*UserTag, error) {

	// Retrieve the bar chart from the data store
	userTagForUpdate, getErr := getUserTag(propUpdater.getParentTableID(), propUpdater.getUserTagID())
	if getErr != nil {
		return nil, fmt.Errorf("updateUserTagProps: Unable to get existing userTag: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(userTagForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateUserTagProps: Unable to update existing userTag properties: %v", propUpdateErr)
	}

	updatedUserTag, updateErr := updateExistingUserTag(userTagForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateUserTagProps: Unable to update existing userTag properties: datastore update error =  %v", updateErr)
	}

	return updatedUserTag, nil
}

type UserTagLabelFormatParams struct {
	UserTagIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams UserTagLabelFormatParams) updateProps(userTag *UserTag) error {

	// TODO - Validate format is well-formed.

	userTag.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type UserTagPermissionParams struct {
	UserTagIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams UserTagPermissionParams) updateProps(userTag *UserTag) error {

	userTag.Properties.Permissions = updateParams.Permissions

	return nil
}

type UserTagValidationParams struct {
	UserTagIDHeader
	Validation ValidationProperties `json:"validation"`
}

func (updateParams UserTagValidationParams) updateProps(userTag *UserTag) error {

	userTag.Properties.Validation = updateParams.Validation

	return nil
}

type UserTagClearValueSupportedParams struct {
	UserTagIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams UserTagClearValueSupportedParams) updateProps(userTag *UserTag) error {

	userTag.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	UserTagIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(userTag *UserTag) error {

	userTag.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}

type SelectableRoleParams struct {
	UserTagIDHeader
	SelectableRoles []string `json:"selectableRoles"`
}

func (updateParams SelectableRoleParams) updateProps(userTag *UserTag) error {

	// TODO - Validate the role list

	userTag.Properties.SelectableRoles = updateParams.SelectableRoles

	return nil
}

type CurrUserSelectableParams struct {
	UserTagIDHeader
	CurrUserSelectable bool `json:"currUserSelectable"`
}

func (updateParams CurrUserSelectableParams) updateProps(userTag *UserTag) error {

	userTag.Properties.CurrUserSelectable = updateParams.CurrUserSelectable

	return nil
}
