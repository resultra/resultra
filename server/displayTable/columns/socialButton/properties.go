package socialButton

import (
	"fmt"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
)

const socialButtonIconStar string = "star"

type SocialButtonProperties struct {
	FieldID      string                                     `json:"fieldID"`
	Icon         string                                     `json:"icon"`
	LabelFormat  common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions  common.ComponentValuePermissionsProperties `json:"permissions"`
	HelpPopupMsg string                                     `json:"helpPopupMsg"`
}

func (srcProps SocialButtonProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*SocialButtonProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
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
