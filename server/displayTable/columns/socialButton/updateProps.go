// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package socialButton

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/form/components/common"
)

type SocialButtonIDInterface interface {
	getSocialButtonID() string
	getParentTableID() string
}

type SocialButtonIDHeader struct {
	ParentTableID  string `json:"parentTableID"`
	SocialButtonID string `json:"socialButtonID"`
}

func (idHeader SocialButtonIDHeader) getSocialButtonID() string {
	return idHeader.SocialButtonID
}

func (idHeader SocialButtonIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type SocialButtonPropUpdater interface {
	SocialButtonIDInterface
	updateProps(socialButton *SocialButton) error
}

func updateSocialButtonProps(trackerDBHandle *sql.DB, propUpdater SocialButtonPropUpdater) (*SocialButton, error) {

	// Retrieve the bar chart from the data store
	socialButtonForUpdate, getErr := getSocialButton(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getSocialButtonID())
	if getErr != nil {
		return nil, fmt.Errorf("updateSocialButtonProps: Unable to get existing socialButton: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(socialButtonForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateSocialButtonProps: Unable to update existing socialButton properties: %v", propUpdateErr)
	}

	updatedSocialButton, updateErr := updateExistingSocialButton(trackerDBHandle, socialButtonForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateSocialButtonProps: Unable to update existing socialButton properties: datastore update error =  %v", updateErr)
	}

	return updatedSocialButton, nil
}

type SocialButtonIconParams struct {
	SocialButtonIDHeader
	Icon string `json:"icon"`
}

func (updateParams SocialButtonIconParams) updateProps(socialButton *SocialButton) error {

	// TODO - Validate icon is a valid name

	socialButton.Properties.Icon = updateParams.Icon

	return nil
}

type SocialButtonLabelFormatParams struct {
	SocialButtonIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams SocialButtonLabelFormatParams) updateProps(socialButton *SocialButton) error {

	// TODO - Validate format is well-formed.

	socialButton.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type SocialButtonPermissionParams struct {
	SocialButtonIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams SocialButtonPermissionParams) updateProps(socialButton *SocialButton) error {

	// TODO - Validate conditions

	socialButton.Properties.Permissions = updateParams.Permissions

	return nil
}

type HelpPopupMsgParams struct {
	SocialButtonIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(socialButton *SocialButton) error {

	socialButton.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
