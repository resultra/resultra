package textInput

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/trackerDatabase"
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

func (srcProps TextInputProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*TextInputProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	if srcProps.ValueListID != nil {
		destValueListID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(*srcProps.ValueListID)
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
