package userTag

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

type ValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() ValidationProperties {
	return ValidationProperties{ValueRequired: true}
}

type UserTagProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          ValidationProperties                       `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
	SelectableRoles     []string                                   `json:"selectableRoles"`
	CurrUserSelectable  bool                                       `json:"currUserSelectable"`
}

func (srcProps UserTagProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*UserTagProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destProps.SelectableRoles = uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.SelectableRoles)

	return &destProps, nil
}

func newDefaultUserTagProperties() UserTagProperties {
	props := UserTagProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        "",
		SelectableRoles:     []string{},
		CurrUserSelectable:  false}
	return props
}
