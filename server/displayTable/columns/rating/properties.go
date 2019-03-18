// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package rating

import (
	"fmt"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
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
	MinVal              int                                        `json:"minVal"`
	MaxVal              int                                        `json:"maxVal"`
}

func (srcProps RatingProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*RatingProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
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
		HelpPopupMsg:        "",
		MinVal:              0,
		MaxVal:              5}
	return props
}
