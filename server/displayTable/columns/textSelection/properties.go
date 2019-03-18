// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package textSelection

import (
	"fmt"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
)

type TextSelectionValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultTextSelectionValidationProperties() TextSelectionValidationProperties {
	return TextSelectionValidationProperties{false}
}

type TextSelectionProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	ValueListID         *string                                    `json:"valueListID,omitempty"`
	Validation          TextSelectionValidationProperties          `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps TextSelectionProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*TextSelectionProperties, error) {

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

func newDefaultTextSelectionProperties() TextSelectionProperties {
	props := TextSelectionProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultTextSelectionValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
