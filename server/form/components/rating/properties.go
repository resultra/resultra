package rating

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

const ratingIconStar string = "star"

type RatingProperties struct {
	FieldID     string                                `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry        `json:"geometry"`
	Tooltips    []string                              `json:"tooltips"`
	Icon        string                                `json:"icon"`
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (srcProps RatingProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*RatingProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultRatingProperties() RatingProperties {
	props := RatingProperties{
		LabelFormat: common.NewDefaultLabelFormatProperties(),
		Tooltips:    []string{},
		Icon:        ratingIconStar}
	return props
}
