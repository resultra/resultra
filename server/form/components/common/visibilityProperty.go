package common

import (
	"resultra/datasheet/server/recordFilter"
)

type ComponentVisibilityProperties struct {
	VisibilityConditions []recordFilter.RecordFilterRule `json:"visibilityConditions"`
}

func NewDefaultComponentVisibilityProperties() ComponentVisibilityProperties {

	props := ComponentVisibilityProperties{
		VisibilityConditions: []recordFilter.RecordFilterRule{}}

	return props
}
