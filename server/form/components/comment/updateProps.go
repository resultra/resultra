package comment

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
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

func updateCommentProps(propUpdater CommentPropUpdater) (*Comment, error) {

	// Retrieve the bar chart from the data store
	commentForUpdate, getErr := getComment(propUpdater.getParentFormID(), propUpdater.getCommentID())
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
