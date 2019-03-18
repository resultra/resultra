// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package file

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form/components/common"
)

type FileIDInterface interface {
	getFileID() string
	getParentTableID() string
}

type FileIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	FileID        string `json:"fileID"`
}

func (idHeader FileIDHeader) getFileID() string {
	return idHeader.FileID
}

func (idHeader FileIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type FilePropUpdater interface {
	FileIDInterface
	updateProps(file *File) error
}

func updateFileProps(trackerDBHandle *sql.DB, propUpdater FilePropUpdater) (*File, error) {

	// Retrieve the bar chart from the data store
	fileForUpdate, getErr := getFile(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getFileID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateFileProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(fileForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateFileProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	file, updateErr := updateExistingFile(trackerDBHandle, propUpdater.getFileID(), fileForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateFileProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return file, nil
}

type FileLabelFormatParams struct {
	FileIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams FileLabelFormatParams) updateProps(file *File) error {

	// TODO - Validate format is well-formed.

	file.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type FilePermissionParams struct {
	FileIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams FilePermissionParams) updateProps(file *File) error {

	file.Properties.Permissions = updateParams.Permissions

	return nil
}

type FileValidationParams struct {
	FileIDHeader
	Validation FileValidationProperties `json:"validation"`
}

func (updateParams FileValidationParams) updateProps(file *File) error {

	file.Properties.Validation = updateParams.Validation

	return nil
}

type FileClearValueSupportedParams struct {
	FileIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams FileClearValueSupportedParams) updateProps(file *File) error {

	file.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	FileIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(file *File) error {

	file.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
