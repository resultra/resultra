// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formButton

import (
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/common/inputProps"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/record"
	"resultra/tracker/server/trackerDatabase"
)

const buttonSizeMedium string = "medium"
const colorSchemeDefault string = "default"
const buttonIconNone string = "none"

type ButtonProperties struct {
	Geometry          componentLayout.LayoutGeometry             `json:"geometry"`
	LinkedFormID      string                                     `json:"linkedFormID"`
	PopupBehavior     inputProps.ButtonPopupBehavior             `json:"popupBehavior"`
	Size              string                                     `json:"size"`
	ColorScheme       string                                     `json:"colorScheme"`
	Icon              string                                     `json:"icon"`
	DefaultValues     []record.DefaultFieldValue                 `json:"defaultValues"`
	ButtonLabelFormat inputProps.FormButtonLabelFormatProperties `json:"buttonLabelFormat"`

	common.ComponentVisibilityProperties
}

func (srcProps ButtonProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*ButtonProperties, error) {

	destProps := srcProps

	destProps.LinkedFormID = cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcProps.LinkedFormID)

	destPopupProps, cloneErr := srcProps.PopupBehavior.Clone(cloneParams.IDRemapper)
	if cloneErr != nil {
		return nil, fmt.Errorf("ButtonProperties.Clone: %v", cloneErr)
	}
	destProps.PopupBehavior = *destPopupProps

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	destDefaultVals, cloneErr := record.CloneDefaultFieldValues(cloneParams.IDRemapper, srcProps.DefaultValues)
	if cloneErr != nil {
		return nil, fmt.Errorf("ButtonPopupBehavior.Clone: %v", cloneErr)
	}
	destProps.DefaultValues = destDefaultVals

	return &destProps, nil
}

func newDefaultButtonProperties() ButtonProperties {

	return ButtonProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		PopupBehavior:                 inputProps.NewDefaultPopupBehavior(),
		ButtonLabelFormat:             inputProps.NewDefaultFormButtonLabelFormatProperties(),
		Size:                          buttonSizeMedium,
		ColorScheme:                   colorSchemeDefault,
		DefaultValues:                 []record.DefaultFieldValue{},
		Icon:                          buttonIconNone}
}
