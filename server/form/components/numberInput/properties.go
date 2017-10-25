package numberInput

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/inputProps"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/numberFormat"
	"resultra/datasheet/server/trackerDatabase"
)

type NumberInputValidationProperties struct {
	Rule       string   `json:"rule"`
	MinVal     *float64 `json:"minVal,omitempty"`
	MaxVal     *float64 `json:"maxVal,omitempty"`
	CompareVal *float64 `json:"compareVal,omitempty"`
}

func newDefaultValidationProperties() NumberInputValidationProperties {
	return NumberInputValidationProperties{Rule: "required"}
}

type NumberInputProperties struct {
	FieldID              string                                     `json:"fieldID"`
	Geometry             componentLayout.LayoutGeometry             `json:"geometry"`
	ValueFormat          numberFormat.NumberFormatProperties        `json:"valueFormat"`
	LabelFormat          common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions          common.ComponentValuePermissionsProperties `json:"permissions"`
	ShowValueSpinner     bool                                       `json:"showValueSpinner"`
	ValueSpinnerStepSize float64                                    `json:"valueSpinnerStepSize"`
	common.ComponentVisibilityProperties
	Validation          NumberInputValidationProperties      `json:"validation"`
	ClearValueSupported bool                                 `json:"clearValueSupported"`
	HelpPopupMsg        string                               `json:"helpPopupMsg"`
	ConditionalFormats  []inputProps.NumberConditionalFormat `json:"conditionalFormats"`
}

func (srcProps NumberInputProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*NumberInputProperties, error) {

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

func newDefaultNumberInputProperties() NumberInputProperties {
	props := NumberInputProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		ValueFormat:                   numberFormat.DefaultNumberFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		ShowValueSpinner:              true,
		ValueSpinnerStepSize:          1.0,
		Validation:                    newDefaultValidationProperties(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  "",
		ConditionalFormats:            []inputProps.NumberConditionalFormat{}}
	return props
}
