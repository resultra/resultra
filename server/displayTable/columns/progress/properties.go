// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package progress

import (
	"fmt"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic/numberFormat"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type ThresholdValues struct {
	StartingVal float64 `json:"startingVal"`
	ColorScheme string  `json:"colorScheme"`
}

type ProgressProperties struct {
	FieldID       string                                `json:"fieldID"`
	MinVal        float64                               `json:"minVal"`
	MaxVal        float64                               `json:"maxVal"`
	ThresholdVals []ThresholdValues                     `json:"thresholdVals"`
	LabelFormat   common.ComponentLabelFormatProperties `json:"labelFormat"`
	ValueFormat   numberFormat.NumberFormatProperties   `json:"valueFormat"`
	HelpPopupMsg  string                                `json:"helpPopupMsg"`
}

func newDefaultProgressProperties() ProgressProperties {
	props := ProgressProperties{
		FieldID:       "",
		MinVal:        0.0,
		MaxVal:        100.0,
		ThresholdVals: []ThresholdValues{},
		LabelFormat:   common.NewDefaultLabelFormatProperties(),
		ValueFormat:   numberFormat.DefaultNumberFormatProperties(),
		HelpPopupMsg:  ""}
	return props

}

func (srcProps ProgressProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*ProgressProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}
