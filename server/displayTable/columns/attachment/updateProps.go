// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package attachment

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/form/components/common"
)

type AttachmentIDInterface interface {
	getAttachmentID() string
	getParentTableID() string
}

type AttachmentIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	AttachmentID  string `json:"attachmentID"`
}

func (idHeader AttachmentIDHeader) getAttachmentID() string {
	return idHeader.AttachmentID
}

func (idHeader AttachmentIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type AttachmentPropUpdater interface {
	AttachmentIDInterface
	updateProps(attachment *Attachment) error
}

func updateAttachmentProps(trackerDBHandle *sql.DB, propUpdater AttachmentPropUpdater) (*Attachment, error) {

	// Retrieve the bar chart from the data store
	attachmentForUpdate, getErr := getAttachment(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getAttachmentID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateAttachmentProps: Unable to get existing attachment: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(attachmentForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateAttachmentProps: Unable to update existing attachment properties: %v", propUpdateErr)
	}

	attachment, updateErr := updateExistingAttachment(trackerDBHandle, propUpdater.getAttachmentID(), attachmentForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateAttachmentProps: Unable to update existing attachment properties: datastore update error =  %v", updateErr)
	}

	return attachment, nil
}

type AttachmentLabelFormatParams struct {
	AttachmentIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams AttachmentLabelFormatParams) updateProps(attachment *Attachment) error {

	// TODO - Validate format is well-formed.

	attachment.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type AttachmentPermissionParams struct {
	AttachmentIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams AttachmentPermissionParams) updateProps(attachment *Attachment) error {

	attachment.Properties.Permissions = updateParams.Permissions

	return nil
}

type AttachmentValidationParams struct {
	AttachmentIDHeader
	Validation ValidationProperties `json:"validation"`
}

func (updateParams AttachmentValidationParams) updateProps(attachment *Attachment) error {

	attachment.Properties.Validation = updateParams.Validation

	return nil
}

type HelpPopupMsgParams struct {
	AttachmentIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(attachment *Attachment) error {

	attachment.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
