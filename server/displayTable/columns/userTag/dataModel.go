package userTag

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
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

func saveUserTag(newUserTag UserTag) error {
	if saveErr := common.SaveNewTableColumn(userTagEntityKind,
		newUserTag.ParentTableID, newUserTag.UserTagID, newUserTag.Properties); saveErr != nil {
		return fmt.Errorf("saveNewUserTag: Unable to save userTag: error = %v", saveErr)
	}
	return nil
}

func saveNewUserTag(params NewUserTagParams) (*UserTag, error) {

	if fieldErr := field.ValidateField(params.FieldID, validUserTagFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: %v", fieldErr)
	}

	properties := newDefaultUserTagProperties()
	properties.FieldID = params.FieldID

	userTagID := uniqueID.GenerateSnowflakeID()
	newUserTag := UserTag{ParentTableID: params.ParentTableID,
		UserTagID:  userTagID,
		ColumnID:   userTagID,
		ColType:    userTagEntityKind,
		Properties: properties}

	if saveErr := saveUserTag(newUserTag); saveErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: Unable to save userTag with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New UserTag: Created new userTag component:  %+v", newUserTag)

	return &newUserTag, nil

}

func getUserTag(parentTableID string, userTagID string) (*UserTag, error) {

	userTagProps := newDefaultUserTagProperties()
	if getErr := common.GetTableColumn(userTagEntityKind, parentTableID,
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

func GetUserTags(parentTableID string) ([]UserTag, error) {

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
	if getErr := common.GetTableColumns(userTagEntityKind, parentTableID, addUserTag); getErr != nil {
		return nil, fmt.Errorf("GetUserTags: Can't get userTags: %v")
	}

	return userTags, nil
}

func CloneUserTags(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	srcUserTags, err := GetUserTags(parentTableID)
	if err != nil {
		return fmt.Errorf("CloneUserTags: %v", err)
	}

	for _, srcUserTag := range srcUserTags {
		remappedUserTagID := remappedIDs.AllocNewOrGetExistingRemappedID(srcUserTag.UserTagID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcUserTag.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destProperties, err := srcUserTag.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destUserTag := UserTag{
			ParentTableID: remappedFormID,
			UserTagID:     remappedUserTagID,
			ColumnID:      remappedUserTagID,
			ColType:       userTagEntityKind,
			Properties:    *destProperties}
		if err := saveUserTag(destUserTag); err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
	}

	return nil
}

func updateExistingUserTag(updatedUserTag *UserTag) (*UserTag, error) {

	if updateErr := common.UpdateTableColumn(userTagEntityKind, updatedUserTag.ParentTableID,
		updatedUserTag.UserTagID, updatedUserTag.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUserTag: failure updating userTag: %v", updateErr)
	}
	return updatedUserTag, nil

}
