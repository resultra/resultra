// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userTag

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
	"log"
)

const userTagEntityKind string = "userTag"

type UserTag struct {
	ParentFormID string            `json:"parentFormID"`
	UserTagID    string            `json:"userTagID"`
	Properties   UserTagProperties `json:"properties"`
}

type NewUserTagParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validUserTagFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUsers {
		return true
	} else {
		return false
	}
}

func saveUserTag(destDBHandle *sql.DB, newUserTag UserTag) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, userTagEntityKind,
		newUserTag.ParentFormID, newUserTag.UserTagID, newUserTag.Properties); saveErr != nil {
		return fmt.Errorf("saveNewUserTag: Unable to save userTag: error = %v", saveErr)
	}
	return nil
}

func saveNewUserTag(trackerDBHandle *sql.DB, params NewUserTagParams) (*UserTag, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validUserTagFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: %v", fieldErr)
	}

	properties := newDefaultUserTagProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newUserTag := UserTag{ParentFormID: params.ParentFormID,
		UserTagID:  uniqueID.GenerateUniqueID(),
		Properties: properties}

	if saveErr := saveUserTag(trackerDBHandle, newUserTag); saveErr != nil {
		return nil, fmt.Errorf("saveNewUserTag: Unable to save userTag with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New UserTag: Created new userTag component:  %+v", newUserTag)

	return &newUserTag, nil

}

func getUserTag(trackerDBHandle *sql.DB, parentFormID string, userTagID string) (*UserTag, error) {

	userTagProps := newDefaultUserTagProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, userTagEntityKind, parentFormID,
		userTagID, &userTagProps); getErr != nil {
		return nil, fmt.Errorf("getUserTag: Unable to retrieve userTag: %v", getErr)
	}

	userTag := UserTag{
		ParentFormID: parentFormID,
		UserTagID:    userTagID,
		Properties:   userTagProps}

	return &userTag, nil
}

func getUserTagsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]UserTag, error) {

	userTags := []UserTag{}
	addUserTag := func(userTagID string, encodedProps string) error {

		userTagProps := newDefaultUserTagProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &userTagProps); decodeErr != nil {
			return fmt.Errorf("GetUserTags: can't decode properties: %v", encodedProps)
		}

		currUserTag := UserTag{
			ParentFormID: parentFormID,
			UserTagID:    userTagID,
			Properties:   userTagProps}
		userTags = append(userTags, currUserTag)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, userTagEntityKind, parentFormID, addUserTag); getErr != nil {
		return nil, fmt.Errorf("GetUserTags: Can't get userTags: %v")
	}

	return userTags, nil
}

func GetUserTags(trackerDBHandle *sql.DB, parentFormID string) ([]UserTag, error) {
	return getUserTagsFromSrc(trackerDBHandle, parentFormID)
}

func CloneUserTags(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcUserTags, err := getUserTagsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneUserTags: %v", err)
	}

	for _, srcUserTag := range srcUserTags {
		remappedUserTagID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcUserTag.UserTagID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcUserTag.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destProperties, err := srcUserTag.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
		destUserTag := UserTag{
			ParentFormID: remappedFormID,
			UserTagID:    remappedUserTagID,
			Properties:   *destProperties}
		if err := saveUserTag(cloneParams.DestDBHandle, destUserTag); err != nil {
			return fmt.Errorf("CloneUserTags: %v", err)
		}
	}

	return nil
}

func updateExistingUserTag(trackerDBHandle *sql.DB, updatedUserTag *UserTag) (*UserTag, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, userTagEntityKind, updatedUserTag.ParentFormID,
		updatedUserTag.UserTagID, updatedUserTag.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUserTag: failure updating userTag: %v", updateErr)
	}
	return updatedUserTag, nil

}
