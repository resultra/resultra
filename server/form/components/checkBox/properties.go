package checkBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

const CheckboxColorSchemeDefault string = "default"

type CheckBoxProperties struct {
	FieldID                string                         `json:"fieldID"`
	Geometry               componentLayout.LayoutGeometry `json:"geometry"`
	ColorScheme            string                         `json:"colorScheme"`
	StrikethroughCompleted bool                           `json:"strikethroughCompleted"`
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
		ColorScheme:            CheckboxColorSchemeDefault,
		StrikethroughCompleted: false}
	return props
}
