package urlLink

import (
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
)

type UrlLinkValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultUrlLinkValidationProperties() UrlLinkValidationProperties {
	return UrlLinkValidationProperties{false}
}

type UrlLinkProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
	Validation          UrlLinkValidationProperties `json:"validation"`
	ClearValueSupported bool                        `json:"clearValueSupported"`
	HelpPopupMsg        string                      `json:"helpPopupMsg"`
}

func (srcProps UrlLinkProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*UrlLinkProperties, error) {

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

func newDefaultUrlLinkProperties() UrlLinkProperties {
	props := UrlLinkProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		Validation:                    newDefaultUrlLinkValidationProperties(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  ""}
	return props
}
