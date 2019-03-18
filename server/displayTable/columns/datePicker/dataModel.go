// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package datePicker

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/displayTable/columns/common"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

const datePickerEntityKind string = "datePicker"

type DatePicker struct {
	ParentTableID string               `json:"parentTableID"`
	DatePickerID  string               `json:"datePickerID"`
	ColType       string               `json:"colType"`
	ColumnID      string               `json:"columnID"`
	Properties    DatePickerProperties `json:"properties"`
}

type NewDatePickerParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validDatePickerFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeTime {
		return true
	} else {
		return false
	}
}

func saveDatePicker(destDBHandle *sql.DB, newDatePicker DatePicker) error {

	if saveErr := common.SaveNewTableColumn(destDBHandle, datePickerEntityKind,
		newDatePicker.ParentTableID, newDatePicker.DatePickerID, newDatePicker.Properties); saveErr != nil {
		return fmt.Errorf("saveNewDatePicker: Unable to save date picker: error = %v", saveErr)
	}
	return nil

}

func saveNewDatePicker(trackerDBHandle *sql.DB, params NewDatePickerParams) (*DatePicker, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validDatePickerFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewDatePicker: %v", fieldErr)
	}

	properties := newDefaultDatePickerProperties()
	properties.FieldID = params.FieldID

	datePickerID := uniqueID.GenerateUniqueID()
	newDatePicker := DatePicker{ParentTableID: params.ParentTableID,
		DatePickerID: datePickerID,
		ColumnID:     datePickerID,
		ColType:      datePickerEntityKind,
		Properties:   properties}

	if err := saveDatePicker(trackerDBHandle, newDatePicker); err != nil {
		return nil, fmt.Errorf("saveNewDatePicker: %v", err)
	}

	log.Printf("INFO: API: New DatePicker: Created new date picker container: %+v", newDatePicker)

	return &newDatePicker, nil

}

func getDatePicker(trackerDBHandle *sql.DB, parentTableID string, datePickerID string) (*DatePicker, error) {

	datePickerProps := newDefaultDatePickerProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, datePickerEntityKind, parentTableID, datePickerID, &datePickerProps); getErr != nil {
		return nil, fmt.Errorf("getDatePicker: Unable to retrieve date picker: %v", getErr)
	}

	datePicker := DatePicker{
		ParentTableID: parentTableID,
		DatePickerID:  datePickerID,
		ColumnID:      datePickerID,
		ColType:       datePickerEntityKind,
		Properties:    datePickerProps}

	return &datePicker, nil
}

func getDatePickersFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]DatePicker, error) {

	datePickers := []DatePicker{}
	addDatePicker := func(datePickerID string, encodedProps string) error {

		datePickerProps := newDefaultDatePickerProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &datePickerProps); decodeErr != nil {
			return fmt.Errorf("GetDatePickers: can't decode properties: %v", encodedProps)
		}

		currDatePicker := DatePicker{
			ParentTableID: parentTableID,
			DatePickerID:  datePickerID,
			ColumnID:      datePickerID,
			ColType:       datePickerEntityKind,
			Properties:    datePickerProps}
		datePickers = append(datePickers, currDatePicker)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, datePickerEntityKind, parentTableID, addDatePicker); getErr != nil {
		return nil, fmt.Errorf("GetDatePickers: Can't get date pickers: %v")
	}

	return datePickers, nil

}

func GetDatePickers(trackerDBHandle *sql.DB, parentTableID string) ([]DatePicker, error) {
	return getDatePickersFromSrc(trackerDBHandle, parentTableID)
}

func CloneDatePickers(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcDatePickers, err := getDatePickersFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneDatePickers: %v", err)
	}

	for _, srcDatePicker := range srcDatePickers {
		remappedDatePickerID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcDatePicker.DatePickerID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcDatePicker.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneDatePickers: %v", err)
		}
		destProperties, err := srcDatePicker.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneDatePickers: %v", err)
		}
		destDatePicker := DatePicker{
			ParentTableID: remappedFormID,
			DatePickerID:  remappedDatePickerID,
			ColumnID:      remappedDatePickerID,
			ColType:       datePickerEntityKind,
			Properties:    *destProperties}
		if err := saveDatePicker(cloneParams.DestDBHandle, destDatePicker); err != nil {
			return fmt.Errorf("CloneDatePickers: %v", err)
		}
	}

	return nil
}

func updateExistingDatePicker(trackerDBHandle *sql.DB, datePickerID string, updatedDatePicker *DatePicker) (*DatePicker, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, datePickerEntityKind, updatedDatePicker.ParentTableID,
		updatedDatePicker.DatePickerID, updatedDatePicker.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingDatePicker: error updating existing date picker: %v", updateErr)
	}

	return updatedDatePicker, nil

}
