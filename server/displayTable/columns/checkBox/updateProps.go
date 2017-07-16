package checkBox

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type CheckboxIDInterface interface {
	getCheckBoxID() string
	getParentTableID() string
}

type CheckboxIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	CheckBoxID    string `json:"checkBoxID"`
}

func (idHeader CheckboxIDHeader) getCheckBoxID() string {
	return idHeader.CheckBoxID
}

func (idHeader CheckboxIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type CheckBoxPropUpdater interface {
	CheckboxIDInterface
	updateProps(checkBox *CheckBox) error
}

func updateCheckBoxProps(propUpdater CheckBoxPropUpdater) (*CheckBox, error) {

	// Retrieve the bar chart from the data store
	checkBoxForUpdate, getErr := getCheckBox(propUpdater.getParentTableID(), propUpdater.getCheckBoxID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to get existing check box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(checkBoxForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: %v", propUpdateErr)
	}

	checkBox, updateErr := updateExistingCheckBox(checkBoxForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: datastore update error =  %v", updateErr)
	}

	return checkBox, nil
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
