package common

import (
	"fmt"
	"resultra/datasheet/server/common/datastoreWrapper"
)

// A LayoutGeometry stores the size and position information for dashboard and form objects.
// It is intended to be a member of other structs which are specific to the given
// dashboard or form object.
type LayoutGeometry struct {
	PositionTop  int `json:"positionTop"`
	PositionLeft int `json:"positionLeft"`
	SizeWidth    int `json:"sizeWidth"`
	SizeHeight   int `json:"sizeHeight"`
}

func NewUnitializedLayoutGeometry() LayoutGeometry {
	return LayoutGeometry{-1, -1, -1, -1}
}

func ValidGeometry(geom LayoutGeometry) bool {
	if (geom.PositionTop >= 0) && (geom.PositionLeft >= 0) &&
		(geom.SizeWidth > 0) && (geom.SizeHeight > 0) {
		return true
	} else {
		return false
	}
}

func (geom *LayoutGeometry) SetPosition(pos LayoutPosition) error {
	if !ValidPosition(pos) {
		return fmt.Errorf("Error setting position for object's geomery: invalid position = %+v", pos)
	}
	geom.PositionTop = pos.Top
	geom.PositionLeft = pos.Left
	return nil
}

type ObjectDimensionsParams struct {
	datastoreWrapper.UniqueIDHeader
	Geometry LayoutGeometry `json:"geometry"`
}

type LayoutPosition struct {
	Top  int `json:"top"`
	Left int `json:"left"`
}

type ObjectRepositionParams struct {
	datastoreWrapper.UniqueIDHeader
	Position LayoutPosition `json:"position"`
}

func ValidPosition(pos LayoutPosition) bool {
	if (pos.Top >= 0) && (pos.Left >= 0) {
		return true
	} else {
		return false
	}
}
