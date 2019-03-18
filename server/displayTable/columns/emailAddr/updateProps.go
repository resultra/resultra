// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package emailAddr

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form/components/common"
)

type EmailAddrIDInterface interface {
	getEmailAddrID() string
	getParentTableID() string
}

type EmailAddrIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	EmailAddrID   string `json:"emailAddrID"`
}

func (idHeader EmailAddrIDHeader) getEmailAddrID() string {
	return idHeader.EmailAddrID
}

func (idHeader EmailAddrIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type EmailAddrPropUpdater interface {
	EmailAddrIDInterface
	updateProps(emailAddr *EmailAddr) error
}

func updateEmailAddrProps(trackerDBHandle *sql.DB, propUpdater EmailAddrPropUpdater) (*EmailAddr, error) {

	// Retrieve the bar chart from the data store
	emailAddrForUpdate, getErr := getEmailAddr(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getEmailAddrID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateEmailAddrProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(emailAddrForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateEmailAddrProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	emailAddr, updateErr := updateExistingEmailAddr(trackerDBHandle, propUpdater.getEmailAddrID(), emailAddrForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateEmailAddrProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return emailAddr, nil
}

type EmailAddrLabelFormatParams struct {
	EmailAddrIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams EmailAddrLabelFormatParams) updateProps(emailAddr *EmailAddr) error {

	// TODO - Validate format is well-formed.

	emailAddr.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type EmailAddrPermissionParams struct {
	EmailAddrIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams EmailAddrPermissionParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.Permissions = updateParams.Permissions

	return nil
}

type EmailAddrValidationParams struct {
	EmailAddrIDHeader
	Validation EmailAddrValidationProperties `json:"validation"`
}

func (updateParams EmailAddrValidationParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.Validation = updateParams.Validation

	return nil
}

type EmailAddrClearValueSupportedParams struct {
	EmailAddrIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams EmailAddrClearValueSupportedParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	EmailAddrIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
