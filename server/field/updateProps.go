// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package field

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
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
	UpdateProps(trackerDBHandle *sql.DB, fieldForUpdate *Field) error
}

func processFieldPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater FieldPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedField, err := UpdateFieldProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedField)
	}
}

func UpdateFieldProps(trackerDBHandle *sql.DB, propUpdater FieldPropUpdater) (*Field, error) {

	fieldForUpdate, getErr := GetField(trackerDBHandle, propUpdater.GetFieldID())
	if getErr != nil {
		return nil, getErr
	}

	// Do the actual property update through the FieldPropUpdater interface
	if propUpdateErr := propUpdater.UpdateProps(trackerDBHandle, fieldForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateFieldProps: Unable to update existing field properties: %v", propUpdateErr)
	}

	updatedField, updateErr := UpdateExistingField(trackerDBHandle, fieldForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateFieldProps: error updating field: %v", updateErr)
	}

	return updatedField, nil

}

type SetFieldNameParams struct {
	FieldIDHeader
	NewFieldName string `json:"newFieldName"`
}

func (updateParams SetFieldNameParams) UpdateProps(trackerDBHandle *sql.DB, field *Field) error {

	// TODO - Need some kind of mutex/transaction around the name update
	if validateErr := validateExistingFieldName(trackerDBHandle, field.FieldID, updateParams.NewFieldName); validateErr != nil {
		return validateErr
	}

	field.Name = updateParams.NewFieldName

	return nil
}

type SetFieldRefNameParams struct {
	FieldIDHeader
	NewFieldRefName string `json:"newFieldRefName"`
}

func (updateParams SetFieldRefNameParams) UpdateProps(trackerDBHandle *sql.DB, field *Field) error {

	// TODO - Need some kind of mutex/transaction around the name update
	if validateErr := validateExistingFieldRefName(trackerDBHandle, field.FieldID,
		updateParams.NewFieldRefName); validateErr != nil {
		return validateErr
	}

	field.RefName = updateParams.NewFieldRefName

	return nil
}
