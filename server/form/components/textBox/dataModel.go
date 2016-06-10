package textBox

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	geometry "resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
)

const textBoxEntityKind string = "textbox"

type TextBoxProperties struct {
	FieldID  string                  `json:"fieldID"`
	Geometry geometry.LayoutGeometry `json:"geometry"`
}

type TextBox struct {
	ParentFormID string            `json:"parentID"`
	TextBoxID    string            `json:"textBoxID"`
	Properties   TextBoxProperties `json:"properties"`
}

type NewTextBoxParams struct {
	ParentFormID       string                  `json:"parentFormID"`
	FieldParentTableID string                  `json:"fieldParentTableID"`
	FieldID            string                  `json:"fieldID"`
	Geometry           geometry.LayoutGeometry `json:"geometry"`
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

func saveNewTextBox(appEngContext appengine.Context, params NewTextBoxParams) (*TextBox, error) {

	if !geometry.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	field, fieldErr := field.GetField(appEngContext, params.FieldParentTableID, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validTextBoxFieldType(field.Type) {
		return nil, fmt.Errorf("NewTextBox: Invalid field type: expecting text or number field, got %v", field.Type)
	}

	properties := TextBoxProperties{
		Geometry: params.Geometry,
		FieldID:  params.FieldID}

	newTextBox := TextBox{ParentFormID: params.ParentFormID,
		TextBoxID:  gocql.TimeUUID().String(),
		Properties: properties}

	if saveErr := common.SaveNewFormComponent(textBoxEntityKind,
		newTextBox.ParentFormID, newTextBox.TextBoxID, newTextBox.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: Unable to save text box with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newTextBox)

	return &newTextBox, nil

}

func getTextBox(appEngContext appengine.Context, parentFormID string, textBoxID string) (*TextBox, error) {

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

func GetTextBoxes(appEngContext appengine.Context, parentFormID string) ([]TextBox, error) {

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

func updateExistingTextBox(appEngContext appengine.Context, textBoxID string, updatedTextBox *TextBox) (*TextBox, error) {

	if updateErr := common.UpdateFormComponent(textBoxEntityKind, updatedTextBox.ParentFormID,
		updatedTextBox.TextBoxID, updatedTextBox.Properties); updateErr != nil {
	}

	return updatedTextBox, nil

}
