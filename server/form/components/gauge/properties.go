package gauge

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/numberFormat"
	"resultra/datasheet/server/generic/uniqueID"
)

type ThresholdValues struct {
	StartingVal float64 `json:"startingVal"`
	ColorScheme string  `json:"colorScheme"`
}

type GaugeProperties struct {
	FieldID       string                                `json:"fieldID"`
	Geometry      componentLayout.LayoutGeometry        `json:"geometry"`
	MinVal        float64                               `json:"minVal"`
	MaxVal        float64                               `json:"maxVal"`
	ThresholdVals []ThresholdValues                     `json:"thresholdVals"`
	LabelFormat   common.ComponentLabelFormatProperties `json:"labelFormat"`
	ValueFormat   numberFormat.NumberFormatProperties   `json:"valueFormat"`
	common.ComponentVisibilityProperties
}

func newDefaultGaugeProperties() GaugeProperties {
	props := GaugeProperties{
		FieldID:                       "",
		MinVal:                        0.0,
		MaxVal:                        100.0,
		ThresholdVals:                 []ThresholdValues{},
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		ValueFormat:                   numberFormat.DefaultNumberFormatProperties()}
	return props

}

func (srcProps GaugeProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*GaugeProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}
