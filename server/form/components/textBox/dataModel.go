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

type TextBoxProperties struct {
	LinkedValType string                         `json:"linkedValType"`
	FieldID       string                         `json:"fieldID"`
	GlobalID      string                         `json:"globalID"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

type TextBox struct {
	ParentFormID string            `json:"parentID"`
	TextBoxID    string            `json:"textBoxID"`
	Properties   TextBoxProperties `json:"properties"`
}

type NewTextBoxParams struct {
	ParentFormID       string                         `json:"parentFormID"`
	FieldParentTableID string                         `json:"fieldParentTableID"`
	LinkedValType      string                         `json:"linkedValType"`
	FieldID            string                         `json:"fieldID"`
	GlobalID           string                         `json:"globalID"`
	Geometry           componentLayout.LayoutGeometry `json:"geometry"`
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

func validLinkedValType(valType string) bool {
	if valType == "global" || valType == "field" {
		return true
	} else {
		return false
	}
}

func saveNewTextBox(params NewTextBoxParams) (*TextBox, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if !validLinkedValType(params.LinkedValType) {
		return nil, fmt.Errorf("Invalid text box parameters (linked value type): %v", params.LinkedValType)
	}

	field, fieldErr := field.GetField(params.FieldParentTableID, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validTextBoxFieldType(field.Type) {
		return nil, fmt.Errorf("NewTextBox: Invalid field type: expecting text or number field, got %v", field.Type)
	}

	properties := TextBoxProperties{
		Geometry:      params.Geometry,
		LinkedValType: params.LinkedValType,
		FieldID:       params.FieldID,
		GlobalID:      params.GlobalID}

	newTextBox := TextBox{ParentFormID: params.ParentFormID,
		TextBoxID:  uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := common.SaveNewFormComponent(textBoxEntityKind,
		newTextBox.ParentFormID, newTextBox.TextBoxID, newTextBox.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: Unable to save text box with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newTextBox)

	return &newTextBox, nil

}

func getTextBox(parentFormID string, textBoxID string) (*TextBox, error) {

	textBoxProps := TextBoxProperties{}
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

		var textBoxProps TextBoxProperties
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

func updateExistingTextBox(textBoxID string, updatedTextBox *TextBox) (*TextBox, error) {

	if updateErr := common.UpdateFormComponent(textBoxEntityKind, updatedTextBox.ParentFormID,
		updatedTextBox.TextBoxID, updatedTextBox.Properties); updateErr != nil {
	}

	return updatedTextBox, nil

}
