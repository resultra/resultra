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
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

type TextBox struct {
	ParentFormID string            `json:"parentID"`
	TextBoxID    string            `json:"textBoxID"`
	Properties   TextBoxProperties `json:"properties"`
}

type NewTextBoxParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
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

func saveNewTextBox(params NewTextBoxParams) (*TextBox, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validTextBoxFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: %v", compLinkErr)
	}

	properties := TextBoxProperties{
		Geometry:      params.Geometry,
		ComponentLink: params.ComponentLink}

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
