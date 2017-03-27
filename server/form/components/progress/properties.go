package progress

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

type ProgressProperties struct {
	FieldID       string                                `json:"fieldID"`
	Geometry      componentLayout.LayoutGeometry        `json:"geometry"`
	MinVal        float64                               `json:"minVal"`
	MaxVal        float64                               `json:"maxVal"`
	ThresholdVals []ThresholdValues                     `json:"thresholdVals"`
	LabelFormat   common.ComponentLabelFormatProperties `json:"labelFormat"`
	ValueFormat   numberFormat.NumberFormatProperties   `json:"valueFormat"`
	common.ComponentVisibilityProperties
}

func newDefaultProgressProperties() ProgressProperties {
	props := ProgressProperties{
		FieldID:                       "",
		MinVal:                        0.0,
		MaxVal:                        100.0,
		ThresholdVals:                 []ThresholdValues{},
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		ValueFormat:                   numberFormat.DefaultNumberFormatProperties()}
	return props

}

func (srcProps ProgressProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ProgressProperties, error) {

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
