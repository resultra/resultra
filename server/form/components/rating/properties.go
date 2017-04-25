package rating

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

const ratingIconStar string = "star"

type RatingValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() RatingValidationProperties {
	return RatingValidationProperties{
		ValueRequired: true}
}

type RatingProperties struct {
	FieldID     string                                `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry        `json:"geometry"`
	Tooltips    []string                              `json:"tooltips"`
	Icon        string                                `json:"icon"`
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
	common.ComponentVisibilityProperties
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation  RatingValidationProperties                 `json:"validation"`
}

func (srcProps RatingProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*RatingProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultRatingProperties() RatingProperties {
	props := RatingProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Tooltips:                      []string{},
		Icon:                          ratingIconStar,
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		Validation:                    newDefaultValidationProperties()}
	return props
}
