package formButton

import (
	"fmt"
	"log"
	"resultra/datasheet/server/record"
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
	updateProps(button *FormButton) error
}

func updateButtonProps(propUpdater ButtonPropUpdater) (*FormButton, error) {

	// Retrieve the bar chart from the data store
	buttonForUpdate, getErr := getButton(propUpdater.getParentTableID(), propUpdater.getButtonID())
	if getErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to get existing button: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(buttonForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to update existing button properties: %v", propUpdateErr)
	}

	updatedButton, updateErr := updateExistingButton(buttonForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateButtonProps: Unable to update existing button properties: datastore update error =  %v", updateErr)
	}

	return updatedButton, nil
}

type ButtonBehaviorParams struct {
	ButtonIDHeader
	PopupBehavior ButtonPopupBehavior `json:"popupBehavior"`
}

func (updateParams ButtonBehaviorParams) updateProps(buttonForUpdate *FormButton) error {

	if err := updateParams.PopupBehavior.validateWellFormed(); err != nil {
		return err
	}

	buttonForUpdate.Properties.PopupBehavior = updateParams.PopupBehavior

	return nil
}

type ButtonDefaultValParams struct {
	ButtonIDHeader
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func (updateParams ButtonDefaultValParams) updateProps(buttonForUpdate *FormButton) error {

	if validateErr := record.ValidateWellFormedDefaultValues(updateParams.DefaultValues); validateErr != nil {
		return fmt.Errorf("updateProps: invalid default value(s): %v", validateErr)
	}

	log.Printf("Setting default values: %+v", updateParams.DefaultValues)

	buttonForUpdate.Properties.PopupBehavior.DefaultValues = updateParams.DefaultValues

	return nil
}

type ButtonSizeParams struct {
	ButtonIDHeader
	Size string `json:"size"`
}

func (updateParams ButtonSizeParams) updateProps(buttonForUpdate *FormButton) error {

	// TODO - Validate valid size

	buttonForUpdate.Properties.Size = updateParams.Size

	return nil
}

type ButtonColorSchemeParams struct {
	ButtonIDHeader
	ColorScheme string `json:"colorScheme"`
}

func (updateParams ButtonColorSchemeParams) updateProps(buttonForUpdate *FormButton) error {

	// TODO - Validate scheme name

	buttonForUpdate.Properties.ColorScheme = updateParams.ColorScheme

	return nil
}

type ButtonIconParams struct {
	ButtonIDHeader
	Icon string `json:"icon"`
}

func (updateParams ButtonIconParams) updateProps(buttonForUpdate *FormButton) error {

	// TODO - Validate icon name

	buttonForUpdate.Properties.Icon = updateParams.Icon

	return nil
}
