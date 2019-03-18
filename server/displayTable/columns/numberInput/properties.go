// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package numberInput

import (
	"fmt"
	"resultra/tracker/server/common/inputProps"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/generic/numberFormat"
	"resultra/tracker/server/trackerDatabase"
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

func (srcProps NumberInputProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*NumberInputProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
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
