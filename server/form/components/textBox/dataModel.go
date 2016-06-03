package textBox

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const textBoxEntityKind string = "TextBox"

type TextBox struct {
	ParentFormID string `json:"parentID"`
	TextBoxID    string `json:"textBoxID"`
	FieldID      string `json:"fieldID"`
	Geometry     common.LayoutGeometry
}

const textBoxParentFormIDFieldName string = "ParentFormID"
const textBoxTextBoxIDFieldName string = "TextBoxID"

type NewTextBoxParams struct {
	ParentID string                `json:"parentID"`
	FieldID  string                `json:"fieldID"`
	Geometry common.LayoutGeometry `json:"geometry"`
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

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	field, fieldErr := field.GetField(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validTextBoxFieldType(field.Type) {
		return nil, fmt.Errorf("NewTextBox: Invalid field type: expecting text or number field, got %v", field.Type)
	}

	newTextBox := TextBox{ParentFormID: params.ParentID,
		FieldID:   params.FieldID,
		TextBoxID: uniqueID.GenerateUniqueID(),
		Geometry:  params.Geometry}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, textBoxEntityKind, &newTextBox)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new text box component: error inserting into datastore: %v", insertErr)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newTextBox)

	return &newTextBox, nil

}

func getTextBox(appEngContext appengine.Context, textBoxID string) (*TextBox, error) {

	var textBox TextBox

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, textBoxEntityKind,
		textBoxTextBoxIDFieldName, textBoxID, &textBox); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to image container from datastore: error = %v", getErr)
	}

	return &textBox, nil
}

func GetTextBoxes(appEngContext appengine.Context, parentFormID string) ([]TextBox, error) {

	var textBoxes []TextBox

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentFormID,
		textBoxEntityKind, textBoxParentFormIDFieldName, &textBoxes)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve text box components: form id=%v", parentFormID)
	}

	return textBoxes, nil

}

func updateExistingTextBox(appEngContext appengine.Context, textBoxID string, updatedTextBox *TextBox) (*TextBox, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		textBoxID, textBoxEntityKind, textBoxTextBoxIDFieldName, updatedTextBox); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating text box: error = %v", updateErr)
	}

	return updatedTextBox, nil

}
