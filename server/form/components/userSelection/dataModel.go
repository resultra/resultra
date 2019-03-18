// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userSelection

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/field"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

const userSelectionEntityKind string = "userSelection"

type UserSelection struct {
	ParentFormID    string                  `json:"parentFormID"`
	UserSelectionID string                  `json:"userSelectionID"`
	Properties      UserSelectionProperties `json:"properties"`
}

type NewUserSelectionParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validUserSelectionFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUser {
		return true
	} else {
		return false
	}
}

func saveUserSelection(destDBHandle *sql.DB, newUserSelection UserSelection) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, userSelectionEntityKind,
		newUserSelection.ParentFormID, newUserSelection.UserSelectionID, newUserSelection.Properties); saveErr != nil {
		return fmt.Errorf("saveNewUserSelection: Unable to save userSelection: error = %v", saveErr)
	}
	return nil
}

func saveNewUserSelection(trackerDBHandle *sql.DB, params NewUserSelectionParams) (*UserSelection, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validUserSelectionFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUserSelection: %v", fieldErr)
	}

	properties := newDefaultUserSelectionProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newUserSelection := UserSelection{ParentFormID: params.ParentFormID,
		UserSelectionID: uniqueID.GenerateUniqueID(),
		Properties:      properties}

	if saveErr := saveUserSelection(trackerDBHandle, newUserSelection); saveErr != nil {
		return nil, fmt.Errorf("saveNewUserSelection: Unable to save userSelection with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New UserSelection: Created new userSelection component:  %+v", newUserSelection)

	return &newUserSelection, nil

}

func getUserSelection(trackerDBHandle *sql.DB, parentFormID string, userSelectionID string) (*UserSelection, error) {

	userSelectionProps := newDefaultUserSelectionProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, userSelectionEntityKind, parentFormID,
		userSelectionID, &userSelectionProps); getErr != nil {
		return nil, fmt.Errorf("getUserSelection: Unable to retrieve userSelection: %v", getErr)
	}

	userSelection := UserSelection{
		ParentFormID:    parentFormID,
		UserSelectionID: userSelectionID,
		Properties:      userSelectionProps}

	return &userSelection, nil
}

func getUserSelectionsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]UserSelection, error) {

	userSelections := []UserSelection{}
	addUserSelection := func(userSelectionID string, encodedProps string) error {

		userSelectionProps := newDefaultUserSelectionProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &userSelectionProps); decodeErr != nil {
			return fmt.Errorf("GetUserSelections: can't decode properties: %v", encodedProps)
		}

		currUserSelection := UserSelection{
			ParentFormID:    parentFormID,
			UserSelectionID: userSelectionID,
			Properties:      userSelectionProps}
		userSelections = append(userSelections, currUserSelection)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, userSelectionEntityKind, parentFormID, addUserSelection); getErr != nil {
		return nil, fmt.Errorf("GetUserSelections: Can't get userSelections: %v")
	}

	return userSelections, nil
}

func GetUserSelections(trackerDBHandle *sql.DB, parentFormID string) ([]UserSelection, error) {
	return getUserSelectionsFromSrc(trackerDBHandle, parentFormID)
}

func CloneUserSelections(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcUserSelections, err := getUserSelectionsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneUserSelections: %v", err)
	}

	for _, srcUserSelection := range srcUserSelections {
		remappedUserSelectionID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcUserSelection.UserSelectionID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcUserSelection.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
		destProperties, err := srcUserSelection.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
		destUserSelection := UserSelection{
			ParentFormID:    remappedFormID,
			UserSelectionID: remappedUserSelectionID,
			Properties:      *destProperties}
		if err := saveUserSelection(cloneParams.DestDBHandle, destUserSelection); err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
	}

	return nil
}

func updateExistingUserSelection(trackerDBHandle *sql.DB, updatedUserSelection *UserSelection) (*UserSelection, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, userSelectionEntityKind, updatedUserSelection.ParentFormID,
		updatedUserSelection.UserSelectionID, updatedUserSelection.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUserSelection: failure updating userSelection: %v", updateErr)
	}
	return updatedUserSelection, nil

}
