package numberInput

import (
	"fmt"
	"resultra/datasheet/server/common/inputProps"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/numberFormat"
	"resultra/datasheet/server/generic/uniqueID"
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
	ValueFormat          numberFormat.NumberFormatProperties        `json:"valueFormat"`
	LabelFormat          common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions          common.ComponentValuePermissionsProperties `json:"permissions"`
	ShowValueSpinner     bool                                       `json:"showValueSpinner"`
	ValueSpinnerStepSize float64                                    `json:"valueSpinnerStepSize"`
	Validation           NumberInputValidationProperties            `json:"validation"`
	ClearValueSupported  bool                                       `json:"clearValueSupported"`
	HelpPopupMsg         string                                     `json:"helpPopupMsg"`
	ConditionalFormats   []inputProps.NumberConditionalFormat       `json:"conditionalFormats"`
}

func (srcProps NumberInputProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*NumberInputProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultNumberInputProperties() NumberInputProperties {
	props := NumberInputProperties{
		LabelFormat:          common.NewDefaultLabelFormatProperties(),
		ValueFormat:          numberFormat.DefaultNumberFormatProperties(),
		Permissions:          common.NewDefaultComponentValuePermissionsProperties(),
		ShowValueSpinner:     true,
		ValueSpinnerStepSize: 1.0,
		Validation:           newDefaultValidationProperties(),
		ClearValueSupported:  false,
		HelpPopupMsg:         "",
		ConditionalFormats:   []inputProps.NumberConditionalFormat{}}
	return props
}
