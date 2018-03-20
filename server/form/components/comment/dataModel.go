package comment

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const commentEntityKind string = "comment"

type Comment struct {
	ParentFormID string            `json:"parentFormID"`
	CommentID    string            `json:"commentID"`
	Properties   CommentProperties `json:"properties"`
}

type NewCommentParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:fieldID`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validCommentFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeComment {
		return true
	} else {
		return false
	}
}

func saveComment(destDBHandle *sql.DB, newComment Comment) error {

	if saveErr := common.SaveNewFormComponent(destDBHandle, commentEntityKind,
		newComment.ParentFormID, newComment.CommentID, newComment.Properties); saveErr != nil {
		return fmt.Errorf("saveComment: Unable to save comment box: error = %v", saveErr)
	}

	return nil
}

func saveNewComment(trackerDBHandle *sql.DB, params NewCommentParams) (*Comment, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validCommentFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewComment: %v", fieldErr)
	}

	properties := newDefaultCommentProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newComment := Comment{ParentFormID: params.ParentFormID,
		CommentID:  uniqueID.GenerateUniqueID(),
		Properties: properties}

	if saveErr := saveComment(trackerDBHandle, newComment); saveErr != nil {
		return nil, fmt.Errorf("saveNewComment: Unable to save comment box with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Comment box: Created new comment box component:  %+v", newComment)

	return &newComment, nil

}

func getComment(trackerDBHandle *sql.DB, parentFormID string, commentID string) (*Comment, error) {

	commentProps := newDefaultCommentProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, commentEntityKind, parentFormID, commentID, &commentProps); getErr != nil {
		return nil, fmt.Errorf("getComment: Unable to retrieve comment box: %v", getErr)
	}

	comment := Comment{
		ParentFormID: parentFormID,
		CommentID:    commentID,
		Properties:   commentProps}

	return &comment, nil
}

func getCommentsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Comment, error) {

	comments := []Comment{}
	addComment := func(commentID string, encodedProps string) error {

		commentProps := newDefaultCommentProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &commentProps); decodeErr != nil {
			return fmt.Errorf("GetRatings: can't decode properties: %v", encodedProps)
		}

		currComment := Comment{
			ParentFormID: parentFormID,
			CommentID:    commentID,
			Properties:   commentProps}
		comments = append(comments, currComment)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, commentEntityKind, parentFormID, addComment); getErr != nil {
		return nil, fmt.Errorf("GetComments: Can't get comment boxes: %v")
	}

	return comments, nil
}

func GetComments(trackerDBHandle *sql.DB, parentFormID string) ([]Comment, error) {
	return getCommentsFromSrc(trackerDBHandle, parentFormID)
}

func CloneComments(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcComments, err := getCommentsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneComments: %v", err)
	}

	for _, srcComment := range srcComments {
		remappedCommentID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcComment.CommentID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcComment.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
		destProperties, err := srcComment.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
		destComment := Comment{
			ParentFormID: remappedFormID,
			CommentID:    remappedCommentID,
			Properties:   *destProperties}
		if err := saveComment(cloneParams.DestDBHandle, destComment); err != nil {
			return fmt.Errorf("CloneComments: %v", err)
		}
	}

	return nil
}

func updateExistingComment(trackerDBHandle *sql.DB, updatedComment *Comment) (*Comment, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, commentEntityKind, updatedComment.ParentFormID,
		updatedComment.CommentID, updatedComment.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingComment: failure updating comment: %v", updateErr)
	}
	return updatedComment, nil

}
