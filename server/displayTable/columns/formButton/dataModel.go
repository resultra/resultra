package formButton

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
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

func saveButton(destDBHandle *sql.DB, newButton FormButton) error {

	if saveErr := common.SaveNewTableColumn(destDBHandle, buttonEntityKind,
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
func validateFormExists(trackerDBHandle *sql.DB, formID string) error {
	var retrievedFormID string
	getErr := trackerDBHandle.QueryRow(`SELECT form_id FROM forms
		 WHERE form_id=$1 LIMIT 1`, formID).Scan(&retrievedFormID)
	if getErr != nil {
		return fmt.Errorf("validateFormExists: Unabled to get form: form ID = %v: datastore err=%v",
			formID, getErr)
	}
	return nil
}

func saveNewButton(trackerDBHandle *sql.DB, params NewButtonParams) (*FormButton, error) {

	if validateErr := validateFormExists(trackerDBHandle, params.LinkedFormID); validateErr != nil {
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

	if err := saveButton(trackerDBHandle, newButton); err != nil {
		return nil, fmt.Errorf("saveNewButton: Unable to save button with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New form button: Created new form button: %+v", newButton)

	return &newButton, nil

}

func getButton(trackerDBHandle *sql.DB, parentTableID string, buttonID string) (*FormButton, error) {

	buttonProps := newDefaultButtonProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, buttonEntityKind, parentTableID, buttonID, &buttonProps); getErr != nil {
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

func getButtonFromButtonID(trackerDBHandle *sql.DB, buttonID string) (*FormButton, error) {

	parentTableID, err := common.GetTableColumnTableID(trackerDBHandle, buttonID)
	if err != nil {
		return nil, fmt.Errorf("getButtonFromButtonID: Unable to retrieve button: %v", err)
	}
	return getButton(trackerDBHandle, parentTableID, buttonID)
}

func getButtonsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]FormButton, error) {

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
	if getErr := common.GetTableColumns(srcDBHandle, buttonEntityKind, parentTableID, addButton); getErr != nil {
		return nil, fmt.Errorf("GetButtons: Can't get buttons: %v")
	}

	return buttons, nil

}

func GetButtons(trackerDBHandle *sql.DB, parentTableID string) ([]FormButton, error) {
	return getButtonsFromSrc(trackerDBHandle, parentTableID)
}

func CloneButtons(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcButtons, err := getButtonsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneButtons: %v", err)
	}

	for _, srcButton := range srcButtons {
		remappedButtonID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcButton.ButtonID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcButton.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
		destProperties, err := srcButton.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
		destButton := FormButton{
			ParentTableID: remappedFormID,
			ButtonID:      remappedButtonID,
			ColumnID:      remappedButtonID,
			ColType:       buttonEntityKind,
			Properties:    *destProperties}
		if err := saveButton(cloneParams.DestDBHandle, destButton); err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
	}

	return nil
}

func updateExistingButton(trackerDBHandle *sql.DB, updatedButton *FormButton) (*FormButton, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, buttonEntityKind, updatedButton.ParentTableID,
		updatedButton.ButtonID, updatedButton.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingButton: failure updating button: %v", updateErr)
	}
	return updatedButton, nil

}
