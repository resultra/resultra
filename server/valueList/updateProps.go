// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package valueList

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic/stringValidation"
)

type ValueListIDInterface interface {
	getValueListID() string
}

type ValueListIDHeader struct {
	ValueListID string `json:"valueListID"`
}

func (idHeader ValueListIDHeader) getValueListID() string {
	return idHeader.ValueListID
}

type ValueListPropUpdater interface {
	ValueListIDInterface
	updateProps(valueList *ValueList) error
}

func updateValueListProps(trackerDBHandle *sql.DB, propUpdater ValueListPropUpdater) (*ValueList, error) {

	// Retrieve the bar chart from the data store
	valueListForUpdate, getErr := GetValueList(trackerDBHandle, propUpdater.getValueListID())
	if getErr != nil {
		return nil, fmt.Errorf("updateValueListProps: Unable to get existing button: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(valueListForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateValueListProps: Unable to update existing form link properties: %v", propUpdateErr)
	}

	updatedValueList, updateErr := updateExistingValueList(trackerDBHandle, valueListForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf(
			"updateValueListProps: Unable to update existing value list properties: datastore update error =  %v", updateErr)
	}

	return updatedValueList, nil
}

type ValueListNameParams struct {
	ValueListIDHeader
	NewName string `json:"newName"`
}

func (updateParams ValueListNameParams) updateProps(valueListForUpdate *ValueList) error {

	if !stringValidation.WellFormedItemLabel(updateParams.NewName) {
		return fmt.Errorf("Can't update form link name: invalid name: %v", updateParams.NewName)
	}

	valueListForUpdate.Name = updateParams.NewName

	return nil
}

type ValueListValuesParams struct {
	ValueListIDHeader
	Values []ValueListValue `json:"values"`
}

func (updateParams ValueListValuesParams) updateProps(valueListForUpdate *ValueList) error {

	valueListForUpdate.Properties.Values = updateParams.Values

	return nil
}
