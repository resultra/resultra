package header

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type HeaderProps struct {
	Geometry   componentLayout.LayoutGeometry `json:"geometry"`
	Title      string                         `json:"title"`
	Size       string                         `json:"size"`
	Underlined bool                           `json:"underlined"`
}

func (srcProps HeaderProps) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*HeaderProps, error) {

	destProps := srcProps
	return &destProps, nil

}

func newDefaultHeaderProps() HeaderProps {
	props := HeaderProps{
		Title:      "Header",
		Size:       "medium",
		Underlined: false}
	return props
}
