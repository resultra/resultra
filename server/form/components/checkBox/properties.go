package checkBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/trackerDatabase"
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
	FieldID                string                                `json:"fieldID"`
	Geometry               componentLayout.LayoutGeometry        `json:"geometry"`
	ColorScheme            string                                `json:"colorScheme"`
	StrikethroughCompleted bool                                  `json:"strikethroughCompleted"`
	LabelFormat            common.ComponentLabelFormatProperties `json:"labelFormat"`
	common.ComponentVisibilityProperties
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          CheckBoxValidationProperties               `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps CheckBoxProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*CheckBoxProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultCheckBoxProperties() CheckBoxProperties {

	props := CheckBoxProperties{
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		ColorScheme:                   CheckboxColorSchemeDefault,
		StrikethroughCompleted:        false,
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		Validation:                    newDefaultValidationProperties(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  ""}
	return props
}
