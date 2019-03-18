// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package datePicker

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/common/inputProps"
	"resultra/tracker/server/form/components/common"
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

func updateDatePickerProps(trackerDBHandle *sql.DB, propUpdater DatePickerPropUpdater) (*DatePicker, error) {

	// Retrieve the bar chart from the data store
	datePickerForUpdate, getErr := getDatePicker(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getDatePickerID())
	if getErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to get existing date picker: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(datePickerForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateDatePickerProps: Unable to update existing date picker properties: %v", propUpdateErr)
	}

	datePicker, updateErr := updateExistingDatePicker(trackerDBHandle, propUpdater.getDatePickerID(), datePickerForUpdate)
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

type DatePickerFormatParams struct {
	DatePickerIDHeader
	DateFormat string `json:"dateFormat"`
}

func (updateParams DatePickerFormatParams) updateProps(datePicker *DatePicker) error {

	// TODO - Validate format

	datePicker.Properties.DateFormat = updateParams.DateFormat

	return nil
}

type DatePickerLabelFormatParams struct {
	DatePickerIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams DatePickerLabelFormatParams) updateProps(datePicker *DatePicker) error {

	// TODO - Validate format is well-formed.

	datePicker.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type DatePickerVisibilityParams struct {
	DatePickerIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams DatePickerVisibilityParams) updateProps(datePicker *DatePicker) error {

	// TODO - Validate conditions

	datePicker.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type DatePickerPermissionParams struct {
	DatePickerIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams DatePickerPermissionParams) updateProps(datePicker *DatePicker) error {

	datePicker.Properties.Permissions = updateParams.Permissions

	return nil
}

type DatePickerValidationParams struct {
	DatePickerIDHeader
	Validation DatePickerValidationProperties `json:"validation"`
}

func (updateParams DatePickerValidationParams) updateProps(datePicker *DatePicker) error {

	datePicker.Properties.Validation = updateParams.Validation

	return nil
}

type DatePickerClearValueSupportedParams struct {
	DatePickerIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams DatePickerClearValueSupportedParams) updateProps(datePicker *DatePicker) error {

	datePicker.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	DatePickerIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(datePicker *DatePicker) error {

	datePicker.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}

type ConditionalFormatParams struct {
	DatePickerIDHeader
	ConditionalFormats []inputProps.DateConditionalFormat `json:"conditionalFormats"`
}

func (updateParams ConditionalFormatParams) updateProps(datePicker *DatePicker) error {

	datePicker.Properties.ConditionalFormats = updateParams.ConditionalFormats

	return nil
}
