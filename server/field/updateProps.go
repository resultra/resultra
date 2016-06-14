package field

import (
	"fmt"
)

type FieldIDInterface interface {
	GetFieldID() string
	GetParentTableID() string
}

type FieldIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func (idHeader FieldIDHeader) GetFieldID() string {
	return idHeader.FieldID
}

func (idHeader FieldIDHeader) GetParentTableID() string {
	return idHeader.ParentTableID
}

type FieldPropUpdater interface {
	FieldIDInterface

	// Normally, UpdateProps would be named updateProps if all the property updaters were in the same
	// pacakge. However, in this case, the calculated field formula is updated in the CalcField package
	// so the function name needs to start with an upper case, so a FieldPropUpdater defined
	// in the CalcField package can be used.
	UpdateProps(fieldForUpdate *Field) error
}

func UpdateFieldProps(propUpdater FieldPropUpdater) (*Field, error) {

	fieldForUpdate, getErr := GetField(propUpdater.GetParentTableID(), propUpdater.GetFieldID())
	if getErr != nil {
		return nil, getErr
	}

	// Do the actual property update through the FieldPropUpdater interface
	if propUpdateErr := propUpdater.UpdateProps(fieldForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateFieldProps: Unable to update existing field properties: %v", propUpdateErr)
	}

	updatedField, updateErr := UpdateExistingField(fieldForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateFieldProps: error updating field: %v", updateErr)
	}

	return updatedField, nil

}
