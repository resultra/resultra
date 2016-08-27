package componentLayout

// A LayoutGeometry stores the size and position information for dashboard and form objects.
// It is intended to be a member of other structs which are specific to the given
// dashboard or form object.
type LayoutGeometry struct {
	SizeWidth  int `json:"sizeWidth"`
	SizeHeight int `json:"sizeHeight"`
}

func NewUnitializedLayoutGeometry() LayoutGeometry {
	return LayoutGeometry{-1, -1}
}

func ValidGeometry(geom LayoutGeometry) bool {
	if (geom.SizeWidth > 0) && (geom.SizeHeight > 0) {
		return true
	} else {
		return false
	}
}
