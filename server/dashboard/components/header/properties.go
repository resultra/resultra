package header

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type HeaderProps struct {
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
	Title    string                         `json:"title"`
}

func (srcProps HeaderProps) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*HeaderProps, error) {

	destProps := srcProps
	return &destProps, nil

}

func newDefaultHeaderProps() HeaderProps {
	props := HeaderProps{
		Title: "Header"}
	return props
}
