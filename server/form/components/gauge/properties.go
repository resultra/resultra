package gauge

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/numberFormat"
	"resultra/datasheet/server/generic/threshold"
	"resultra/datasheet/server/trackerDatabase"
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
