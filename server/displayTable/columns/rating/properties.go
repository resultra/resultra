package rating

import (
	"fmt"
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
	FieldID             string                                     `json:"fieldID"`
	Tooltips            []string                                   `json:"tooltips"`
	Icon                string                                     `json:"icon"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          RatingValidationProperties                 `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
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
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Tooltips:            []string{},
		Icon:                ratingIconStar,
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
