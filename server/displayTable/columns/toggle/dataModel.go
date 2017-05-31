package toggle

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const toggleEntityKind string = "toggle"

type Toggle struct {
	ParentTableID string           `json:"parentTableID"`
	ToggleID      string           `json:"toggleID"`
	ColumnID      string           `json:"columnID"`
	ColType       string           `json:"colType"`
	Properties    ToggleProperties `json:"properties"`
}

type NewToggleParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validToggleFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveToggle(newToggle Toggle) error {
	if saveErr := common.SaveNewTableColumn(toggleEntityKind,
		newToggle.ParentTableID, newToggle.ToggleID, newToggle.Properties); saveErr != nil {
		return fmt.Errorf("saveToggle: Unable to save bar chart with error = %v", saveErr)
	}
	return nil
}

func saveNewToggle(params NewToggleParams) (*Toggle, error) {

	if fieldErr := field.ValidateField(params.FieldID, validToggleFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewToggle: %v", fieldErr)
	}

	properties := newDefaultToggleProperties()
	properties.FieldID = params.FieldID

	toggleID := uniqueID.GenerateSnowflakeID()
	newToggle := Toggle{ParentTableID: params.ParentTableID,
		ToggleID:   toggleID,
		ColumnID:   toggleID,
		ColType:    toggleEntityKind,
		Properties: properties}

	if err := saveToggle(newToggle); err != nil {
		return nil, fmt.Errorf("saveNewToggle: Unable to save bar chart with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New Toggle: Created new check box container:  %+v", newToggle)

	return &newToggle, nil

}

func getToggle(parentTableID string, toggleID string) (*Toggle, error) {

	toggleProps := newDefaultToggleProperties()
	if getErr := common.GetTableColumn(toggleEntityKind, parentTableID, toggleID, &toggleProps); getErr != nil {
		return nil, fmt.Errorf("getToggle: Unable to retrieve check box: %v", getErr)
	}

	toggle := Toggle{
		ParentTableID: parentTableID,
		ToggleID:      toggleID,
		ColumnID:      toggleID,
		ColType:       toggleEntityKind,
		Properties:    toggleProps}

	return &toggle, nil
}

func GetToggles(parentTableID string) ([]Toggle, error) {

	toggles := []Toggle{}
	addToggle := func(toggleID string, encodedProps string) error {

		toggleProps := newDefaultToggleProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &toggleProps); decodeErr != nil {
			return fmt.Errorf("GetTogglees: can't decode properties: %v", encodedProps)
		}

		currToggle := Toggle{
			ParentTableID: parentTableID,
			ToggleID:      toggleID,
			ColumnID:      toggleID,
			ColType:       toggleEntityKind,
			Properties:    toggleProps}
		toggles = append(toggles, currToggle)

		return nil
	}
	if getErr := common.GetTableColumns(toggleEntityKind, parentTableID, addToggle); getErr != nil {
		return nil, fmt.Errorf("GetTogglees: Can't get togglees: %v")
	}

	return toggles, nil
}

func CloneToggles(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	srcTogglees, err := GetToggles(parentTableID)
	if err != nil {
		return fmt.Errorf("CloneTogglees: %v", err)
	}

	for _, srcToggle := range srcTogglees {
		remappedToggleID := remappedIDs.AllocNewOrGetExistingRemappedID(srcToggle.ToggleID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcToggle.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
		destProperties, err := srcToggle.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
		destToggle := Toggle{
			ParentTableID: remappedFormID,
			ToggleID:      remappedToggleID,
			ColumnID:      remappedToggleID,
			ColType:       toggleEntityKind,
			Properties:    *destProperties}
		if err := saveToggle(destToggle); err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
	}

	return nil
}

func updateExistingToggle(updatedToggle *Toggle) (*Toggle, error) {

	if updateErr := common.UpdateTableColumn(toggleEntityKind, updatedToggle.ParentTableID,
		updatedToggle.ToggleID, updatedToggle.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingToggle: failure updating toggle: %v", updateErr)
	}
	return updatedToggle, nil

}
