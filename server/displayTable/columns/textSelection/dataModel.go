package textSelection

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/displayTable/columns/common"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

const textSelectionEntityKind string = "textSelection"

type TextSelection struct {
	ParentTableID string                  `json:"parentTableID"`
	SelectionID   string                  `json:"selectionID"`
	ColType       string                  `json:"colType"`
	ColumnID      string                  `json:"columnID"`
	Properties    TextSelectionProperties `json:"properties"`
}

type NewTextSelectionParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validTextSelectionFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeText {
		return true
	} else {
		return false
	}
}

func saveTextSelection(destDBHandle *sql.DB, newTextSelection TextSelection) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, textSelectionEntityKind,
		newTextSelection.ParentTableID, newTextSelection.SelectionID, newTextSelection.Properties); saveErr != nil {
		return fmt.Errorf("saveTextInput: Unable to save text selection: %v", saveErr)
	}
	return nil
}

func saveNewTextSelection(trackerDBHandle *sql.DB, params NewTextSelectionParams) (*TextSelection, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validTextSelectionFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewTextSelection: %v", fieldErr)
	}

	properties := newDefaultTextSelectionProperties()
	properties.FieldID = params.FieldID

	selectionID := uniqueID.GenerateUniqueID()
	newTextSelection := TextSelection{ParentTableID: params.ParentTableID,
		SelectionID: selectionID,
		ColumnID:    selectionID,
		Properties:  properties,
		ColType:     textSelectionEntityKind}

	if err := saveTextSelection(trackerDBHandle, newTextSelection); err != nil {
		return nil, fmt.Errorf("saveNewTextInput: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newTextSelection)

	return &newTextSelection, nil

}

func getTextSelection(trackerDBHandle *sql.DB, parentTableID string, selectionID string) (*TextSelection, error) {

	textSelectionProps := newDefaultTextSelectionProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, textSelectionEntityKind,
		parentTableID, selectionID, &textSelectionProps); getErr != nil {
		return nil, fmt.Errorf("getTextInput: Unable to retrieve text box: %v", getErr)
	}

	textSelection := TextSelection{
		ParentTableID: parentTableID,
		SelectionID:   selectionID,
		ColumnID:      selectionID,
		Properties:    textSelectionProps,
		ColType:       textSelectionEntityKind}

	return &textSelection, nil
}

func getTextSelectionsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]TextSelection, error) {

	textSelections := []TextSelection{}
	addTextSelection := func(selectionID string, encodedProps string) error {

		textSelectionProps := newDefaultTextSelectionProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &textSelectionProps); decodeErr != nil {
			return fmt.Errorf("GetTextInputs: can't decode properties: %v", encodedProps)
		}

		currTextSelection := TextSelection{
			ParentTableID: parentTableID,
			SelectionID:   selectionID,
			ColumnID:      selectionID,
			Properties:    textSelectionProps,
			ColType:       textSelectionEntityKind}
		textSelections = append(textSelections, currTextSelection)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, textSelectionEntityKind, parentTableID, addTextSelection); getErr != nil {
		return nil, fmt.Errorf("GetTextInputs: Can't get text boxes: %v")
	}

	return textSelections, nil

}

func GetTextSelections(trackerDBHandle *sql.DB, parentTableID string) ([]TextSelection, error) {
	return getTextSelectionsFromSrc(trackerDBHandle, parentTableID)
}

func CloneTextSelections(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcTextInputes, err := getTextSelectionsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneTextInputes: %v", err)
	}

	for _, srcTextInput := range srcTextInputes {
		remappedSelectionID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcTextInput.SelectionID)
		remappedTableID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcTextInput.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneTextInputs: %v", err)
		}
		destProperties, err := srcTextInput.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneTextInputs: %v", err)
		}
		destTextSelection := TextSelection{
			ParentTableID: remappedTableID,
			SelectionID:   remappedSelectionID,
			ColumnID:      remappedSelectionID,
			Properties:    *destProperties,
			ColType:       textSelectionEntityKind}
		if err := saveTextSelection(cloneParams.DestDBHandle, destTextSelection); err != nil {
			return fmt.Errorf("CloneTextSelections: %v", err)
		}
	}

	return nil
}

func updateExistingTextSelection(trackerDBHandle *sql.DB,
	selectionID string, updatedTextSelection *TextSelection) (*TextSelection, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, textSelectionEntityKind, updatedTextSelection.ParentTableID,
		updatedTextSelection.SelectionID, updatedTextSelection.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingTextInput: error updating existing text box component: %v", updateErr)
	}

	return updatedTextSelection, nil

}
