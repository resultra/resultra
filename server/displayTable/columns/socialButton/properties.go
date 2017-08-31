package socialButton

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

const socialButtonIconStar string = "star"

type SocialButtonProperties struct {
	FieldID      string                                     `json:"fieldID"`
	Icon         string                                     `json:"icon"`
	LabelFormat  common.ComponentLabelFormatProperties      `json:"labelFormat"`
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

	return &destProps, nil
}

func newDefaultSocialButtonProperties() SocialButtonProperties {
	props := SocialButtonProperties{
		LabelFormat:  common.NewDefaultLabelFormatProperties(),
		Icon:         socialButtonIconStar,
		Permissions:  common.NewDefaultComponentValuePermissionsProperties(),
		HelpPopupMsg: ""}
	return props
}
