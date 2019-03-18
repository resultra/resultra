package displayTable

import "resultra/tracker/server/generic/uniqueID"

type DisplayTableProperties struct {
	OrderedColumns []string `json:"orderedColumns"`
}

func newDefaultDisplayTableProperties() DisplayTableProperties {
	props := DisplayTableProperties{
		OrderedColumns: []string{}}
	return props
}

func (srcProps DisplayTableProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DisplayTableProperties, error) {

	destProps := srcProps

	destProps.OrderedColumns = uniqueID.CloneIDList(remappedIDs, srcProps.OrderedColumns)

	return &destProps, nil
}
