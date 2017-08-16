package alert

import (
	"fmt"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

type AlertCondition struct {
	FieldID     string     `json:"fieldID"`
	ConditionID string     `json:"conditionID"`
	DateParam   *time.Time `json:"dateParam,omitempty"`
	NumberParam *float64   `json:"numberParam,omitempty"`
}

type AlertProperties struct {
	FormID         string           `json:"formID"`
	SummaryFieldID string           `json:"summaryFieldID"`
	Conditions     []AlertCondition `json:"conditions"`
}

func (srcProps AlertProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*AlertProperties, error) {

	destProps := AlertProperties{}

	destConditions := []AlertCondition{}
	for _, srcCondition := range srcProps.Conditions {
		destCondition := srcCondition

		remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcCondition.FieldID)
		if err != nil {
			return nil, fmt.Errorf("AlertProperties.Clone: %v", err)
		}
		destCondition.FieldID = remappedFieldID

		destConditions = append(destConditions, destCondition)
	}
	destProps.Conditions = destConditions

	destFormID, formIDErr := remappedIDs.GetExistingRemappedID(srcProps.FormID)
	if formIDErr != nil {
		return nil, fmt.Errorf("AlertProperties.Clone: %v", formIDErr)
	}
	destProps.FormID = destFormID

	destFieldID, fieldIDErr := remappedIDs.GetExistingRemappedID(srcProps.SummaryFieldID)
	if fieldIDErr != nil {
		return nil, fmt.Errorf("AlertProperties.Clone: %v", fieldIDErr)
	}
	destProps.SummaryFieldID = destFieldID

	return &destProps, nil
}

func newDefaultAlertProperties() AlertProperties {
	defaultProps := AlertProperties{
		Conditions: []AlertCondition{}}

	return defaultProps
}
