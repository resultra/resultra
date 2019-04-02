// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package textInput

import (
	"fmt"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type TextInputValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultTextInputValidationProperties() TextInputValidationProperties {
	return TextInputValidationProperties{false}
}

type TextInputProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	ValueListID         *string                                    `json:"valueListID,omitempty"`
	Validation          TextInputValidationProperties              `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps TextInputProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*TextInputProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	if srcProps.ValueListID != nil {
		destValueListID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(*srcProps.ValueListID)
		destProps.ValueListID = &destValueListID
	}

	return &destProps, nil
}

func newDefaultTextInputProperties() TextInputProperties {
	props := TextInputProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultTextInputValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
