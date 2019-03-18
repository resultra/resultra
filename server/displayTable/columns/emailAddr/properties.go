package emailAddr

import (
	"fmt"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
)

type EmailAddrValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultEmailAddrValidationProperties() EmailAddrValidationProperties {
	return EmailAddrValidationProperties{false}
}

type EmailAddrProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          EmailAddrValidationProperties              `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps EmailAddrProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*EmailAddrProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultEmailAddrProperties() EmailAddrProperties {
	props := EmailAddrProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultEmailAddrValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
