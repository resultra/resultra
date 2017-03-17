package common

import (
	"resultra/datasheet/server/recordFilter"
)

const LabelFormatField string = "field"

type ComponentLabelFormatProperties struct {
	LabelType   string `json:"labelType"`
	CustomLabel string `json:"customLabel"`
}

func NewDefaultLabelFormatProperties() ComponentLabelFormatProperties {
	return ComponentLabelFormatProperties{
		LabelType:   LabelFormatField,
		CustomLabel: ""}
}

type ComponentVisibilityProperties struct {
	VisibilityConditions []recordFilter.RecordFilterRule `json:"visibilityConditions"`
}

func NewDefaultComponentVisibilityProperties() ComponentVisibilityProperties {

	props := ComponentVisibilityProperties{
		VisibilityConditions: []recordFilter.RecordFilterRule{}}

	return props
}
