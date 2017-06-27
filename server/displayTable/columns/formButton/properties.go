package formButton

import (
	"fmt"
	"resultra/datasheet/server/common/inputProps"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/record"
)

const popupBehaviorModeless string = "modeless"
const popupBehaviorModal string = "modal"
const buttonSizeMedium string = "medium"
const colorSchemeDefault string = "default"
const buttonIconNone string = "none"

type ButtonProperties struct {
	LinkedFormID  string                         `json:"linkedFormID"`
	PopupBehavior inputProps.ButtonPopupBehavior `json:"popupBehavior"`
	Size          string                         `json:"size"`
	ColorScheme   string                         `json:"colorScheme"`
	Icon          string                         `json:"icon"`
	DefaultValues []record.DefaultFieldValue     `json:"defaultValues"`
}

func (srcProps ButtonProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ButtonProperties, error) {

	destProps := srcProps

	destProps.LinkedFormID = remappedIDs.AllocNewOrGetExistingRemappedID(srcProps.LinkedFormID)

	destPopupProps, cloneErr := srcProps.PopupBehavior.Clone(remappedIDs)
	if cloneErr != nil {
		return nil, fmt.Errorf("ButtonProperties.Clone: %v", cloneErr)
	}
	destProps.PopupBehavior = *destPopupProps

	destDefaultVals, cloneErr := record.CloneDefaultFieldValues(remappedIDs, srcProps.DefaultValues)
	if cloneErr != nil {
		return nil, fmt.Errorf("ButtonPopupBehavior.Clone: %v", cloneErr)
	}
	destProps.DefaultValues = destDefaultVals

	return &destProps, nil
}

func newDefaultButtonProperties() ButtonProperties {

	return ButtonProperties{
		PopupBehavior: inputProps.NewDefaultPopupBehavior(),
		Size:          buttonSizeMedium,
		ColorScheme:   colorSchemeDefault,
		Icon:          buttonIconNone,
		DefaultValues: []record.DefaultFieldValue{}}
}
