package selection

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const selectionEntityKind string = "selection"

type SelectionSelectableVal struct {
	Val   string `json:"val"`
	Label string `json:"label"`
}

type SelectionProperties struct {
	ComponentLink  common.ComponentLink           `json:"componentLink"`
	Geometry       componentLayout.LayoutGeometry `json:"geometry"`
	SelectableVals []SelectionSelectableVal       `json:"selectableVals"`
}

type Selection struct {
	ParentFormID string              `json:"parentFormID"`
	SelectionID  string              `json:"selectionID"`
	Properties   SelectionProperties `json:"properties"`
}

type NewSelectionParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
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

func saveNewSelection(params NewSelectionParams) (*Selection, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validSelectionFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewSelection: %v", compLinkErr)
	}

	properties := SelectionProperties{
		Geometry:      params.Geometry,
		ComponentLink: params.ComponentLink}

	newSelection := Selection{ParentFormID: params.ParentFormID,
		SelectionID: uniqueID.GenerateSnowflakeID(),
		Properties:  properties}

	if saveErr := common.SaveNewFormComponent(selectionEntityKind,
		newSelection.ParentFormID, newSelection.SelectionID, newSelection.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewSelection: Unable to save text box with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newSelection)

	return &newSelection, nil

}

func getSelection(parentFormID string, selectionID string) (*Selection, error) {

	selectionProps := SelectionProperties{}
	if getErr := common.GetFormComponent(selectionEntityKind, parentFormID, selectionID, &selectionProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	selection := Selection{
		ParentFormID: parentFormID,
		SelectionID:  selectionID,
		Properties:   selectionProps}

	return &selection, nil
}

func GetSelections(parentFormID string) ([]Selection, error) {

	selections := []Selection{}
	addSelection := func(selectionID string, encodedProps string) error {

		var selectionProps SelectionProperties
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
	if getErr := common.GetFormComponents(selectionEntityKind, parentFormID, addSelection); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return selections, nil

}

func updateExistingSelection(selectionID string, updatedSelection *Selection) (*Selection, error) {

	if updateErr := common.UpdateFormComponent(selectionEntityKind, updatedSelection.ParentFormID,
		updatedSelection.SelectionID, updatedSelection.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingSelection: error updating existing selection component: %v", updateErr)
	}

	return updatedSelection, nil

}
