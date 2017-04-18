package valueList

import (
	"resultra/datasheet/server/generic/uniqueID"
)

type ValueListValue struct {
	NumValue   *float64 `json:"numValue:omitempty"`
	TextValue  *string  `json:"textValue:omitempty"`
	ValueLabel string   `json:"valueLabel"`
}

type ValueListProperties struct {
	Values []ValueListValue `json:"values"`
}

func newDefaultValueListProperties() ValueListProperties {
	defaultProps := ValueListProperties{
		Values: []ValueListValue{}}
	return defaultProps
}

func (srcProps ValueListProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ValueListProperties, error) {

	destProps := srcProps

	return &destProps, nil
}
