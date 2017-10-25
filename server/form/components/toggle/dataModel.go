package toggle

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
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

func saveToggle(destDBHandle *sql.DB, newToggle Toggle) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, toggleEntityKind,
		newToggle.ParentFormID, newToggle.ToggleID, newToggle.Properties); saveErr != nil {
		return fmt.Errorf("saveToggle: Unable to save bar chart with error = %v", saveErr)
	}
	return nil
}

func saveNewToggle(params NewToggleParams) (*Toggle, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validToggleFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewToggle: %v", fieldErr)
	}

	properties := newDefaultToggleProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newToggle := Toggle{ParentFormID: params.ParentFormID,
		ToggleID:   uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveToggle(databaseWrapper.DBHandle(), newToggle); err != nil {
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

func getTogglesFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Toggle, error) {

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
	if getErr := common.GetFormComponents(srcDBHandle, toggleEntityKind, parentFormID, addToggle); getErr != nil {
		return nil, fmt.Errorf("GetTogglees: Can't get togglees: %v")
	}

	return togglees, nil
}

func GetToggles(parentFormID string) ([]Toggle, error) {
	return getTogglesFromSrc(databaseWrapper.DBHandle(), parentFormID)
}

func CloneToggles(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcTogglees, err := getTogglesFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneTogglees: %v", err)
	}

	for _, srcToggle := range srcTogglees {
		remappedToggleID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcToggle.ToggleID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcToggle.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
		destProperties, err := srcToggle.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneTogglees: %v", err)
		}
		destToggle := Toggle{
			ParentFormID: remappedFormID,
			ToggleID:     remappedToggleID,
			Properties:   *destProperties}
		if err := saveToggle(cloneParams.DestDBHandle, destToggle); err != nil {
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
