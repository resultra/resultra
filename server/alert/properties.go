// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package alert

import (
	"fmt"
	"resultra/tracker/server/recordFilter"
	"resultra/tracker/server/trackerDatabase"
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

func (srcProps AlertProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*AlertProperties, error) {

	destProps := AlertProperties{}

	destTriggerConditions, err := srcProps.TriggerConditions.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("AlertProperties.Clone: %v", err)
	}
	destProps.TriggerConditions = *destTriggerConditions

	// TODO - Remap the any field references within the caption message.

	destProps.CaptionMessage = cloneAlertCaptionMsg(cloneParams, srcProps.CaptionMessage)

	srcCondition := srcProps.Condition
	if srcCondition != nil {
		remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcCondition.FieldID)
		if err != nil {
			return nil, fmt.Errorf("AlertProperties.Clone: %v", err)
		}
		destCondition := *srcCondition
		destCondition.FieldID = remappedFieldID
		destProps.Condition = &destCondition
	} else {
		destProps.Condition = nil
	}

	destFormID, formIDErr := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FormID)
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
