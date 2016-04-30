package textBox

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/dataModel"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const textBoxEntityKind string = "TextBox"

func textBoxChildParentEntityRel() datastoreWrapper.ChildParentEntityRel {
	return datastoreWrapper.ChildParentEntityRel{ParentEntityKind: dataModel.LayoutEntityKind, ChildEntityKind: textBoxEntityKind}
}

type TextBox struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

type TextBoxRef struct {
	TextBoxID string                `json:"textBoxID"`
	FieldRef  field.FieldRef        `json:"fieldRef"`
	Geometry  common.LayoutGeometry `json:"geometry"`
}

type NewTextBoxParams struct {
	// ContainerID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
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

func saveNewTextBox(appEngContext appengine.Context, params NewTextBoxParams) (*TextBoxRef, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewTextBox: Can't text box with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validTextBoxFieldType(fieldRef.FieldInfo.Type) {
		return nil, fmt.Errorf("NewTextBox: Invalid field type: expecting text or number field, got %v", fieldRef.FieldInfo.Type)
	}

	newTextBox := TextBox{Field: fieldKey, Geometry: params.Geometry}

	textBoxID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext, params.ParentID, textBoxChildParentEntityRel(), &newTextBox)
	if insertErr != nil {
		return nil, insertErr
	}

	textBoxRef := TextBoxRef{
		TextBoxID: textBoxID,
		FieldRef:  *fieldRef,
		Geometry:  params.Geometry}

	log.Printf("INFO: API: NewLayout: Created new Layout container: id=%v params=%+v", textBoxID, params)

	return &textBoxRef, nil

}

func getTextBox(appEngContext appengine.Context, textBoxID string) (*TextBox, error) {

	var textBox TextBox
	if getErr := datastoreWrapper.GetChildEntity(appEngContext, textBoxID, textBoxChildParentEntityRel(), &textBox); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to get bar chart from datastore: error = %v", getErr)
	}
	return &textBox, nil
}

func GetTextBoxes(appEngContext appengine.Context, parentFormID string) ([]TextBoxRef, error) {

	var textBoxes []TextBox
	textBoxIDs, getErr := datastoreWrapper.GetAllChildEntities(appEngContext, parentFormID, textBoxChildParentEntityRel(), &textBoxes)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve layout containers: form id=%v", parentFormID)
	}

	textBoxRefs := make([]TextBoxRef, len(textBoxes))
	for textBoxIter, currTextBox := range textBoxes {

		textBoxID := textBoxIDs[textBoxIter]

		fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, currTextBox.Field)
		if fieldErr != nil {
			return nil, fmt.Errorf("GetTextBoxes: Error retrieving field for text box: error = %v", fieldErr)
		}

		textBoxRefs[textBoxIter] = TextBoxRef{
			TextBoxID: textBoxID,
			FieldRef:  *fieldRef,
			Geometry:  currTextBox.Geometry}

	} // for each text box
	return textBoxRefs, nil

}

func updateExistingTextBox(appEngContext appengine.Context, textBoxID string, updatedTextBox *TextBox) (*TextBoxRef, error) {

	if updateErr := datastoreWrapper.UpdateExistingChildEntity(appEngContext, textBoxID,
		textBoxChildParentEntityRel(), updatedTextBox); updateErr != nil {
		return nil, fmt.Errorf("updateExistingTextBox: Error updating text box: error = %v", updateErr)
	}

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, updatedTextBox.Field)
	if fieldErr != nil {
		return nil, fmt.Errorf("updateExistingTextBox: Error retrieving field for text box: error = %v", fieldErr)
	}

	textBoxRef := TextBoxRef{
		TextBoxID: textBoxID,
		FieldRef:  *fieldRef,
		Geometry:  updatedTextBox.Geometry}

	return &textBoxRef, nil

}
