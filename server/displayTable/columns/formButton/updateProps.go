// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formButton

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/resultra/resultra/server/common/inputProps"
	"github.com/resultra/resultra/server/record"
)

type ButtonIDInterface interface {
	getButtonID() string
	getParentTableID() string
}

type ButtonIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	ButtonID      string `json:"buttonID"`
}

func (idHeader ButtonIDHeader) getButtonID() string {
	return idHeader.ButtonID
}

func (idHeader ButtonIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type ButtonPropUpdater interface {
	ButtonIDInterface
	updateProps(trackerDBHandle *sql.DB, button *FormButton) error
}

func updateButtonProps(trackerDBHandle *sql.DB, propUpdater ButtonPropUpdater) (*FormButton, error) {

	// Retrieve the bar chart from the data store
	buttonForUpdate, getErr := getButton(trackerDBHandle, propUpdater.getParentTableID(),
		propUpdater.getButtonID())
	if getErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to get existing button: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(trackerDBHandle, buttonForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to update existing button properties: %v", propUpdateErr)
	}

	updatedButton, updateErr := updateExistingButton(trackerDBHandle, buttonForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to update existing button properties: datastore update error =  %v", updateErr)
	}

	return updatedButton, nil
}

type ButtonBehaviorParams struct {
	ButtonIDHeader
	PopupBehavior inputProps.ButtonPopupBehavior `json:"popupBehavior"`
}

func (updateParams ButtonBehaviorParams) updateProps(trackerDBHandle *sql.DB, buttonForUpdate *FormButton) error {

	if err := updateParams.PopupBehavior.ValidateWellFormed(); err != nil {
		return err
	}

	buttonForUpdate.Properties.PopupBehavior = updateParams.PopupBehavior

	return nil
}

type ButtonDefaultValParams struct {
	ButtonIDHeader
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func (updateParams ButtonDefaultValParams) updateProps(trackerDBHandle *sql.DB,
	buttonForUpdate *FormButton) error {

	if validateErr := record.ValidateWellFormedDefaultValues(trackerDBHandle, updateParams.DefaultValues); validateErr != nil {
		return fmt.Errorf("updateProps: invalid default value(s): %v", validateErr)
	}

	log.Printf("Setting default values: %+v", updateParams.DefaultValues)

	buttonForUpdate.Properties.DefaultValues = updateParams.DefaultValues

	return nil
}

type ButtonSizeParams struct {
	ButtonIDHeader
	Size string `json:"size"`
}

func (updateParams ButtonSizeParams) updateProps(trackerDBHandle *sql.DB, buttonForUpdate *FormButton) error {

	// TODO - Validate valid size

	buttonForUpdate.Properties.Size = updateParams.Size

	return nil
}

type ButtonColorSchemeParams struct {
	ButtonIDHeader
	ColorScheme string `json:"colorScheme"`
}

func (updateParams ButtonColorSchemeParams) updateProps(trackerDBHandle *sql.DB,
	buttonForUpdate *FormButton) error {

	// TODO - Validate scheme name

	buttonForUpdate.Properties.ColorScheme = updateParams.ColorScheme

	return nil
}

type ButtonIconParams struct {
	ButtonIDHeader
	Icon string `json:"icon"`
}

func (updateParams ButtonIconParams) updateProps(trackerDBHandle *sql.DB, buttonForUpdate *FormButton) error {

	// TODO - Validate icon name

	buttonForUpdate.Properties.Icon = updateParams.Icon

	return nil
}

type ButtonLabelFormatParams struct {
	ButtonIDHeader
	ButtonLabelFormat inputProps.FormButtonLabelFormatProperties `json:"buttonLabelFormat"`
}

func (updateParams ButtonLabelFormatParams) updateProps(trackerDBHandle *sql.DB, buttonForUpdate *FormButton) error {

	buttonForUpdate.Properties.ButtonLabelFormat = updateParams.ButtonLabelFormat

	return nil
}
