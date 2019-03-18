// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package datePicker

import (
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/common/inputProps"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
	"time"
)

const validationRuleNone string = "none"
const validationRuleRequired string = "required"
const validationRuleFuture string = "future"
const validationRulePast string = "past"
const validationRuleBefore string = "before"
const validationRuleAfter string = "after"
const validationRuleBetween string = "between"

type DatePickerValidationProperties struct {
	Rule        string     `json:"rule"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	CompareDate *time.Time `json:"compareDate,omitempty"`
}

func newDefaultDatePickerProps() DatePickerValidationProperties {
	return DatePickerValidationProperties{Rule: validationRuleRequired}
}

type DatePickerProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	DateFormat  string                                     `json:"dateFormat"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
	Validation          DatePickerValidationProperties     `json:"validation"`
	ClearValueSupported bool                               `json:"clearValueSupported"`
	HelpPopupMsg        string                             `json:"helpPopupMsg"`
	ConditionalFormats  []inputProps.DateConditionalFormat `json:"conditionalFormats"`
}

const dateFormatDefault string = "date"

func (srcProps DatePickerProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*DatePickerProperties, error) {

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

func newDefaultDatePickerProperties() DatePickerProperties {
	props := DatePickerProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		DateFormat:                    dateFormatDefault,
		Validation:                    newDefaultDatePickerProps(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  "",
		ConditionalFormats:            []inputProps.DateConditionalFormat{}}
	return props
}
