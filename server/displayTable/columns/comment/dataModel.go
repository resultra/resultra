package comment

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const commentEntityKind string = "comment"

type Comment struct {
	ParentTableID string            `json:"parentTableID"`
	CommentID     string            `json:"commentID"`
	ColumnID      string            `json:"columnID"`
	ColType       string            `json:"colType"`
	Properties    CommentProperties `json:"properties"`
}

type NewCommentParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:fieldID`
}

func validCommentFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeComment {
		return true
	} else {
		return false
	}
}

func saveComment(newComment Comment) error {

	if saveErr := common.SaveNewTableColumn(commentEntityKind,
		newComment.ParentTableID, newComment.CommentID, newComment.Properties); saveErr != nil {
		return fmt.Errorf("saveComment: Unable to save comment box: error = %v", saveErr)
	}

	return nil
}

func saveNewComment(params NewCommentParams) (*Comment, error) {

	if fieldErr := field.ValidateField(params.FieldID, validCommentFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewComment: %v", fieldErr)
	}

	properties := newDefaultCommentProperties()
	properties.FieldID = params.FieldID

	commentID := uniqueID.GenerateSnowflakeID()
	newComment := Comment{ParentTableID: params.ParentTableID,
		CommentID:  commentID,
		ColumnID:   commentID,
		ColType:    commentEntityKind,
		Properties: properties}

	if saveErr := saveComment(newComment); saveErr != nil {
		return nil, fmt.Errorf("saveNewComment: Unable to save comment box with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Comment box: Created new comment box component:  %+v", newComment)

	return &newComment, nil

}

func getComment(parentTableID string, commentID string) (*Comment, error) {

	commentProps := newDefaultCommentProperties()
	if getErr := common.GetTableColumn(commentEntityKind, parentTableID, commentID, &commentProps); getErr != nil {
		return nil, fmt.Errorf("getComment: Unable to retrieve comment box: %v", getErr)
	}

	comment := Comment{
		ParentTableID: parentTableID,
		CommentID:     commentID,
		ColumnID:      commentID,
		ColType:       commentEntityKind,
		Properties:    commentProps}

	return &comment, nil
}

func GetComments(parentTableID string) ([]Comment, error) {

	comments := []Comment{}
	addComment := func(commentID string, encodedProps string) error {

		commentProps := newDefaultCommentProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &commentProps); decodeErr != nil {
			return fmt.Errorf("GetRatings: can't decode properties: %v", encodedProps)
		}

		currComment := Comment{
			ParentTableID: parentTableID,
			CommentID:     commentID,
			ColumnID:      commentID,
			ColType:       commentEntityKind,
			Properties:    commentProps}
		comments = append(comments, currComment)

		return nil
	}
	if getErr := common.GetTableColumns(commentEntityKind, parentTableID, addComment); getErr != nil {
		return nil, fmt.Errorf("GetComments: Can't get comment boxes: %v")
	}

	return comments, nil
}

func CloneComments(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	srcComments, err := GetComments(parentTableID)
	if err != nil {
		return fmt.Errorf("CloneComments: %v", err)
	}

	for _, srcComment := range srcComments {
		remappedCommentID := remappedIDs.AllocNewOrGetExistingRemappedID(srcComment.CommentID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcComment.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
		destProperties, err := srcComment.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
		destComment := Comment{
			ParentTableID: remappedFormID,
			CommentID:     remappedCommentID,
			ColumnID:      remappedCommentID,
			ColType:       commentEntityKind,
			Properties:    *destProperties}
		if err := saveComment(destComment); err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
	}

	return nil
}

func updateExistingComment(updatedComment *Comment) (*Comment, error) {

	if updateErr := common.UpdateTableColumn(commentEntityKind, updatedComment.ParentTableID,
		updatedComment.CommentID, updatedComment.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingComment: failure updating comment: %v", updateErr)
	}
	return updatedComment, nil

}
