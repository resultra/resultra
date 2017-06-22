package formButton

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const buttonEntityKind string = "button"

type FormButton struct {
	ParentTableID string           `json:"parentTableID"`
	ButtonID      string           `json:"buttonID"`
	ColumnID      string           `json:"columnID"`
	ColType       string           `json:"colType"`
	Properties    ButtonProperties `json:"properties"`
}

type NewButtonParams struct {
	ParentTableID string `json:"parentTableID"`
	LinkedFormID  string `json:"linkedFormID"`
}

func saveButton(newButton FormButton) error {

	if saveErr := common.SaveNewTableColumn(buttonEntityKind,
		newButton.ParentTableID, newButton.ButtonID, newButton.Properties); saveErr != nil {
		return fmt.Errorf("saveButton: Unable to save button: error = %v", saveErr)
	}
	return nil

}

// This function somewhat duplicates the same functino in the form package. However,
// since the form package already depends on this package, a circular package reference cannot
// be created.
// TODO - Move the datamodel specific functions in the form package to a lower level package
// (which doesn't depend on this package), but
// keep the controller-level functionality in a higher level package.
func validateFormExists(formID string) error {
	var retrievedFormID string
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT form_id FROM forms
		 WHERE form_id=$1 LIMIT 1`, formID).Scan(&retrievedFormID)
	if getErr != nil {
		return fmt.Errorf("validateFormExists: Unabled to get form: form ID = %v: datastore err=%v",
			formID, getErr)
	}
	return nil
}

func saveNewButton(params NewButtonParams) (*FormButton, error) {

	if validateErr := validateFormExists(params.LinkedFormID); validateErr != nil {
		return nil, validateErr
	}

	properties := newDefaultButtonProperties()
	properties.LinkedFormID = params.LinkedFormID

	buttonID := uniqueID.GenerateSnowflakeID()
	newButton := FormButton{ParentTableID: params.ParentTableID,
		ButtonID:   buttonID,
		ColumnID:   buttonID,
		ColType:    buttonEntityKind,
		Properties: properties}

	if err := saveButton(newButton); err != nil {
		return nil, fmt.Errorf("saveNewButton: Unable to save button with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New form button: Created new form button: %+v", newButton)

	return &newButton, nil

}

func getButton(parentTableID string, buttonID string) (*FormButton, error) {

	buttonProps := newDefaultButtonProperties()
	if getErr := common.GetTableColumn(buttonEntityKind, parentTableID, buttonID, &buttonProps); getErr != nil {
		return nil, fmt.Errorf("getButton: Unable to retrieve button: %v", getErr)
	}

	button := FormButton{
		ParentTableID: parentTableID,
		ButtonID:      buttonID,
		ColumnID:      buttonID,
		ColType:       buttonEntityKind,
		Properties:    buttonProps}

	return &button, nil
}

func GetButtons(parentTableID string) ([]FormButton, error) {

	buttons := []FormButton{}
	addButton := func(buttonID string, encodedProps string) error {

		buttonProps := newDefaultButtonProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &buttonProps); decodeErr != nil {
			return fmt.Errorf("GetButtons: can't decode properties: %v", encodedProps)
		}

		currButton := FormButton{
			ParentTableID: parentTableID,
			ButtonID:      buttonID,
			ColumnID:      buttonID,
			ColType:       buttonEntityKind,
			Properties:    buttonProps}
		buttons = append(buttons, currButton)

		return nil
	}
	if getErr := common.GetTableColumns(buttonEntityKind, parentTableID, addButton); getErr != nil {
		return nil, fmt.Errorf("GetButtons: Can't get buttons: %v")
	}

	return buttons, nil

}

func CloneButtons(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	srcButtons, err := GetButtons(parentTableID)
	if err != nil {
		return fmt.Errorf("CloneButtons: %v", err)
	}

	for _, srcButton := range srcButtons {
		remappedButtonID := remappedIDs.AllocNewOrGetExistingRemappedID(srcButton.ButtonID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcButton.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
		destProperties, err := srcButton.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
		destButton := FormButton{
			ParentTableID: remappedFormID,
			ButtonID:      remappedButtonID,
			ColumnID:      remappedButtonID,
			ColType:       buttonEntityKind,
			Properties:    *destProperties}
		if err := saveButton(destButton); err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
	}

	return nil
}

func updateExistingButton(updatedButton *FormButton) (*FormButton, error) {

	if updateErr := common.UpdateTableColumn(buttonEntityKind, updatedButton.ParentTableID,
		updatedButton.ButtonID, updatedButton.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingButton: failure updating button: %v", updateErr)
	}
	return updatedButton, nil

}
