// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package numberInput

import (
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/common/inputProps"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic/numberFormat"
	"github.com/resultra/resultra/server/trackerDatabase"
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
