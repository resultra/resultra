package field

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/api"
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
	UpdateProps(fieldForUpdate *Field) error
}

func processFieldPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater FieldPropUpdater) {
	if updatedField, err := UpdateFieldProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedField)
	}
}

func UpdateFieldProps(propUpdater FieldPropUpdater) (*Field, error) {

	fieldForUpdate, getErr := GetField(propUpdater.GetFieldID())
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

type SetFieldNameParams struct {
	FieldIDHeader
	NewFieldName string `json:"newFieldName"`
}

func (updateParams SetFieldNameParams) UpdateProps(field *Field) error {

	if validateErr := validateExistingFieldName(field.FieldID, updateParams.NewFieldName); validateErr != nil {
		return validateErr
	}

	field.Name = updateParams.NewFieldName

	return nil
}
