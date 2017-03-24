package checkBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

const CheckboxColorSchemeDefault string = "default"

type CheckBoxProperties struct {
	FieldID                string                                `json:"fieldID"`
	Geometry               componentLayout.LayoutGeometry        `json:"geometry"`
	ColorScheme            string                                `json:"colorScheme"`
	StrikethroughCompleted bool                                  `json:"strikethroughCompleted"`
	LabelFormat            common.ComponentLabelFormatProperties `json:"labelFormat"`
	common.ComponentVisibilityProperties
	ReadOnly bool `json:"readOnly"`
}

func (srcProps CheckBoxProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*CheckBoxProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.VisibilityConditions)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = destVisibilityConditions

	return &destProps, nil
}

func newDefaultCheckBoxProperties() CheckBoxProperties {

	props := CheckBoxProperties{
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		ColorScheme:                   CheckboxColorSchemeDefault,
		StrikethroughCompleted:        false,
		ReadOnly:                      false}
	return props
}
