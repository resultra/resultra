package common

import (
	"resultra/datasheet/server/recordFilter"
)

type ComponentVisibilityProperties struct {
	VisibilityConditions recordFilter.RecordFilterRuleSet `json:"visibilityConditions"`
}

func NewDefaultComponentVisibilityProperties() ComponentVisibilityProperties {

	props := ComponentVisibilityProperties{
		VisibilityConditions: recordFilter.NewDefaultRecordFilterRuleSet()}

	return props
}
