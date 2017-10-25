package file

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/trackerDatabase"
)

type FileValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultFileValidationProperties() FileValidationProperties {
	return FileValidationProperties{false}
}

type FileProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          FileValidationProperties                   `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps FileProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*FileProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultFileProperties() FileProperties {
	props := FileProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultFileValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
