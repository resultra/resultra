package textInput

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type TextInputValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultTextInputValidationProperties() TextInputValidationProperties {
	return TextInputValidationProperties{false}
}

type TextInputProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	ValueListID         *string                                    `json:"valueListID,omitempty"`
	Validation          TextInputValidationProperties              `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps TextInputProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*TextInputProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	if srcProps.ValueListID != nil {
		destValueListID := remappedIDs.AllocNewOrGetExistingRemappedID(*srcProps.ValueListID)
		destProps.ValueListID = &destValueListID
	}

	return &destProps, nil
}

func newDefaultTextInputProperties() TextInputProperties {
	props := TextInputProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultTextInputValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
