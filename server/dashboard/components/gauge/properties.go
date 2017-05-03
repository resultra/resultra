package gauge

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type GaugeProps struct {
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
	Title    string                         `json:"title"`
}

func (srcProps GaugeProps) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*GaugeProps, error) {

	destProps := srcProps
	return &destProps, nil

}

func newDefaultGaugeProps() GaugeProps {
	props := GaugeProps{
		Title: "Gauge"}
	return props
}
