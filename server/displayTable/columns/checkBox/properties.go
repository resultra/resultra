package checkBox

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

const CheckboxColorSchemeDefault string = "default"

type CheckBoxValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() CheckBoxValidationProperties {
	return CheckBoxValidationProperties{
		ValueRequired: true}
}

type CheckBoxProperties struct {
	FieldID                string                                     `json:"fieldID"`
	ColorScheme            string                                     `json:"colorScheme"`
	StrikethroughCompleted bool                                       `json:"strikethroughCompleted"`
	LabelFormat            common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions            common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation             CheckBoxValidationProperties               `json:"validation"`
	ClearValueSupported    bool                                       `json:"clearValueSupported"`
}

func (srcProps CheckBoxProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*CheckBoxProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultCheckBoxProperties() CheckBoxProperties {

	props := CheckBoxProperties{
		LabelFormat:            common.NewDefaultLabelFormatProperties(),
		ColorScheme:            CheckboxColorSchemeDefault,
		StrikethroughCompleted: false,
		Permissions:            common.NewDefaultComponentValuePermissionsProperties(),
		Validation:             newDefaultValidationProperties(),
		ClearValueSupported:    false}
	return props
}
