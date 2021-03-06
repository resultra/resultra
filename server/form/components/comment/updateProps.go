// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package comment

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
)

type CommentIDInterface interface {
	getCommentID() string
	getParentFormID() string
}

type CommentIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	CommentID    string `json:"commentID"`
}

func (idHeader CommentIDHeader) getCommentID() string {
	return idHeader.CommentID
}

func (idHeader CommentIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type CommentPropUpdater interface {
	CommentIDInterface
	updateProps(comment *Comment) error
}

func updateCommentProps(trackerDBHandle *sql.DB, propUpdater CommentPropUpdater) (*Comment, error) {

	// Retrieve the bar chart from the data store
	commentForUpdate, getErr := getComment(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getCommentID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCommentProps: Unable to get existing comment box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(commentForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCommentProps: Unable to update existing comment box properties: %v", propUpdateErr)
	}

	updatedComment, updateErr := updateExistingComment(trackerDBHandle, commentForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCommentProps: Unable to update existing comment box properties: datastore update error =  %v", updateErr)
	}

	return updatedComment, nil
}

type CommentResizeParams struct {
	CommentIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams CommentResizeParams) updateProps(comment *Comment) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set comment box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	comment.Properties.Geometry = updateParams.Geometry

	return nil
}

type CommentLabelFormatParams struct {
	CommentIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams CommentLabelFormatParams) updateProps(comment *Comment) error {

	// TODO - Validate format is well-formed.

	comment.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type CommentVisibilityParams struct {
	CommentIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams CommentVisibilityParams) updateProps(comment *Comment) error {

	// TODO - Validate conditions

	comment.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type CommentPermissionParams struct {
	CommentIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams CommentPermissionParams) updateProps(comment *Comment) error {

	comment.Properties.Permissions = updateParams.Permissions

	return nil
}

type HelpPopupMsgParams struct {
	CommentIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(comment *Comment) error {

	comment.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
