// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package socialButton

import (
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/trackerDatabase"
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

func (srcProps SocialButtonProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*SocialButtonProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(cloneParams)
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
