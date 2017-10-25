package header

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/trackerDatabase"
)

type HeaderProps struct {
	Geometry   componentLayout.LayoutGeometry `json:"geometry"`
	Title      string                         `json:"title"`
	Size       string                         `json:"size"`
	Underlined bool                           `json:"underlined"`
}

func (srcProps HeaderProps) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*HeaderProps, error) {

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
