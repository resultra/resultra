package valueList

import (
	"resultra/datasheet/server/trackerDatabase"
)

type ValueListValue struct {
	NumValue  *float64 `json:"numValue,omitempty"`
	TextValue *string  `json:"textValue,omitempty"`
}

type ValueListProperties struct {
	Values []ValueListValue `json:"values"`
}

func newDefaultValueListProperties() ValueListProperties {
	defaultProps := ValueListProperties{
		Values: []ValueListValue{}}
	return defaultProps
}

func (srcProps ValueListProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*ValueListProperties, error) {

	destProps := srcProps

	return &destProps, nil
}
