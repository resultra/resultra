package alert

import (
	"fmt"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
	"time"
)

type AlertCondition struct {
	FieldID     string     `json:"fieldID"`
	ConditionID string     `json:"conditionID"`
	DateParam   *time.Time `json:"dateParam,omitempty"`
	NumberParam *float64   `json:"numberParam,omitempty"`
}

type AlertProperties struct {
	FormID            string                           `json:"formID"`
	Condition         *AlertCondition                  `json:"condition"`
	CaptionMessage    string                           `json:"captionMessage"`
	TriggerConditions recordFilter.RecordFilterRuleSet `json:"triggerConditions"`
}

func (srcProps AlertProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*AlertProperties, error) {

	destProps := AlertProperties{}

	destTriggerConditions, err := srcProps.TriggerConditions.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("AlertProperties.Clone: %v")
	}
	destProps.TriggerConditions = *destTriggerConditions

	// TODO - Remap the any field references within the caption message.

	srcCondition := srcProps.Condition
	if srcCondition != nil {
		remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcCondition.FieldID)
		if err != nil {
			return nil, fmt.Errorf("AlertProperties.Clone: %v", err)
		}
		destCondition := *srcCondition
		destCondition.FieldID = remappedFieldID
		destProps.Condition = &destCondition
	} else {
		destProps.Condition = nil
	}

	destFormID, formIDErr := remappedIDs.GetExistingRemappedID(srcProps.FormID)
	if formIDErr != nil {
		return nil, fmt.Errorf("AlertProperties.Clone: %v", formIDErr)
	}
	destProps.FormID = destFormID

	return &destProps, nil
}

func newDefaultAlertProperties() AlertProperties {
	defaultProps := AlertProperties{
		TriggerConditions: recordFilter.NewDefaultRecordFilterRuleSet()}

	return defaultProps
}
