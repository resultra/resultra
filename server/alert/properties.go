package alert

import (
	"resultra/datasheet/server/generic/uniqueID"
)

type AlertProperties struct {
	DummyProp bool `json:"dummyProp"`
}

func (srcProps AlertProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*AlertProperties, error) {

	destProps := AlertProperties{}

	return &destProps, nil
}

func newDefaultAlertProperties() AlertProperties {
	defaultProps := AlertProperties{}

	return defaultProps
}
