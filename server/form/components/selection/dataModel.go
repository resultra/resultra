// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package selection

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
	"log"
)

const selectionEntityKind string = "selection"

type Selection struct {
	ParentFormID string              `json:"parentFormID"`
	SelectionID  string              `json:"selectionID"`
	Properties   SelectionProperties `json:"properties"`
}

type NewSelectionParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validSelectionFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeText {
		return true
	} else if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveSelection(destDBHandle *sql.DB, newSelection Selection) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, selectionEntityKind,
		newSelection.ParentFormID, newSelection.SelectionID, newSelection.Properties); saveErr != nil {
		return fmt.Errorf("saveNewSelection: Unable to save selection: error = %v", saveErr)
	}
	return nil
}

func saveNewSelection(trackerDBHandle *sql.DB, params NewSelectionParams) (*Selection, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := field.ValidateField(trackerDBHandle, params.FieldID, validSelectionFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewSelection: %v", compLinkErr)
	}

	properties := newDefaultSelectionProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newSelection := Selection{ParentFormID: params.ParentFormID,
		SelectionID: uniqueID.GenerateUniqueID(),
		Properties:  properties}

	if saveErr := saveSelection(trackerDBHandle, newSelection); saveErr != nil {
		return nil, fmt.Errorf("saveNewSelection: Unable to save text box with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newSelection)

	return &newSelection, nil

}

func getSelection(trackerDBHandle *sql.DB, parentFormID string, selectionID string) (*Selection, error) {

	selectionProps := newDefaultSelectionProperties()
	if getErr := common.GetFormComponent(trackerDBHandle,
		selectionEntityKind, parentFormID, selectionID, &selectionProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	selection := Selection{
		ParentFormID: parentFormID,
		SelectionID:  selectionID,
		Properties:   selectionProps}

	return &selection, nil
}

func getSelectionsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Selection, error) {

	selections := []Selection{}
	addSelection := func(selectionID string, encodedProps string) error {

		selectionProps := newDefaultSelectionProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &selectionProps); decodeErr != nil {
			return fmt.Errorf("GetSelectiones: can't decode properties: %v", encodedProps)
		}

		currSelection := Selection{
			ParentFormID: parentFormID,
			SelectionID:  selectionID,
			Properties:   selectionProps}
		selections = append(selections, currSelection)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, selectionEntityKind, parentFormID, addSelection); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return selections, nil

}

func GetSelections(trackerDBHandle *sql.DB, parentFormID string) ([]Selection, error) {
	return getSelectionsFromSrc(trackerDBHandle, parentFormID)
}

func CloneSelections(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcSelections, err := getSelectionsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneSelections: %v", err)
	}

	for _, srcSelection := range srcSelections {
		remappedSelectionID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcSelection.SelectionID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcSelection.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneSelections: %v", err)
		}
		destProperties, err := srcSelection.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneSelections: %v", err)
		}
		destSelection := Selection{
			ParentFormID: remappedFormID,
			SelectionID:  remappedSelectionID,
			Properties:   *destProperties}
		if err := saveSelection(cloneParams.DestDBHandle, destSelection); err != nil {
			return fmt.Errorf("CloneSelections: %v", err)
		}
	}

	return nil
}

func updateExistingSelection(trackerDBHandle *sql.DB, selectionID string, updatedSelection *Selection) (*Selection, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, selectionEntityKind, updatedSelection.ParentFormID,
		updatedSelection.SelectionID, updatedSelection.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingSelection: error updating existing selection component: %v", updateErr)
	}

	return updatedSelection, nil

}
