package userTag

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const userTagEntityKind string = "userTag"

type UserTag struct {
	ParentFormID    string                  `json:"parentFormID"`
	UserTagID string                  `json:"userTagID"`
	Properties      UserTagProperties `json:"properties"`
}

type NewUserTagParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validUserTagFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUser {
		return true
	} else {
		return false
	}
}

func saveUserTag(newUserTag UserTag) error {
	if saveErr := common.SaveNewFormComponent(userTagEntityKind,
		newUserTag.ParentFormID, newUserTag.UserTagID, newUserTag.Properties); saveErr != nil {
		return fmt.Errorf("saveNewUserTag: Unable to save userTag: error = %v", saveErr)
	}
	return nil
}

func saveNewUserTag(params NewUserTagParams) (*UserTag, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validUserTagFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: %v", fieldErr)
	}

	properties := newDefaultUserTagProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newUserTag := UserTag{ParentFormID: params.ParentFormID,
		UserTagID: uniqueID.GenerateSnowflakeID(),
		Properties:      properties}

	if saveErr := saveUserTag(newUserTag); saveErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: Unable to save userTag with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New UserTag: Created new userTag component:  %+v", newUserTag)

	return &newUserTag, nil

}

func getUserTag(parentFormID string, userTagID string) (*UserTag, error) {

	userTagProps := newDefaultUserTagProperties()
	if getErr := common.GetFormComponent(userTagEntityKind, parentFormID,
		userTagID, &userTagProps); getErr != nil {
		return nil, fmt.Errorf("getUserTag: Unable to retrieve userTag: %v", getErr)
	}

	userTag := UserTag{
		ParentFormID:    parentFormID,
		UserTagID: userTagID,
		Properties:      userTagProps}

	return &userTag, nil
}

func GetUserTags(parentFormID string) ([]UserTag, error) {

	userTags := []UserTag{}
	addUserTag := func(userTagID string, encodedProps string) error {

		userTagProps := newDefaultUserTagProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &userTagProps); decodeErr != nil {
			return fmt.Errorf("GetUserTags: can't decode properties: %v", encodedProps)
		}

		currUserTag := UserTag{
			ParentFormID:    parentFormID,
			UserTagID: userTagID,
			Properties:      userTagProps}
		userTags = append(userTags, currUserTag)

		return nil
	}
	if getErr := common.GetFormComponents(userTagEntityKind, parentFormID, addUserTag); getErr != nil {
		return nil, fmt.Errorf("GetUserTags: Can't get userTags: %v")
	}

	return userTags, nil
}

func CloneUserTags(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcUserTags, err := GetUserTags(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneUserTags: %v", err)
	}

	for _, srcUserTag := range srcUserTags {
		remappedUserTagID := remappedIDs.AllocNewOrGetExistingRemappedID(srcUserTag.UserTagID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcUserTag.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destProperties, err := srcUserTag.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destUserTag := UserTag{
			ParentFormID:    remappedFormID,
			UserTagID: remappedUserTagID,
			Properties:      *destProperties}
		if err := saveUserTag(destUserTag); err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
	}

	return nil
}

func updateExistingUserTag(updatedUserTag *UserTag) (*UserTag, error) {

	if updateErr := common.UpdateFormComponent(userTagEntityKind, updatedUserTag.ParentFormID,
		updatedUserTag.UserTagID, updatedUserTag.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUserTag: failure updating userTag: %v", updateErr)
	}
	return updatedUserTag, nil

}
