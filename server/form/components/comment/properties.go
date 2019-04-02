// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package comment

import (
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type CommentProperties struct {
	FieldID     string                                `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry        `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
	common.ComponentVisibilityProperties
	Permissions  common.ComponentValuePermissionsProperties `json:"permissions"`
	HelpPopupMsg string                                     `json:"helpPopupMsg"`
}

func (srcProps CommentProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*CommentProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("CommentProperties.Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultCommentProperties() CommentProperties {
	props := CommentProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		HelpPopupMsg:                  ""}
	return props
}
