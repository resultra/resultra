package datePicker

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
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

func saveNewDatePicker(params NewDatePickerParams) (*DatePicker, error) {

	if fieldErr := field.ValidateField(params.FieldID, validDatePickerFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewDatePicker: %v", fieldErr)
	}

	properties := newDefaultDatePickerProperties()
	properties.FieldID = params.FieldID

	datePickerID := uniqueID.GenerateSnowflakeID()
	newDatePicker := DatePicker{ParentTableID: params.ParentTableID,
		DatePickerID: datePickerID,
		ColumnID:     datePickerID,
		ColType:      datePickerEntityKind,
		Properties:   properties}

	if err := saveDatePicker(databaseWrapper.DBHandle(), newDatePicker); err != nil {
		return nil, fmt.Errorf("saveNewDatePicker: %v", err)
	}

	log.Printf("INFO: API: New DatePicker: Created new date picker container: %+v", newDatePicker)

	return &newDatePicker, nil

}

func getDatePicker(parentTableID string, datePickerID string) (*DatePicker, error) {

	datePickerProps := newDefaultDatePickerProperties()
	if getErr := common.GetTableColumn(datePickerEntityKind, parentTableID, datePickerID, &datePickerProps); getErr != nil {
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

func GetDatePickers(parentTableID string) ([]DatePicker, error) {
	return getDatePickersFromSrc(databaseWrapper.DBHandle(), parentTableID)
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

func updateExistingDatePicker(datePickerID string, updatedDatePicker *DatePicker) (*DatePicker, error) {

	if updateErr := common.UpdateTableColumn(datePickerEntityKind, updatedDatePicker.ParentTableID,
		updatedDatePicker.DatePickerID, updatedDatePicker.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingDatePicker: error updating existing date picker: %v", updateErr)
	}

	return updatedDatePicker, nil

}
