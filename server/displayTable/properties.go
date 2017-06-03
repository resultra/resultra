package displayTable

import "resultra/datasheet/server/generic/uniqueID"

type DisplayTableProperties struct {
	OrderedColumns []string           `json:"orderedColumns"`
	ColWidths      map[string]float64 `json:"colWidths"`
}

func newDefaultDisplayTableProperties() DisplayTableProperties {
	props := DisplayTableProperties{
		OrderedColumns: []string{},
		ColWidths:      map[string]float64{}}
	return props
}

func (srcProps DisplayTableProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DisplayTableProperties, error) {

	destProps := srcProps

	destProps.OrderedColumns = uniqueID.CloneIDList(remappedIDs, srcProps.OrderedColumns)

	destColWidths := map[string]float64{}
	for srcColID, colWidth := range srcProps.ColWidths {
		destColID := remappedIDs.AllocNewOrGetExistingRemappedID(srcColID)
		destColWidths[destColID] = colWidth
	}
	destProps.ColWidths = destColWidths

	return &destProps, nil
}
