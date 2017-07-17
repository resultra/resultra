package toggle

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

const ToggleColorSchemeDefault string = "default"

type ToggleValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() ToggleValidationProperties {
	return ToggleValidationProperties{
		ValueRequired: true}
}

type ToggleProperties struct {
	FieldID             string                                     `json:"fieldID"`
	OffColorScheme      string                                     `json:"offColorScheme"`
	OnColorScheme       string                                     `json:"onColorScheme"`
	OffLabel            string                                     `json:"offLabel"`
	OnLabel             string                                     `json:"onLabel"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          ToggleValidationProperties                 `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
}

func (srcProps ToggleProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ToggleProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultToggleProperties() ToggleProperties {

	props := ToggleProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		OffColorScheme:      ToggleColorSchemeDefault,
		OnColorScheme:       ToggleColorSchemeDefault,
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultValidationProperties(),
		OffLabel:            "No",
		OnLabel:             "Yes",
		ClearValueSupported: false}
	return props
}
