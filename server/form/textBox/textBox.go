package textBox

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/common/datastoreWrapper"
	"resultra/datasheet/server/dataModel"
	"resultra/datasheet/server/field"
)

const textBoxEntityKind string = "TextBox"

type TextBox struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

type TextBoxRef struct {
	datastoreWrapper.UniqueIDHeader
	FieldRef field.FieldRef        `json:"fieldRef"`
	Geometry common.LayoutGeometry `json:"geometry"`
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

func NewTextBox(appEngContext appengine.Context, params NewTextBoxParams) (*TextBoxRef, error) {
	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	parentKey, err := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.LayoutEntityKind,
		params.ParentID)
	if err != nil {
		return nil, err
	}

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, field.GetFieldParams{params.FieldID})
	if fieldErr != nil {
		return nil, fmt.Errorf("NewTextBox: Can't text box with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validTextBoxFieldType(fieldRef.FieldInfo.Type) {
		return nil, fmt.Errorf("NewTextBox: Invalid field type: expecting text or number field, got %v", fieldRef.FieldInfo.Type)
	}

	newTextBox := TextBox{Field: fieldKey, Geometry: params.Geometry}

	textBoxID, insertErr := datastoreWrapper.InsertNewEntity(appEngContext, textBoxEntityKind,
		parentKey, &newTextBox)
	if insertErr != nil {
		return nil, insertErr
	}

	textBoxRef := TextBoxRef{
		UniqueIDHeader: datastoreWrapper.NewUniqueIDHeader(params.ParentID, textBoxID),
		FieldRef:       *fieldRef,
		Geometry:       params.Geometry}

	log.Printf("INFO: API: NewLayout: Created new Layout container: id=%v params=%+v", textBoxID, params)

	return &textBoxRef, nil

}

func getTextBox(appEngContext appengine.Context, textBoxID datastoreWrapper.UniqueID) (*TextBox, error) {

	parentKey, parentKeyErr := datastoreWrapper.NewRootEntityKey(appEngContext, dataModel.LayoutEntityKind, textBoxID.ParentID)
	if parentKeyErr != nil {
		return nil, fmt.Errorf("getTextBox: unable to retrieve parent key for dashboard = %v", textBoxID.ParentID)
	}

	var textBox TextBox
	if getErr := datastoreWrapper.GetChildEntityByID(textBoxID.ObjectID, appEngContext, textBoxEntityKind, parentKey, &textBox); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to get bar chart from datastore: error = %v", getErr)
	}

	return &textBox, nil
}

func GetTextBoxes(appEngContext appengine.Context, parentFormID datastoreWrapper.UniqueRootID) ([]TextBoxRef, error) {

	parentKey, parentErr := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.LayoutEntityKind, parentFormID.ObjectID)
	if parentErr != nil {
		return nil, fmt.Errorf("GetTextBoxes: Unable to retrieve text boxes: Unable to retrieve parent form: error = %v", parentErr)
	}

	textBoxQuery := datastore.NewQuery(textBoxEntityKind).Ancestor(parentKey)
	var textBoxes []TextBox
	keys, getErr := textBoxQuery.GetAll(appEngContext, &textBoxes)

	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve layout containers: form id=%v", parentFormID.ObjectID)
	} else {

		textBoxRefs := make([]TextBoxRef, len(textBoxes))
		for textBoxIter, currTextBox := range textBoxes {

			textBoxKey := keys[textBoxIter]
			textBoxID, encodeErr := datastoreWrapper.EncodeUniqueEntityIDToStr(textBoxKey)
			if encodeErr != nil {
				return nil, fmt.Errorf("Failed to encode unique ID for record: key=%+v, encode err=%v", textBoxKey, encodeErr)
			}

			fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, currTextBox.Field)
			if fieldErr != nil {
				return nil, fmt.Errorf("GetTextBoxes: Error retrieving field for text box: error = %v", fieldErr)
			}

			textBoxRefs[textBoxIter] = TextBoxRef{
				UniqueIDHeader: datastoreWrapper.NewUniqueIDHeader(parentFormID.ObjectID, textBoxID),
				FieldRef:       *fieldRef,
				Geometry:       currTextBox.Geometry}

		} // for each text box
		return textBoxRefs, nil
	}

}

func updateExistingTextBox(appEngContext appengine.Context, uniqueID datastoreWrapper.UniqueID, updatedTextBox *TextBox) (*TextBoxRef, error) {

	parentKey, getErr := datastoreWrapper.GetExistingRootEntityKey(appEngContext, dataModel.LayoutEntityKind, uniqueID.ParentID)
	if getErr != nil {
		return nil, fmt.Errorf("updateExistingTextBox: Invalid parent for text box: %v", getErr)
	}

	if updateErr := datastoreWrapper.UpdateExistingEntity(appEngContext,
		uniqueID.ObjectID, textBoxEntityKind, parentKey, updatedTextBox); updateErr != nil {
		return nil, updateErr
	}

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, updatedTextBox.Field)
	if fieldErr != nil {
		return nil, fmt.Errorf("updateExistingTextBox: Error retrieving field for text box: error = %v", fieldErr)
	}

	textBoxRef := TextBoxRef{
		UniqueIDHeader: datastoreWrapper.UniqueIDHeader{uniqueID},
		FieldRef:       *fieldRef,
		Geometry:       updatedTextBox.Geometry}

	return &textBoxRef, nil

}
