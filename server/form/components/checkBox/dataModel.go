package checkBox

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const checkBoxEntityKind string = "CheckBox"

type CheckBox struct {
	ParentFormID string `json:"parentID"`
	CheckBoxID   string `json:"checkBoxID"`
	FieldID      string `json:"fieldID"`
	Geometry     common.LayoutGeometry
}

const checkBoxIDFieldName string = "CheckBoxID"
const checkBoxParentFormIDFieldName string = "ParentFormID"

type NewCheckBoxParams struct {
	ParentID string                `json:"parentID"`
	FieldID  string                `json:"fieldID"`
	Geometry common.LayoutGeometry `json:"geometry"`
}

func validCheckBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveNewCheckBox(appEngContext appengine.Context, params NewCheckBoxParams) (*CheckBox, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	field, fieldErr := field.GetField(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validCheckBoxFieldType(field.Type) {
		return nil, fmt.Errorf("saveNewCheckBox: Invalid field type: expecting bool field, got %v", field.Type)
	}

	newCheckBox := CheckBox{ParentFormID: params.ParentID,
		FieldID:    params.FieldID,
		CheckBoxID: uniqueID.GenerateUniqueID(),
		Geometry:   params.Geometry}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, checkBoxEntityKind, &newCheckBox)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new image component: error inserting into datastore: %v", insertErr)
	}

	log.Printf("INFO: API: New Checkbox: Created new check box container:  %+v", newCheckBox)

	return &newCheckBox, nil

}

func getCheckBox(appEngContext appengine.Context, checkBoxID string) (*CheckBox, error) {

	var checkBox CheckBox

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, checkBoxEntityKind,
		checkBoxIDFieldName, checkBoxID, &checkBox); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to checkbox container from datastore: error = %v", getErr)
	}

	return &checkBox, nil
}

func GetCheckBoxes(appEngContext appengine.Context, parentFormID string) ([]CheckBox, error) {

	var checkBoxes []CheckBox

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentFormID,
		checkBoxEntityKind, checkBoxParentFormIDFieldName, &checkBoxes)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve html editors: form id=%v", parentFormID)
	}

	return checkBoxes, nil
}

func updateExistingCheckBox(appEngContext appengine.Context, checkBoxID string, updatedCheckBox *CheckBox) (*CheckBox, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		checkBoxID, checkBoxEntityKind, checkBoxIDFieldName, updatedCheckBox); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating html editor: error = %v", updateErr)
	}

	return updatedCheckBox, nil

}
