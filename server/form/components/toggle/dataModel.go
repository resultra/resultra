package toggle

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const toggleEntityKind string = "toggle"

type Toggle struct {
	ParentFormID string           `json:"parentFormID"`
	ToggleID     string           `json:"toggleID"`
	Properties   ToggleProperties `json:"properties"`
}

type NewToggleParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validToggleFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveToggle(newToggle Toggle) error {
	if saveErr := common.SaveNewFormComponent(toggleEntityKind,
		newToggle.ParentFormID, newToggle.ToggleID, newToggle.Properties); saveErr != nil {
		return fmt.Errorf("saveToggle: Unable to save bar chart with error = %v", saveErr)
	}
	return nil
}

func saveNewToggle(params NewToggleParams) (*Toggle, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := common.ValidateField(params.FieldID, validToggleFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewToggle: %v", fieldErr)
	}

	properties := newDefaultToggleProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newToggle := Toggle{ParentFormID: params.ParentFormID,
		ToggleID:   uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveToggle(newToggle); err != nil {
		return nil, fmt.Errorf("saveNewToggle: Unable to save bar chart with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New Toggle: Created new check box container:  %+v", newToggle)

	return &newToggle, nil

}

func getToggle(parentFormID string, toggleID string) (*Toggle, error) {

	toggleProps := newDefaultToggleProperties()
	if getErr := common.GetFormComponent(toggleEntityKind, parentFormID, toggleID, &toggleProps); getErr != nil {
		return nil, fmt.Errorf("getToggle: Unable to retrieve check box: %v", getErr)
	}

	toggle := Toggle{
		ParentFormID: parentFormID,
		ToggleID:     toggleID,
		Properties:   toggleProps}

	return &toggle, nil
}

func GetToggles(parentFormID string) ([]Toggle, error) {

	togglees := []Toggle{}
	addToggle := func(toggleID string, encodedProps string) error {

		toggleProps := newDefaultToggleProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &toggleProps); decodeErr != nil {
			return fmt.Errorf("GetTogglees: can't decode properties: %v", encodedProps)
		}

		currToggle := Toggle{
			ParentFormID: parentFormID,
			ToggleID:     toggleID,
			Properties:   toggleProps}
		togglees = append(togglees, currToggle)

		return nil
	}
	if getErr := common.GetFormComponents(toggleEntityKind, parentFormID, addToggle); getErr != nil {
		return nil, fmt.Errorf("GetTogglees: Can't get togglees: %v")
	}

	return togglees, nil
}

func CloneToggles(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcTogglees, err := GetToggles(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneTogglees: %v", err)
	}

	for _, srcToggle := range srcTogglees {
		remappedToggleID := remappedIDs.AllocNewOrGetExistingRemappedID(srcToggle.ToggleID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcToggle.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
		destProperties, err := srcToggle.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
		destToggle := Toggle{
			ParentFormID: remappedFormID,
			ToggleID:     remappedToggleID,
			Properties:   *destProperties}
		if err := saveToggle(destToggle); err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
	}

	return nil
}

func updateExistingToggle(updatedToggle *Toggle) (*Toggle, error) {

	if updateErr := common.UpdateFormComponent(toggleEntityKind, updatedToggle.ParentFormID,
		updatedToggle.ToggleID, updatedToggle.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingToggle: failure updating toggle: %v", updateErr)
	}
	return updatedToggle, nil

}
