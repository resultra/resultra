package textInput

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const textInputEntityKind string = "textInput"

type TextInput struct {
	ParentTableID string              `json:"parentTableID"`
	TextInputID   string              `json:"textInputID"`
	ColType       string              `json:"colType"`
	ColumnID      string              `json:"columnID"`
	Properties    TextInputProperties `json:"properties"`
}

type NewTextInputParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validTextInputFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeText {
		return true
	} else {
		return false
	}
}

func saveTextInput(newTextInput TextInput) error {
	if saveErr := common.SaveNewTableColumn(textInputEntityKind,
		newTextInput.ParentTableID, newTextInput.TextInputID, newTextInput.Properties); saveErr != nil {
		return fmt.Errorf("saveTextInput: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewTextInput(params NewTextInputParams) (*TextInput, error) {

	if fieldErr := field.ValidateField(params.FieldID, validTextInputFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewTextInput: %v", fieldErr)
	}

	properties := newDefaultTextInputProperties()
	properties.FieldID = params.FieldID

	textInputID := uniqueID.GenerateSnowflakeID()
	newTextInput := TextInput{ParentTableID: params.ParentTableID,
		TextInputID: textInputID,
		ColumnID:    textInputID,
		Properties:  properties,
		ColType:     textInputEntityKind}

	if err := saveTextInput(newTextInput); err != nil {
		return nil, fmt.Errorf("saveNewTextInput: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newTextInput)

	return &newTextInput, nil

}

func getTextInput(parentTableID string, textInputID string) (*TextInput, error) {

	textInputProps := newDefaultTextInputProperties()
	if getErr := common.GetTableColumn(textInputEntityKind, parentTableID, textInputID, &textInputProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	textInput := TextInput{
		ParentTableID: parentTableID,
		TextInputID:   textInputID,
		ColumnID:      textInputID,
		Properties:    textInputProps,
		ColType:       textInputEntityKind}

	return &textInput, nil
}

func GetTextInputs(parentTableID string) ([]TextInput, error) {

	textInputs := []TextInput{}
	addTextInput := func(textInputID string, encodedProps string) error {

		textInputProps := newDefaultTextInputProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &textInputProps); decodeErr != nil {
			return fmt.Errorf("GetTextInputs: can't decode properties: %v", encodedProps)
		}

		currTextInput := TextInput{
			ParentTableID: parentTableID,
			TextInputID:   textInputID,
			ColumnID:      textInputID,
			Properties:    textInputProps,
			ColType:       textInputEntityKind}
		textInputs = append(textInputs, currTextInput)

		return nil
	}
	if getErr := common.GetTableColumns(textInputEntityKind, parentTableID, addTextInput); getErr != nil {
		return nil, fmt.Errorf("GetTextInputs: Can't get text boxes: %v")
	}

	return textInputs, nil

}

func CloneTextInputs(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcTextInputes, err := GetTextInputs(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneTextInputes: %v", err)
	}

	for _, srcTextInput := range srcTextInputes {
		remappedTextInputID := remappedIDs.AllocNewOrGetExistingRemappedID(srcTextInput.TextInputID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcTextInput.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneTextInputs: %v", err)
		}
		destProperties, err := srcTextInput.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneTextInputs: %v", err)
		}
		destTextInput := TextInput{
			ParentTableID: remappedFormID,
			TextInputID:   remappedTextInputID,
			ColumnID:      remappedTextInputID,
			Properties:    *destProperties,
			ColType:       textInputEntityKind}
		if err := saveTextInput(destTextInput); err != nil {
			return fmt.Errorf("CloneTextInputs: %v", err)
		}
	}

	return nil
}

func updateExistingTextInput(textInputID string, updatedTextInput *TextInput) (*TextInput, error) {

	if updateErr := common.UpdateTableColumn(textInputEntityKind, updatedTextInput.ParentTableID,
		updatedTextInput.TextInputID, updatedTextInput.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingTextInput: error updating existing text box component: %v", updateErr)
	}

	return updatedTextInput, nil

}
