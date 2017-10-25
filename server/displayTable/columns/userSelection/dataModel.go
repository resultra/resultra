package userSelection

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const userSelectionEntityKind string = "userSelection"

type UserSelection struct {
	ParentTableID   string                  `json:"parentTableID"`
	UserSelectionID string                  `json:"userSelectionID"`
	ColumnID        string                  `json:"columnID"`
	ColType         string                  `json:"colType"`
	Properties      UserSelectionProperties `json:"properties"`
}

type NewUserSelectionParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validUserSelectionFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUser {
		return true
	} else {
		return false
	}
}

func saveUserSelection(destDBHandle *sql.DB, newUserSelection UserSelection) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, userSelectionEntityKind,
		newUserSelection.ParentTableID, newUserSelection.UserSelectionID, newUserSelection.Properties); saveErr != nil {
		return fmt.Errorf("saveNewUserSelection: Unable to save userSelection: error = %v", saveErr)
	}
	return nil
}

func saveNewUserSelection(params NewUserSelectionParams) (*UserSelection, error) {

	if fieldErr := field.ValidateField(params.FieldID, validUserSelectionFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUserSelection: %v", fieldErr)
	}

	properties := newDefaultUserSelectionProperties()
	properties.FieldID = params.FieldID

	userSelectionID := uniqueID.GenerateSnowflakeID()
	newUserSelection := UserSelection{ParentTableID: params.ParentTableID,
		UserSelectionID: userSelectionID,
		ColumnID:        userSelectionID,
		ColType:         userSelectionEntityKind,
		Properties:      properties}

	if saveErr := saveUserSelection(databaseWrapper.DBHandle(), newUserSelection); saveErr != nil {
		return nil, fmt.Errorf("saveNewUserSelection: Unable to save userSelection with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New UserSelection: Created new userSelection component:  %+v", newUserSelection)

	return &newUserSelection, nil

}

func getUserSelection(parentTableID string, userSelectionID string) (*UserSelection, error) {

	userSelectionProps := newDefaultUserSelectionProperties()
	if getErr := common.GetTableColumn(userSelectionEntityKind, parentTableID,
		userSelectionID, &userSelectionProps); getErr != nil {
		return nil, fmt.Errorf("getUserSelection: Unable to retrieve userSelection: %v", getErr)
	}

	userSelection := UserSelection{
		ParentTableID:   parentTableID,
		UserSelectionID: userSelectionID,
		ColumnID:        userSelectionID,
		ColType:         userSelectionEntityKind,
		Properties:      userSelectionProps}

	return &userSelection, nil
}

func getUserSelectionsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]UserSelection, error) {

	userSelections := []UserSelection{}
	addUserSelection := func(userSelectionID string, encodedProps string) error {

		userSelectionProps := newDefaultUserSelectionProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &userSelectionProps); decodeErr != nil {
			return fmt.Errorf("GetUserSelections: can't decode properties: %v", encodedProps)
		}

		currUserSelection := UserSelection{
			ParentTableID:   parentTableID,
			UserSelectionID: userSelectionID,
			ColumnID:        userSelectionID,
			ColType:         userSelectionEntityKind,
			Properties:      userSelectionProps}
		userSelections = append(userSelections, currUserSelection)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, userSelectionEntityKind, parentTableID, addUserSelection); getErr != nil {
		return nil, fmt.Errorf("GetUserSelections: Can't get userSelections: %v")
	}

	return userSelections, nil
}

func GetUserSelections(parentTableID string) ([]UserSelection, error) {
	return getUserSelectionsFromSrc(databaseWrapper.DBHandle(), parentTableID)
}

func CloneUserSelections(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcUserSelections, err := getUserSelectionsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneUserSelections: %v", err)
	}

	for _, srcUserSelection := range srcUserSelections {
		remappedUserSelectionID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcUserSelection.UserSelectionID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcUserSelection.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
		destProperties, err := srcUserSelection.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
		destUserSelection := UserSelection{
			ParentTableID:   remappedFormID,
			UserSelectionID: remappedUserSelectionID,
			ColumnID:        remappedUserSelectionID,
			ColType:         userSelectionEntityKind,
			Properties:      *destProperties}
		if err := saveUserSelection(cloneParams.DestDBHandle, destUserSelection); err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
	}

	return nil
}

func updateExistingUserSelection(updatedUserSelection *UserSelection) (*UserSelection, error) {

	if updateErr := common.UpdateTableColumn(userSelectionEntityKind, updatedUserSelection.ParentTableID,
		updatedUserSelection.UserSelectionID, updatedUserSelection.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUserSelection: failure updating userSelection: %v", updateErr)
	}
	return updatedUserSelection, nil

}
