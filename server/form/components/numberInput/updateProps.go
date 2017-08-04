package numberInput

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/inputProps"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/numberFormat"
)

type NumberInputIDInterface interface {
	getNumberInputID() string
	getParentFormID() string
}

type NumberInputIDHeader struct {
	ParentFormID  string `json:"parentFormID"`
	NumberInputID string `json:"numberInputID"`
}

func (idHeader NumberInputIDHeader) getNumberInputID() string {
	return idHeader.NumberInputID
}

func (idHeader NumberInputIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type NumberInputPropUpdater interface {
	NumberInputIDInterface
	updateProps(numberInput *NumberInput) error
}

func updateNumberInputProps(propUpdater NumberInputPropUpdater) (*NumberInput, error) {

	// Retrieve the bar chart from the data store
	numberInputForUpdate, getErr := getNumberInput(propUpdater.getParentFormID(), propUpdater.getNumberInputID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateNumberInputProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(numberInputForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateNumberInputProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	numberInput, updateErr := updateExistingNumberInput(propUpdater.getNumberInputID(), numberInputForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateNumberInputProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return numberInput, nil
}

type NumberInputResizeParams struct {
	NumberInputIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams NumberInputResizeParams) updateProps(numberInput *NumberInput) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set text box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	numberInput.Properties.Geometry = updateParams.Geometry

	return nil
}

type NumberInputValueFormatParams struct {
	NumberInputIDHeader
	ValueFormat numberFormat.NumberFormatProperties `json:"valueFormat"`
}

func (updateParams NumberInputValueFormatParams) updateProps(numberInput *NumberInput) error {

	numberInput.Properties.ValueFormat = updateParams.ValueFormat

	return nil
}

type NumberInputLabelFormatParams struct {
	NumberInputIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams NumberInputLabelFormatParams) updateProps(numberInput *NumberInput) error {

	// TODO - Validate format is well-formed.

	numberInput.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type NumberInputVisibilityParams struct {
	NumberInputIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams NumberInputVisibilityParams) updateProps(numberInput *NumberInput) error {

	// TODO - Validate conditions

	numberInput.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type NumberInputPermissionParams struct {
	NumberInputIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams NumberInputPermissionParams) updateProps(numberInput *NumberInput) error {

	numberInput.Properties.Permissions = updateParams.Permissions

	return nil
}

type ShowValueSpinnerParams struct {
	NumberInputIDHeader
	ShowValueSpinner bool `json:"showValueSpinner"`
}

func (updateParams ShowValueSpinnerParams) updateProps(numberInput *NumberInput) error {

	numberInput.Properties.ShowValueSpinner = updateParams.ShowValueSpinner

	return nil
}

type ValueSpinnerStepSizeParams struct {
	NumberInputIDHeader
	ValueSpinnerStepSize float64 `json:"valueSpinnerStepSize"`
}

func (updateParams ValueSpinnerStepSizeParams) updateProps(numberInput *NumberInput) error {

	if updateParams.ValueSpinnerStepSize <= 0.0 {
		return fmt.Errorf("Invalid spinner step size %v, must be > 0.0", updateParams.ValueSpinnerStepSize)
	}

	numberInput.Properties.ValueSpinnerStepSize = updateParams.ValueSpinnerStepSize

	return nil
}

type NumberInputValidationParams struct {
	NumberInputIDHeader
	Validation NumberInputValidationProperties `json:"validation"`
}

func (updateParams NumberInputValidationParams) updateProps(numberInput *NumberInput) error {

	numberInput.Properties.Validation = updateParams.Validation

	return nil
}

type NumberInputClearValueSupportedParams struct {
	NumberInputIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams NumberInputClearValueSupportedParams) updateProps(numberInput *NumberInput) error {

	numberInput.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	NumberInputIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(numberInput *NumberInput) error {

	numberInput.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}

type ConditionalFormatParams struct {
	NumberInputIDHeader
	ConditionalFormats []inputProps.NumberConditionalFormat `json:"conditionalFormats"`
}

func (updateParams ConditionalFormatParams) updateProps(numberInput *NumberInput) error {

	numberInput.Properties.ConditionalFormats = updateParams.ConditionalFormats

	return nil
}
