package urlLink

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type UrlLinkValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultUrlLinkValidationProperties() UrlLinkValidationProperties {
	return UrlLinkValidationProperties{false}
}

type UrlLinkProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          UrlLinkValidationProperties              `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps UrlLinkProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*UrlLinkProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultUrlLinkProperties() UrlLinkProperties {
	props := UrlLinkProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultUrlLinkValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
