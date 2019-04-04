// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package datePicker

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
	"log"
)

const datePickerEntityKind string = "datepicker"

type DatePicker struct {
	ParentFormID string               `json:"parentFormID"`
	DatePickerID string               `json:"datePickerID"`
	Properties   DatePickerProperties `json:"properties"`
}

type NewDatePickerParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validDatePickerFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeTime {
		return true
	} else {
		return false
	}
}

func saveDatePicker(destDBHandle *sql.DB, newDatePicker DatePicker) error {

	if saveErr := common.SaveNewFormComponent(destDBHandle, datePickerEntityKind,
		newDatePicker.ParentFormID, newDatePicker.DatePickerID, newDatePicker.Properties); saveErr != nil {
		return fmt.Errorf("saveNewDatePicker: Unable to save date picker: error = %v", saveErr)
	}
	return nil

}

func saveNewDatePicker(trackerDBHandle *sql.DB, params NewDatePickerParams) (*DatePicker, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validDatePickerFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewDatePicker: %v", fieldErr)
	}

	properties := newDefaultDatePickerProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newDatePicker := DatePicker{ParentFormID: params.ParentFormID,
		DatePickerID: uniqueID.GenerateUniqueID(),
		Properties:   properties}

	if err := saveDatePicker(trackerDBHandle, newDatePicker); err != nil {
		return nil, fmt.Errorf("saveNewDatePicker: %v", err)
	}

	log.Printf("INFO: API: New DatePicker: Created new date picker container: %+v", newDatePicker)

	return &newDatePicker, nil

}

func getDatePicker(trackerDBHandle *sql.DB, parentFormID string, datePickerID string) (*DatePicker, error) {

	datePickerProps := newDefaultDatePickerProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, datePickerEntityKind, parentFormID, datePickerID, &datePickerProps); getErr != nil {
		return nil, fmt.Errorf("getDatePicker: Unable to retrieve date picker: %v", getErr)
	}

	datePicker := DatePicker{
		ParentFormID: parentFormID,
		DatePickerID: datePickerID,
		Properties:   datePickerProps}

	return &datePicker, nil
}

func getDatePickersFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]DatePicker, error) {

	datePickers := []DatePicker{}
	addDatePicker := func(datePickerID string, encodedProps string) error {

		datePickerProps := newDefaultDatePickerProperties()
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
	if getErr := common.GetFormComponents(srcDBHandle, datePickerEntityKind, parentFormID, addDatePicker); getErr != nil {
		return nil, fmt.Errorf("GetDatePickers: Can't get date pickers: %v")
	}

	return datePickers, nil

}

func GetDatePickers(trackerDBHandle *sql.DB, parentFormID string) ([]DatePicker, error) {
	return getDatePickersFromSrc(trackerDBHandle, parentFormID)
}

func CloneDatePickers(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcDatePickers, err := getDatePickersFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneDatePickers: %v", err)
	}

	for _, srcDatePicker := range srcDatePickers {
		remappedDatePickerID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcDatePicker.DatePickerID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcDatePicker.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneDatePickers: %v", err)
		}
		destProperties, err := srcDatePicker.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneDatePickers: %v", err)
		}
		destDatePicker := DatePicker{
			ParentFormID: remappedFormID,
			DatePickerID: remappedDatePickerID,
			Properties:   *destProperties}
		if err := saveDatePicker(cloneParams.DestDBHandle, destDatePicker); err != nil {
			return fmt.Errorf("CloneDatePickers: %v", err)
		}
	}

	return nil
}

func updateExistingDatePicker(trackerDBHandle *sql.DB, datePickerID string, updatedDatePicker *DatePicker) (*DatePicker, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, datePickerEntityKind, updatedDatePicker.ParentFormID,
		updatedDatePicker.DatePickerID, updatedDatePicker.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingDatePicker: error updating existing date picker: %v", updateErr)
	}

	return updatedDatePicker, nil

}
