// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package tag

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form/components/common"
)

type TagIDInterface interface {
	getTagID() string
	getParentTableID() string
}

type TagIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	TagID         string `json:"tagID"`
}

func (idHeader TagIDHeader) getTagID() string {
	return idHeader.TagID
}

func (idHeader TagIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type TagPropUpdater interface {
	TagIDInterface
	updateProps(tag *Tag) error
}

func updateTagProps(trackerDBHandle *sql.DB, propUpdater TagPropUpdater) (*Tag, error) {

	// Retrieve the bar chart from the data store
	tagForUpdate, getErr := getTag(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getTagID())
	if getErr != nil {
		return nil, fmt.Errorf("updateTagProps: Unable to get existing tag: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(tagForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateTagProps: Unable to update existing tag properties: %v", propUpdateErr)
	}

	updatedTag, updateErr := updateExistingTag(trackerDBHandle, tagForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateTagProps: Unable to update existing tag properties: datastore update error =  %v", updateErr)
	}

	return updatedTag, nil
}

type TagLabelFormatParams struct {
	TagIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams TagLabelFormatParams) updateProps(tag *Tag) error {

	// TODO - Validate format is well-formed.

	tag.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type TagPermissionParams struct {
	TagIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams TagPermissionParams) updateProps(tag *Tag) error {

	tag.Properties.Permissions = updateParams.Permissions

	return nil
}

type TagValidationParams struct {
	TagIDHeader
	Validation ValidationProperties `json:"validation"`
}

func (updateParams TagValidationParams) updateProps(tag *Tag) error {

	tag.Properties.Validation = updateParams.Validation

	return nil
}

type HelpPopupMsgParams struct {
	TagIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(tag *Tag) error {

	tag.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
