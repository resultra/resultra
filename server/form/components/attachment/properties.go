// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package attachment

import (
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
)

type ValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() ValidationProperties {
	return ValidationProperties{
		ValueRequired: true}
}

type ImageProperties struct {
	FieldID      string                                     `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat  common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions  common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation   ValidationProperties                       `json:"validation"`
	HelpPopupMsg string                                     `json:"helpPopupMsg"`
}

func (srcProps ImageProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*ImageProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultAttachmentProperties() ImageProperties {
	props := ImageProperties{
		LabelFormat:  common.NewDefaultLabelFormatProperties(),
		Permissions:  common.NewDefaultComponentValuePermissionsProperties(),
		Validation:   newDefaultValidationProperties(),
		HelpPopupMsg: ""}
	return props
}
