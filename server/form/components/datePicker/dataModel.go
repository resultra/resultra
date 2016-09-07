package datePicker

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const datePickerEntityKind string = "datepicker"

type DatePickerProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

type DatePicker struct {
	ParentFormID string               `json:"parentFormID"`
	DatePickerID string               `json:"datePickerID"`
	Properties   DatePickerProperties `json:"properties"`
}

type NewDatePickerParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func validDatePickerFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeTime {
		return true
	} else {
		return false
	}
}

func saveNewDatePicker(params NewDatePickerParams) (*DatePicker, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validDatePickerFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewDatePicker: %v", compLinkErr)
	}

	properties := DatePickerProperties{
		Geometry:      params.Geometry,
		ComponentLink: params.ComponentLink}

	newDatePicker := DatePicker{ParentFormID: params.ParentFormID,
		DatePickerID: uniqueID.GenerateSnowflakeID(),
		Properties:   properties}

	if saveErr := common.SaveNewFormComponent(datePickerEntityKind,
		newDatePicker.ParentFormID, newDatePicker.DatePickerID, newDatePicker.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewDatePicker: Unable to save date picker with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New DatePicker: Created new date picker container: %+v", newDatePicker)

	return &newDatePicker, nil

}

func getDatePicker(parentFormID string, datePickerID string) (*DatePicker, error) {

	datePickerProps := DatePickerProperties{}
	if getErr := common.GetFormComponent(datePickerEntityKind, parentFormID, datePickerID, &datePickerProps); getErr != nil {
		return nil, fmt.Errorf("getDatePicker: Unable to retrieve date picker: %v", getErr)
	}

	datePicker := DatePicker{
		ParentFormID: parentFormID,
		DatePickerID: datePickerID,
		Properties:   datePickerProps}

	return &datePicker, nil
}

func GetDatePickers(parentFormID string) ([]DatePicker, error) {

	datePickers := []DatePicker{}
	addDatePicker := func(datePickerID string, encodedProps string) error {

		var datePickerProps DatePickerProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &datePickerProps); decodeErr != nil {
			return fmt.Errorf("GetDatePickers: can't decode properties: %v", encodedProps)
		}

		currDatePicker := DatePicker{
			ParentFormID: parentFormID,
			DatePickerID: datePickerID,
			Properties:   datePickerProps}
		datePickers = append(datePickers, currDatePicker)

		return nil
	}
	if getErr := common.GetFormComponents(datePickerEntityKind, parentFormID, addDatePicker); getErr != nil {
		return nil, fmt.Errorf("GetDatePickers: Can't get date pickers: %v")
	}

	return datePickers, nil

}

func updateExistingDatePicker(datePickerID string, updatedDatePicker *DatePicker) (*DatePicker, error) {

	if updateErr := common.UpdateFormComponent(datePickerEntityKind, updatedDatePicker.ParentFormID,
		updatedDatePicker.DatePickerID, updatedDatePicker.Properties); updateErr != nil {
	}

	return updatedDatePicker, nil

}
