package textInput

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type TextInputIDInterface interface {
	getTextInputID() string
	getParentTableID() string
}

type TextInputIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	TextInputID   string `json:"textInputID"`
}

func (idHeader TextInputIDHeader) getTextInputID() string {
	return idHeader.TextInputID
}

func (idHeader TextInputIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type TextInputPropUpdater interface {
	TextInputIDInterface
	updateProps(textInput *TextInput) error
}

func updateTextInputProps(propUpdater TextInputPropUpdater) (*TextInput, error) {

	// Retrieve the bar chart from the data store
	textInputForUpdate, getErr := getTextInput(propUpdater.getParentTableID(), propUpdater.getTextInputID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateTextInputProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(textInputForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateTextInputProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	textInput, updateErr := updateExistingTextInput(propUpdater.getTextInputID(), textInputForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateTextInputProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return textInput, nil
}

type TextInputLabelFormatParams struct {
	TextInputIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams TextInputLabelFormatParams) updateProps(textInput *TextInput) error {

	// TODO - Validate format is well-formed.

	textInput.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type TextInputPermissionParams struct {
	TextInputIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams TextInputPermissionParams) updateProps(textInput *TextInput) error {

	textInput.Properties.Permissions = updateParams.Permissions

	return nil
}

type TextInputValueListParams struct {
	TextInputIDHeader
	ValueListID *string `json:"valueListID"`
}

func (updateParams TextInputValueListParams) updateProps(textInput *TextInput) error {

	if updateParams.ValueListID != nil {
		textInput.Properties.ValueListID = updateParams.ValueListID
	} else {
		textInput.Properties.ValueListID = nil

	}

	return nil
}

type TextInputValidationParams struct {
	TextInputIDHeader
	Validation TextInputValidationProperties `json:"validation"`
}

func (updateParams TextInputValidationParams) updateProps(textInput *TextInput) error {

	textInput.Properties.Validation = updateParams.Validation

	return nil
}
