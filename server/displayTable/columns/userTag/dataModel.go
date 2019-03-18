package userTag

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/displayTable/columns/common"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

const userTagEntityKind string = "userTag"

type UserTag struct {
	ParentTableID string            `json:"parentTableID"`
	UserTagID     string            `json:"userTagID"`
	ColumnID      string            `json:"columnID"`
	ColType       string            `json:"colType"`
	Properties    UserTagProperties `json:"properties"`
}

type NewUserTagParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validUserTagFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUsers {
		return true
	} else {
		return false
	}
}

func saveUserTag(destDBHandle *sql.DB, newUserTag UserTag) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, userTagEntityKind,
		newUserTag.ParentTableID, newUserTag.UserTagID, newUserTag.Properties); saveErr != nil {
		return fmt.Errorf("saveNewUserTag: Unable to save userTag: error = %v", saveErr)
	}
	return nil
}

func saveNewUserTag(trackerDBHandle *sql.DB, params NewUserTagParams) (*UserTag, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validUserTagFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: %v", fieldErr)
	}

	properties := newDefaultUserTagProperties()
	properties.FieldID = params.FieldID

	userTagID := uniqueID.GenerateUniqueID()
	newUserTag := UserTag{ParentTableID: params.ParentTableID,
		UserTagID:  userTagID,
		ColumnID:   userTagID,
		ColType:    userTagEntityKind,
		Properties: properties}

	if saveErr := saveUserTag(trackerDBHandle, newUserTag); saveErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: Unable to save userTag with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New UserTag: Created new userTag component:  %+v", newUserTag)

	return &newUserTag, nil

}

func getUserTag(trackerDBHandle *sql.DB, parentTableID string, userTagID string) (*UserTag, error) {

	userTagProps := newDefaultUserTagProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, userTagEntityKind, parentTableID,
		userTagID, &userTagProps); getErr != nil {
		return nil, fmt.Errorf("getUserTag: Unable to retrieve userTag: %v", getErr)
	}

	userTag := UserTag{
		ParentTableID: parentTableID,
		UserTagID:     userTagID,
		ColumnID:      userTagID,
		ColType:       userTagEntityKind,
		Properties:    userTagProps}

	return &userTag, nil
}

func getUserTagsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]UserTag, error) {

	userTags := []UserTag{}
	addUserTag := func(userTagID string, encodedProps string) error {

		userTagProps := newDefaultUserTagProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &userTagProps); decodeErr != nil {
			return fmt.Errorf("GetUserTags: can't decode properties: %v", encodedProps)
		}

		currUserTag := UserTag{
			ParentTableID: parentTableID,
			UserTagID:     userTagID,
			ColumnID:      userTagID,
			ColType:       userTagEntityKind,
			Properties:    userTagProps}
		userTags = append(userTags, currUserTag)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, userTagEntityKind, parentTableID, addUserTag); getErr != nil {
		return nil, fmt.Errorf("GetUserTags: Can't get userTags: %v")
	}

	return userTags, nil
}

func GetUserTags(trackerDBHandle *sql.DB, parentTableID string) ([]UserTag, error) {
	return getUserTagsFromSrc(trackerDBHandle, parentTableID)
}

func CloneUserTags(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcUserTags, err := getUserTagsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneUserTags: %v", err)
	}

	for _, srcUserTag := range srcUserTags {
		remappedUserTagID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcUserTag.UserTagID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcUserTag.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destProperties, err := srcUserTag.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destUserTag := UserTag{
			ParentTableID: remappedFormID,
			UserTagID:     remappedUserTagID,
			ColumnID:      remappedUserTagID,
			ColType:       userTagEntityKind,
			Properties:    *destProperties}
		if err := saveUserTag(cloneParams.DestDBHandle, destUserTag); err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
	}

	return nil
}

func updateExistingUserTag(trackerDBHandle *sql.DB, updatedUserTag *UserTag) (*UserTag, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, userTagEntityKind, updatedUserTag.ParentTableID,
		updatedUserTag.UserTagID, updatedUserTag.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUserTag: failure updating userTag: %v", updateErr)
	}
	return updatedUserTag, nil

}
