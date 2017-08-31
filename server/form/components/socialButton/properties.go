package socialButton

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

const socialButtonIconStar string = "star"

type SocialButtonProperties struct {
	FieldID     string                                `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry        `json:"geometry"`
	Tooltips    []string                              `json:"tooltips"`
	Icon        string                                `json:"icon"`
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
	common.ComponentVisibilityProperties
	Permissions  common.ComponentValuePermissionsProperties `json:"permissions"`
	HelpPopupMsg string                                     `json:"helpPopupMsg"`
}

func (srcProps SocialButtonProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*SocialButtonProperties, error) {

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

func newDefaultSocialButtonProperties() SocialButtonProperties {
	props := SocialButtonProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Tooltips:                      []string{},
		Icon:                          socialButtonIconStar,
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		HelpPopupMsg:                  ""}
	return props
}
