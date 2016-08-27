package datePicker

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
)

type DatePickerIDInterface interface {
	getDatePickerID() string
	getParentFormID() string
}

type DatePickerIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	DatePickerID string `json:"datePickerID"`
}

func (idHeader DatePickerIDHeader) getDatePickerID() string {
	return idHeader.DatePickerID
}

func (idHeader DatePickerIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type DatePickerPropUpdater interface {
	DatePickerIDInterface
	updateProps(datePicker *DatePicker) error
}

func updateDatePickerProps(propUpdater DatePickerPropUpdater) (*DatePicker, error) {

	// Retrieve the bar chart from the data store
	datePickerForUpdate, getErr := getDatePicker(propUpdater.getParentFormID(), propUpdater.getDatePickerID())
	if getErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to get existing date picker: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(datePickerForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to update existing date picker properties: %v", propUpdateErr)
	}

	datePicker, updateErr := updateExistingDatePicker(propUpdater.getDatePickerID(), datePickerForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to update existing date picker properties: datastore update error =  %v", updateErr)
	}

	return datePicker, nil
}

type DatePickerResizeParams struct {
	DatePickerIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams DatePickerResizeParams) updateProps(datePicker *DatePicker) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set date picker dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	datePicker.Properties.Geometry = updateParams.Geometry

	return nil
}
