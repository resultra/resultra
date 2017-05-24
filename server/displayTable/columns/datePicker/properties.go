package datePicker

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
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
	DateFormat  string                                     `json:"dateFormat"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation  DatePickerValidationProperties             `json:"validation"`
}

const dateFormatDefault string = "date"

func (srcProps DatePickerProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DatePickerProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultDatePickerProperties() DatePickerProperties {
	props := DatePickerProperties{
		LabelFormat: common.NewDefaultLabelFormatProperties(),
		Permissions: common.NewDefaultComponentValuePermissionsProperties(),
		DateFormat:  dateFormatDefault}
	return props
}
