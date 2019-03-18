package common

import (
	"resultra/tracker/server/recordFilter"
)

type ComponentVisibilityProperties struct {
	VisibilityConditions recordFilter.RecordFilterRuleSet `json:"visibilityConditions"`
}

func NewDefaultComponentVisibilityProperties() ComponentVisibilityProperties {

	props := ComponentVisibilityProperties{
		VisibilityConditions: recordFilter.NewDefaultRecordFilterRuleSet()}

	return props
}
