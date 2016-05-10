package datePicker

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/common"
)

type DatePickerIDInterface interface {
	getDatePickerID() string
}

type DatePickerIDHeader struct {
	DatePickerID string `json:"datePickerID"`
}

func (idHeader DatePickerIDHeader) getDatePickerID() string {
	return idHeader.DatePickerID
}

type DatePickerPropUpdater interface {
	DatePickerIDInterface
	updateProps(datePicker *DatePicker) error
}

func updateDatePickerProps(appEngContext appengine.Context, propUpdater DatePickerPropUpdater) (*DatePickerRef, error) {

	// Retrieve the bar chart from the data store
	datePickerForUpdate, getErr := getDatePicker(appEngContext, propUpdater.getDatePickerID())
	if getErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to get existing date picker: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(datePickerForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to update existing date picker properties: %v", propUpdateErr)
	}

	datePickerRef, updateErr := updateExistingDatePicker(appEngContext, propUpdater.getDatePickerID(), datePickerForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to update existing date picker properties: datastore update error =  %v", updateErr)
	}

	return datePickerRef, nil
}

type DatePickerResizeParams struct {
	DatePickerIDHeader
	Geometry common.LayoutGeometry `json:"geometry"`
}

func (updateParams DatePickerResizeParams) updateProps(datePicker *DatePicker) error {

	if !common.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set date picker dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	datePicker.Geometry = updateParams.Geometry

	return nil
}

type DatePickerRepositionParams struct {
	DatePickerIDHeader
	Position common.LayoutPosition `json:"position"`
}

func (updateParams DatePickerRepositionParams) updateProps(datePicker *DatePicker) error {

	if err := datePicker.Geometry.SetPosition(updateParams.Position); err != nil {
		return fmt.Errorf("Error setting position for date picker: Invalid geometry: %v", err)
	}

	return nil
}
