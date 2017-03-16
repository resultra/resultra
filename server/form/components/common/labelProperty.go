package common

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
