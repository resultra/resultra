package valueList

import (
	"fmt"
	"resultra/datasheet/server/generic/stringValidation"
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

func updateValueListProps(propUpdater ValueListPropUpdater) (*ValueList, error) {

	// Retrieve the bar chart from the data store
	valueListForUpdate, getErr := GetValueList(propUpdater.getValueListID())
	if getErr != nil {
		return nil, fmt.Errorf("updateValueListProps: Unable to get existing button: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(valueListForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateValueListProps: Unable to update existing form link properties: %v", propUpdateErr)
	}

	updatedValueList, updateErr := updateExistingValueList(valueListForUpdate)
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
