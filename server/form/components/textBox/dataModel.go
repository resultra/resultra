package textBox

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const textBoxEntityKind string = "textbox"

type TextBox struct {
	ParentFormID string            `json:"parentFormID"`
	TextBoxID    string            `json:"textBoxID"`
	Properties   TextBoxProperties `json:"properties"`
}

type NewTextBoxParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validTextBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeText {
		return true
	} else if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveTextBox(newTextBox TextBox) error {
	if saveErr := common.SaveNewFormComponent(textBoxEntityKind,
		newTextBox.ParentFormID, newTextBox.TextBoxID, newTextBox.Properties); saveErr != nil {
		return fmt.Errorf("saveTextBox: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewTextBox(params NewTextBoxParams) (*TextBox, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := common.ValidateField(params.FieldID, validTextBoxFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: %v", fieldErr)
	}

	properties := newDefaultTextBoxProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newTextBox := TextBox{ParentFormID: params.ParentFormID,
		TextBoxID:  uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveTextBox(newTextBox); err != nil {
		return nil, fmt.Errorf("saveNewTextBox: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newTextBox)

	return &newTextBox, nil

}

func getTextBox(parentFormID string, textBoxID string) (*TextBox, error) {

	textBoxProps := newDefaultTextBoxProperties()
	if getErr := common.GetFormComponent(textBoxEntityKind, parentFormID, textBoxID, &textBoxProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	textBox := TextBox{
		ParentFormID: parentFormID,
		TextBoxID:    textBoxID,
		Properties:   textBoxProps}

	return &textBox, nil
}

func GetTextBoxes(parentFormID string) ([]TextBox, error) {

	textBoxes := []TextBox{}
	addTextBox := func(textBoxID string, encodedProps string) error {

		textBoxProps := newDefaultTextBoxProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &textBoxProps); decodeErr != nil {
			return fmt.Errorf("GetTextBoxes: can't decode properties: %v", encodedProps)
		}

		currTextBox := TextBox{
			ParentFormID: parentFormID,
			TextBoxID:    textBoxID,
			Properties:   textBoxProps}
		textBoxes = append(textBoxes, currTextBox)

		return nil
	}
	if getErr := common.GetFormComponents(textBoxEntityKind, parentFormID, addTextBox); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return textBoxes, nil

}

func CloneTextBoxes(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcTextBoxes, err := GetTextBoxes(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneTextBoxes: %v", err)
	}

	for _, srcTextBox := range srcTextBoxes {
		remappedTextBoxID := remappedIDs.AllocNewOrGetExistingRemappedID(srcTextBox.TextBoxID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcTextBox.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneTextBoxes: %v", err)
		}
		destProperties, err := srcTextBox.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneTextBoxes: %v", err)
		}
		destTextBox := TextBox{
			ParentFormID: remappedFormID,
			TextBoxID:    remappedTextBoxID,
			Properties:   *destProperties}
		if err := saveTextBox(destTextBox); err != nil {
			return fmt.Errorf("CloneTextBoxes: %v", err)
		}
	}

	return nil
}

func updateExistingTextBox(textBoxID string, updatedTextBox *TextBox) (*TextBox, error) {

	if updateErr := common.UpdateFormComponent(textBoxEntityKind, updatedTextBox.ParentFormID,
		updatedTextBox.TextBoxID, updatedTextBox.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingTextBox: error updating existing text box component: %v", updateErr)
	}

	return updatedTextBox, nil

}
