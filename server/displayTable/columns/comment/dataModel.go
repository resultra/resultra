package comment

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
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

func saveComment(destDBHandle *sql.DB, newComment Comment) error {

	if saveErr := common.SaveNewTableColumn(destDBHandle, commentEntityKind,
		newComment.ParentTableID, newComment.CommentID, newComment.Properties); saveErr != nil {
		return fmt.Errorf("saveComment: Unable to save comment box: error = %v", saveErr)
	}

	return nil
}

func saveNewComment(trackerDBHandle *sql.DB, params NewCommentParams) (*Comment, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validCommentFieldType); fieldErr != nil {
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

	if saveErr := saveComment(trackerDBHandle, newComment); saveErr != nil {
		return nil, fmt.Errorf("saveNewComment: Unable to save comment box with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Comment box: Created new comment box component:  %+v", newComment)

	return &newComment, nil

}

func getComment(trackerDBHandle *sql.DB, parentTableID string, commentID string) (*Comment, error) {

	commentProps := newDefaultCommentProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, commentEntityKind, parentTableID, commentID, &commentProps); getErr != nil {
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

func getCommentsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]Comment, error) {

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
	if getErr := common.GetTableColumns(srcDBHandle, commentEntityKind, parentTableID, addComment); getErr != nil {
		return nil, fmt.Errorf("GetComments: Can't get comment boxes: %v")
	}

	return comments, nil
}

func GetComments(trackerDBHandle *sql.DB, parentTableID string) ([]Comment, error) {
	return getCommentsFromSrc(trackerDBHandle, parentTableID)
}

func CloneComments(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcComments, err := getCommentsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneComments: %v", err)
	}

	for _, srcComment := range srcComments {
		remappedCommentID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcComment.CommentID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcComment.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
		destProperties, err := srcComment.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
		destComment := Comment{
			ParentTableID: remappedFormID,
			CommentID:     remappedCommentID,
			ColumnID:      remappedCommentID,
			ColType:       commentEntityKind,
			Properties:    *destProperties}
		if err := saveComment(cloneParams.DestDBHandle, destComment); err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
	}

	return nil
}

func updateExistingComment(trackerDBHandle *sql.DB, updatedComment *Comment) (*Comment, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, commentEntityKind, updatedComment.ParentTableID,
		updatedComment.CommentID, updatedComment.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingComment: failure updating comment: %v", updateErr)
	}
	return updatedComment, nil

}
