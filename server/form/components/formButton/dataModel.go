package formButton

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const buttonEntityKind string = "formButton"

type FormButton struct {
	ParentFormID string           `json:"parentFormID"`
	ButtonID     string           `json:"buttonID"`
	Properties   ButtonProperties `json:"properties"`
}

type NewButtonParams struct {
	ParentFormID string                         `json:"parentFormID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
	LinkedFormID string                         `json:"linkedFormID"`
}

func saveButton(destDBHandle *sql.DB, newButton FormButton) error {

	if saveErr := common.SaveNewFormComponent(destDBHandle, buttonEntityKind,
		newButton.ParentFormID, newButton.ButtonID, newButton.Properties); saveErr != nil {
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

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid form component layout parameters: %+v", params)
	}

	if validateErr := validateFormExists(params.LinkedFormID); validateErr != nil {
		return nil, validateErr
	}

	properties := newDefaultButtonProperties()
	properties.Geometry = params.Geometry
	properties.LinkedFormID = params.LinkedFormID

	newButton := FormButton{ParentFormID: params.ParentFormID,
		ButtonID:   uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveButton(databaseWrapper.DBHandle(), newButton); err != nil {
		return nil, fmt.Errorf("saveNewButton: Unable to save button with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New form button: Created new form button: %+v", newButton)

	return &newButton, nil

}

func getButtonFromButtonID(buttonID string) (*FormButton, error) {

	parentFormID, err := common.GetFormComponentFormID(buttonID)
	if err != nil {
		return nil, fmt.Errorf("getButtonFromButtonID: Unable to retrieve button: %v", err)
	}
	return getButton(parentFormID, buttonID)
}

func getButton(parentFormID string, buttonID string) (*FormButton, error) {

	buttonProps := newDefaultButtonProperties()
	if getErr := common.GetFormComponent(buttonEntityKind, parentFormID, buttonID, &buttonProps); getErr != nil {
		return nil, fmt.Errorf("getButton: Unable to retrieve button: %v", getErr)
	}

	button := FormButton{
		ParentFormID: parentFormID,
		ButtonID:     buttonID,
		Properties:   buttonProps}

	return &button, nil
}

func getButtonsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]FormButton, error) {

	buttons := []FormButton{}
	addButton := func(datePickerID string, encodedProps string) error {

		buttonProps := newDefaultButtonProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &buttonProps); decodeErr != nil {
			return fmt.Errorf("GetButtons: can't decode properties: %v", encodedProps)
		}

		currButton := FormButton{
			ParentFormID: parentFormID,
			ButtonID:     datePickerID,
			Properties:   buttonProps}
		buttons = append(buttons, currButton)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, buttonEntityKind, parentFormID, addButton); getErr != nil {
		return nil, fmt.Errorf("GetButtons: Can't get buttons: %v")
	}

	return buttons, nil

}

func GetButtons(parentFormID string) ([]FormButton, error) {
	return getButtonsFromSrc(databaseWrapper.DBHandle(), parentFormID)
}

func CloneButtons(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcButtons, err := getButtonsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneButtons: %v", err)
	}

	for _, srcButton := range srcButtons {
		remappedButtonID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcButton.ButtonID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcButton.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
		destProperties, err := srcButton.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
		destButton := FormButton{
			ParentFormID: remappedFormID,
			ButtonID:     remappedButtonID,
			Properties:   *destProperties}
		if err := saveButton(cloneParams.DestDBHandle, destButton); err != nil {
			return fmt.Errorf("CloneButtons: %v", err)
		}
	}

	return nil
}

func updateExistingButton(updatedButton *FormButton) (*FormButton, error) {

	if updateErr := common.UpdateFormComponent(buttonEntityKind, updatedButton.ParentFormID,
		updatedButton.ButtonID, updatedButton.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingButton: failure updating button: %v", updateErr)
	}
	return updatedButton, nil

}
