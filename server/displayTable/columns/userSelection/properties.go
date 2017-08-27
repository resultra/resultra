package userSelection

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type ValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() ValidationProperties {
	return ValidationProperties{ValueRequired: true}
}

type UserSelectionProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          ValidationProperties                       `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
	SelectableRoles     []string                                   `json:"selectableRoles"`
	CurrUserSelectable  bool                                       `json:"currUserSelectable"`
}

func (srcProps UserSelectionProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*UserSelectionProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destProps.SelectableRoles = uniqueID.CloneIDList(remappedIDs, srcProps.SelectableRoles)

	return &destProps, nil
}

func newDefaultUserSelectionProperties() UserSelectionProperties {
	props := UserSelectionProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        "",
		SelectableRoles:     []string{},
		CurrUserSelectable:  false}
	return props
}
