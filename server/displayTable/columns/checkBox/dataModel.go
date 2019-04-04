// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package checkBox

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/displayTable/columns/common"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
	"log"
)

const checkBoxEntityKind string = "checkbox"

type CheckBox struct {
	ParentTableID string             `json:"parentTableID"`
	CheckBoxID    string             `json:"checkBoxID"`
	ColumnID      string             `json:"columnID"`
	Properties    CheckBoxProperties `json:"properties"`
	ColType       string             `json:"colType"`
}

type NewCheckBoxParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validCheckBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveCheckbox(destDBHandle *sql.DB, newCheckBox CheckBox) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, checkBoxEntityKind,
		newCheckBox.ParentTableID, newCheckBox.CheckBoxID, newCheckBox.Properties); saveErr != nil {
		return fmt.Errorf("saveCheckbox: Unable to save bar chart with error = %v", saveErr)
	}
	return nil
}

func saveNewCheckBox(trackerDBHandle *sql.DB, params NewCheckBoxParams) (*CheckBox, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validCheckBoxFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", fieldErr)
	}

	properties := newDefaultCheckBoxProperties()
	properties.FieldID = params.FieldID

	checkBoxID := uniqueID.GenerateUniqueID()
	newCheckBox := CheckBox{ParentTableID: params.ParentTableID,
		CheckBoxID: checkBoxID,
		ColumnID:   checkBoxID,
		Properties: properties,
		ColType:    checkBoxEntityKind}

	if err := saveCheckbox(trackerDBHandle, newCheckBox); err != nil {
		return nil, fmt.Errorf("saveNewCheckBox: Unable to save bar chart with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New Checkbox: Created new check box container:  %+v", newCheckBox)

	return &newCheckBox, nil

}

func getCheckBox(trackerDBHandle *sql.DB, parentTableID string, checkBoxID string) (*CheckBox, error) {

	checkBoxProps := newDefaultCheckBoxProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, checkBoxEntityKind, parentTableID, checkBoxID, &checkBoxProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve check box: %v", getErr)
	}

	checkBox := CheckBox{
		ParentTableID: parentTableID,
		CheckBoxID:    checkBoxID,
		ColumnID:      checkBoxID,
		Properties:    checkBoxProps,
		ColType:       checkBoxEntityKind}

	return &checkBox, nil
}

func getCheckBoxesFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]CheckBox, error) {

	checkBoxes := []CheckBox{}
	addCheckbox := func(checkboxID string, encodedProps string) error {

		checkBoxProps := newDefaultCheckBoxProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &checkBoxProps); decodeErr != nil {
			return fmt.Errorf("GetCheckBoxes: can't decode properties: %v", encodedProps)
		}

		currCheckBox := CheckBox{
			ParentTableID: parentTableID,
			CheckBoxID:    checkboxID,
			ColumnID:      checkboxID,
			Properties:    checkBoxProps,
			ColType:       checkBoxEntityKind}
		checkBoxes = append(checkBoxes, currCheckBox)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, checkBoxEntityKind, parentTableID, addCheckbox); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get checkboxes: %v")
	}

	return checkBoxes, nil
}

func GetCheckBoxes(trackerDBHandle *sql.DB, parentTableID string) ([]CheckBox, error) {
	return getCheckBoxesFromSrc(trackerDBHandle, parentTableID)
}

func CloneCheckBoxes(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcCheckBoxes, err := getCheckBoxesFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneCheckBoxes: %v", err)
	}

	for _, srcCheckBox := range srcCheckBoxes {
		remappedCheckBoxID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcCheckBox.CheckBoxID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcCheckBox.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destProperties, err := srcCheckBox.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destCheckBox := CheckBox{
			ParentTableID: remappedFormID,
			CheckBoxID:    remappedCheckBoxID,
			ColumnID:      remappedCheckBoxID,
			Properties:    *destProperties,
			ColType:       checkBoxEntityKind}
		if err := saveCheckbox(cloneParams.DestDBHandle, destCheckBox); err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
	}

	return nil
}

func updateExistingCheckBox(trackerDBHandle *sql.DB, updatedCheckBox *CheckBox) (*CheckBox, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, checkBoxEntityKind, updatedCheckBox.ParentTableID,
		updatedCheckBox.CheckBoxID, updatedCheckBox.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: failure updating checkbox: %v", updateErr)
	}
	return updatedCheckBox, nil

}
