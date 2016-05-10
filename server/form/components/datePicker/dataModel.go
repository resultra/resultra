package datePicker

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const datePickerEntityKind string = "DatePicker"

type DatePicker struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

type DatePickerRef struct {
	DatePickerID string                `json:"datePickerID"`
	FieldRef     field.FieldRef        `json:"fieldRef"`
	Geometry     common.LayoutGeometry `json:"geometry"`
}

type NewDatePickerParams struct {
	// ContainerID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	ParentID string                `json:"parentID"`
	FieldID  string                `json:"fieldID"`
	Geometry common.LayoutGeometry `json:"geometry"`
}

func validDatePickerFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeTime {
		return true
	} else {
		return false
	}
}

func saveNewDatePicker(appEngContext appengine.Context, params NewDatePickerParams) (*DatePickerRef, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("saveNewDatePicker: Can't text box with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validDatePickerFieldType(fieldRef.FieldInfo.Type) {
		return nil, fmt.Errorf("saveNewDatePicker: Invalid field type: expecting time field, got %v", fieldRef.FieldInfo.Type)
	}

	newDatePicker := DatePicker{Field: fieldKey, Geometry: params.Geometry}

	datePickerID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext, params.ParentID, datePickerEntityKind, &newDatePicker)
	if insertErr != nil {
		return nil, insertErr
	}

	datePickerRef := DatePickerRef{
		DatePickerID: datePickerID,
		FieldRef:     *fieldRef,
		Geometry:     params.Geometry}

	log.Printf("INFO: API: New DatePicker: Created new date picker container: id=%v params=%+v", datePickerID, params)

	return &datePickerRef, nil

}

func getDatePicker(appEngContext appengine.Context, datePickerID string) (*DatePicker, error) {

	var datePicker DatePicker
	if getErr := datastoreWrapper.GetEntity(appEngContext, datePickerID, &datePicker); getErr != nil {
		return nil, fmt.Errorf("getDatePicker: Unable to get date picker from datastore: error = %v", getErr)
	}
	return &datePicker, nil
}

func GetDatePickers(appEngContext appengine.Context, parentFormID string) ([]DatePickerRef, error) {

	var datePickers []DatePicker
	datePickerIDs, getErr := datastoreWrapper.GetAllChildEntities(appEngContext, parentFormID, datePickerEntityKind, &datePickers)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve date pickers: form id=%v", parentFormID)
	}

	datePickerRefs := make([]DatePickerRef, len(datePickers))
	for datePickerIter, currDatePicker := range datePickers {

		datePickerID := datePickerIDs[datePickerIter]

		fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, currDatePicker.Field)
		if fieldErr != nil {
			return nil, fmt.Errorf("GetDatePickers: Error retrieving field for date picker: error = %v", fieldErr)
		}

		datePickerRefs[datePickerIter] = DatePickerRef{
			DatePickerID: datePickerID,
			FieldRef:     *fieldRef,
			Geometry:     currDatePicker.Geometry}

	} // for each check box
	return datePickerRefs, nil

}

func updateExistingDatePicker(appEngContext appengine.Context, datePickerID string, updatedDatePicker *DatePicker) (*DatePickerRef, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntity(appEngContext, datePickerID, updatedDatePicker); updateErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: Error updating check box: error = %v", updateErr)
	}

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, updatedDatePicker.Field)
	if fieldErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: Error retrieving field for check box: error = %v", fieldErr)
	}

	datePickerRef := DatePickerRef{
		DatePickerID: datePickerID,
		FieldRef:     *fieldRef,
		Geometry:     updatedDatePicker.Geometry}

	return &datePickerRef, nil

}
