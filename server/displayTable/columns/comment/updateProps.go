package comment

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type CommentIDInterface interface {
	getCommentID() string
	getParentTableID() string
}

type CommentIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	CommentID     string `json:"commentID"`
}

func (idHeader CommentIDHeader) getCommentID() string {
	return idHeader.CommentID
}

func (idHeader CommentIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type CommentPropUpdater interface {
	CommentIDInterface
	updateProps(comment *Comment) error
}

func updateCommentProps(propUpdater CommentPropUpdater) (*Comment, error) {

	// Retrieve the bar chart from the data store
	commentForUpdate, getErr := getComment(propUpdater.getParentTableID(), propUpdater.getCommentID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCommentProps: Unable to get existing comment box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(commentForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCommentProps: Unable to update existing comment box properties: %v", propUpdateErr)
	}

	updatedComment, updateErr := updateExistingComment(commentForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCommentProps: Unable to update existing comment box properties: datastore update error =  %v", updateErr)
	}

	return updatedComment, nil
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

type CommentPermissionParams struct {
	CommentIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams CommentPermissionParams) updateProps(comment *Comment) error {

	comment.Properties.Permissions = updateParams.Permissions

	return nil
}
