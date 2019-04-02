// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package checkBox

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
)

const checkBoxEntityKind string = "checkbox"

type CheckBox struct {
	ParentFormID string             `json:"parentFormID"`
	CheckBoxID   string             `json:"checkBoxID"`
	Properties   CheckBoxProperties `json:"properties"`
}

type NewCheckBoxParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validCheckBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveCheckbox(destDBHandle *sql.DB, newCheckBox CheckBox) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, checkBoxEntityKind,
		newCheckBox.ParentFormID, newCheckBox.CheckBoxID, newCheckBox.Properties); saveErr != nil {
		return fmt.Errorf("saveCheckbox: Unable to save bar chart with error = %v", saveErr)
	}
	return nil
}

func saveNewCheckBox(trackerDBHandle *sql.DB, params NewCheckBoxParams) (*CheckBox, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validCheckBoxFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", fieldErr)
	}

	properties := newDefaultCheckBoxProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newCheckBox := CheckBox{ParentFormID: params.ParentFormID,
		CheckBoxID: uniqueID.GenerateUniqueID(),
		Properties: properties}

	if err := saveCheckbox(trackerDBHandle, newCheckBox); err != nil {
		return nil, fmt.Errorf("saveNewCheckBox: Unable to save bar chart with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New Checkbox: Created new check box container:  %+v", newCheckBox)

	return &newCheckBox, nil

}

func getCheckBox(trackerDBHandle *sql.DB, parentFormID string, checkBoxID string) (*CheckBox, error) {

	checkBoxProps := newDefaultCheckBoxProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, checkBoxEntityKind, parentFormID, checkBoxID, &checkBoxProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve check box: %v", getErr)
	}

	checkBox := CheckBox{
		ParentFormID: parentFormID,
		CheckBoxID:   checkBoxID,
		Properties:   checkBoxProps}

	return &checkBox, nil
}

func getCheckBoxesFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]CheckBox, error) {

	checkBoxes := []CheckBox{}
	addCheckbox := func(checkboxID string, encodedProps string) error {

		checkBoxProps := newDefaultCheckBoxProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &checkBoxProps); decodeErr != nil {
			return fmt.Errorf("GetCheckBoxes: can't decode properties: %v", encodedProps)
		}

		currCheckBox := CheckBox{
			ParentFormID: parentFormID,
			CheckBoxID:   checkboxID,
			Properties:   checkBoxProps}
		checkBoxes = append(checkBoxes, currCheckBox)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, checkBoxEntityKind, parentFormID, addCheckbox); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get checkboxes: %v")
	}

	return checkBoxes, nil
}

func GetCheckBoxes(trackerDBHandle *sql.DB, parentFormID string) ([]CheckBox, error) {
	return getCheckBoxesFromSrc(trackerDBHandle, parentFormID)
}

func CloneCheckBoxes(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcCheckBoxes, err := getCheckBoxesFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneCheckBoxes: %v", err)
	}

	for _, srcCheckBox := range srcCheckBoxes {
		remappedCheckBoxID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcCheckBox.CheckBoxID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcCheckBox.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destProperties, err := srcCheckBox.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destCheckBox := CheckBox{
			ParentFormID: remappedFormID,
			CheckBoxID:   remappedCheckBoxID,
			Properties:   *destProperties}
		if err := saveCheckbox(cloneParams.DestDBHandle, destCheckBox); err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
	}

	return nil
}

func updateExistingCheckBox(trackerDBHandle *sql.DB, updatedCheckBox *CheckBox) (*CheckBox, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, checkBoxEntityKind, updatedCheckBox.ParentFormID,
		updatedCheckBox.CheckBoxID, updatedCheckBox.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: failure updating checkbox: %v", updateErr)
	}
	return updatedCheckBox, nil

}
