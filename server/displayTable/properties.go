package displayTable

import "resultra/datasheet/server/generic/uniqueID"

type DisplayTableProperties struct {
}

func newDefaultDisplayTableProperties() DisplayTableProperties {
	props := DisplayTableProperties{}
	return props
}

func (srcProps DisplayTableProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DisplayTableProperties, error) {

	destProps := srcProps

	return &destProps, nil
}
