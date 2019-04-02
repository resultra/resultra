// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package gauge

import (
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic/numberFormat"
	"github.com/resultra/resultra/server/generic/threshold"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type GaugeProperties struct {
	FieldID       string                                `json:"fieldID"`
	Geometry      componentLayout.LayoutGeometry        `json:"geometry"`
	MinVal        float64                               `json:"minVal"`
	MaxVal        float64                               `json:"maxVal"`
	ThresholdVals []threshold.ThresholdValues           `json:"thresholdVals"`
	LabelFormat   common.ComponentLabelFormatProperties `json:"labelFormat"`
	ValueFormat   numberFormat.NumberFormatProperties   `json:"valueFormat"`
	common.ComponentVisibilityProperties
	HelpPopupMsg string `json:"helpPopupMsg"`
}

func newDefaultGaugeProperties() GaugeProperties {
	props := GaugeProperties{
		FieldID:                       "",
		MinVal:                        0.0,
		MaxVal:                        100.0,
		ThresholdVals:                 []threshold.ThresholdValues{},
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		ValueFormat:                   numberFormat.DefaultNumberFormatProperties(),
		HelpPopupMsg:                  ""}
	return props

}

func (srcProps GaugeProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*GaugeProperties, error) {

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
