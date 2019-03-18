// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userTag

import (
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

type ValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() ValidationProperties {
	return ValidationProperties{ValueRequired: true}
}

type UserTagProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
	Validation          ValidationProperties `json:"validation"`
	ClearValueSupported bool                 `json:"clearValueSupported"`
	HelpPopupMsg        string               `json:"helpPopupMsg"`
	SelectableRoles     []string             `json:"selectableRoles"`
	CurrUserSelectable  bool                 `json:"currUserSelectable"`
}

func (srcProps UserTagProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*UserTagProperties, error) {

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

	destProps.SelectableRoles = uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.SelectableRoles)

	return &destProps, nil
}

func newDefaultUserTagProperties() UserTagProperties {
	props := UserTagProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		Validation:                    newDefaultValidationProperties(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  "",
		SelectableRoles:               []string{},
		CurrUserSelectable:            false}
	return props
}
