// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package textBox

import (
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type TextBoxValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultTextBoxValidationProperties() TextBoxValidationProperties {
	return TextBoxValidationProperties{false}
}

type TextBoxProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	ValueListID *string                                    `json:"valueListID,omitempty"`
	common.ComponentVisibilityProperties
	Validation          TextBoxValidationProperties `json:"validation"`
	ClearValueSupported bool                        `json:"clearValueSupported"`
	HelpPopupMsg        string                      `json:"helpPopupMsg"`
}

func (srcProps TextBoxProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*TextBoxProperties, error) {

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

	if srcProps.ValueListID != nil {
		destValueListID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(*srcProps.ValueListID)
		destProps.ValueListID = &destValueListID
	}

	return &destProps, nil
}

func newDefaultTextBoxProperties() TextBoxProperties {
	props := TextBoxProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		Validation:                    newDefaultTextBoxValidationProperties(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  ""}
	return props
}
