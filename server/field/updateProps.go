package field

import (
	"appengine"
	"fmt"
)

type FieldIDInterface interface {
	GetFieldID() string
}

type FieldIDHeader struct {
	FieldID string `json:"fieldID"`
}

func (idHeader FieldIDHeader) GetFieldID() string {
	return idHeader.FieldID
}

type FieldPropUpdater interface {
	FieldIDInterface

	// Normally, UpdateProps would be named updateProps if all the property updaters were in the same
	// pacakge. However, in this case, the calculated field formula is updated in the CalcField package
	// so the function name needs to start with an upper case, so a FieldPropUpdater defined
	// in the CalcField package can be used.
	UpdateProps(appEngContext appengine.Context, fieldForUpdate *Field) error
}

func UpdateFieldProps(appEngContext appengine.Context, propUpdater FieldPropUpdater) (*FieldRef, error) {

	fieldForUpdate, getErr := GetField(appEngContext, propUpdater.GetFieldID())
	if getErr != nil {
		return nil, getErr
	}

	// Do the actual property update through the FieldPropUpdater interface
	if propUpdateErr := propUpdater.UpdateProps(appEngContext, fieldForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateFieldProps: Unable to update existing field properties: %v", propUpdateErr)
	}

	updatedFieldRef, updateErr := UpdateExistingField(appEngContext, propUpdater.GetFieldID(), fieldForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateFieldProps: error updating field: %v", updateErr)
	}

	return updatedFieldRef, nil

}
