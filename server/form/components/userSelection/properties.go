package userSelection

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
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

type UserSelectionProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
	Validation          ValidationProperties `json:"validation"`
	ClearValueSupported bool                 `json:"clearValueSupported"`
	HelpPopupMsg        string               `json:"helpPopupMsg"`
	SelectableRoles     []string             `json:"selectableRoles"`
	CurrUserSelectable  bool                 `json:"currUserSelectable"`
}

func (srcProps UserSelectionProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*UserSelectionProperties, error) {

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

	destProps.SelectableRoles = uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.SelectableRoles)

	return &destProps, nil
}

func newDefaultUserSelectionProperties() UserSelectionProperties {
	props := UserSelectionProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		Validation:                    newDefaultValidationProperties(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  "",
		SelectableRoles:               []string{},
		CurrUserSelectable:            false}
	return props
}
