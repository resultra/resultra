package userSelection

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const userSelectionEntityKind string = "userSelection"

type UserSelection struct {
	ParentFormID    string                  `json:"parentFormID"`
	UserSelectionID string                  `json:"userSelectionID"`
	Properties      UserSelectionProperties `json:"properties"`
}

type NewUserSelectionParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func validUserSelectionFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUser {
		return true
	} else {
		return false
	}
}

func saveUserSelection(newUserSelection UserSelection) error {
	if saveErr := common.SaveNewFormComponent(userSelectionEntityKind,
		newUserSelection.ParentFormID, newUserSelection.UserSelectionID, newUserSelection.Properties); saveErr != nil {
		return fmt.Errorf("saveNewUserSelection: Unable to save userSelection: error = %v", saveErr)
	}
	return nil
}

func saveNewUserSelection(params NewUserSelectionParams) (*UserSelection, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validUserSelectionFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewUserSelection: %v", compLinkErr)
	}

	properties := UserSelectionProperties{
		ComponentLink: params.ComponentLink,
		Geometry:      params.Geometry}

	newUserSelection := UserSelection{ParentFormID: params.ParentFormID,
		UserSelectionID: uniqueID.GenerateSnowflakeID(),
		Properties:      properties}

	if saveErr := saveUserSelection(newUserSelection); saveErr != nil {
		return nil, fmt.Errorf("saveNewUserSelection: Unable to save userSelection with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New UserSelection: Created new userSelection component:  %+v", newUserSelection)

	return &newUserSelection, nil

}

func getUserSelection(parentFormID string, userSelectionID string) (*UserSelection, error) {

	userSelectionProps := UserSelectionProperties{}
	if getErr := common.GetFormComponent(userSelectionEntityKind, parentFormID,
		userSelectionID, &userSelectionProps); getErr != nil {
		return nil, fmt.Errorf("getUserSelection: Unable to retrieve userSelection: %v", getErr)
	}

	userSelection := UserSelection{
		ParentFormID:    parentFormID,
		UserSelectionID: userSelectionID,
		Properties:      userSelectionProps}

	return &userSelection, nil
}

func GetUserSelections(parentFormID string) ([]UserSelection, error) {

	userSelections := []UserSelection{}
	addUserSelection := func(userSelectionID string, encodedProps string) error {

		var userSelectionProps UserSelectionProperties
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
	if getErr := common.GetFormComponents(userSelectionEntityKind, parentFormID, addUserSelection); getErr != nil {
		return nil, fmt.Errorf("GetUserSelections: Can't get userSelections: %v")
	}

	return userSelections, nil
}

func CloneUserSelections(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcUserSelections, err := GetUserSelections(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneUserSelections: %v", err)
	}

	for _, srcUserSelection := range srcUserSelections {
		remappedUserSelectionID := remappedIDs.AllocNewOrGetExistingRemappedID(srcUserSelection.UserSelectionID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcUserSelection.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
		destProperties, err := srcUserSelection.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
		destUserSelection := UserSelection{
			ParentFormID:    remappedFormID,
			UserSelectionID: remappedUserSelectionID,
			Properties:      *destProperties}
		if err := saveUserSelection(destUserSelection); err != nil {
			return fmt.Errorf("CloneUserSelections: %v", err)
		}
	}

	return nil
}

func updateExistingUserSelection(updatedUserSelection *UserSelection) (*UserSelection, error) {

	if updateErr := common.UpdateFormComponent(userSelectionEntityKind, updatedUserSelection.ParentFormID,
		updatedUserSelection.UserSelectionID, updatedUserSelection.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUserSelection: failure updating userSelection: %v", updateErr)
	}
	return updatedUserSelection, nil

}
