// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package checkBox

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
)

type CheckboxIDInterface interface {
	getCheckBoxID() string
	getParentFormID() string
}

type CheckboxIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	CheckBoxID   string `json:"checkBoxID"`
}

func (idHeader CheckboxIDHeader) getCheckBoxID() string {
	return idHeader.CheckBoxID
}

func (idHeader CheckboxIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type CheckBoxPropUpdater interface {
	CheckboxIDInterface
	updateProps(checkBox *CheckBox) error
}

func updateCheckBoxProps(trackerDBHandle *sql.DB, propUpdater CheckBoxPropUpdater) (*CheckBox, error) {

	// Retrieve the bar chart from the data store
	checkBoxForUpdate, getErr := getCheckBox(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getCheckBoxID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to get existing check box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(checkBoxForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: %v", propUpdateErr)
	}

	checkBox, updateErr := updateExistingCheckBox(trackerDBHandle, checkBoxForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: datastore update error =  %v", updateErr)
	}

	return checkBox, nil
}

type CheckBoxResizeParams struct {
	CheckboxIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams CheckBoxResizeParams) updateProps(checkBox *CheckBox) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	checkBox.Properties.Geometry = updateParams.Geometry

	return nil
}

type CheckBoxColorSchemeParams struct {
	CheckboxIDHeader
	ColorScheme string `json:"colorScheme"`
}

func (updateParams CheckBoxColorSchemeParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate against list of valid color schemes

	checkBox.Properties.ColorScheme = updateParams.ColorScheme

	return nil
}

type CheckBoxStrikethroughParams struct {
	CheckboxIDHeader
	StrikethroughCompleted bool `json:"strikethroughCompleted"`
}

func (updateParams CheckBoxStrikethroughParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate against list of valid color schemes

	checkBox.Properties.StrikethroughCompleted = updateParams.StrikethroughCompleted

	return nil
}

type CheckBoxLabelFormatParams struct {
	CheckboxIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams CheckBoxLabelFormatParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate format is well-formed.

	checkBox.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type CheckBoxVisibilityParams struct {
	CheckboxIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams CheckBoxVisibilityParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate conditions

	checkBox.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type CheckBoxPermissionParams struct {
	CheckboxIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams CheckBoxPermissionParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate conditions

	checkBox.Properties.Permissions = updateParams.Permissions

	return nil
}

type CheckBoxValidationParams struct {
	CheckboxIDHeader
	Validation CheckBoxValidationProperties `json:"validation"`
}

func (updateParams CheckBoxValidationParams) updateProps(checkBox *CheckBox) error {

	checkBox.Properties.Validation = updateParams.Validation

	return nil
}

type CheckBoxClearValueSupportedParams struct {
	CheckboxIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams CheckBoxClearValueSupportedParams) updateProps(checkBox *CheckBox) error {

	checkBox.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	CheckboxIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(checkBox *CheckBox) error {

	checkBox.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
