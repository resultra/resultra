// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package checkBox

import (
	"fmt"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/trackerDatabase"
)

const CheckboxColorSchemeDefault string = "default"

type CheckBoxValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() CheckBoxValidationProperties {
	return CheckBoxValidationProperties{
		ValueRequired: true}
}

type CheckBoxProperties struct {
	FieldID                string                                     `json:"fieldID"`
	ColorScheme            string                                     `json:"colorScheme"`
	StrikethroughCompleted bool                                       `json:"strikethroughCompleted"`
	LabelFormat            common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions            common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation             CheckBoxValidationProperties               `json:"validation"`
	ClearValueSupported    bool                                       `json:"clearValueSupported"`
	HelpPopupMsg           string                                     `json:"helpPopupMsg"`
}

func (srcProps CheckBoxProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*CheckBoxProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultCheckBoxProperties() CheckBoxProperties {

	props := CheckBoxProperties{
		LabelFormat:            common.NewDefaultLabelFormatProperties(),
		ColorScheme:            CheckboxColorSchemeDefault,
		StrikethroughCompleted: false,
		Permissions:            common.NewDefaultComponentValuePermissionsProperties(),
		Validation:             newDefaultValidationProperties(),
		ClearValueSupported:    false,
		HelpPopupMsg:           ""}
	return props
}
