package formButton

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/record"
)

const popupBehaviorModeless string = "modeless"
const popupBehaviorModal string = "modal"
const buttonSizeMedium string = "medium"
const colorSchemeDefault string = "default"

type ButtonPopupBehavior struct {
	PopupMode            string                     `json:"popupMode"`
	CustomLabelModalSave string                     `json:"customLabelModalSave"`
	DefaultValues        []record.DefaultFieldValue `json:"defaultValues"`
}

func (srcProps ButtonPopupBehavior) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ButtonPopupBehavior, error) {

	destProps := srcProps

	destDefaultVals, cloneErr := record.CloneDefaultFieldValues(remappedIDs, srcProps.DefaultValues)
	if cloneErr != nil {
		return nil, fmt.Errorf("ButtonPopupBehavior.Clone: %v", cloneErr)
	}

	destProps.DefaultValues = destDefaultVals

	return &destProps, nil

}

func newDefaultPopupBehavior() ButtonPopupBehavior {
	defaultPopupBehavior := ButtonPopupBehavior{
		PopupMode:            popupBehaviorModeless,
		CustomLabelModalSave: "",
		DefaultValues:        []record.DefaultFieldValue{}}
	return defaultPopupBehavior
}

func (buttonPopupBehavior ButtonPopupBehavior) validateWellFormed() error {

	if !(buttonPopupBehavior.PopupMode == popupBehaviorModeless ||
		buttonPopupBehavior.PopupMode == popupBehaviorModal) {
		return fmt.Errorf("Invalid form popup mode: %v", buttonPopupBehavior.PopupMode)
	}

	if validLabelErr := stringValidation.ValidateOptionalItemLabel(buttonPopupBehavior.CustomLabelModalSave); validLabelErr != nil {
		return validLabelErr
	}

	return nil
}

type ButtonProperties struct {
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
	LinkedFormID  string                         `json:"linkedFormID"`
	PopupBehavior ButtonPopupBehavior            `json:"popupBehavior"`
	Size          string                         `json:"size"`
	ColorScheme   string                         `json:"colorScheme"`
}

func (srcProps ButtonProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ButtonProperties, error) {

	destProps := srcProps

	destProps.LinkedFormID = remappedIDs.AllocNewOrGetExistingRemappedID(srcProps.LinkedFormID)

	destPopupProps, cloneErr := srcProps.PopupBehavior.Clone(remappedIDs)
	if cloneErr != nil {
		return nil, fmt.Errorf("ButtonProperties.Clone: %v", cloneErr)
	}
	destProps.PopupBehavior = *destPopupProps

	return &destProps, nil
}

func newDefaultButtonProperties() ButtonProperties {

	return ButtonProperties{
		PopupBehavior: newDefaultPopupBehavior(),
		Size:          buttonSizeMedium,
		ColorScheme:   colorSchemeDefault}
}
