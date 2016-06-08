package datePicker

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const datePickerEntityKind string = "DatePicker"

type DatePicker struct {
	ParentFormID string                `json:"parentFormID"`
	DatePickerID string                `json:"datePickerID"`
	FieldID      string                `json:"fieldID"`
	Geometry     common.LayoutGeometry `json:"geometry"`
}

const datePickerIDFieldName string = "DatePickerID"
const datePickerParentFormIDFieldName string = "ParentFormID"

type NewDatePickerParams struct {
	FieldParentTableID string                `json:"fieldParentTableID"`
	ParentID           string                `json:"parentID"`
	FieldID            string                `json:"fieldID"`
	Geometry           common.LayoutGeometry `json:"geometry"`
}

func validDatePickerFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeTime {
		return true
	} else {
		return false
	}
}

func saveNewDatePicker(appEngContext appengine.Context, params NewDatePickerParams) (*DatePicker, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	field, fieldErr := field.GetField(appEngContext, params.FieldParentTableID, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validDatePickerFieldType(field.Type) {
		return nil, fmt.Errorf("saveNewDatePicker: Invalid field type: expecting time field, got %v", field.Type)
	}

	newDatePicker := DatePicker{ParentFormID: params.ParentID,
		FieldID:      params.FieldID,
		DatePickerID: uniqueID.GenerateUniqueID(),
		Geometry:     params.Geometry}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, datePickerEntityKind, &newDatePicker)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new image component: error inserting into datastore: %v", insertErr)
	}

	log.Printf("INFO: API: New DatePicker: Created new date picker container: %+v", newDatePicker)

	return &newDatePicker, nil

}

func getDatePicker(appEngContext appengine.Context, datePickerID string) (*DatePicker, error) {

	var datePicker DatePicker

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, datePickerEntityKind,
		datePickerIDFieldName, datePickerID, &datePicker); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to checkbox container from datastore: error = %v", getErr)
	}

	return &datePicker, nil
}

func GetDatePickers(appEngContext appengine.Context, parentFormID string) ([]DatePicker, error) {

	var datePickers []DatePicker

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentFormID,
		datePickerEntityKind, datePickerParentFormIDFieldName, &datePickers)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve date picker components: form id=%v", parentFormID)
	}

	return datePickers, nil

}

func updateExistingDatePicker(appEngContext appengine.Context, datePickerID string, updatedDatePicker *DatePicker) (*DatePicker, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		datePickerID, datePickerEntityKind, datePickerIDFieldName, updatedDatePicker); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating date picker: error = %v", updateErr)
	}

	return updatedDatePicker, nil

}
